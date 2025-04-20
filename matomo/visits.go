package matomo

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jalavosus/matomogql/graph/model"
)

func (c clientImpl) GetVisitorProfile(
	ctx context.Context, idSite int, visitorId string,
) (*model.VisitorProfile, error) {

	params := c.buildRequestParams(idSite, "Live.getVisitorProfile")
	params.Set("visitorId", visitorId)

	var result *model.VisitorProfile
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c clientImpl) GetVisitorProfiles(
	ctx context.Context, idSite int, visitorIds []string,
) ([]*model.VisitorProfile, error) {

	idSiteStr := strconv.Itoa(idSite)
	queries := make([][2]string, len(visitorIds))
	for i, id := range visitorIds {
		queries[i] = [2]string{idSiteStr, id}
	}

	return c.GetVisitorProfilesBulk(ctx, queries...)
}

func (c clientImpl) GetVisitorProfilesBulk(ctx context.Context, queries ...[2]string) ([]*model.VisitorProfile, error) {
	params := c.buildRequestParams(-1, "API.getBulkRequest")

	for i, query := range queries {
		idSite, _ := strconv.Atoi(query[0])
		moreParams := c.buildRequestParams(idSite, "Live.getVisitorProfile")
		moreParams.Set("visitorId", query[1])

		params.Set(
			fmt.Sprintf("urls[%d]", i), // urls[0], urls[1], etc.
			moreParams.Encode(),
		)
	}

	var result []*model.VisitorProfile
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c clientImpl) GetLastVisits(ctx context.Context, idSite int, opts *model.LastVisitsOpts) ([]*model.Visit, error) {
	params := c.buildRequestParams(idSite, "Live.getLastVisitsDetails")
	params.Set("expanded", "1")
	params.Set("filterLimit", "-1")

	if opts == nil {
		opts = new(model.LastVisitsOpts)
	}

	var querySegments []string

	if segments, ok := opts.Segments.ValueOK(); ok && len(segments) > 0 {
		querySegments = append(querySegments, segments...)
	}

	if goalIds, ok := opts.GoalIds.ValueOK(); ok && len(goalIds) > 0 {
		for _, idGoal := range goalIds {
			querySegments = append(querySegments, "visitConvertedGoalId=="+strconv.Itoa(idGoal))
		}
	}

	if len(querySegments) > 0 {
		params.Set("segment", strings.Join(querySegments, ";"))
	}

	if dateOpts, ok := opts.Date.ValueOK(); ok {
		params.Set("period", strings.ToLower(dateOpts.Period.String()))

		var date = dateOpts.StartDate
		if endDate, ok := dateOpts.EndDate.ValueOK(); ok && *endDate != "" {
			date += "," + *endDate
		}

		params.Set("date", date)
	}

	var result []*model.Visit
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	orderByOpts, orderByOptsSet := opts.OrderBy.ValueOK()
	if orderByOptsSet {
		orderBy, orderBySet := orderByOpts.Timestamp.ValueOK()
		if orderBySet && orderBy != nil {
			sort.Slice(result, func(i, j int) bool {
				if *orderBy == model.OrderByDesc {
					return result[i].ServerTimestamp > result[j].ServerTimestamp
				}

				return result[i].ServerTimestamp < result[j].ServerTimestamp
			})
		}
	}

	return result, nil
}
