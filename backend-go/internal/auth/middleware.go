package auth

import (
	"backend-go/internal/pkg/token"
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type ContextKey string

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // It makes sense how the request needs to be a pointer. This makes sure the value of this request is accessible to multiple functions.
			header := r.Header.Get("Authorization")

			// next feels like Express
			// If there's not an Authorization entry in the header of this request, the user is not logged in.
			if header == "" {
				// If header is empty, no subsequent operations on token or text will be exected,
				// user_id key in context will remain 0. That's why we use 0 for unloggedin users.
				next.ServeHTTP(w, r)
				return
			}
			tokenString := strings.TrimPrefix(header, "Bearer ")
			username, err := token.ParseToken(tokenString)
			// fmt.Println("token parsed: ", username)
			if err != nil {
				fmt.Println("This should be an error.")
			}
			id, err2 := GetUserId(username)
			fmt.Println("parsed id from token: ", id)
			fmt.Println(reflect.TypeOf(id))
			if err2 != nil {
				fmt.Println("This should be an error.")
			} else if id < 1 {
				fmt.Println("This should be an error.")
			}

			// id is an int, is it ok to use it in context?
			contextKey := ContextKey("user_id")
			ctx := context.WithValue(r.Context(), contextKey, id)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func FromContext(ctx context.Context, key ContextKey) int {
	contextKey := ContextKey(key)
	value, _ := ctx.Value(contextKey).(int)
	return value
}
