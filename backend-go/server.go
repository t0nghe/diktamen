package main

import (
	"backend-go/graph"
	"backend-go/graph/generated"
	"backend-go/internal/auth"
	dbconn "backend-go/internal/pkg/database"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "9090"

func main() {
	log.Println("[diktamen] trying to get PORT var")
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("[diktamen] didn't get PORT var, using 9090")
		port = defaultPort
	}

	log.Println("[diktamen] creating chi router")
	router := chi.NewRouter()
	// router.Use(cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000", "https://6f0e-194-103-157-156.ngrok.io"},
	// 	AllowCredentials: true,
	// 	AllowedHeaders:   []string{"Authorization"},
	// 	AllowedMethods:   []string{http.MethodGet, http.MethodPost},
	// 	Debug:            true,
	// }).Handler)
	log.Println("[diktamen] allowing all cors")
	router.Use(cors.AllowAll().Handler) // Try if this solves the problem.

	log.Println("[diktamen] adding auth middleware")
	router.Use(auth.Middleware())

	log.Println("[diktamen] init db")
	dbconn.InitDb()
	defer dbconn.CloseDb()

	log.Print("[diktamen] let's try serving first")
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	log.Println("[diktamen] starting playground handler")
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Println("[diktamen] trying to start playground")
	log.Printf("connect to :%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
