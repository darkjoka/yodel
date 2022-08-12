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
	"github.com/darkjoka/yodel/graph/model"
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

	DB := db.New(dsn)
	defer DB.Close()

	cfg := &graph.Resolver{
		UserScheme: model.UserScheme{DB: DB},
		PostScheme: model.PostScheme{DB: DB},
		CommentScheme: model.CommentScheme{DB:DB},
		VoteScheme: model.VoteScheme{DB:DB}
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: cfg}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
