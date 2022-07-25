package queries

import (
	dbconn "backend-go/internal/pkg/database"
	"backend-go/internal/pkg/score"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type userDictationRecord struct {
	nextTryTime time.Time
	sentId      int
	tryCount    int
	tryScore    float64
	tryText     string
	tryTime     time.Time
	userId      int
}

type userDictationWordRecord struct {
	indexInSent       int
	sentId            int
	userDictationId   int64
	userId            int
	userInputScore    float64
	userInputWordform string
}

func BumpUpFinishedIndex(userId int, sentId int) error {
	fmt.Println("BumpUpFinishedIndex")
	// get article_id; index_in_article
	stmt, err := dbconn.Db.Prepare("SELECT article_id, index_in_article FROM sent WHERE id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	var articleId, indexInArticle int // article_id; index_in_article
	err = stmt.QueryRow(sentId).Scan(&articleId, &indexInArticle)
	if err != nil {
		return err // This error includes the case where sent.id doesn't exist.
	}

	// ALAS... MySQL doens't support IF NOT EXISTS () ... ELSE...

	stmtExists, err := dbconn.Db.Prepare("SELECT IF (EXISTS (SELECT id FROM user_article WHERE user_id=? AND article_id = ?), 1, 0);")
	if err != nil {
		return err
	}
	defer stmtExists.Close()
	var exists int
	_ = stmtExists.QueryRow(userId, articleId).Scan(&exists)

	if exists == 1 {
		stmtBumpUp, err := dbconn.Db.Prepare("UPDATE user_article SET finished_sent_index=? WHERE finished_sent_index<? AND user_id=? AND article_id=?;")
		if err != nil {
			return err
		}
		defer stmtBumpUp.Close()
		result, err := stmtBumpUp.Exec(indexInArticle, indexInArticle, userId, articleId)
		if err != nil {
			return err
		}
		rc, _ := result.RowsAffected() // rows count
		fmt.Println(rc)
	} else {
		stmtInsert, _ := dbconn.Db.Prepare("INSERT INTO user_article (article_id, finished_sent_index, user_id) VALUES (?, ?, ?);")
		defer stmtInsert.Close()
		resultInsert, err := stmtInsert.Exec(articleId, indexInArticle, userId)
		if err != nil {
			return err
		}
		rcInsert, _ := resultInsert.RowsAffected()
		fmt.Println(rcInsert)
	}

	return nil
}

// GetThisTryCount returns n, where n indicates the n-th time this is that a user tries a sentence. n is 1-based. If there is an error, 0 is returned.
func GetThisTryCount(userId int, sentId int) (int, error) {
	// Query SQL to get previous trycount+1.
	stmtTryCount, err := dbconn.Db.Prepare("SELECT IF (EXISTS (SELECT try_count FROM user_dictation WHERE user_id=? AND sent_id=?), MAX(try_count) + 1, 1 ) AS this_try_count FROM user_dictation WHERE user_id=? AND sent_id=?;")
	if err != nil {
		return 0, err
	}
	defer stmtTryCount.Close()

	var thisTryCount int
	err = stmtTryCount.QueryRow(userId, sentId, userId, sentId).Scan(&thisTryCount)
	if err != nil {
		return 0, err
	}
	return thisTryCount, nil
}

func GetWordforms(sentId int) ([]int, []string, []bool, []int, error) {
	stmtWordforms, err := dbconn.Db.Prepare("SELECT index_in_sent, wordform, is_cloze, length FROM sent_word WHERE sent_id=?;")
	if err != nil {
		return nil, nil, nil, nil, err // haven't thought through what the return values look like
	}
	defer stmtWordforms.Close()

	rows, err := stmtWordforms.Query(sentId)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	var indexesInSent []int
	var wordforms []string
	var areClozes []bool
	var lengths []int

	for rows.Next() {
		var index int
		var wordform string
		var isCloze bool
		var length int
		// 220724 my understanding is Scan() method assigns values it gets to these destinations.
		// Next on line 105 when we append `index`, we are appending a value.
		// On the next iteration, Scan() places new value to that mem addr, the value changes
		// and again when appending, we'll append the new value.
		err := rows.Scan(&index, &wordform, &isCloze, &length)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		indexesInSent = append(indexesInSent, index)
		wordforms = append(wordforms, wordform)
		areClozes = append(areClozes, isCloze)
		lengths = append(lengths, length)
	}
	return indexesInSent, wordforms, areClozes, lengths, nil
}

// InsertUserDication inserts a record to user_dictation table and returns ID of the inserted record and a possible error.
func InsertUserDictation(input userDictationRecord) (int64, error) {
	stmtUD, err := dbconn.Db.Prepare("INSERT INTO user_dictation (next_try_time, sent_id, try_count, try_score, try_text, try_time, user_id) VALUES (?,?,?,?,?,?,?);")
	if err != nil {
		return 0, err
	}
	defer stmtUD.Close()

	result, err := stmtUD.Exec(input.nextTryTime.Format(time.RFC3339), input.sentId, input.tryCount, input.tryScore, input.tryText, input.tryTime.Format(time.RFC3339), input.userId)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

// InsertUserDictationWord inserts a record to user_dictation_word, each user_dictation_word is one of the words associated with a user_dictation record.
func InsertUserDictationWord(input userDictationWordRecord) (int64, error) {
	stmtUDW, err := dbconn.Db.Prepare("INSERT INTO user_dictation_word (index_in_sent, sent_id, user_dictation_id, user_id, user_input_score, user_input_wordform) VALUES (?, ?, ?, ?, ?, ?);")
	if err != nil {
		return 0, err
	}
	defer stmtUDW.Close()

	result, err := stmtUDW.Exec(input.indexInSent, input.sentId, input.userDictationId, input.userId, input.userInputScore, input.userInputWordform)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func TrySentence(userId int, sentId int, userInput string) (float64, error) {
	/*
		STEPS
		1. SELECT article_id, index_in_article FROM sent WHERE sent_id;
			1.1. IN user_article, bump up `finished_sent_index` by one
			1.2. If no row, create record:
				article_id, user_id, finished_sent_index (index_in_article)
		2. SELECT MAX(try_count) AS prev_try_count FROM user_dictation WHERE user_id AND sent_id;
			2.1. this_try_count = prev_try_count + 1 // If no prev_try_count, this_try_count = 1
		3. SELECT index_in_sent, wordform, is_cloze, length FROM sent_word WHERE sent_id
			3.1. create slices for index_in_sent, wordform, is_cloze, length
			3.2. create slice for userInputWords (DONE);
			3.3. create float slice for userInputScores;
		4. try_text = concatinate userInputWords leave a space between tokens (for display only)
		5. try_score = [concat wordform]/[concat userInputWords] if is_cloze;
			5.1. and calculate edit distance of the two big strings
		6. calcualte userInputScores, also by edit distance.
			6.2. (length - editDistance) / length
		7. calculate next_try_time:
			7.1. if try score > 99.8%: now + 200 years
			7.2. else: 48hr * try_score * this_try_count
		8. INSERT INTO user_dictation:
			1. next_try_time (step 7), sent_id (input), try_count (step 2.1), try_score (step 5), try_text (step 4), try_time (now), user_id (input)
			2. we get ID from this INSERT operation (How do we get that?)
		9. for each word in userInputWords, if is_cloze:
			1. INSERT (index_in_sent, sent_id, user_dictation_id, user_id, user_input_score, user_input_wordform)
			* Not all words will be saved to the table. Only those is_cloze are.
		10. Return `try_score`
	*/

	/* STEP 1 - BUMP UP finished_sent_index */
	err := BumpUpFinishedIndex(userId, sentId)
	if err != nil {
		return 0, err
	}

	/* STEP 2 - GET TRY COUNT */
	thisTryCount, err := GetThisTryCount(userId, sentId)
	if err != nil {
		return 0, err
	}
	fmt.Println(thisTryCount) // To prevent errors.

	/* STEP 3 - CREATE SLICES */

	indexesInSent, wordforms, areClozes, _, err := GetWordforms(sentId)

	if err != nil {
		return 0, err
	}

	var userInputWords []string
	err = json.Unmarshal([]byte(userInput), &userInputWords)
	if err != nil {
		return 0, err
	}
	tryText := strings.Join(userInputWords, " ") // STEP 4 - try_text

	/* STEPS 5, 6 - calculate scores */
	textForScoring := ""
	userInputForScoring := ""
	userInputScores := []float64{}

	// DATA STRUCTURE TO SAVE TO TABEL?
	// WE CAN STILL USE A FULL LENGTH userInputWords. LEAVE THE VALUE AS AN EMPTY STRING IF ISCLOZE IS FALSE.
	// WHEN SAVING THE RECORDS TO user_dictation_word TABLE, WE ITER AGAIN OVER THESE SLICES, AND SKIP ITEMS WHERE is_cloze IS FALSE.

	for i := 0; i < len(indexesInSent); i++ {
		if areClozes[i] {
			textForScoring += wordforms[i]
			userInputForScoring += userInputWords[i]
		}
		userInputScores = append(userInputScores, score.Score(wordforms[i], userInputWords[i]))
	}
	tryScore := score.Score(textForScoring, userInputForScoring)

	/* STEP 7 - next_try_time */
	var nextTryTime time.Time
	if tryScore > 0.998 {
		// 10 years from now
		nextTryTime = time.Now().Add(10 * 365 * 24 * time.Hour)
	} else {
		nextTryTime = time.Now().Add(time.Duration(float64(thisTryCount*48*60)*tryScore) * time.Minute)
	}

	/* STEP 8 - INSERT INTO user_dictation */
	userDictationToInsert := userDictationRecord{
		nextTryTime: nextTryTime,
		sentId:      sentId,
		tryCount:    thisTryCount,
		tryScore:    tryScore,
		tryText:     tryText,
		tryTime:     time.Now(),
		userId:      userId,
	}

	userDictationId, err := InsertUserDictation(userDictationToInsert)
	if err != nil {
		return 0, err
	}

	/* STEP 9 - INSERT RECORDS TO user_dication_word */
	for j := 0; j < len(indexesInSent); j++ {
		if areClozes[j] {
			var udw userDictationWordRecord
			udw.indexInSent = indexesInSent[j]
			udw.sentId = sentId
			udw.userDictationId = userDictationId
			udw.userId = userId
			udw.userInputScore = userInputScores[j]
			udw.userInputWordform = userInputWords[j]

			_, err := InsertUserDictationWord(udw)
			if err != nil {
				return 0.1, err
			}
		}
	}

	/* STEP 10 - RETURN tryScore */
	return tryScore, nil
}
