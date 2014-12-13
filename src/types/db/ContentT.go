// Тип Контент
package db

import (
	`types`
)

// Контент
type Content struct {
	Uri_Id      uint64 // Uri
	Lang        string // Язык
	Title       string // Заголовок
	Keywords    string // Ключи
	Description string // Описание
	Content     []byte // Контент
	Block       string // Блок
}

func init() {
	// Набор сценариев для типа
	types.SetScenario(`Content`, map[string]types.Scenario{
		`root`: types.Scenario{
			Name:        `Контент`,
			Description: `Базовая конфигурация всех свойств для всех сценарией указанного типа`,
			Sample:      new(Content),
			Property: []types.Property{{
				Name:     `Uri_Id`,
				Title:    `Uri`,
				Required: `yes`,
				FormType: `link`,
			}, {
				Name:     `Lang`,
				Title:    `Язык`,
				Required: `yes`,
				FormType: `text`,
			}, {
				Name:     `Title`,
				Title:    `Заголовок`,
				FormType: `text`,
			}, {
				Name:     `Keywords`,
				Title:    `Ключи`,
				FormType: `text`,
			}, {
				Name:     `Description`,
				Title:    `Описание`,
				FormType: `text`,
			}, {
				Name:     `Content`,
				Title:    `Контент`,
				FormType: `file`,
			}, {
				Name:     `Block`,
				Title:    `Блок`,
				Required: `yes`,
				FormType: `text`,
			}},
		},
		`All`: types.Scenario{
			Description: `Все свойства`,
			Property: []types.Property{{
				Name: `Uri_Id`,
			}, {
				Name: `Lang`,
			}, {
				Name: `Title`,
			}, {
				Name: `Keywords`,
			}, {
				Name: `Description`,
			}, {
				Name: `Content`,
			}, {
				Name: `Block`,
			}},
		},
	})
}
