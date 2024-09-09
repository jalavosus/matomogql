package handlers

import (
	"crypto/subtle"
	"net/http"
	"sync"

	"github.com/jalavosus/matomogql/utils"
)

var (
	httpAuthUsername = sync.OnceValue(utils.MakeEnvFunc("HTTP_AUTH_USERNAME"))
	httpAuthPassword = sync.OnceValue(utils.MakeEnvFunc("HTTP_AUTH_PASSWORD"))
)

func HandleAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := checkRequestAuth(r); err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="FerretParty", charset="UTF-8"`)
			sendError(w, http.StatusUnauthorized, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkRequestAuth(r *http.Request) (err error) {
	authUsername, authPassword, ok := r.BasicAuth()
	if !ok {
		err = ErrMissingAuth
	} else if !checkAuthValues(authUsername, httpAuthUsername()) || !checkAuthValues(authPassword, httpAuthPassword()) {
		err = ErrInvalidAuth
	}

	return
}

func checkAuthValues(want, got string) bool {
	wantHash := utils.Sha256Sum(want)
	gotHash := utils.Sha256Sum(got)

	return subtle.ConstantTimeCompare(wantHash, gotHash) == 1
}
