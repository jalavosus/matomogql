package handlers

import (
	"net/http"
)

func CheckOptions(r *http.Request, w http.ResponseWriter) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")

	return r.Method == http.MethodOptions
}

func HandleOptions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if CheckOptions(r, w) {
			return
		}

		next.ServeHTTP(w, r)
	})
}
