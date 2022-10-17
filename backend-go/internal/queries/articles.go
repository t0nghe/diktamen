package queries

import (
	"backend-go/graph/model"
	dbconn "backend-go/internal/pkg/database"
	"database/sql"
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
	articleRow := dbconn.Db.QueryRow("SELECT id, title, sent_count, description FROM article WHERE id=$1;", articleId)

	err := articleRow.Scan(&record.aid, &record.attl, &record.asc, &record.adesc)
	if err != nil {
		log.Println(err)
		log.Println("SELECT id, title, sent_count, description FROM article WHERE id=$1;")
		return nil, err
	}

	// Get user `finished_sent_index`
	finishedSentIdxRow := dbconn.Db.QueryRow("SELECT finished_sent_index FROM user_article WHERE user_id=$1 AND article_id=$2;", userId, articleId)

	err = finishedSentIdxRow.Scan(&record.ufi)
	if err == sql.ErrNoRows {
		record.ufi = 0
	} else if err != nil {
		log.Println(err)
		log.Println("SELECT finished_sent_index FROM user_article WHERE user_id=$1 AND article_id=$2;")
		return nil, err
	}

	return &model.UserArticle{ArticleID: &record.aid, ArticleTitle: record.attl, ArticleSentCount: record.asc, ArticleDescription: &record.adesc, UserFinishedIndex: &record.ufi}, nil
}

func GetUserArticlesList(userId int) ([]*model.UserArticle, error) {
	rows, err := dbconn.Db.Query("SELECT a.id, a.title, a.sent_count, a.description, COALESCE (ua.finished_sent_index, 0) FROM article a LEFT JOIN user_article ua ON a.id = ua.article_id WHERE ua.user_id = $1;", userId)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var articleRecords []*model.UserArticle

	for rows.Next() {
		var row articleRow // Earlier I defined it as *articleRow type and it didn't work.
		err := rows.Scan(&row.aid, &row.attl, &row.asc, &row.adesc, &row.ufi)

		if err != nil {
			log.Println("!!! inside GetUserArticlesList")
			log.Println(row)
			log.Println(err)
		}

		articleRecords = append(articleRecords, &model.UserArticle{ArticleID: &row.aid, ArticleTitle: row.attl, ArticleSentCount: row.asc, ArticleDescription: &row.adesc, UserFinishedIndex: &row.ufi})
	}

	return articleRecords, nil
}

type debugRow struct {
	one   int    // article id
	two   string // article title
	three int    // article sentence count
	four  string // article description
}

func GetUserUnseenArticlesList(userId int) ([]*model.UserArticle, error) {
	rows, err := dbconn.Db.Query("SELECT DISTINCT a.id, a.title, a.sent_count, a.description FROM article a LEFT JOIN user_article ua ON a.id = ua.article_id WHERE a.id NOT IN (SELECT article_id FROM user_article WHERE user_id=$1);", userId)

	if err != nil {
		log.Println("SELECT DISTINCT a.id, a.title, a.sent_count, a.description FROM article a LEFT JOIN user_article ua ON a.id = ua.article_id WHERE a.id NOT IN (SELECT article_id FROM user_article WHERE user_id=$1);")
		log.Println(rows)
		log.Println("userId", userId)
		log.Println(err)
	}
	defer rows.Close()

	// var articleRecords []*model.UserArticle

	for rows.Next() {
		// var row articleRow // Earlier I defined it as *articleRow type and it didn't work.
		var row debugRow // Earlier I defined it as *articleRow type and it didn't work.
		err := rows.Scan(&row.one, &row.two, &row.three, &row.four)

		log.Println("debug line 96", row)

		if err != nil {
			log.Print("SELECT DISTINCT a.id, a.title, a.sent_count, a.description FROM article a LEFT JOIN user_article ua ON a.id = ua.article_id WHERE a.id NOT IN (SELECT article_id FROM user_article WHERE user_id=$1); - second half")
			log.Println(err)
		}

		// zero := 0
		// articleRecords = append(articleRecords, &model.UserArticle{ArticleID: &row.aid, ArticleTitle: row.attl, ArticleSentCount: row.asc, ArticleDescription: &row.adesc, UserFinishedIndex: &zero})
	}

	// return articleRecords, nil
	return nil, nil
}
