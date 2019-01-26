package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/lexicforlxd/backend-reloaded/graphql"
	"github.com/lexicforlxd/backend-reloaded/host"
	_hostRest "github.com/lexicforlxd/backend-reloaded/host/delivery/rest"
	_hostRepo "github.com/lexicforlxd/backend-reloaded/host/repository"
	_hostUsecase "github.com/lexicforlxd/backend-reloaded/host/usecase"
	"github.com/lexicforlxd/backend-reloaded/models"
	"github.com/lexicforlxd/backend-reloaded/resolvers"
)

const defaultPort = "8080"
const defaultTimeout = "30"

var hostUsecase host.Usecase

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/hosts", _hostRest.NewHostHandler(hostUsecase))
	})

	router.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolvers.NewResolver(hostUsecase)})))
	return router
}

func main() {

	/*
		Database
	*/
	db, err := models.CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	db.AutoMigrate(models.User{}, models.Container{}, models.Host{})
	defer db.Close()

	/*
		Timeout
	*/
	timeout := os.Getenv("TIMEOUT")
	if timeout == "" {
		timeout = defaultTimeout
	}
	timeoutInt, _ := strconv.Atoi(timeout)
	timeoutContext := time.Duration(timeoutInt) * time.Second

	/*
		HostInit
	*/
	hostRepo := _hostRepo.NewHostRepository(db)
	hostUsecase = _hostUsecase.NewHostUsecase(hostRepo, timeoutContext)

	/*
		Router
	*/
	router := Routes()
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("connect to http://localhost:%s/query ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
