package matomo

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql"

	"github.com/jalavosus/matomogql/graph/model"
)

func (c clientImpl) GetGoal(ctx context.Context, idSite, idGoal int) (*model.Goal, error) {
	params := c.buildRequestParams(idSite, "Goals.getGoal")
	params.Set("idGoal", strconv.Itoa(idGoal))

	var result *model.Goal
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c clientImpl) GetGoals(
	ctx context.Context, idSite int, goalIds []int, opts *model.GetGoalsOptions,
) ([]*model.Goal, error) {

	params := c.buildRequestParams(-1, "API.getBulkRequest")

	if opts == nil {
		opts = new(model.GetGoalsOptions)
	}

	for i, id := range goalIds {
		moreParams := c.buildRequestParams(idSite, "Goals.getGoal")
		moreParams.Set("idGoal", strconv.Itoa(id))
		params.Set(
			fmt.Sprintf("urls[%d]", i), // urls[0], urls[1], etc.
			moreParams.Encode(),
		)
	}

	var result []*model.Goal
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	if val, ok := opts.OrderByName.ValueOK(); ok && *val {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Name < result[j].Name
		})
	} else {
		sort.Slice(result, func(i, j int) bool {
			return result[i].IDGoal < result[j].IDGoal
		})
	}

	return result, nil
}

func (c clientImpl) GetAllGoals(ctx context.Context, idSite int, opts *model.GetGoalsOptions) ([]*model.Goal, error) {
	if opts == nil {
		opts = new(model.GetGoalsOptions)
	}

	params := c.buildRequestParams(idSite, "Goals.getGoals")
	if val, ok := opts.OrderByName.ValueOK(); ok && *val {
		params.Set("orderByName", "true")
	}

	var result []*model.Goal
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	if val, ok := opts.OrderByName.ValueOK(); ok && *val {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Name < result[j].Name
		})
	} else {
		sort.Slice(result, func(i, j int) bool {
			return result[i].IDGoal < result[j].IDGoal
		})
	}

	return result, nil
}

func (c clientImpl) GetConvertedVisits(
	ctx context.Context, idSite, idGoal int, opts *model.ConvertedVisitsOptions,
) ([]*model.Visit, error) {

	visitsOpts := &model.LastVisitsOpts{
		Date:     opts.Date,
		Segments: graphql.OmittableOf([]string{fmt.Sprintf("visitConvertedGoalId==%d", idGoal)}),
		OrderBy:  opts.OrderBy,
	}

	return c.GetLastVisits(ctx, idSite, visitsOpts)
}

func (c clientImpl) GetConvertedVisitsBulk(ctx context.Context, queries ...[6]string) ([][]*model.Visit, error) {
	params := c.buildRequestParams(noIdSite, "API.getBulkRequest")

	for i, query := range queries {
		parsedQuery := parseConvertedVisitsQuery(query)
		moreParams := c.buildRequestParams(parsedQuery.idSite, "Live.getLastVisitsDetails")
		moreParams.Set("expanded", "1")
		moreParams.Set("filterLimit", "-1")
		if parsedQuery.searchSegment != "" {
			moreParams.Set("segment", parsedQuery.searchSegment)
		} else {
			moreParams.Set("segment", fmt.Sprintf("visitConvertedGoalId==%d", parsedQuery.idGoal))
		}
		if parsedQuery.period != "" {
			moreParams.Set("period", strings.ToLower(parsedQuery.period))
		}
		if parsedQuery.date != "" {
			moreParams.Set("date", parsedQuery.date)
		}

		params.Set(
			fmt.Sprintf("urls[%d]", i), // urls[0], urls[1], etc.
			moreParams.Encode(),
		)
	}

	var result [][]*model.Visit
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	return result, nil
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
