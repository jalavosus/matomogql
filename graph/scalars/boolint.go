package scalars

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

type BoolInt bool

func (b BoolInt) MarshalGQL(w io.Writer) {
	graphql.MarshalBoolean(bool(b)).MarshalGQL(w)
}

func (b BoolInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(b))
}

func (b *BoolInt) UnmarshalGQL(v any) error {
	switch x := v.(type) {
	case bool:
		*b = BoolInt(x)
	case int:
		*b = x == 1
	default:
		return fmt.Errorf("unable to parse %[1]v as BoolInt, has type %[1]T", v)
	}

	return nil
}

func (b *BoolInt) UnmarshalJSON(v []byte) error {
	var raw int
	if err := json.Unmarshal(v, &raw); err != nil {
		return err
	}

	*b = raw == 1

	return nil
}
