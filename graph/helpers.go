package graph

import (
	"sort"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jalavosus/matomogql/graph/model"
	"github.com/jalavosus/matomogql/utils"
)

func (r *visitorProfileResolver) getVisitBySortOrder(obj *model.VisitorProfile, order model.OrderBy) (*model.Visit, error) {
	visits := obj.LastVisits
	if len(visits) == 0 {
		return nil, nil // or return an appropriate error
	}

	sortedVisits := orderLastVisits(visits, &model.OrderByOptions{Timestamp: graphql.OmittableOf(utils.ToPointer(order))})
	return sortedVisits[0], nil
}

// orderLastVisits sorts lastVisits in-place by ServerTimestamp.
// It defaults to descending order unless a valid ordering option is provided.
func orderLastVisits(lastVisits []*model.Visit, order *model.OrderByOptions) []*model.Visit {
	orderBy := getOrderBy(order)
	asc := orderBy == model.OrderByAsc

	sort.Slice(lastVisits, func(i, j int) bool {
		if asc {
			return lastVisits[i].ServerTimestamp < lastVisits[j].ServerTimestamp
		}
		return lastVisits[i].ServerTimestamp > lastVisits[j].ServerTimestamp
	})

	return lastVisits
}

// getOrderBy safely extracts the desired ordering, falling back to DESC.
func getOrderBy(order *model.OrderByOptions) model.OrderBy {
	if order == nil {
		return model.OrderByDesc
	}
	if ob, ok := order.Timestamp.ValueOK(); ok {
		return *ob
	}
	return model.OrderByDesc
}

func visitActionByType(visit *model.Visit, eventType string) (ad *model.VisitActionDetails) {
	for _, a := range visit.ActionDetails {
		if a.Type == eventType {
			ad = a
			break
		}
	}

	return
}
