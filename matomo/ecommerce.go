package matomo

import (
	"context"
	"strings"

	"github.com/jalavosus/matomogql/graph/model"
)

func GetEcommerceItemsName(ctx context.Context, idSite int, opts *model.GetEcommerceGoalsOptions) ([]*model.EcommerceGoal, error) {
	return getEcommerceItems(ctx, idSite, opts, "Name")
}

func GetEcommerceItemsSku(ctx context.Context, idSite int, opts *model.GetEcommerceGoalsOptions) ([]*model.EcommerceGoal, error) {
	return getEcommerceItems(ctx, idSite, opts, "Sku")
}

func getEcommerceItems(ctx context.Context, idSite int, opts *model.GetEcommerceGoalsOptions, searchType string) ([]*model.EcommerceGoal, error) {
	params, endpoint := buildRequestParams(idSite, "Goals.getItems"+searchType)
	params.Set("period", strings.ToLower(opts.Date.Period.String()))

	var date = opts.Date.StartDate
	if endDate, ok := opts.Date.EndDate.ValueOK(); ok && *endDate != "" {
		date += "," + *endDate
	}

	params.Set("date", date)

	var result []*model.EcommerceGoal
	if err := httpGet(ctx, endpoint, params, &result); err != nil {
		return nil, err
	}

	for i := range result {
		result[i].IDSite = idSite
	}

	return result, nil
}
