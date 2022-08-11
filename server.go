package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/darkjoka/yodel/graph"
	"github.com/darkjoka/yodel/graph/db"
	"github.com/darkjoka/yodel/graph/generated"
	"github.com/joho/godotenv"
)

const defaultPort = "8000"

func main() {
	godotenv.Load(".env")
	port, portOk := os.LookupEnv("PORT")
	dsn, dsnOk := os.LookupEnv("DATABASE_URI")

	if !portOk {
		port = defaultPort
	}

	if !dsnOk {
		panic("Database not setup")
	}

	db := db.New(dsn)
	defer db.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
