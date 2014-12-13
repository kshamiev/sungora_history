// Тип Группы
package db

import (
	"types"
)

// Группы
type Groups struct {
	Id          uint64 // Id
	Name        string // Наименование
	Description string // Описание группы
	IsDefault   bool   // Группа по умолчанию
	Del         bool   // Запись удалена
	Hash        string // Контрольная сумма для синхронизации (SHA256)
}

func init() {
	// Набор сценариев для типа
	types.SetScenario(`Groups`, map[string]types.Scenario{
		`root`: types.Scenario{
			Name:        `Группы`,
			Description: `Базовая конфигурация всех свойств для всех сценарией указанного типа`,
			Sample:      new(Groups),
			Property: []types.Property{{
				Name:     `Id`,
				Title:    `Id`,
				Readonly: `yes`,
				Required: `yes`,
				FormType: `hidden`,
			}, {
				Name:     `Name`,
				Title:    `Наименование`,
				Required: `yes`,
				FormType: `text`,
			}, {
				Name:     `Description`,
				Title:    `Описание группы`,
				FormType: `textarea`,
			}, {
				Name:     `IsDefault`,
				Title:    `Группа по умолчанию`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `Del`,
				Title:    `Запись удалена`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `Hash`,
				Title:    `Контрольная сумма для синхронизации (SHA256)`,
				FormType: `hidden`,
			}},
		},
		`All`: types.Scenario{
			Description: `Все свойства`,
			Property: []types.Property{{
				Name: `Name`,
			}, {
				Name: `Description`,
			}, {
				Name: `IsDefault`,
			}, {
				Name: `Del`,
			}, {
				Name: `Hash`,
			}},
		},
		`GridAdmin`: types.Scenario{
			Description: `Список групп в админке`,
			Property: []types.Property{{
				Name: `Id`,
			}, {
				Name: `Name`,
			}},
		},
	})
}
