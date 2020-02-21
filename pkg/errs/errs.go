package errs

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotFound       = errors.New(http.StatusText(http.StatusNotFound))
	ErrUnauthorized   = errors.New(http.StatusText(http.StatusUnauthorized))
	ErrInternalServer = errors.New(http.StatusText(http.StatusInternalServerError))
	ErrForbidden      = errors.New(http.StatusText(http.StatusForbidden))
	ErrBadRequest     = errors.New(http.StatusText(http.StatusBadRequest))
)

// NewUnauthorized new error type
func NewUnauthorized(err error, msg ...string) *Errs {
	if err == nil {
		err = ErrUnauthorized
	}

	return &Errs{
		codeHTTP: http.StatusUnauthorized,
		err:      err,
		kind:     trace(2),
		trace:    Traces(err),
		message:  msg,
	}
}

// NewUnauthorizedCode new error type
func NewUnauthorizedCode(err error, code int, msg ...interface{}) *Errs {
	if err == nil {
		err = ErrUnauthorized
	}

	return &Errs{
		codeHTTP: http.StatusUnauthorized,
		err:      err,
		kind:     trace(2),
		trace:    Traces(err),
		message:  messageGet(code, msg...),
	}
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
		trace:    Traces(err),
		message:  msg,
	}
}

// NewNotFoundCode new error type
func NewNotFoundCode(err error, code int, msg ...interface{}) *Errs {
	if err == nil {
		err = ErrNotFound
	}

	return &Errs{
		codeHTTP: http.StatusNotFound,
		err:      err,
		kind:     trace(2),
		trace:    Traces(err),
		message:  messageGet(code, msg...),
	}
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
		trace:    Traces(err),
		message:  msg,
	}
}

// NewInternalServerCode new error type
func NewInternalServerCode(err error, code int, msg ...interface{}) *Errs {
	if err == nil {
		err = ErrInternalServer
	}

	return &Errs{
		codeHTTP: http.StatusInternalServerError,
		err:      err,
		kind:     trace(2),
		trace:    Traces(err),
		message:  messageGet(code, msg...),
	}
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
		trace:    Traces(err),
		message:  msg,
	}
}

// NewForbiddenCode new error type
func NewForbiddenCode(err error, code int, msg ...interface{}) *Errs {
	if err == nil {
		err = ErrForbidden
	}

	return &Errs{
		codeHTTP: http.StatusForbidden,
		err:      err,
		kind:     trace(2),
		trace:    Traces(err),
		message:  messageGet(code, msg...),
	}
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
		trace:    Traces(err),
		message:  msg,
	}
}

// NewBadRequestCode new error type
func NewBadRequestCode(err error, code int, msg ...interface{}) *Errs {
	if err == nil {
		err = ErrBadRequest
	}

	return &Errs{
		codeHTTP: http.StatusBadRequest,
		err:      err,
		kind:     trace(2),
		trace:    Traces(err),
		message:  messageGet(code, msg...),
	}
}

func messageGet(code int, msg ...interface{}) []string {
	if _, ok := messageCode[code]; ok {
		return []string{fmt.Sprintf(messageCode[code], msg...)}
	}

	return []string{fmt.Sprintf("message %d not implemented", code)}
}

type Errs struct {
	codeHTTP int      // код http
	err      error    // сама ошибка от внешнего сервиса или либы
	kind     string   // где произошла ошибка
	trace    []string // трассировка ошибки
	message  []string // сообщение для пользователя
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
