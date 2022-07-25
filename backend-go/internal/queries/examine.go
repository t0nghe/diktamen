package queries

import (
	"backend-go/graph/model"
	dbconn "backend-go/internal/pkg/database"
)

// Queries from DB to get sentences in an article that the user has heard correctly.
func QueryExamineCorrectSents(userId int, articleId int) ([]*model.SeenSent, error) {

	// Queries the sent_ids, indexes_in_article and user input (try_text) on the last trial, if correct.
	// When a sentence has a try_score, it certainly is seen. We don't need to add a filter for that. Since sentences graduate the first time the user gets it right, if the try_score is greater than 99.7, we know it's their last try on this sentence.
	// Checking for the incorrect sents may not be this straightforward.
	stmtEcs, err := dbconn.Db.Prepare("SELECT s.id, s.index_in_article, ud.try_text FROM user_dictation ud LEFT JOIN sent s ON ud.sent_id=s.id WHERE ud.try_score > .997 AND ud.user_id=? AND s.article_id=?;")
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

func QueryExamineIncorrectSents() {
	panic("to implement")
}
