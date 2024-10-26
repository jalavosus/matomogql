package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/jalavosus/matomogql/graph/model"
)

// GetVisitorProfile is the resolver for the getVisitorProfile field.
func (r *queryResolver) GetVisitorProfile(ctx context.Context, idSite int, visitorID string) (*model.VisitorProfile, error) {
	return r.matomoClient.GetVisitorProfile(ctx, idSite, visitorID)
}

// GetVisitorProfiles is the resolver for the getVisitorProfiles field.
func (r *queryResolver) GetVisitorProfiles(ctx context.Context, idSite int, visitorIds []string) ([]*model.VisitorProfile, error) {
	return r.matomoClient.GetVisitorProfiles(ctx, idSite, visitorIds)
}
