package auth

import (
	dbconn "backend-go/internal/pkg/database"
	"backend-go/internal/pkg/token"
	"database/sql"
	"fmt"
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// validatePasswordNaive checks if the password contains at least chars, in which there are at least 1 letter, 1 digit and 1 symbol.
func validatePasswordNaive(password string) bool {
	// re_letter := regexp.MustCompile(`[A-Za-z]`)
	// re_digit := regexp.MustCompile(`[0-9]`)

	if len(password) < 4 {
		return false
	}
	return true
}

// CreateUser takes username and password, creates a new record in user table, returns true if succeeds, false if fails. This function works with `userSignUp` mutation.
func CreateUser(username string, email string, password string) error {

	pattern := "[a-z][a-z0-9_]{3,}"
	re, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("regex error")
	}
	username_ok := re.MatchString(username)
	if !username_ok {
		return fmt.Errorf("username needs to start with a letter and contain four characters (letter, number, _)")
	}
	if !validatePasswordNaive(password) {
		// return fmt.Errorf("password needs to contain six chars, including at least 1 letter, 1 digit and 1 symbol")
		return fmt.Errorf("password needs to contain at least 4 chars")
	}

	stmt, err := dbconn.Db.Prepare("INSERT INTO dktmuser (username, email, pwhash) VALUES ($1, $2, $3);")
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer stmt.Close()

	pwhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	res, err := stmt.Exec(username, email, pwhash)

	if err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println(res)

	return nil
}

// UserLogin validates username and password, and if valid, returns a token to GraphQL client.
func UserLogin(username string, password string) (string, error) {

	var pwhash []byte
	row := dbconn.Db.QueryRow("SELECT pwhash FROM dktmuser WHERE username=$1;", username)

	err := row.Scan(&pwhash)
	if err != nil {
		log.Println(err)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(pwhash, []byte(password))
	if err != nil {
		return "", fmt.Errorf("password is wrong")
	}

	// Generate token
	tokenString, err := token.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserId(username string) (int, error) {
	row := dbconn.Db.QueryRow("SELECT id FROM dktmuser WHERE username=$1;", username)

	var userId int
	err := row.Scan(&userId)
	fmt.Println("userId in GetUserId", userId)
	fmt.Println("&userId in GetUserId", &userId)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return userId, nil
}
