package matomo

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
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
	endpoint += "?" + params.Encode()

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

	var result []*model.EcommerceGoal
	if err := json.Unmarshal(bodyRaw, &result); err != nil {
		return nil, err
	}

	for i := range result {
		result[i].IDSite = idSite
	}

	return result, nil
}
