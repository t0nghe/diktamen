package token

import (
	"log"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// GenerateToken generates a jwt token that includes username and expiry date in its claims.
func GenerateToken(username string) (string, error) {
	SecretKey := []byte(os.Getenv("BACKEND_GO_SECRETKEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (string, error) {
	SecretKey := []byte(os.Getenv("BACKEND_GO_SECRETKEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
