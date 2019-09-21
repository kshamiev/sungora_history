package errs

import (
	"fmt"
	"net/http"
	"path"
	"runtime"
	"strings"
)

// BadRequest httpStatus: 400
type BadRequest struct {
	err     error    // сама ошибка от внешнего сервиса или либы
	kind    string   // где произошла ошибка
	trace   []string // трассировка ошибки
	message string   // сообщение пользователю
}

// Error for logs
func (e *BadRequest) Error() string {
	if e.err != nil {
		return e.err.Error() + "; " + e.kind
	}
	return http.StatusText(http.StatusBadRequest) + "; " + e.kind
}

// Trace for logs
func (e *BadRequest) Trace() []string {
	return e.trace
}

// Response response message to user
func (e *BadRequest) Response() string {
	return e.message
}

// HTTPCode http status response
func (e *BadRequest) HTTPCode() int {
	return http.StatusBadRequest
}

// NewBadRequest new error type
func NewBadRequest(err error, message string) *BadRequest {
	return &BadRequest{
		err:     err,
		kind:    trace(2),
		trace:   traces(err),
		message: message,
	}
}

func traces(err error) (tr []string) {
	kind := ""
	if err != nil {
		kind = err.Error() + "; "
	}
	for i := 4; true; i++ {
		t := trace(i)
		if t == "" {
			break
		}
		if strings.Contains(t, "/src/") {
			continue // LIBRARY GOPATH
		}
		if strings.Contains(t, "/mod/") {
			continue // LIBRARY MOD
		}
		if strings.Contains(t, "/vendor/") {
			continue // LIBRARY VENDOR
		}
		tr = append(tr, kind+t)
	}
	return tr
}

func trace(step int) string {
	pc, file, line, ok := runtime.Caller(step)
	if line == 0 {
		return ""
	}
	kind := fmt.Sprintf("%s:%d", file, line)
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			kind += ":" + path.Base(fn.Name())
		}
	}
	return kind
}
