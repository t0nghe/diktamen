package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDb() {
	log.Println("before getting env var for db string")
	dbstr := os.Getenv("DIKTAMEN_DB_STRING")
	log.Println("[diktamen] obtained dbstr")
	log.Println("[diktamen] ", dbstr)
	if len(dbstr) == 0 {
		log.Panic("No dbstring")
	}

	db, err := sql.Open("postgres", dbstr)
	if err != nil {
		log.Println("[diktamen] connection to db failed")
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
