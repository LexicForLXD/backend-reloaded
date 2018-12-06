package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"gitlab.com/lexicforlxd/backend-reloaded/graphql"
	"gitlab.com/lexicforlxd/backend-reloaded/models"
	"gitlab.com/lexicforlxd/backend-reloaded/resolvers"
)

const defaultPort = "8080"

func main() {

	db, err := models.CreateConnection()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	db.AutoMigrate(models.User{}, models.Container{}, models.Host{})

	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &resolvers.Resolver{db}})))

	log.Printf("connect to http://localhost:%s/query ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
