package loaders

import (
	"context"
	"sort"
	"strconv"
	"strings"

	"github.com/jalavosus/matomogql/graph/model"
	"github.com/jalavosus/matomogql/matomo"
)

func getGoalConvertedVisits(ctx context.Context, queries []string) (rets [][]*model.Visit, errs []error) {
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

func GetGoalConvertedVisits(ctx context.Context, idSite int, idGoal int, opts *model.ConvertedVisitsOptions, orderBy *model.OrderByOptions) ([]*model.Visit, error) {
	var b strings.Builder
	b.WriteString(strconv.Itoa(idSite) + ":")
	b.WriteString(strconv.Itoa(idGoal) + ":")
	b.WriteString(strings.ToLower(opts.Period.String()) + ":")
	b.WriteString(opts.StartDate + ":")
	if endDate, ok := opts.EndDate.ValueOK(); ok && len(*endDate) > 0 {
		b.WriteString(*endDate)
	}

	loaders := For(ctx)
	res, err := loaders.GoalConvertedVisitsLoader.Load(ctx, b.String())
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
