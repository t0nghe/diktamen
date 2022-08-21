package queries

import (
	"backend-go/graph/model"
	dbconn "backend-go/internal/pkg/database"
	"database/sql"
	"fmt"
	"log"
)

type articleRow struct {
	aid   int    // article id
	attl  string // article title
	asc   int    // article sentence count
	adesc string // article description
	ufi   int    // user finished index
}

func GetSingleArticle(userId int, articleId int) (*model.UserArticle, error) {
	var record articleRow
	// Get info about article
	stmt, err := dbconn.Db.Prepare("SELECT id, title, sent_count, description FROM article WHERE id=?;")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(articleId).Scan(&record.aid, &record.attl, &record.asc, &record.adesc)
	if err != nil {
		return nil, err
	}

	// Get user `finished_sent_index`
	stmt2, err := dbconn.Db.Prepare("SELECT finished_sent_index FROM user_article WHERE user_id=? AND article_id=?;")
	if err != nil {
		return nil, err
	}

	err = stmt2.QueryRow(userId, articleId).Scan(&record.ufi)
	if err == sql.ErrNoRows {
		record.ufi = 0
	} else if err != nil {
		return nil, err
	}

	// Note for a new user, `user_article` table will return an empty row
	// stmt, err := dbconn.Db.Prepare("SELECT a.id, a.title, a.sent_count, a.description, ua.finished_sent_index FROM article a LEFT JOIN user_article ua ON a.id = ua.article_id WHERE ua.user_id=? AND a.id=?;")
	return &model.UserArticle{ArticleID: &record.aid, ArticleTitle: record.attl, ArticleSentCount: record.asc, ArticleDescription: &record.adesc, UserFinishedIndex: &record.ufi}, nil
}

func GetUserArticlesList(userId int) ([]*model.UserArticle, error) {
	stmt, err := dbconn.Db.Prepare("SELECT a.id, a.title, a.sent_count, a.description, ua.finished_sent_index FROM article a LEFT JOIN user_article ua ON a.id = ua.article_id WHERE ua.user_id = ?;")
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
		err := rows.Scan(&row.aid, &row.attl, &row.asc, &row.adesc, &row.ufi)
		fmt.Println(row)

		if err != nil {
			log.Fatal(err)
		}

		articleRecords = append(articleRecords, &model.UserArticle{ArticleID: &row.aid, ArticleTitle: row.attl, ArticleSentCount: row.asc, ArticleDescription: &row.adesc, UserFinishedIndex: &row.ufi})
	}

	return articleRecords, nil
}

func GetUserUnseenArticlesList(userId int) ([]*model.UserArticle, error) {
	stmt, err := dbconn.Db.Prepare("SELECT DISTINCT a.id, a.title, a.sent_count, a.description FROM article a LEFT JOIN user_article ua ON a.id = ua.article_id WHERE a.id NOT IN (SELECT article_id FROM user_article WHERE user_id=?);")

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
		err := rows.Scan(&row.aid, &row.attl, &row.asc, &row.adesc)
		fmt.Println(row)

		if err != nil {
			log.Fatal(err)
		}

		zero := 0
		articleRecords = append(articleRecords, &model.UserArticle{ArticleID: &row.aid, ArticleTitle: row.attl, ArticleSentCount: row.asc, ArticleDescription: &row.adesc, UserFinishedIndex: &zero})
	}

	return articleRecords, nil
}
