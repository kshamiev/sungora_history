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
	if u.ID() > 0 {
		_, _ = io.WriteString(w, strconv.Quote(u.String()))
		return
	}
	_, _ = io.WriteString(w, `""`)
}

// Decimal

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

// NullTime

func MarshalNullTime(t null.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		if t.Valid {
			_, _ = io.WriteString(w, strconv.Quote(t.Time.String()))
			return
		}
		_, _ = io.WriteString(w, `""`)
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

// NullString

func MarshalNullString(s null.String) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(s.String))
	})
}

func UnmarshalNullString(v interface{}) (null.String, error) {
	switch val := v.(type) {
	case null.String:
		return val, nil
	case string:
		if val == "" {
			return null.StringFrom(val), nil
		}
		return null.String{}, nil
	default:
		return null.String{}, fmt.Errorf("%T is not a null.String", v)
	}
}
