package handlers

import (
	"net/http"
)

func sendError(w http.ResponseWriter, code int, err error) {
	data := (HttpError{
		Message: err.Error(),
		Code:    code,
	}).Serialize()

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}
