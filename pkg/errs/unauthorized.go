// nolint:dupl
package errs

import (
	"net/http"
)

// Unauthorized httpStatus: 401
type Unauthorized struct {
	err   error    // сама ошибка от внешнего сервиса или либы
	kind  string   // где произошла ошибка
	trace []string // трассировка ошибки
}

// Error for logs
func (e *Unauthorized) Error() string {
	if e.err != nil {
		return e.err.Error() + "; " + e.kind
	}
	return http.StatusText(http.StatusUnauthorized) + "; " + e.kind
}

// Trace for logs
func (e *Unauthorized) Trace() []string {
	return e.trace
}

// Response response message to user
func (e *Unauthorized) Response() string {
	return http.StatusText(http.StatusUnauthorized)
}

// HTTPCode http status response
func (e *Unauthorized) HTTPCode() int {
	return http.StatusUnauthorized
}

// NewUnauthorized new error type
func NewUnauthorized(err error) *Unauthorized {
	return &Unauthorized{
		err:   err,
		kind:  trace(2),
		trace: traces(err),
	}
}
