package matomo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/jalavosus/matomogql/graph/model"
	"github.com/jalavosus/matomogql/utils"
)

func GetGoal(ctx context.Context, idSite int, idGoal int) (*model.Goal, error) {
	params, endpoint := buildRequestParams(idSite, "Goals.getGoal")
	params.Set("idGoal", strconv.Itoa(idGoal))
	endpoint = endpoint + "?" + params.Encode()

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(res.Body)

	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var goal *model.Goal
	if err := json.Unmarshal(bodyRaw, &goal); err != nil {
		return nil, err
	}

	return goal, nil
}

func GetGoals(ctx context.Context, idSite int, opts *model.GetGoalsOptions) ([]*model.Goal, error) {
	params, endpoint := buildRequestParams(idSite, "Goals.getGoals")
	if opts != nil && utils.CheckBoolPointer(opts.OrderByName) {
		params.Set("orderByName", "true")
	}
	endpoint = endpoint + "?" + params.Encode()

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(res.Body)

	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var goals []*model.Goal
	if err := json.Unmarshal(bodyRaw, &goals); err != nil {
		return nil, err
	}

	return goals, nil
}

func GetConvertedVisits(ctx context.Context, idSite, idGoal int, opts *model.ConvertedVisitsOptions) ([]*model.VisitDetails, error) {
	params, endpoint := buildRequestParams(idSite, "Live.getLastVisitsDetails")
	params.Set("segment", fmt.Sprintf("visitConvertedGoalId==%d", idGoal))
	params.Set("expanded", "1")
	params.Set("filterLimit", "-1")
	params.Set("period", strings.ToLower(opts.Period.String()))

	var date = opts.StartDate
	if opts.EndDate != nil && len(*opts.EndDate) > 0 {
		date += "," + *opts.EndDate
	}
	params.Set("date", date)

	endpoint = endpoint + "?" + params.Encode()

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(res.Body)

	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var visitDetails []*model.VisitDetails
	if err := json.Unmarshal(bodyRaw, &visitDetails); err != nil {
		return nil, err
	}

	return visitDetails, nil
}

func GetConvertedVisitsBulk(ctx context.Context, queries ...string) ([][]*model.VisitDetails, error) {
	var (
		vals, endpoint = buildRequestParams(-1, "API.getBulkRequest")
	)

	for i, query := range queries {
		parsedQuery := parseConvertedVisitsQuery(query)
		queryVals, _ := buildRequestParams(parsedQuery.idSite, "Live.getLastVisitsDetails")
		queryVals.Set("expanded", "1")
		queryVals.Set("filterLimit", "-1")
		queryVals.Set("segment", fmt.Sprintf("visitConvertedGoalId==%d", parsedQuery.idGoal))
		queryVals.Set("period", strings.ToLower(parsedQuery.period))
		queryVals.Set("date", parsedQuery.date)

		vals.Set(
			fmt.Sprintf("urls[%d]", i), // urls[0], urls[1], etc.
			queryVals.Encode(),
		)
	}

	endpoint = endpoint + "?" + vals.Encode()

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(res.Body)

	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var visitDetails [][]*model.VisitDetails
	if err := json.Unmarshal(bodyRaw, &visitDetails); err != nil {
		return nil, err
	}

	return visitDetails, nil
}

type convertedVisitsQuery struct {
	idSite    int
	idGoal    int
	period    string
	startDate string
	endDate   string
	date      string
}

// format: idSite:idGoal:period:startDate:endDate
func parseConvertedVisitsQuery(query string) (parsedQuery convertedVisitsQuery) {
	parsedQuery = convertedVisitsQuery{}
	split := strings.Split(query, ":")

	parsedQuery.idSite, _ = strconv.Atoi(split[0])
	parsedQuery.idGoal, _ = strconv.Atoi(split[1])
	parsedQuery.period = split[2]
	parsedQuery.startDate = split[3]
	parsedQuery.endDate = split[4]

	parsedQuery.date = parsedQuery.startDate
	if len(parsedQuery.endDate) > 0 {
		parsedQuery.date = parsedQuery.date + "," + parsedQuery.endDate
	}

	return
}
