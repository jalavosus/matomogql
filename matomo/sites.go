package matomo

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/jalavosus/matomogql/graph/model"
)

func GetSiteFromID(ctx context.Context, idSite int) (*model.Site, error) {
	params, endpoint := buildRequestParams(idSite, "SitesManager.getSiteFromId")
	endpoint = endpoint + "?" + params.Encode()

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result *model.Site
	if err := json.Unmarshal(bodyRaw, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetSiteURLsFromID(ctx context.Context, idSite int) ([]string, error) {
	params, endpoint := buildRequestParams(idSite, "SitesManager.getSiteUrlsFromId")
	endpoint = endpoint + "?" + params.Encode()

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result []string
	if err := json.Unmarshal(bodyRaw, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetSitesWithViewAccess(ctx context.Context) ([]*model.Site, error) {
	params, endpoint := buildRequestParams(noIdSite, "SitesManager.getSitesWithViewAccess")
	endpoint = endpoint + "?" + params.Encode()

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result []*model.Site
	if err := json.Unmarshal(bodyRaw, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetSitesWithAtLeastViewAccess(ctx context.Context) ([]*model.Site, error) {
	params, endpoint := buildRequestParams(noIdSite, "SitesManager.getSitesWithAtLeastViewAccess")
	endpoint = endpoint + "?" + params.Encode()

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result []*model.Site
	if err := json.Unmarshal(bodyRaw, &result); err != nil {
		return nil, err
	}

	return result, nil
}
