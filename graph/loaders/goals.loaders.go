package loaders

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/jalavosus/matomogql/graph/model"
	"github.com/jalavosus/matomogql/matomo"
)

func getGoalConvertedVisits(matomoClient matomo.Client) func(
	ctx context.Context, queries [][6]string,
) (rets [][]*model.Visit, errs []error) {

	return func(ctx context.Context, queries [][6]string) ([][]*model.Visit, []error) {
		n := len(queries)

		// Pre-allocate result slices.
		results := make([][]*model.Visit, n)
		errs := make([]error, n)

		// Matomo bulk call.
		visits, err := matomoClient.GetConvertedVisitsBulk(ctx, queries...)
		if err != nil {
			// Surface the same error for every requested element.
			for i := range errs {
				errs[i] = err
			}
			return results, errs
		}

		// Defensive: Matomo should return exactly n entries.
		if len(visits) != n {
			e := fmt.Errorf("matomo: unexpected response length %d (want %d)", len(visits), n)
			for i := range errs {
				errs[i] = e
			}
			return results, errs
		}

		// Happy path – copy slice-of-slices.
		copy(results, visits)
		return results, errs
	}

}

// GetGoalConvertedVisits returns the visits that converted a given goal.
// The optional parameters allow date-range filtering, segmentation and in-memory
// ordering by server timestamp (ASC/DESC). Any ordering is applied client-side
// after the Matomo bulk API response has been received.
func GetGoalConvertedVisits(
	ctx context.Context,
	idSite int,
	idGoal, segment string,
	opts *model.ConvertedVisitsOptions,
	orderBy *model.OrderByOptions,
) ([]*model.Visit, error) {

	const (
		qSite = iota
		qGoal
		qPeriod
		qStart
		qEnd
		qSegment
	)

	var query [6]string
	query[qSite] = strconv.Itoa(idSite)
	query[qGoal] = idGoal
	query[qSegment] = segment

	if opts != nil && opts.Date.IsSet() {
		if date := opts.Date.Value(); date != nil {
			query[qPeriod] = date.Period.String()
			query[qStart] = date.StartDate
			if end, ok := date.EndDate.ValueOK(); ok && *end != "" {
				query[qEnd] = *end
			}
		}
	}

	loaders := For(ctx)
	res, err := loaders.GoalConvertedVisitsLoader.Load(ctx, query)
	if err != nil {
		return nil, err
	}

	if orderBy != nil {
		if ts, ok := orderBy.Timestamp.ValueOK(); ok {
			asc := ts.String() == "ASC"
			sort.SliceStable(res, func(i, j int) bool {
				if asc {
					return res[i].ServerTimestamp < res[j].ServerTimestamp
				}
				return res[i].ServerTimestamp > res[j].ServerTimestamp
			})
		}
	}

	return res, nil
}
