package matomo

import (
	"context"
	"fmt"

	"github.com/jalavosus/matomogql/graph/model"
)

func (c clientImpl) GetSiteFromID(ctx context.Context, idSite int) (*model.Site, error) {
	params := c.buildRequestParams(idSite, "SitesManager.getSiteFromId")

	var result *model.Site
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c clientImpl) GetSitesFromIDs(ctx context.Context, siteIds ...int) ([]*model.Site, error) {
	params := c.buildRequestParams(-1, "API.getBulkRequest")

	for i, idSite := range siteIds {
		moreParams := c.buildRequestParams(idSite, "SitesManager.getSiteFromId")
		params.Set(
			fmt.Sprintf("urls[%d]", i), // urls[0], urls[1], etc.
			moreParams.Encode(),
		)
	}

	var result []*model.Site
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c clientImpl) GetSiteURLsFromID(ctx context.Context, idSite int) ([]string, error) {
	params := c.buildRequestParams(idSite, "SitesManager.getSiteUrlsFromId")

	var result []string
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c clientImpl) GetSitesWithViewAccess(ctx context.Context) ([]*model.Site, error) {
	params := c.buildRequestParams(noIdSite, "SitesManager.getSitesWithViewAccess")

	var result []*model.Site
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c clientImpl) GetSitesWithAtLeastViewAccess(ctx context.Context) ([]*model.Site, error) {
	params := c.buildRequestParams(noIdSite, "SitesManager.getSitesWithAtLeastViewAccess")

	var result []*model.Site
	if err := c.httpGet(ctx, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}
