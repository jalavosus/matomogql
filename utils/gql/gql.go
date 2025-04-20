package gql

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/jalavosus/matomogql/graph"
	"github.com/jalavosus/matomogql/graph/loaders"
	"github.com/jalavosus/matomogql/matomo"
)

func MakeServer(enablePlayground bool) http.Handler {
	matomoClient := matomo.NewClient(matomo.GetEnv())
	execSchema := graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(matomoClient)})

	var srv = loaders.Middleware(newServer(execSchema), matomoClient)

	mux := http.NewServeMux()
	mux.Handle("/query", srv)
	if enablePlayground {
		mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}

	return mux
}

var transports = []graphql.Transport{
	transport.Options{},
	transport.GET{},
	transport.POST{},
	transport.MultipartForm{},
	transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	},
}

func newServer(schema graphql.ExecutableSchema) http.Handler {
	srv := handler.New(schema)

	for _, t := range transports {
		srv.AddTransport(t)
	}

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv
}
