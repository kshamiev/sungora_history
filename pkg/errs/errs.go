package errs

import "net/http"

// NewUnauthorized new error type
func NewUnauthorized(err error) *Errs {
	return &Errs{
		codeHTTP: http.StatusUnauthorized,
		err:      err,
		kind:     trace(2),
		trace:    Traces(err),
	}
}

// NewNotFound new error type
func NewNotFound(err error) *Errs {
	return &Errs{
		codeHTTP: http.StatusNotFound,
		err:      err,
		kind:     trace(2),
		trace:    Traces(err),
	}
}

// NewInternalServer new error type
func NewInternalServer(err error) *Errs {
	return &Errs{
		codeHTTP: http.StatusInternalServerError,
		err:      err,
		kind:     trace(2),
		trace:    Traces(err),
	}
}

// NewForbidden new error type
func NewForbidden(err error) *Errs {
	return &Errs{
		codeHTTP: http.StatusForbidden,
		err:      err,
		kind:     trace(2),
		trace:    Traces(err),
	}
}

// Errs new error type
func NewBadRequest(err error) *Errs {
	return &Errs{
		codeHTTP: http.StatusBadRequest,
		err:      err,
		kind:     trace(2),
		trace:    Traces(err),
	}
}

type Errs struct {
	codeHTTP int      // код http
	err      error    // сама ошибка от внешнего сервиса или либы
	kind     string   // где произошла ошибка
	trace    []string // трассировка ошибки
}

// Error for logs
func (e *Errs) Error() string {
	if e.err != nil {
		return e.err.Error() + "; " + e.kind
	}
	return http.StatusText(e.codeHTTP) + "; " + e.kind
}

// Trace for logs
func (e *Errs) Trace() []string {
	return e.trace
}

// Response response message to user
func (e *Errs) Response() string {
	if e.codeHTTP != http.StatusInternalServerError && e.err != nil {
		return e.err.Error()
	}
	return http.StatusText(e.codeHTTP)
}

// HTTPCode http status response
func (e *Errs) HTTPCode() int {
	return e.codeHTTP
}
