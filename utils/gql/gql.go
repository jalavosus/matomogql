package gql

import (
	"net/http"
	
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	
	"github.com/jalavosus/matomogql/graph"
	"github.com/jalavosus/matomogql/graph/loaders"
	"github.com/jalavosus/matomogql/matomo"
)

func MakeServer(enablePlayground bool) http.Handler {
	matomoClient := matomo.NewClient(matomo.GetEnv())
	
	execSchema := graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(matomoClient)})
	
	var srv = loaders.Middleware(handler.New(execSchema), matomoClient)
	
	mux := http.NewServeMux()
	mux.Handle("/query", srv)
	if enablePlayground {
		mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}
	
	return mux
}
