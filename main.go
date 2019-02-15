package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/go-chi/chi"
	"github.com/lexicforlxd/backend-reloaded/config"
	rollbar "github.com/rollbar/rollbar-go"
)

const defaultPort = "8080"

func init() {
	config.InitViper()
	config.InitRollbar()
}

func main() {

	/*
		Router
	*/
	router := config.InitRouter()
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	rollbar.Wait()

	log.Printf("connect to http://localhost:%s", viper.GetString("port"))
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), router))
}
