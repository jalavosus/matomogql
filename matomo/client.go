package matomo

import (
	"context"

	"github.com/jalavosus/matomogql/graph/model"
)

type Client interface {
	GetEcommerceItemsName(ctx context.Context, idSite int, opts *model.GetEcommerceGoalsOptions) ([]*model.EcommerceGoal, error)
	GetEcommerceItemsSku(ctx context.Context, idSite int, opts *model.GetEcommerceGoalsOptions) ([]*model.EcommerceGoal, error)

	GetGoal(ctx context.Context, idSite, idGoal int) (*model.Goal, error)
	GetGoals(ctx context.Context, idSite int, goalIds []int, opts *model.GetGoalsOptions) ([]*model.Goal, error)
	GetAllGoals(ctx context.Context, idSite int, opts *model.GetGoalsOptions) ([]*model.Goal, error)

	GetConvertedVisits(ctx context.Context, idSite, idGoal int, opts *model.ConvertedVisitsOptions) ([]*model.Visit, error)
	GetConvertedVisitsBulk(ctx context.Context, queries ...[6]string) ([][]*model.Visit, error)

	GetSiteFromID(ctx context.Context, idSite int) (*model.Site, error)
	GetSitesFromIDs(ctx context.Context, siteIds ...int) ([]*model.Site, error)
	GetSiteURLsFromID(ctx context.Context, idSite int) ([]string, error)
	GetSitesWithViewAccess(ctx context.Context) ([]*model.Site, error)
	GetSitesWithAtLeastViewAccess(ctx context.Context) ([]*model.Site, error)

	GetVisitorProfile(ctx context.Context, idSite int, visitorId string) (*model.VisitorProfile, error)
	GetVisitorProfiles(ctx context.Context, idSite int, visitorIds []string) ([]*model.VisitorProfile, error)
	GetVisitorProfilesBulk(ctx context.Context, queries ...[2]string) ([]*model.VisitorProfile, error)
	GetLastVisits(ctx context.Context, idSite int, opts *model.LastVisitsOpts) ([]*model.Visit, error)
}

type clientImpl struct {
	apiKey      string
	apiEndpoint string
}

func NewClient(apiKey, apiEndpoint string) Client {
	return clientImpl{apiKey, apiEndpoint}
}
