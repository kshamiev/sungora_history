// Тип Контроллеры
package db

import (
	"time"
	"types"
)

// Контроллеры
type Controllers struct {
	Id          uint64    // Id
	Name        string    // Человеческое название контроллера
	Path        string    // Контроллер (модуль/контроллер/метод)
	IsBefore    bool      // Порядок выполнения контроллеров по умолчанию (до или после)
	IsInternal  bool      // Внутренний контроллер
	IsDefault   bool      // Контроллер по умолчанию
	IsHidden    bool      // Скрытый контроллер
	Position    int32     // Сортировка (приоритет выполнения)
	Date        time.Time // Дата регистрации контроллера
	Domain      string    // Домен или regexp описывающий домен
	Content     string    // Контент контроллера (текстовые или бинарные данные)
	ContentTime time.Time // Дата и время копии файла контента в файловой системе
	Groups      []uint64  `db:"-"` // Группы контроллера (права)
}

func init() {
	// Набор сценариев для типа
	types.SetScenario(`Controllers`, map[string]types.Scenario{
		`root`: types.Scenario{
			Name:        `Контроллеры`,
			Description: `Базовая конфигурация всех свойств для всех сценарией указанного типа`,
			Sample:      new(Controllers),
			Property: []types.Property{{
				Name:     `Id`,
				Title:    `Id`,
				Readonly: `yes`,
				Required: `yes`,
				FormType: `hidden`,
			}, {
				Name:     `Name`,
				Title:    `Человеческое название контроллера`,
				FormType: `text`,
			}, {
				Name:     `Path`,
				Title:    `Контроллер (модуль/контроллер/метод)`,
				Required: `yes`,
				FormType: `text`,
			}, {
				Name:     `IsBefore`,
				Title:    `Порядок выполнения контроллеров по умолчанию (до или после)`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `IsInternal`,
				Title:    `Внутренний контроллер`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `IsDefault`,
				Title:    `Контроллер по умолчанию`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `IsHidden`,
				Title:    `Скрытый контроллер`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `Position`,
				Title:    `Сортировка (приоритет выполнения)`,
				Required: `yes`,
				Default:  `1`,
				FormType: `hidden`,
			}, {
				Name:     `Date`,
				Title:    `Дата регистрации контроллера`,
				FormType: `datetime`,
			}, {
				Name:     `Domain`,
				Title:    `Домен или regexp описывающий домен`,
				FormType: `text`,
			}, {
				Name:     `Content`,
				Title:    `Контент`,
				FormType: `content`,
			}, {
				Name:     `ContentTime`,
				Title:    `Дата и время копии файла контента в файловой системе`,
				FormType: `datetime`,
			}, {
				Name:     `Groups`,
				Title:    `Группы контроллера (права)`,
				FormType: `linkcross`,
				Uri:      `groups`,
			}},
		},
		`All`: types.Scenario{
			Description: `Все свойства`,
			Property: []types.Property{{
				Name: `Name`,
			}, {
				Name: `Path`,
			}, {
				Name: `IsBefore`,
			}, {
				Name: `IsInternal`,
			}, {
				Name: `IsDefault`,
			}, {
				Name: `IsHidden`,
			}, {
				Name: `Position`,
			}, {
				Name: `Date`,
			}, {
				Name: `Domain`,
			}, {
				Name: `Content`,
			}, {
				Name: `ContentTime`,
			}, {
				Name: `Groups`,
			}},
		},
		`GridAdmin`: types.Scenario{
			Description: `Список контроллеров в админке`,
			Property: []types.Property{{
				Name: `Id`,
			}, {
				Name: `Name`,
			}, {
				Name: `Path`,
			}, {
				Name: `Date`,
			}},
		},
	})
}
