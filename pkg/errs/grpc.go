package errs

import (
	"errors"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const delimiter = "!===!"

// GRPC new error type
func GRPC(code codes.Code, err error, msg string, p ...interface{}) error {
	return status.Errorf(code, err.Error()+"; "+trace(2)+delimiter+msg, p...)
}

// NewGRPC new error type
func NewGRPC(err error) *Errs {
	e := &Errs{}
	if s, ok := status.FromError(err); ok {
		l := strings.Split(s.Message(), delimiter)
		e.codeHTTP = 400 // TODO mapping code GRPC to HTTP
		e.codeGRPC = s.Code()
		e.err = errors.New(l[0])
		if len(l) > 1 {
			e.message = []string{l[1]}
		}
	}
	return e
}
