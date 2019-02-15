package config

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/lexicforlxd/backend-reloaded/graphql"
	"github.com/lexicforlxd/backend-reloaded/host"
	_hostRest "github.com/lexicforlxd/backend-reloaded/host/delivery/rest"
	_graphUtil "github.com/lexicforlxd/backend-reloaded/util/delivery/graphql"
)

func InitRouter() *chi.Mux {
	diContainer := BuildContainer()

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
		diContainer.Invoke(func(hostUsecase host.Usecase) {
			r.Mount("/hosts", _hostRest.NewHostHandler(hostUsecase))
		})
	})

	if err := diContainer.Invoke(func(resolver graphql.ResolverRoot) {
		router.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}), handler.ErrorPresenter(_graphUtil.CustomErrorHandler)))
	}); err != nil {
		log.Fatalf("Error while init: %v", err)
	}
	return router
}
