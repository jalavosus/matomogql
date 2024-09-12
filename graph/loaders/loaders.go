package loaders

import (
	"context"
	"net/http"
	"time"

	"github.com/vikstrous/dataloadgen"

	"github.com/jalavosus/matomogql/graph/model"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type Loaders struct {
	GoalConvertedVisitsLoader *dataloadgen.Loader[[6]string, []*model.Visit]
	VisitorProfilesLoader     *dataloadgen.Loader[[2]string, *model.VisitorProfile]
}

func NewLoaders() *Loaders {
	return &Loaders{
		GoalConvertedVisitsLoader: dataloadgen.NewLoader(
			getGoalConvertedVisits,
			dataloadgen.WithWait(8*time.Millisecond),
			dataloadgen.WithBatchCapacity(8),
		),
		VisitorProfilesLoader: dataloadgen.NewLoader(
			getVisitorProfiles,
			dataloadgen.WithWait(8*time.Millisecond),
			dataloadgen.WithBatchCapacity(10),
		),
	}
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// Middleware injects data loaders into the context
func Middleware(next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders()
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
		next.ServeHTTP(w, r)
	})
}
