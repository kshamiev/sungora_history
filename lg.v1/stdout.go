package lg

import "fmt"

func saveStdout(com comand) {
	fmt.Println(com.message)

	d, err := getCallerInfo(1)
	fmt.Println(d, err)




}

