package scalars

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
)

func MarshalDecimal(d decimal.Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, d.String())
	})
}

func UnmarshalDecimal(v any) (decimal.Decimal, error) {
	var d decimal.Decimal

	str, ok := v.(string)
	if !ok {
		return d, fmt.Errorf("UnmarshalDecimal: expected string, got %T", v)
	}

	return decimal.NewFromString(str)
}
