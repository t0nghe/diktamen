package queries

import (
	"backend-go/graph/model"
	dbconn "backend-go/internal/pkg/database"
	"log"
)

// QuerySeenSentIds returns the sent_ids of sentences that the user has tried in an articlee.
func QuerySeenSentIds(userId int, articleId int) ([]int, []int, error) {
	// We need to query user_article to get `finished_sent_index`

	// Then query sent, to get `sent_id` field where `index_in_article` is lower than or equal to `finished_sent_index`
	rows, err := dbconn.Db.Query("SELECT DISTINCT s.id, s.index_in_article FROM sent s LEFT JOIN user_article ua ON s.article_id=ua.article_id WHERE ua.user_id=$1 AND ua.article_id=$2 AND ua.finished_sent_index >= s.index_in_article;", userId, articleId)

	if err != nil {
		return []int{0}, []int{0}, err
	}
	defer rows.Close()

	var sentIds []int
	var indexesInArticle []int
	for rows.Next() {
		var sentId int
		var indInArt int
		err = rows.Scan(&sentId, &indInArt)
		if err != nil {
			return []int{0}, []int{0}, nil
		}
		sentIds = append(sentIds, sentId)
		indexesInArticle = append(indexesInArticle, indInArt)
	}

	return sentIds, indexesInArticle, nil
}

func QueryLastTryText(userId int, sentId int) (string, error) {
	row := dbconn.Db.QueryRow("SELECT try_text FROM user_dictation WHERE (try_count) IN (SELECT MAX(try_count) FROM user_dictation WHERE user_id=$1 AND sent_id=$2) AND user_id=$3 AND sent_id=$4;", userId, sentId, userId, sentId)

	var result string
	err := row.Scan(&result)
	if err != nil {
		log.Println(row)
		return "", err
	}

	return result, nil
}

func QuerySeenSents(userId int, articleId int) ([]*model.SeenSent, error) {
	// Get sentence IDs that are seen in an article
	seenSentIds, indexesInArticle, err := QuerySeenSentIds(userId, articleId)
	if err != nil {
		return nil, err
	}

	var result []*model.SeenSent
	for k := 0; k < len(seenSentIds); k++ {
		var seenSent model.SeenSent
		seenSent.SentID = &seenSentIds[k] // Could this be buggy?
		seenSent.IndexInArticle = &indexesInArticle[k]
		tryText, err := QueryLastTryText(userId, seenSentIds[k])
		if err != nil {
			return nil, err
		}
		seenSent.TryText = tryText

		result = append(result, &seenSent)
	}

	return result, nil
}
