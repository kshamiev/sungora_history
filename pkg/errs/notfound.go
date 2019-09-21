// nolint:dupl
package errs

import (
	"net/http"
)

// NotFound httpStatus: 404
type NotFound struct {
	err     error    // сама ошибка от внешнего сервиса или либы
	kind    string   // где произошла ошибка
	trace   []string // трассировка ошибки
	message string   // сообщение пользователю
}

// Error for logs
func (e *NotFound) Error() string {
	if e.err != nil {
		return e.err.Error() + "; " + e.kind
	}
	return http.StatusText(http.StatusNotFound) + "; " + e.kind
}

// Trace for logs
func (e *NotFound) Trace() []string {
	return e.trace
}

// Response response message to user
func (e *NotFound) Response() string {
	return e.message
}

// HTTPCode http status response
func (e *NotFound) HTTPCode() int {
	return http.StatusNotFound
}

// NewNotFound new error type
func NewNotFound(err error, message string) *NotFound {
	return &NotFound{
		err:     err,
		kind:    trace(2),
		trace:   traces(err),
		message: message,
	}
}
