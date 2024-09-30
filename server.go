package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"

	"github.com/garfiny/gqlgen-todos/graph"
	"github.com/garfiny/gqlgen-todos/postgres"
)

const defaultPort = "8080"

func main() {

	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "passwd",
		Database: "todos",
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		TodoList: graph.Todos,
		Wishes:   graph.Wishes,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
