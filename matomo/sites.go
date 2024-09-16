package matomo

import (
	"context"

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
