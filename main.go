package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/fsnotify/fsnotify"
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
	_graphUtil "github.com/lexicforlxd/backend-reloaded/util/delivery/graphql"
	rollbar "github.com/rollbar/rollbar-go"
	"github.com/spf13/viper"
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

	router.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolvers.NewResolver(hostUsecase)}), handler.ErrorPresenter(_graphUtil.CustomErrorHandler)))
	return router
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		readCerts()
	})

	readCerts()
	viper.SetDefault("rollbar.codeVersion", "v1")
	viper.SetDefault("rollbar.serverRoot", "github.com/lexicforlxd/backend-reloaded")
	viper.SetDefault("rollbar.environment", "development")
	viper.SetDefault("database.host", "localhost")

	initRollbar()
}

func initRollbar() {
	rollbar.SetToken(viper.GetString("rollbar.token"))
	rollbar.SetEnvironment(viper.GetString("rollbar.environment")) // defaults to "development"
	rollbar.SetCodeVersion(viper.GetString("rollbar.codeVersion")) // optional Git hash/branch/tag (required for GitHub integration)
	// rollbar.SetServerHost("web.1")                       // optional override; defaults to hostname
	rollbar.SetServerRoot(viper.GetString("rollbar.serverRoot")) // path of project (required for GitHub integration and non-project stacktrace collapsing)

}

func readCerts() {
	cert, err := ioutil.ReadFile(viper.GetString("tls.certFile"))
	key, err := ioutil.ReadFile(viper.GetString("tls.keyFile"))
	if err != nil {
		log.Fatal(err)
	}
	viper.Set("tls.cert", string(cert))
	viper.Set("tls.key", string(key))
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

	rollbar.Wait()

	log.Printf("connect to http://localhost:%s/query ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
