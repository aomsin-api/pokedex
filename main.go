package main

import (
	"log"
	"net/http"
	"os"
	"pokedex/database"
	"pokedex/graph"
	"pokedex/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	arg := os.Args[1]
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db, err := database.PokedexInit(&arg)
	if err != nil {
		log.Fatalln(err)
	}
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					Pokedex: &database.PokedexOp{
						Db: db,
					},
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
