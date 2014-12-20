// qwerty project qwerty.go
package qwerty

import (
	"lib/logs"
	"runtime"
)

func QwertyTwo() {

	logs.Dumper(runtime.Caller(0))
	logs.Dumper(runtime.Caller(1))
	logs.Dumper(runtime.Caller(2))
	logs.Dumper(runtime.Caller(3))
	logs.Dumper(runtime.Caller(4))
	logs.Dumper(runtime.Caller(5))

	//var fn *runtime.Func
	//var fName string
	//n := 1
	//for pc, file, line, ok := runtime.Caller(n); ok == true; {

	//	fn = runtime.FuncForPC(pc)
	//	if fn != nil {
	//		fName = fn.Name()
	//	}
	//	logs.Dumper(fName, file, line, n)
	//	n++
	//	if n > 8 {
	//		break
	//	}
	//}

}
