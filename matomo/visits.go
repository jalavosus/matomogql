package matomo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/jalavosus/matomogql/graph/model"
)

func GetVisitorProfile(ctx context.Context, idSite int, visitorId string) (*model.VisitorProfile, error) {
	params, endpoint := buildRequestParams(idSite, "Live.getVisitorProfile")
	params.Set("visitorId", visitorId)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	endpoint = endpoint + "?" + params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result *model.VisitorProfile
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetVisitorProfiles(ctx context.Context, idSite int, visitorIds []string) ([]*model.VisitorProfile, error) {
	idSiteStr := strconv.Itoa(idSite)
	queries := make([][2]string, len(visitorIds))
	for i, id := range visitorIds {
		queries[i] = [2]string{idSiteStr, id}
	}

	return GetVisitorProfilesBulk(ctx, queries...)
}

func GetVisitorProfilesBulk(ctx context.Context, queries ...[2]string) ([]*model.VisitorProfile, error) {
	var (
		vals, endpoint = buildRequestParams(-1, "API.getBulkRequest")
	)

	for i, query := range queries {
		idSite, _ := strconv.Atoi(query[0])
		queryVals, _ := buildRequestParams(idSite, "Live.getVisitorProfile")
		queryVals.Set("visitorId", query[1])

		vals.Set(
			fmt.Sprintf("urls[%d]", i), // urls[0], urls[1], etc.
			queryVals.Encode(),
		)
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	endpoint = endpoint + "?" + vals.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result []*model.VisitorProfile
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
