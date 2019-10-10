package errs

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

func Traces(err error) (tr []string) {
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
