package loaders

import (
	"context"
	"sort"
	"strconv"

	"github.com/jalavosus/matomogql/graph/model"
	"github.com/jalavosus/matomogql/matomo"
)

func getGoalConvertedVisits(ctx context.Context, queries [][6]string) (rets [][]*model.Visit, errs []error) {
	rets = make([][]*model.Visit, len(queries))
	errs = make([]error, len(queries))

	res, err := matomo.GetConvertedVisitsBulk(ctx, queries...)
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

func GetGoalConvertedVisits(ctx context.Context, idSite int, idGoal string, opts *model.ConvertedVisitsOptions, orderBy *model.OrderByOptions) ([]*model.Visit, error) {
	var query = [6]string{
		strconv.Itoa(idSite),
		idGoal,
		opts.Period.String(),
		opts.StartDate,
	}

	if endDate, ok := opts.EndDate.ValueOK(); ok && *endDate != "" {
		query[4] = *endDate
	}

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
