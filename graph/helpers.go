package graph

import (
	"sort"
	
	"github.com/jalavosus/matomogql/graph/model"
)

func orderLastVisits(lastVisits []*model.Visit, order *model.OrderByOptions) []*model.Visit {
	orderBy := model.OrderByDesc
	if order != nil {
		if ob, ok := order.Timestamp.ValueOK(); ok {
			orderBy = *ob
		}
	}

	sort.Slice(lastVisits, func(i, j int) bool {
		if orderBy == model.OrderByAsc {
			return lastVisits[i].ServerTimestamp < lastVisits[j].ServerTimestamp
		}

		return lastVisits[i].ServerTimestamp > lastVisits[j].ServerTimestamp
	})

	return lastVisits
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
