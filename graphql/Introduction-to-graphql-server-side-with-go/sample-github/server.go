package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kudoas/graphql-sample/graph"
	"github.com/kudoas/graphql-sample/graph/services"
	"github.com/kudoas/graphql-sample/internal"
	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultPort = "8080"
	dbFile      = "./mygraphql.db"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=on", dbFile))
	if err != nil {
		// TODO: graceful shutdown
		log.Fatal(err)
	}
	defer db.Close()

	service := services.New(db)
	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: &graph.Resolver{
		Srv: service,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
