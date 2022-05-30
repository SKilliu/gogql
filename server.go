package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SKilliu/gogql/middlewares"

	"github.com/SKilliu/gogql/service"

	"github.com/SKilliu/gogql/directives"

	"github.com/go-chi/chi/middleware"

	"github.com/SKilliu/gogql/storage"
	"github.com/go-chi/chi"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SKilliu/gogql/graph"
	"github.com/SKilliu/gogql/graph/generated"
)

const defaultPort = "8080"

func main() {

	storage.InitStorage()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(
		middleware.DefaultLogger,
		middlewares.AuthMiddleware,
	)

	gqlConfig := generated.Config{
		Resolvers: &graph.Resolver{Service: service.NewUserService(storage.GetUsersStorage())}}
	gqlConfig.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(gqlConfig))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(server.ListenAndServe())
}
