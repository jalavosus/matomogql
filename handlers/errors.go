package handlers

import (
	"encoding/json"
	"errors"
)

var (
	ErrMissingAuth = errors.New("missing basic auth")
	ErrInvalidAuth = errors.New("invalid auth")
)

type HttpError struct {
	Message string `json:"msg"`
	Code    int    `json:"code"`
}

func (e HttpError) Serialize() []byte {
	data, _ := json.Marshal(e)
	return data
}
