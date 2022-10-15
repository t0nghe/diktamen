package queries

import (
	"backend-go/graph/model"
	dbconn "backend-go/internal/pkg/database"
	"time"
)

func QueryDueSents(userId int, daysAhead int) ([]*model.DueSent, error) {
	dueTime := time.Now().Add(time.Duration(daysAhead*24) * time.Hour)

	stmtDue, err := dbconn.Db.Prepare(`SELECT s.id, s.media_uri FROM (SELECT sent_id FROM (SELECT sent_id AS sid, MAX(try_count) AS mtc FROM user_dictation WHERE user_id=$1 GROUP BY sent_id) gauche LEFT JOIN user_dictation droid ON gauche.sid=droid.sent_id WHERE gauche.mtc=droid.try_count AND user_id=$2 AND next_try_time < $3) due JOIN sent s ON due.sent_id=s.id;`)
	if err != nil {
		return nil, err
	}
	defer stmtDue.Close()

	rows, err := stmtDue.Query(userId, userId, dueTime.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	var dueSents []*model.DueSent
	for rows.Next() {
		var sentId int
		var mediaUri string
		err = rows.Scan(&sentId, &mediaUri)
		if err != nil {
			return nil, err
		}

		sentWords, err := GetTriedSentWords(userId, sentId)
		if err != nil {
			return nil, err
		}

		dueSents = append(dueSents, &model.DueSent{SentID: &sentId, MediaURI: &mediaUri, SentWords: sentWords})
	}
	return dueSents, nil
}
