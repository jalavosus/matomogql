package scalars

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

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
		if x != "" {
			*s = strings.Split(x, ",")
		} else {
			*s = []string{}
		}
	case []string:
		*s = x
	default:
		return fmt.Errorf("unable to parse %[1]v as StringList, has type %[1]T", v)
	}

	return nil
}

func (s *StringList) UnmarshalJSON(v []byte) error {
	var raw string
	if err := json.Unmarshal(v, &raw); err != nil {
		return err
	}

	if raw != "" {
		*s = strings.Split(raw, ",")
	} else {
		*s = []string{}
	}

	return nil
}
