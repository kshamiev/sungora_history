// запуск теста
// SET GOPATH=C:\Work\zegota
// go test -v types/db | go test -v -bench . types/db
package db_test

import (
	"testing"

	"types"
	_ "types/db"
)

func TestTypes(t *testing.T) {
	var sourceList = []string{
		`Content`,
		`Controllers`,
		`Groups`,
		`GroupsUri`,
		`Uri`,
		`Users`,
	}
	for _, source := range sourceList {
		scenario, err := types.GetScenario(source, `All`)
		if err != nil {
			t.Log(err.Error())
		}
		for _, p := range scenario.Property {
			if p.Title == `` || p.FormType == `` || p.AliasDb == `` {
				t.Errorf(`Ошибка свойства: %s в источнике: %s`, source, p.Name)
			}
		}
	}
}
