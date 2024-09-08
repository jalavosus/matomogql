package gql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/jalavosus/matomogql/graph"
	"github.com/jalavosus/matomogql/graph/loaders"
)

func MakeServer() *http.ServeMux {
	execSchema := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})

	var srv = loaders.Middleware(handler.NewDefaultServer(execSchema))

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	return mux
}
