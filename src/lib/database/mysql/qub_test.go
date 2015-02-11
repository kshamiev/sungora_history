// Запуск теста
// SET GOPATH=C:\Work\zegota
// go test -v lib/database/mysql
// go test -v -bench . lib/database/mysql
package mysql_test

import (
	"testing"

	"lib/database"
)

// Тестирование конструктора запросов
func TestQub(t *testing.T) {
	var query = database.NewQub(1).SelectScenario(`Users`, `All`)
	var sql = query.From(`Users as z`).Where(`AND Id < 100`).Order(`Name ASC`).Get()
	t.Log(sql)
}
