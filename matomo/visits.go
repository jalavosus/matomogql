package matomo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jalavosus/matomogql/graph/model"
)

func GetVisitorProfile(ctx context.Context, idSite int, visitorId string) (*model.VisitorProfile, error) {
	params, endpoint := buildRequestParams(idSite, "Live.getVisitorProfile")
	params.Set("visitorId", visitorId)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	endpoint = endpoint + "?" + params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer closeResBody(resp.Body)

	body, err := io.ReadAll(resp.Body)
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
	queries := make([]string, len(visitorIds))
	for i, id := range visitorIds {
		queries[i] = fmt.Sprintf("%d:%s", idSite, id)
	}

	return GetVisitorProfilesBulk(ctx, queries...)
}

func GetVisitorProfilesBulk(ctx context.Context, queries ...string) ([]*model.VisitorProfile, error) {
	var (
		vals, endpoint = buildRequestParams(-1, "API.getBulkRequest")
	)

	for i, query := range queries {
		parsedQuery := parseGetVisitorProfilesQuery(query)
		queryVals, _ := buildRequestParams(parsedQuery.idSite, "Live.getVisitorProfile")
		queryVals.Set("visitorId", parsedQuery.visitorId)

		vals.Set(
			fmt.Sprintf("urls[%d]", i), // urls[0], urls[1], etc.
			queryVals.Encode(),
		)
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	endpoint = endpoint + "?" + vals.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer closeResBody(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []*model.VisitorProfile
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

type parsedVisitorProfilesQuery struct {
	idSite    int
	visitorId string
}

func parseGetVisitorProfilesQuery(query string) (parsed parsedVisitorProfilesQuery) {
	parsed = parsedVisitorProfilesQuery{}

	split := strings.Split(query, ":")
	parsed.idSite, _ = strconv.Atoi(split[0])
	parsed.visitorId = split[1]

	return
}
