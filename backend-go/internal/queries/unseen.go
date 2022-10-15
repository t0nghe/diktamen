package queries

import (
	"backend-go/graph/model"
	dbconn "backend-go/internal/pkg/database"
	"database/sql"
)

// QueryUserUnseenSents searches and returns sentences a user hasn't seen in an article.
func QueryUserUnseenSents(userId int, articleId int) ([]*model.UnseenSent, error) {
	var result []*model.UnseenSent

	// Find out sentence index the user has reached in an article.
	stmt1, err := dbconn.Db.Prepare("SELECT finished_sent_index FROM user_article WHERE user_id=$1 AND article_id=$2;")
	if err != nil {
		return nil, err
	}
	defer stmt1.Close()

	var finishedSentIndex int
	err = stmt1.QueryRow(userId, articleId).Scan(&finishedSentIndex)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			finishedSentIndex = 0
		}
	}

	// Get relevant fields in unseen sentences, namely sentences with index > finished index.
	stmt2, err := dbconn.Db.Prepare("SELECT id, index_in_article, media_uri FROM sent WHERE article_id=$1 AND index_in_article>$2;")
	if err != nil {
		return nil, err
	}
	defer stmt2.Close()
	rows, err := stmt2.Query(articleId, finishedSentIndex)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var aid, iia int
		var mu string

		err = rows.Scan(&aid, &iia, &mu)
		if err != nil {
			return nil, err
		}

		result = append(result, &model.UnseenSent{SentID: &aid, IndexInArticle: &iia, MediaURI: &mu})
	}

	// Add wordform and metadata for words in this sentence.
	for i := 0; i < len(result); i++ {
		words, err := GetUnseenSentWords(*result[i].SentID)
		if err != nil {
			return nil, err
		}
		result[i].SentWords = *words
	}

	return result, nil
}

// GetUnseenSentWords gets wordforms and metadata about a word in a sentence out of `sent` table.
func GetUnseenSentWords(sentId int) (*[]*model.UnseenSentWord, error) {
	var result []*model.UnseenSentWord

	stmt, err := dbconn.Db.Prepare("SELECT length, is_cloze, wordform, index_in_sent FROM sent_word WHERE sent_id=$1;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sentId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var wl int    // word length
		var wic bool  // word is cloze
		var wf string // wordform
		var iis int   // index in sent

		err := rows.Scan(&wl, &wic, &wf, &iis)
		if err != nil {
			return nil, err
		}
		result = append(result, &model.UnseenSentWord{Length: wl, IsCloze: wic, Wordform: wf, IndexInSent: iis})
	}

	return &result, nil
}
