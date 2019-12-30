package typ

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

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
