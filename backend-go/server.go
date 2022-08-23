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
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	// router.Use(cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000", "https://6f0e-194-103-157-156.ngrok.io"},
	// 	AllowCredentials: true,
	// 	AllowedHeaders:   []string{"Authorization"},
	// 	AllowedMethods:   []string{http.MethodGet, http.MethodPost},
	// 	Debug:            true,
	// }).Handler)
	router.Use(cors.AllowAll().Handler) // Try if this solves the problem.
	router.Use(auth.Middleware())

	dbconn.InitDb()
	defer dbconn.CloseDb()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to :%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
