package queries

import (
	"backend-go/graph/model"
	dbconn "backend-go/internal/pkg/database"
	"database/sql"
	"log"
)

// QueryUserUnseenSents searches and returns sentences a user hasn't seen in an article.
func QueryUserUnseenSents(userId int, articleId int) ([]*model.UnseenSent, error) {
	var result []*model.UnseenSent

	// Find out sentence index the user has reached in an article.
	row := dbconn.Db.QueryRow("SELECT finished_sent_index FROM user_article WHERE user_id=$1 AND article_id=$2;", userId, articleId)

	var finishedSentIndex int
	err := row.Scan(&finishedSentIndex)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			log.Println("SELECT finished_sent_index FROM user_article WHERE user_id=$1 AND article_id=$2;")
			return nil, err
		} else {
			finishedSentIndex = 0
		}
	}

	// Get relevant fields in unseen sentences, namely sentences with index > finished index.
	rows, err := dbconn.Db.Query("SELECT id, index_in_article, media_uri FROM sent WHERE article_id=$1 AND index_in_article>$2;", articleId, finishedSentIndex)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var aid, iia int
		var mu string

		err = rows.Scan(&aid, &iia, &mu)
		if err != nil {
			log.Println(err)
			log.Println("SELECT id, index_in_article, media_uri FROM sent WHERE article_id=$1 AND index_in_article>$2;")
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

	rows, err := dbconn.Db.Query("SELECT length, is_cloze, wordform, index_in_sent FROM sent_word WHERE sent_id=$1;", sentId)
	if err != nil {
		log.Println("sentId: ", sentId)
		log.Println("SELECT length, is_cloze, wordform, index_in_sent FROM sent_word WHERE sent_id=$1;")
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

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
