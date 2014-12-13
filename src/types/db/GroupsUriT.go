// Тип Права
package db

import (
	"types"
)

// Права
type GroupsUri struct {
	Groups_Id      uint64 // Группа
	Uri_Id         uint64 // Роутинг
	Controllers_Id uint64 // Контроллер
	Get            bool   // Запрос на получение данных
	Post           bool   // Запрос на добавление данных
	Put            bool   // Запрос на изменение данных
	Delete         bool   // Запрос на удаление данных
	Options        bool   // Запрос на получение опций
	Disable        bool   // Запрещен запуск контроллера (для группы)
	Name           string `db:"-"` // Название (временные значение для выборок)
}

func init() {
	// Набор сценариев для типа
	types.SetScenario(`GroupsUri`, map[string]types.Scenario{
		`root`: types.Scenario{
			Name:        `Права`,
			Description: `Базовая конфигурация всех свойств для всех сценарией указанного типа`,
			Sample:      new(GroupsUri),
			Property: []types.Property{{
				Name:     `Groups_Id`,
				Title:    `Группа`,
				FormType: `link`,
			}, {
				Name:     `Uri_Id`,
				Title:    `Роутинг`,
				FormType: `link`,
			}, {
				Name:     `Controllers_Id`,
				Title:    `Контроллер`,
				FormType: `link`,
			}, {
				Name:     `Get`,
				Title:    `Запрос на получение данных`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `Post`,
				Title:    `Запрос на добавление данных`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `Put`,
				Title:    `Запрос на изменение данных`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `Delete`,
				Title:    `Запрос на удаление данных`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `Options`,
				Title:    `Запрос на получение опций`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `Disable`,
				Title:    `Запрещен запуск контроллера (для группы)`,
				Required: `yes`,
				FormType: `bool`,
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
				Name: `Uri_Id`,
			}, {
				Name: `Controllers_Id`,
			}, {
				Name: `Get`,
			}, {
				Name: `Post`,
			}, {
				Name: `Put`,
			}, {
				Name: `Delete`,
			}, {
				Name: `Options`,
			}, {
				Name: `Disable`,
			}, {
				Name: `Name`,
			}},
		},
	})
}
