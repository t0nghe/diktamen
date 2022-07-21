package queries

import (
	"backend-go/graph/model"
	dbconn "backend-go/internal/pkg/database"
	"fmt"
	"log"
	"strconv"
)

type articleRow struct {
	aid  int    // article id
	attl string // article title
	asc  int    // artile sentence count
	ufi  int    // user finished index
}

func GetUserArticlesList(userId int) ([]*model.UserArticle, error) {
	stmt, err := dbconn.Db.Prepare("SELECT a.id, a.title, a.sent_count, ua.finished_sent_index FROM article a LEFT JOIN user_article ua ON a.id = ua.article_id WHERE ua.user_id = ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		log.Fatal(err)
	}

	var articleRecords []*model.UserArticle

	for rows.Next() {
		var row articleRow // Earlier I defined it as *articleRow type and it didn't work.
		err := rows.Scan(&row.aid, &row.attl, &row.asc, &row.ufi)
		fmt.Println(row)

		if err != nil {
			log.Fatal(err)
		}

		articleRecords = append(articleRecords, &model.UserArticle{ArticleID: strconv.Itoa(row.aid), ArticleTitle: row.attl, ArticleSentCount: row.asc, UserFinishedIndex: &row.ufi})
	}

	return articleRecords, nil
}

func GetUserUnseenArticlesList(userId int) ([]*model.UserArticle, error) {
	stmt, err := dbconn.Db.Prepare("SELECT DISTINCT a.id, a.title, a.sent_count FROM article a LEFT JOIN user_article ua ON a.id = ua.article_id WHERE ua.user_id != ? OR ua.user_id IS NULL;")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		log.Fatal(err)
	}

	var articleRecords []*model.UserArticle

	for rows.Next() {
		var row articleRow // Earlier I defined it as *articleRow type and it didn't work.
		err := rows.Scan(&row.aid, &row.attl, &row.asc)
		fmt.Println(row)

		if err != nil {
			log.Fatal(err)
		}

		zero := 0
		articleRecords = append(articleRecords, &model.UserArticle{ArticleID: strconv.Itoa(row.aid), ArticleTitle: row.attl, ArticleSentCount: row.asc, UserFinishedIndex: &zero})
	}

	return articleRecords, nil
}
