package main

import (
	"errors"
	"fmt"
	"reflect"

	"gopkg.in/kshamiev/sungora.v1/lg"
)

func main() {

	AddRoute("/", &Test{})
	// route["/"].Test()

	obj := route["/"]

	objValue := reflect.ValueOf(obj)
	lg.Dumper(obj)
	met := objValue.MethodByName("POST")
	if met.IsValid() == false {
		fmt.Println("NOT FOUND")
	} else {
		fmt.Println("METHOD FOUND")

		// оставлено для примера передачи параметров в метод
		var params []interface{}
		var in = make([]reflect.Value, 0)
		for i := range params {
			in = append(in, reflect.ValueOf(params[i]))
		}

		// вызов
		out := met.Call(in)
		if nil == out[0].Interface() {
			fmt.Println("call ok")
		} else {
			fmt.Println(out[0].Interface().(error))
		}

	}


}

var route = make(map[string]Face)

func AddRoute(uri string, c Face) {
	route[uri] = c
}

type Face interface {
	POST() (err error)
}

type TestBase struct {
}

func (self *TestBase) POST() (err error) {
	fmt.Println("OK")
	return errors.New("RTYgfhfghfghgfhUI")
}

type Test struct {
	TestBase
}
