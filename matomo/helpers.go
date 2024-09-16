package matomo

import (
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	apiFormat = "JSON"
	apiModule = "API"
)

const (
	noIdSite       = -9999
	defaultTimeout = 10 * time.Second
)

var (
	httpClient = new(http.Client)
)

func buildRequestParams(idSite int, method string) (values url.Values, endpoint string) {
	var apiKey string

	values = url.Values{}
	apiKey, endpoint = getEnv()

	values.Set("method", method)
	if idSite != noIdSite {
		values.Set("idSite", strconv.Itoa(idSite))
	}
	values.Set("format", apiFormat)
	values.Set("module", apiModule)
	values.Set("token_auth", apiKey)

	return
}
