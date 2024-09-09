package gql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/jalavosus/matomogql/graph"
	"github.com/jalavosus/matomogql/graph/loaders"
)

func MakeServer(enablePlayground bool) http.Handler {
	execSchema := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})

	var srv = loaders.Middleware(handler.NewDefaultServer(execSchema))

	mux := http.NewServeMux()
	mux.Handle("/query", srv)
	if enablePlayground {
		mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}

	return mux
}
