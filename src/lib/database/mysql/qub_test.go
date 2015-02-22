// Запуск теста
// SET GOPATH=C:\Work\projectName
// go test -v lib/database/mysql
// go test -v -bench . lib/database/mysql
package mysql_test

import (
	"os"
	"path/filepath"
	"testing"

	"lib/database"
	"lib/logs"
)

// Тестирование конструктора запросов
func TestQub(t *testing.T) {

	var cfglogs = new(logs.Cfglogs)
	cfglogs.Debug = true
	cfglogs.DebugDetail = 0
	cfglogs.Level = 6
	cfglogs.Mode = `file`
	cfglogs.Lang = `ru-ru`
	cfglogs.Separator = ` `

	cfglogs.File, _ = os.Getwd()
	cfglogs.File = filepath.Dir(cfglogs.File)
	cfglogs.File = filepath.Dir(cfglogs.File)
	cfglogs.File = filepath.Dir(cfglogs.File) + `/log`
	os.MkdirAll(cfglogs.File, 0777)
	cfglogs.File += `/logs_test.log`

	logs.Init(cfglogs)
	logs.GoStart(nil)

	var query = database.NewQub(1).Select(`Login, Email`)
	var sql = query.From(`Users as z`).Where(`AND Id < 100`).Order(`Name ASC`).Get()
	t.Log(sql)

	if logs.GoClose() == false {
		t.Fatal(`Ошибка остановки службы логирования`)
	}

}
