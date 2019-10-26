package typ

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null"
)

func (u *UUID) UnmarshalGQL(v interface{}) error {
	const errWrongUUID = "wrong uuid"
	switch data := v.(type) {
	case UUID:
		*u = data
	case string:
		value, err := UUIDParse(data)
		if err != nil {
			return errors.New(errWrongUUID)
		}
		*u = value
	default:
		return fmt.Errorf("wrong uuid")
	}
	return nil
}

func (u UUID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(u.String()))
}

func MarshalDecimal(d decimal.Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(d.String()))
	})
}

func UnmarshalDecimal(v interface{}) (decimal.Decimal, error) {
	switch val := v.(type) {
	case decimal.Decimal:
		return val, nil
	case string:
		d, err := decimal.NewFromString(val)
		if err != nil {
			return decimal.Zero, err
		}
		return d, nil
	case float32:
		return decimal.NewFromFloat32(val), nil
	case float64:
		return decimal.NewFromFloat(val), nil
	default:
		return decimal.Zero, fmt.Errorf("%T is not a deciaml", v)
	}
}

func MarshalNullTime(t null.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(t.Time.String()))
	})
}

func UnmarshalNullTime(v interface{}) (null.Time, error) {
	switch val := v.(type) {
	case null.Time:
		return val, nil
	default:
		return null.Time{}, fmt.Errorf("%T is not a null.Time", v)
	}
}
