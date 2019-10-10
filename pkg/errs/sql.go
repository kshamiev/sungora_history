package errs

import (
	"database/sql"
	"errors"
)

func NewSQL(err error, message ...string) error {
	if sql.ErrNoRows == err {
		if len(message) > 0 {
			return NewNotFound(errors.New(message[0]))
		}
		return NewNotFound(nil)
	}
	return NewInternalServer(err)
}
