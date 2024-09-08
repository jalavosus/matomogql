package loaders

import (
	"context"
	"strconv"
	"strings"

	"github.com/jalavosus/matomogql/graph/model"
	"github.com/jalavosus/matomogql/matomo"
)

func getVisitorProfiles(ctx context.Context, queries []string) (rets []*model.VisitorProfile, errs []error) {
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
	var b strings.Builder
	b.WriteString(strconv.Itoa(idSite) + ":")
	b.WriteString(visitorId)

	loaders := For(ctx)
	res, err := loaders.VisitorProfilesLoader.Load(ctx, b.String())
	if err != nil {
		return nil, err
	}

	return res, nil
}
