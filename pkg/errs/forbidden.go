// nolint:dupl
package errs

import (
	"net/http"
)

// Forbidden httpStatus: 403
type Forbidden struct {
	err   error    // сама ошибка от внешнего сервиса или либы
	kind  string   // где произошла ошибка
	trace []string // трассировка ошибки
}

// Error for logs
func (e *Forbidden) Error() string {
	if e.err != nil {
		return e.err.Error() + "; " + e.kind
	}
	return http.StatusText(http.StatusForbidden) + "; " + e.kind
}

// Trace for logs
func (e *Forbidden) Trace() []string {
	return e.trace
}

// Response response message to user
func (e *Forbidden) Response() string {
	return http.StatusText(http.StatusForbidden)
}

// HTTPCode http status response
func (e *Forbidden) HTTPCode() int {
	return http.StatusForbidden
}

// NewForbidden new error type
func NewForbidden(err error) *Forbidden {
	return &Forbidden{
		err:   err,
		kind:  trace(2),
		trace: traces(err),
	}
}
