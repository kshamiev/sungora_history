package typ

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

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
		if val == "" {
			val = "0"
		}

		d, err := decimal.NewFromString(val)

		if err != nil {
			return decimal.Zero, err
		}

		return d, nil
	case int:
		return decimal.New(int64(val), 0), nil
	case int64:
		return decimal.New(val, 0), nil
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
		_, _ = io.WriteString(w, strconv.Quote(t.Time.String()))
	})
}

func UnmarshalNullTime(v interface{}) (null.Time, error) {
	switch val := v.(type) {
	case string:
		t, err := time.Parse(time.RFC3339, val)
		if err != nil {
			return null.Time{}, fmt.Errorf("%T is not a null.Time", val)
		}

		return null.TimeFrom(t), nil
	case null.Time:
		return val, nil
	default:
		return null.Time{}, fmt.Errorf("%T is not a null.Time", v)
	}
}

// NullString

func MarshalNullString(str null.String) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(str.String))
	})
}

func UnmarshalNullString(v interface{}) (null.String, error) {
	switch val := v.(type) {
	case null.String:
		return val, nil
	case string:
		return null.StringFrom(val), nil
	default:
		return null.String{}, fmt.Errorf("%T is not a null.String", v)
	}
}

// Time

func MarshalTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(t.Format(time.RFC3339)))
	})
}

func UnmarshalTime(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(time.RFC3339, tmpStr)
	}

	return time.Time{}, errors.New("time should be RFC3339 formatted string")
}
