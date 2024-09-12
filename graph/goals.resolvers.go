package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/jalavosus/matomogql/graph/loaders"
	"github.com/jalavosus/matomogql/graph/model"
	"github.com/jalavosus/matomogql/matomo"
)

// ConvertedVisits is the resolver for the convertedVisits field.
func (r *goalResolver) ConvertedVisits(ctx context.Context, obj *model.Goal, opts *model.ConvertedVisitsOptions, orderBy *model.OrderByOptions) ([]*model.Visit, error) {
	return loaders.GetGoalConvertedVisits(ctx, obj.IDSite, obj.IDGoal, opts, orderBy)
}

// GetGoal is the resolver for the getGoal field.
func (r *queryResolver) GetGoal(ctx context.Context, idSite int, idGoal int) (*model.Goal, error) {
	return matomo.GetGoal(ctx, idSite, idGoal)
}

// GetGoals is the resolver for the getGoals field.
func (r *queryResolver) GetGoals(ctx context.Context, idSite int, goalIds []int, opts *model.GetGoalsOptions) ([]*model.Goal, error) {
	return matomo.GetGoals(ctx, idSite, goalIds, opts)
}

// GetAllGoals is the resolver for the getAllGoals field.
func (r *queryResolver) GetAllGoals(ctx context.Context, idSite int, opts *model.GetGoalsOptions) ([]*model.Goal, error) {
	return matomo.GetAllGoals(ctx, idSite, opts)
}

// Goal returns GoalResolver implementation.
func (r *Resolver) Goal() GoalResolver { return &goalResolver{r} }

type goalResolver struct{ *Resolver }
