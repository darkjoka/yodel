package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/darkjoka/yodel/graph"
	"github.com/darkjoka/yodel/graph/auth"
	"github.com/darkjoka/yodel/graph/db"
	"github.com/darkjoka/yodel/graph/generated"
	"github.com/darkjoka/yodel/graph/model"
	"github.com/go-chi/chi/v5"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun/extra/bundebug"
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

	// Print all queries to stdout.
	DB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	UserScheme := model.UserScheme{DB: DB}
	router := chi.NewRouter()
	router.Use(auth.Middleware(UserScheme))

	cfg := &graph.Resolver{
		UserScheme:      UserScheme,
		PostScheme:      model.PostScheme{DB: DB},
		CommentScheme:   model.CommentScheme{DB: DB},
		VoteScheme:      model.VoteScheme{DB: DB},
		CommentorScheme: model.CommentorScheme{DB: DB},
	}

	// should be run once
	// opt for cli or find way to check if there's a change from previous migration
	// m, err := migrate.New(
	// 	"file://graph/db/migrations", dsn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := m.Up(); err != nil {
	// 	log.Fatal(err)
	// }

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: cfg}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
