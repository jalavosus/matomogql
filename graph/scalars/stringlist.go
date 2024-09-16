package scalars

import (
	"encoding/json"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

type StringList []string

func (s StringList) MarshalGQL(w io.Writer) {
	graphql.MarshalAny(s).MarshalGQL(w)
}

func (s StringList) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string(s))
}

func (s *StringList) UnmarshalGQL(v any) error {
	switch x := v.(type) {
	case string:
		*s = StringList()
	}
}
