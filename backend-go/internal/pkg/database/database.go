package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func InitDb() {
	db, err := sql.Open("mysql", "root:dbpass@(0.0.0.0:3306)/diktdev")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	Db = db
}

func CloseDb() error {
	return Db.Close()
}
