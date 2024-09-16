package matomo

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jalavosus/matomogql/graph/model"
)

func GetVisitorProfile(ctx context.Context, idSite int, visitorId string) (*model.VisitorProfile, error) {
	params, endpoint := buildRequestParams(idSite, "Live.getVisitorProfile")
	params.Set("visitorId", visitorId)

	var result *model.VisitorProfile
	if err := httpGet(ctx, endpoint, params, &result); err != nil {
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
		params, endpoint = buildRequestParams(-1, "API.getBulkRequest")
	)

	for i, query := range queries {
		idSite, _ := strconv.Atoi(query[0])
		moreParams, _ := buildRequestParams(idSite, "Live.getVisitorProfile")
		moreParams.Set("visitorId", query[1])

		params.Set(
			fmt.Sprintf("urls[%d]", i), // urls[0], urls[1], etc.
			moreParams.Encode(),
		)
	}

	var result []*model.VisitorProfile
	if err := httpGet(ctx, endpoint, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetLastVisits(ctx context.Context, idSite int, opts *model.LastVisitsOpts) ([]*model.Visit, error) {
	params, endpoint := buildRequestParams(idSite, "Live.getLastVisitsDetails")
	params.Set("expanded", "1")
	params.Set("filterLimit", "-1")

	if opts == nil {
		opts = new(model.LastVisitsOpts)
	}

	if segments, ok := opts.Segments.ValueOK(); ok && len(segments) > 0 {
		params.Set("segment", strings.Join(segments, ";"))
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
	if err := httpGet(ctx, endpoint, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}
