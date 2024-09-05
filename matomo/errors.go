package matomo

import (
	"errors"
)

var (
	ErrorMissingAPIKey   = errors.New("missing Matomo API key in env")
	ErrorMissingEndpoint = errors.New("missing Matomo endpoint in env")
)
