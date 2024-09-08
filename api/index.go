package api

import (
	"net/http"

	"github.com/jalavosus/matomogql/utils/gql"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	mux := gql.MakeServer()
	mux.ServeHTTP(w, r)
}
