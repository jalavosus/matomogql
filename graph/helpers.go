package graph

import (
	"sort"

	"github.com/jalavosus/matomogql/graph/model"
)

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
