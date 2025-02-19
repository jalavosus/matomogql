package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"
)

// HelloWorld is the resolver for the helloWorld field.
func (r *queryResolver) HelloWorld(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: HelloWorld - helloWorld"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
