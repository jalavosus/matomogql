package matomo

import (
	"context"
	"strings"

	"github.com/jalavosus/matomogql/graph/model"
)

func (c clientImpl) GetEcommerceItemsName(ctx context.Context, idSite int, opts *model.GetEcommerceGoalsOptions) ([]*model.EcommerceGoal, error) {
	return c.getEcommerceItems(ctx, idSite, opts, "Name")
}

func (c clientImpl) GetEcommerceItemsSku(ctx context.Context, idSite int, opts *model.GetEcommerceGoalsOptions) ([]*model.EcommerceGoal, error) {
	return c.getEcommerceItems(ctx, idSite, opts, "Sku")
}

func (c clientImpl) getEcommerceItems(ctx context.Context, idSite int, opts *model.GetEcommerceGoalsOptions, searchType string) ([]*model.EcommerceGoal, error) {
	params := c.buildRequestParams(idSite, "Goals.getItems"+searchType)
	params.Set("period", strings.ToLower(opts.Date.Period.String()))

	date := opts.Date.StartDate
	if endDate, ok := opts.Date.EndDate.ValueOK(); ok && *endDate != "" {
		date += "," + *endDate
	}
	params.Set("date", date)

	var result []*model.EcommerceGoal
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	for i := range result {
		result[i].IDSite = idSite
	}

	return result, nil
}
