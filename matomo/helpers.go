package matomo

import (
	"context"
	"encoding/json"
	"io"
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

func (c clientImpl) buildRequestParams(idSite int, method string) (values url.Values) {
	values = url.Values{}

	values.Set("method", method)
	if idSite != noIdSite {
		values.Set("idSite", strconv.Itoa(idSite))
	}
	values.Set("format", apiFormat)
	values.Set("module", apiModule)
	values.Set("token_auth", c.apiKey)

	return
}

func (c clientImpl) httpGet(ctx context.Context, params url.Values, out any) error {
	endpoint := c.apiEndpoint + "?" + params.Encode()

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, http.NoBody)
	if err != nil {
		return err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bodyRaw, out); err != nil {
		return err
	}

	return nil
}
