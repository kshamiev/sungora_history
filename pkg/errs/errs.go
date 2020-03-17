package errs

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound       = errors.New(http.StatusText(http.StatusNotFound))
	ErrUnauthorized   = errors.New(http.StatusText(http.StatusUnauthorized))
	ErrInternalServer = errors.New(http.StatusText(http.StatusInternalServerError))
	ErrForbidden      = errors.New(http.StatusText(http.StatusForbidden))
	ErrBadRequest     = errors.New(http.StatusText(http.StatusBadRequest))
)

// New401 new error type
func New401(err error, msg ...string) *Errs {
	return NewUnauthorized(err, msg...)
}

// NewUnauthorized new error type
func NewUnauthorized(err error, msg ...string) *Errs {
	if err == nil {
		err = ErrUnauthorized
	}
	return &Errs{
		codeHTTP: http.StatusUnauthorized,
		err:      err,
		kind:     trace(2),
		message:  msg,
	}
}

// New404 new error type
func New404(err error, msg ...string) *Errs {
	return NewNotFound(err, msg...)
}

// NewNotFound new error type
func NewNotFound(err error, msg ...string) *Errs {
	if err == nil {
		err = ErrNotFound
	}
	return &Errs{
		codeHTTP: http.StatusNotFound,
		err:      err,
		kind:     trace(2),
		message:  msg,
	}
}

// New500 new error type
func New500(err error, msg ...string) *Errs {
	return NewInternalServer(err, msg...)
}

// NewInternalServer new error type
func NewInternalServer(err error, msg ...string) *Errs {
	if err == nil {
		err = ErrInternalServer
	}

	return &Errs{
		codeHTTP: http.StatusInternalServerError,
		err:      err,
		kind:     trace(2),
		message:  msg,
	}
}

// New403 new error type
func New403(err error, msg ...string) *Errs {
	return NewForbidden(err, msg...)
}

// NewForbidden new error type
func NewForbidden(err error, msg ...string) *Errs {
	if err == nil {
		err = ErrForbidden
	}

	return &Errs{
		codeHTTP: http.StatusForbidden,
		err:      err,
		kind:     trace(2),
		message:  msg,
	}
}

// New400 new error type
func New400(err error, msg ...string) *Errs {
	return NewBadRequest(err, msg...)
}

// NewBadRequest new error type
func NewBadRequest(err error, msg ...string) *Errs {
	if err == nil {
		err = ErrBadRequest
	}

	return &Errs{
		codeHTTP: http.StatusBadRequest,
		err:      err,
		kind:     trace(2),
		message:  msg,
	}
}

type Errs struct {
	codeHTTP int      // код http
	err      error    // сама ошибка от внешнего сервиса или либы
	kind     string   // где произошла ошибка
	message  []string // сообщение для пользователя
}

// Error for logs
func (e *Errs) Error() string {
	if e.err != nil {
		return e.err.Error() + "; " + e.kind
	}

	return http.StatusText(e.codeHTTP) + "; " + e.kind
}

// Response response message to user
func (e *Errs) Response() string {
	if len(e.message) > 0 {
		return e.message[0]
	} else if e.err != nil {
		return e.err.Error()
	}

	return http.StatusText(e.codeHTTP)
}

// HTTPCode http status response
func (e *Errs) HTTPCode() int {
	return e.codeHTTP
}
