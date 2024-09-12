package matomo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/jalavosus/matomogql/graph/model"
)

func GetGoal(ctx context.Context, idSite, idGoal int) (*model.Goal, error) {
	params, endpoint := buildRequestParams(idSite, "Goals.getGoal")
	params.Set("idGoal", strconv.Itoa(idGoal))
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

	var goal *model.Goal
	if err := json.Unmarshal(bodyRaw, &goal); err != nil {
		return nil, err
	}

	return goal, nil
}

func GetGoals(ctx context.Context, idSite int, goalIds []int, opts *model.GetGoalsOptions) ([]*model.Goal, error) {
	var (
		vals, endpoint = buildRequestParams(-1, "API.getBulkRequest")
	)

	if opts == nil {
		opts = new(model.GetGoalsOptions)
	}

	for i, id := range goalIds {
		params, _ := buildRequestParams(idSite, "Goals.getGoal")
		params.Set("idGoal", strconv.Itoa(id))
		vals.Set(
			fmt.Sprintf("urls[%d]", i), // urls[0], urls[1], etc.
			params.Encode(),
		)
	}

	endpoint = endpoint + "?" + vals.Encode()

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

	var goals []*model.Goal

	if err := json.Unmarshal(bodyRaw, &goals); err != nil {
		return nil, err
	}

	if val, ok := opts.OrderByName.ValueOK(); ok && *val {
		sort.Slice(goals, func(i, j int) bool {
			return goals[i].Name < goals[j].Name
		})
	} else {
		sort.Slice(goals, func(i, j int) bool {
			return goals[i].IDGoal < goals[j].IDGoal
		})
	}

	return goals, nil
}

func GetAllGoals(ctx context.Context, idSite int, opts *model.GetGoalsOptions) ([]*model.Goal, error) {
	if opts == nil {
		opts = new(model.GetGoalsOptions)
	}

	params, endpoint := buildRequestParams(idSite, "Goals.getGoals")
	if val, ok := opts.OrderByName.ValueOK(); ok && *val {
		params.Set("orderByName", "true")
	}

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

	var goals []*model.Goal
	if err := json.Unmarshal(bodyRaw, &goals); err != nil {
		return nil, err
	}

	if val, ok := opts.OrderByName.ValueOK(); ok && *val {
		sort.Slice(goals, func(i, j int) bool {
			return goals[i].Name < goals[j].Name
		})
	} else {
		sort.Slice(goals, func(i, j int) bool {
			return goals[i].IDGoal < goals[j].IDGoal
		})
	}

	return goals, nil
}

func GetConvertedVisits(ctx context.Context, idSite, idGoal int, opts *model.ConvertedVisitsOptions) ([]*model.Visit, error) {
	params, endpoint := buildRequestParams(idSite, "Live.getLastVisitsDetails")
	params.Set("segment", fmt.Sprintf("visitConvertedGoalId==%d", idGoal))
	params.Set("expanded", "1")
	params.Set("filterLimit", "-1")

	if opts != nil && opts.Date.IsSet() {
		dateOpts := opts.Date.Value()
		params.Set("period", strings.ToLower(dateOpts.Period.String()))

		var date = dateOpts.StartDate
		if endDate, ok := dateOpts.EndDate.ValueOK(); ok && *endDate != "" {
			date += "," + *endDate
		}

		params.Set("date", date)
	}

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

	var visitDetails []*model.Visit
	if err := json.Unmarshal(bodyRaw, &visitDetails); err != nil {
		return nil, err
	}

	return visitDetails, nil
}

func GetConvertedVisitsBulk(ctx context.Context, queries ...[6]string) ([][]*model.Visit, error) {
	var (
		vals, endpoint = buildRequestParams(-1, "API.getBulkRequest")
	)

	for i, query := range queries {
		parsedQuery := parseConvertedVisitsQuery(query)
		queryVals, _ := buildRequestParams(parsedQuery.idSite, "Live.getLastVisitsDetails")
		queryVals.Set("expanded", "1")
		queryVals.Set("filterLimit", "-1")
		if parsedQuery.searchSegment != "" {
			queryVals.Set("segment", parsedQuery.searchSegment)
		} else {
			queryVals.Set("segment", fmt.Sprintf("visitConvertedGoalId==%d", parsedQuery.idGoal))
		}
		if parsedQuery.period != "" {
			queryVals.Set("period", strings.ToLower(parsedQuery.period))
		}
		if parsedQuery.date != "" {
			queryVals.Set("date", parsedQuery.date)
		}

		vals.Set(
			fmt.Sprintf("urls[%d]", i), // urls[0], urls[1], etc.
			queryVals.Encode(),
		)
	}

	endpoint = endpoint + "?" + vals.Encode()

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

	var visitDetails [][]*model.Visit
	if err := json.Unmarshal(bodyRaw, &visitDetails); err != nil {
		fmt.Println(string(bodyRaw))
		return nil, err
	}

	return visitDetails, nil
}

type convertedVisitsQuery struct {
	period        string
	date          string
	searchSegment string
	idGoal        int
	idSite        int
}

//nolint:gocritic // I know it's a large param, I'm not futzing with dataloader and pointer keys.
func parseConvertedVisitsQuery(query [6]string) (parsedQuery convertedVisitsQuery) {
	parsedQuery = convertedVisitsQuery{}

	parsedQuery.idSite, _ = strconv.Atoi(query[0])
	parsedQuery.idGoal, _ = strconv.Atoi(query[1])
	parsedQuery.period = query[2]

	parsedQuery.date = query[3]
	if endDate := query[4]; endDate != "" {
		parsedQuery.date = parsedQuery.date + "," + endDate
	}

	if segment := query[5]; segment != "" {
		parsedQuery.searchSegment = segment
	}

	return
}
