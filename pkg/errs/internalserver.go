// nolint:dupl
package errs

import (
	"net/http"
)

// InternalServer httpStatus: 500
type InternalServer struct {
	err   error    // сама ошибка от внешнего сервиса или либы
	kind  string   // где произошла ошибка
	trace []string // трассировка ошибки
}

// Error for logs
func (e *InternalServer) Error() string {
	if e.err != nil {
		return e.err.Error() + "; " + e.kind
	}
	return http.StatusText(http.StatusInternalServerError) + "; " + e.kind
}

// Trace for logs
func (e *InternalServer) Trace() []string {
	return e.trace
}

// Response response message to user
func (e *InternalServer) Response() string {
	return http.StatusText(http.StatusInternalServerError)
}

// HTTPCode http status response
func (e *InternalServer) HTTPCode() int {
	return http.StatusInternalServerError
}

// NewInternalServer new error type
func NewInternalServer(err error) *InternalServer {
	return &InternalServer{
		err:   err,
		kind:  trace(2),
		trace: traces(err),
	}
}
