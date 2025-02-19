package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jalavosus/matomogql/utils/gql"
)

const (
	defaultPort = 8080
)

var (
	addr = fmt.Sprintf(":%d", defaultPort)
)

func main() {
	handler := gql.MakeServer(true)
	//handler = handlers.HandleAuth(handler)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", defaultPort)

	log.Fatal(http.ListenAndServe(addr, handler))
}
