package lg

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func Info(code int, messages ...interface{}) {
	if conf.Info == false {
		return
	}
	com := comand{
		level:    "Info",
		message:  searchMsg(code, messages...),
		datetime: datetimeLabal(),
	}
	if conf.Traces == true {
		if tr, err := getTrace(); err == nil {
			com.traces = tr
		}
	}
	logCh <- com
}

func Warning(code int, messages ...interface{}) {
	if conf.Warning == false {
		return
	}
	com := comand{
		level:    "Warning",
		message:  searchMsg(code, messages...),
		datetime: datetimeLabal(),
	}
	if conf.Traces == true {
		if tr, err := getTrace(); err == nil {
			com.traces = tr
		}
	}
	logCh <- com
}

func Error(code int, messages ...interface{}) {
	if conf.Error == false {
		return
	}
	com := comand{
		level:    "Error",
		message:  searchMsg(code, messages...),
		datetime: datetimeLabal(),
	}
	if conf.Traces == true {
		if tr, err := getTrace(); err == nil {
			com.traces = tr
		}
	}
	logCh <- com
}

// Dumper Dump all variables to STDOUT
func Dumper(idl ...interface{}) {
	ret := dump(idl...)
	fmt.Print(ret.String())
}

// dump Dump all variables to bytes.Buffer
func dump(idl ...interface{}) bytes.Buffer {
	var buf bytes.Buffer
	var wr io.Writer

	wr = io.MultiWriter(&buf)
	for _, field := range idl {
		fset := token.NewFileSet()
		ast.Fprint(wr, fset, field, ast.NotNilFilter)
	}
	return buf
}
