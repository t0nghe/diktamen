package queries

import (
	"backend-go/graph/model"
	dbconn "backend-go/internal/pkg/database"
)

func GetTriedSentWords(userId int, sentId int) ([]*model.TriedSentWord, error) {
	stmtTsw, err := dbconn.Db.Prepare(`SELECT index_in_sent, wordform, is_cloze, length, IFNULL(user_input_score, 0), IFNULL(user_input_wordform, '') FROM (SELECT wordform, is_cloze, index_in_sent, length FROM sent_word WHERE sent_id=$1) sw
    LEFT JOIN (
        SELECT sent_id, index_in_sent AS iis_right, user_input_score, user_input_wordform
        FROM user_dictation_word
        WHERE user_dictation_id=(
            SELECT id FROM user_dictation WHERE try_count=(
                SELECT MAX(try_count) FROM user_dictation WHERE sent_id=$2 AND user_id=$3
            )
        AND sent_id=$4 AND user_id=$5
    )) qr
    ON sw.index_in_sent=qr.iis_right;`)
	if err != nil {
		return nil, err
	}
	defer stmtTsw.Close()

	rows, err := stmtTsw.Query(sentId, sentId, userId, sentId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var triedSentWords []*model.TriedSentWord
	for rows.Next() {
		var indexInSent, length int
		var wordform, userInputWordform string
		var isCloze bool
		var userInputScore float64
		err := rows.Scan(&indexInSent, &wordform, &isCloze, &length, &userInputScore, &userInputWordform)
		if err != nil {
			return nil, err
		}
		triedSentWords = append(triedSentWords, &model.TriedSentWord{
			Length: length, IsCloze: isCloze, Wordform: wordform, IndexInSent: indexInSent, LastInputText: userInputWordform, LastInputScore: userInputScore,
		})
	}

	return triedSentWords, nil
}

// Queries from DB to get sentences in an article that the user has heard correctly.
func QueryExamineCorrectSents(userId int, articleId int) ([]*model.SeenSent, error) {

	// Queries the sent_ids, indexes_in_article and user input (try_text) on the last trial, if correct.
	// When a sentence has a try_score, it certainly is seen. We don't need to add a filter for that. Since sentences graduate the first time the user gets it right, if the try_score is greater than 99.7, we know it's their last try on this sentence.
	// Checking for the incorrect sents may not be this straightforward.
	stmtEcs, err := dbconn.Db.Prepare("SELECT s.id, s.index_in_article, ud.try_text FROM user_dictation ud LEFT JOIN sent s ON ud.sent_id=s.id WHERE ud.try_score > .998 AND ud.user_id=$1 AND s.article_id=$2;")
	if err != nil {
		return nil, err
	}
	defer stmtEcs.Close()

	rows, err := stmtEcs.Query(userId, articleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*model.SeenSent
	for rows.Next() {
		var sid int
		var iia int
		var txt string
		err = rows.Scan(&sid, &iia, &txt)
		if err != nil {
			return nil, err
		}
		result = append(result, &model.SeenSent{SentID: &sid, IndexInArticle: &iia, TryText: txt})
	}

	return result, nil
}

func QueryExamineIncorrectSents(userId int, articleId int) ([]*model.IncorrectSeenSent, error) {

	stmtIncoSents, err := dbconn.Db.Prepare(`SELECT id, index_in_article FROM sent WHERE id IN (SELECT ud.sent_id FROM user_dictation ud LEFT JOIN (SELECT sent_id, MAX(try_count) AS mtc FROM user_dictation WHERE user_id=$1 AND sent_id IN (SELECT id FROM sent WHERE article_id=$2) GROUP BY sent_id) umtc ON ud.sent_id=umtc.sent_id WHERE ud.try_count=umtc.mtc AND try_score<=0.998 AND user_id=$3);`)
	if err != nil {
		return nil, err
	}
	defer stmtIncoSents.Close()

	rows, err := stmtIncoSents.Query(userId, articleId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incoSeenSents []*model.IncorrectSeenSent
	for rows.Next() {
		var sentId int
		var indexInArticle int
		err := rows.Scan(&sentId, &indexInArticle)
		if err != nil {
			return nil, err
		}
		sentWords, err := GetTriedSentWords(userId, sentId)
		if err != nil {
			return nil, err
		}
		incoSeenSents = append(incoSeenSents, &model.IncorrectSeenSent{
			SentID:         &sentId,
			IndexInArticle: &indexInArticle,
			SentWords:      sentWords,
		})
	}

	return incoSeenSents, nil
}
