package matomo

import (
	"context"
	"fmt"

	"github.com/jalavosus/matomogql/graph/model"
)

func GetSiteFromID(ctx context.Context, idSite int) (*model.Site, error) {
	params, endpoint := buildRequestParams(idSite, "SitesManager.getSiteFromId")

	var result *model.Site
	if err := httpGet(ctx, endpoint, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetSitesFromIDs(ctx context.Context, siteIds ...int) ([]*model.Site, error) {
	params, endpoint := buildRequestParams(-1, "API.getBulkRequest")

	for i, idSite := range siteIds {
		moreParams, _ := buildRequestParams(idSite, "SitesManager.getSiteFromId")
		params.Set(
			fmt.Sprintf("urls[%d]", i), // urls[0], urls[1], etc.
			moreParams.Encode(),
		)
	}

	var result []*model.Site
	if err := httpGet(ctx, endpoint, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetSiteURLsFromID(ctx context.Context, idSite int) ([]string, error) {
	params, endpoint := buildRequestParams(idSite, "SitesManager.getSiteUrlsFromId")

	var result []string
	if err := httpGet(ctx, endpoint, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetSitesWithViewAccess(ctx context.Context) ([]*model.Site, error) {
	params, endpoint := buildRequestParams(noIdSite, "SitesManager.getSitesWithViewAccess")

	var result []*model.Site
	if err := httpGet(ctx, endpoint, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetSitesWithAtLeastViewAccess(ctx context.Context) ([]*model.Site, error) {
	params, endpoint := buildRequestParams(noIdSite, "SitesManager.getSitesWithAtLeastViewAccess")

	var result []*model.Site
	if err := httpGet(ctx, endpoint, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}
