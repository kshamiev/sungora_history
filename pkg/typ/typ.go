package typ

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null"
)

type UUID struct {
	uuid.UUID
}

func UUIDNew() UUID {
	return UUID{UUID: uuid.Must(uuid.NewRandom())}
}

func UUIDParse(s string) (UUID, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return UUID{}, err
	}
	return UUID{u}, nil
}

func UUIDMustParse(s string) UUID {
	return UUID{uuid.MustParse(s)}
}

func (u UUID) Value() (driver.Value, error) {
	if u.ID() == 0 {
		return nil, nil
	}
	return u.String(), nil

}

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

// ////

type SampleJs struct {
	ID    uint64
	Name  string
	Items []Item
}

type Item struct {
	Price    float64
	Quantity int
}

func (m SampleJs) Value() (driver.Value, error) {
	if cmp.Equal(m, SampleJs{}) {
		return nil, nil
	}
	return json.Marshal(m)
}

// Scan scan value into Jsonb
func (m *SampleJs) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, m)
}
