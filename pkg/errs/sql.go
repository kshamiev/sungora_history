package errs

import (
	"database/sql"
)

func NewSQL(err error, message ...string) error {
	if sql.ErrNoRows == err {
		if len(message) > 0 {
			return NewNotFound(err, message...)
		}

		return NewNotFound(err)
	}

	return NewInternalServer(err, message...)
}
