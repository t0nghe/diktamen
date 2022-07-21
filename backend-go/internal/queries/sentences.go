package queries

import (
	dbconn "backend-go/internal/pkg/database"
	"fmt"
	"log"
)

func QueryUserFullArticle(userId int, articleID string) (articleRow, []sentDetail) {
	fmt.Println("userId, articleID: ", userId, articleID)
	stmt, err := dbconn.Db.Prepare("SELECT a.id, a.title, a.sent_count, ua.finished_sent_index FROM article a LEFT OUTER JOIN user_article ua ON a.id = ua.article_id WHERE a.id=? AND ua.user_id=?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var row articleRow
	// INTERESTING articleID here is a string. In db, the id field is an auto-increment int.
	// but the query still works.
	// even in terminal, you can still run
	err = stmt.QueryRow(articleID, userId).Scan(&row.aid, &row.attl, &row.asc, &row.ufi)
	fmt.Println("query row working? ", row)
	if err != nil {
		log.Fatal(err)
	}
	return row, nil
}
