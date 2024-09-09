package api

import (
	"net/http"

	"github.com/jalavosus/matomogql/handlers"
	"github.com/jalavosus/matomogql/utils/gql"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	handlers.HandleOptions(
		handlers.HandleAuth(gql.MakeServer(false)),
	).ServeHTTP(w, r)
}
