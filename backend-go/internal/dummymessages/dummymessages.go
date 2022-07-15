package dummymessages

import (
	dbconn "backend-go/internal/pkg/database"
	"log"
	"math/rand"
)

type DummyMessage struct {
	Success bool
	Message string
}

func (dmsg DummyMessage) Save() int64 {
	stmt, err := dbconn.Db.Prepare("INSERT INTO dummy_message (number, success, message) VALUES (?, ?, ?);")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(rand.Intn(100), dmsg.Success, dmsg.Message)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetDummyMessages() []DummyMessage {
	stmt, err := dbconn.Db.Prepare("SELECT success, message FROM dummy_message;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var dmsgs []DummyMessage

	for rows.Next() {
		var dmsg DummyMessage
		err := rows.Scan(&dmsg.Success, &dmsg.Message)
		if err != nil {
			log.Fatal(err)
		}
		dmsgs = append(dmsgs, dmsg)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return dmsgs
}
