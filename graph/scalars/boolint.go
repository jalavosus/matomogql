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
	var (
		rawInt   int
		rawBool  bool
		val      bool
		ok       bool
		gotValue bool
	)

	rawInt, ok = v.(int)
	if ok {
		gotValue = true
		val = rawInt == 1
		goto Ret
	}

	rawBool, ok = v.(bool)
	if ok {
		gotValue = true
		val = rawBool
	}

Ret:
	if !gotValue {
		return fmt.Errorf("error parsing BoolInt %[1]v, has type %[1]T", v)
	}

	*b = BoolInt(val)

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
