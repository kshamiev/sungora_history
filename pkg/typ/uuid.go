package typ

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type UUID struct {
	uuid.UUID
}

type UUIDS []UUID

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
