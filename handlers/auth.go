package handlers

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"sync"

	"github.com/jalavosus/matomogql/utils"
)

var (
	httpAuthUsername = sync.OnceValue(utils.MakeEnvFunc("HTTP_AUTH_USERNAME"))
	httpAuthPassword = sync.OnceValue(utils.MakeEnvFunc("HTTP_AUTH_PASSWORD"))
	httpAuthRealm    = sync.OnceValue(utils.MakeEnvFuncWithDefault("HTTP_AUTH_REALM", "Restricted"))
)

func HandleAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authenticateHeader := fmt.Sprintf(`Basic realm="%[1]s", charset="UTF-8"`, httpAuthRealm())

		if err := checkRequestAuth(r); err != nil {
			w.Header().Set("WWW-Authenticate", authenticateHeader)
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
