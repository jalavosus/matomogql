package loaders

import (
	"context"
	"sort"
	"strconv"

	"github.com/jalavosus/matomogql/graph/model"
	"github.com/jalavosus/matomogql/matomo"
)

func getGoalConvertedVisits(matomoClient matomo.Client) func(ctx context.Context, queries [][6]string) (rets [][]*model.Visit, errs []error) {
	return func(ctx context.Context, queries [][6]string) (rets [][]*model.Visit, errs []error) {
		rets = make([][]*model.Visit, len(queries))
		errs = make([]error, len(queries))

		res, err := matomoClient.GetConvertedVisitsBulk(ctx, queries...)
		if err != nil {
			for i := range len(queries) {
				rets[i] = nil
				errs[i] = err
			}

			return
		}

		for i := range len(queries) {
			rets[i] = res[i]
		}

		return
	}
}

func GetGoalConvertedVisits(ctx context.Context, idSite int, idGoal, segment string, opts *model.ConvertedVisitsOptions, orderBy *model.OrderByOptions) ([]*model.Visit, error) {
	var dateOpts *model.DateRangeOptions
	if opts != nil && opts.Date.IsSet() {
		dateOpts = opts.Date.Value()
	}

	var query = [6]string{
		strconv.Itoa(idSite),
		idGoal,
	}

	if dateOpts != nil {
		query[2] = dateOpts.Period.String()
		query[3] = dateOpts.StartDate
		if endDate, ok := dateOpts.EndDate.ValueOK(); ok && *endDate != "" {
			query[4] = *endDate
		}
	}

	query[5] = segment

	loaders := For(ctx)
	res, err := loaders.GoalConvertedVisitsLoader.Load(ctx, query)
	if err != nil {
		return nil, err
	}

	if orderBy == nil {
		orderBy = new(model.OrderByOptions)
	}

	if ob, ok := orderBy.Timestamp.ValueOK(); ok {
		sort.Slice(res, func(i, j int) bool {
			if ob.String() == "ASC" {
				return res[i].ServerTimestamp < res[j].ServerTimestamp
			}

			return res[i].ServerTimestamp > res[j].ServerTimestamp
		})
	}

	return res, nil
}
