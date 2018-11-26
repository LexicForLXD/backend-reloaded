package main

import (
	log "log"
	http "net/http"
	os "os"

	handler "github.com/99designs/gqlgen/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	backend_reloaded "gitlab.com/lexicforlxd/backend-reloaded"
)

const defaultPort = "8080"

func main() {
	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	defer db.Close()

	db.AutoMigrate(&User{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(backend_reloaded.NewExecutableSchema(backend_reloaded.Config{Resolvers: &backend_reloaded.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
