// Тип Связь
package db

import (
	"types"
)

// Связь
type UsersGroups struct {
	Users_Id  uint64 // Пользователь
	Groups_Id uint64 // Группа
	Name      string `db:"-"` // Название (временные значение для выборок)
}

func init() {
	// Набор сценариев для типа
	types.SetScenario(`UsersGroups`, map[string]types.Scenario{
		`root`: types.Scenario{
			Name:        `Связь`,
			Description: `Базовая конфигурация всех свойств для всех сценарией указанного типа`,
			Sample:      new(UsersGroups),
			Property: []types.Property{{
				Name:     `Groups_Id`,
				Title:    `Группа`,
				FormType: `link`,
			}, {
				Name:     `Users_Id`,
				Title:    `Пользователь`,
				FormType: `link`,
			}, {
				Name:     `Name`,
				Title:    `Наименование связываемого объекта`,
				FormType: `text`,
			}},
		},
		`All`: types.Scenario{
			Description: `Все свойства`,
			Property: []types.Property{{
				Name: `Groups_Id`,
			}, {
				Name: `Users_Id`,
			}, {
				Name: `Name`,
			}},
		},
	})
}
