package errs

import (
	"database/sql"
	"net/http"
)

func SQLNoRows(err error) error {
	if sql.ErrNoRows == err {
		return NewNotFound(err, http.StatusText(http.StatusNotFound))
	}
	return NewInternalServer(err)
}
