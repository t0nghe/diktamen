package queries

import (
	dbconn "backend-go/internal/pkg/database"
	"fmt"
	"log"
)

func TryInsertSomething() {
	stmt, err := dbconn.Db.Prepare("INSERT INTO article (sent_count, title) VALUES (?, ?)")
	if err != nil {
		log.Fatal("Damn!")
	}
	defer stmt.Close()

	result, err := stmt.Exec(10, "Hello, this is a test")
	if err != nil {
		log.Fatal("Second damn!")
	}
	lastIndex, _ := result.LastInsertId()
	fmt.Println(lastIndex)
}
