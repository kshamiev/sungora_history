// Разработка взаимосвязи с php
//
// Для формирования динамического вывода
package main

import (
	"fmt"
	"net/http"

	"lib/logs"
)

var tpl = `
<!DOCTYPE html>
<html>
  <head>
  </head>
  <body>
    <br>

    <form action="/" method="post">
        <input type="hidden" name="act" value="Login">
        <table width="300px" cellspacing="2" cellpadding="4" border="1" align="center">
            <tr>
                <th colspan="2">Вход</th>
            </tr>
            <tr>
                <td width="50px">Login:</td>
                <td>
                    <input type="text" name="Login">
                </td>
            </tr>
            <tr>
                <td>Password:</td>
                <td>
                    <input type="password" name="Password">
                </td>
            </tr>
            <tr>
                <td colspan="2" align="center">
                    <input type="submit" value="вход">
                </td>
            </tr>
        </table>
    </form>

  </body>
<html>
`

var anError = `<p style="color: red;">%s</p>`
var anSuccess = `<p style="color: green;"><pre>%s</pre></p>`

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		logs.Dumper(err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // Must be called before writing response
	fmt.Fprint(writer, tpl)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		fmt.Fprintf(writer, anSuccess, logs.DumperString(request.Form))
	}
	fmt.Println("/ OK")
}
