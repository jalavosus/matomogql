package loaders

import (
	"context"
	"strconv"

	"github.com/jalavosus/matomogql/graph/model"
	"github.com/jalavosus/matomogql/matomo"
)

func getVisitorProfiles(ctx context.Context, queries [][2]string) (rets []*model.VisitorProfile, errs []error) {
	rets = make([]*model.VisitorProfile, len(queries))
	errs = make([]error, len(queries))

	res, err := matomo.GetVisitorProfilesBulk(ctx, queries...)
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

func GetVisitorProfile(ctx context.Context, idSite int, visitorId string) (*model.VisitorProfile, error) {
	loaders := For(ctx)
	res, err := loaders.VisitorProfilesLoader.Load(ctx, [2]string{strconv.Itoa(idSite), visitorId})
	if err != nil {
		return nil, err
	}

	return res, nil
}
