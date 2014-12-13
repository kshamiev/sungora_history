// Тип Роутинг
package db

import (
	"time"
	"types"
)

// Разделы
type Uri struct {
	Id            uint64    // Id
	Method        []string  // Метод запроса
	Domain        string    // Домен или regexp описывающий домен
	Uri           string    // URI без указания домена и протокола до необязательных параметров
	Name          string    // Название раздела
	Redirect      string    // Если не пусто, то содержит адрес безусловной переадресации
	Layout        string    // Макет сайта
	IsAuthorized  bool      // =1 - доступ к разделу разрешен только авторизованным пользователям
	IsMenuVisible bool      // =1 - раздел отображается в стандартном меню
	IsDisable     bool      // =1 - Раздел отключен, при попытке доступа выдается 404
	Content       []byte    // Шаблон или контент раздела (текстовые или бинарные данные)
	ContentTime   time.Time // Дата и время обновления контента
	ContentType   string    // Mime type контента раздела, может быть переопределён контроллером
	ContentEncode string    // Кодировка контента (по умолчанию utf-8) для заголовка
	Position      int32     // Сортировка, приоритет в роутинге
	Title         string    // Заголовок раздела - title
	KeyWords      string    // Ключевые слова - keywords
	Description   string    // Описание раздела - description
	Controllers   []string  `db:"-"` // Контроллеры урла (Path)
	Groups        []uint64  `db:"-"` // Группы урла (права)
}

func init() {
	// Набор сценариев для типа
	types.SetScenario(`Uri`, map[string]types.Scenario{
		`root`: types.Scenario{
			Name:        `Разделы`,
			Description: `Базовая конфигурация всех свойств для всех сценарией указанного типа`,
			Sample:      new(Uri),
			Property: []types.Property{{
				Name:     `Id`,
				Title:    `Id`,
				Readonly: `yes`,
				Required: `yes`,
				FormType: `hidden`,
			}, {
				Name:     `Method`,
				Title:    `Метод запроса`,
				Default:  `GET`,
				FormType: `checkbox`,
				EnumSet:  map[string]string{`GET`: `GET`, `OPTIONS`: `OPTIONS`, `HEAD`: `HEAD`, `POST`: `POST`, `PUT`: `PUT`, `PATCH`: `PATCH`, `DELETE`: `DELETE`, `TRACE`: `TRACE`, `LINK`: `LINK`, `UNLINK`: `UNLINK`, `CONNECT`: `CONNECT`, `WS`: `WS`},
			}, {
				Name:     `Domain`,
				Title:    `Домен или regexp описывающий домен`,
				FormType: `text`,
			}, {
				Name:     `Uri`,
				Title:    `URI без указания домена и протокола до необязательных параметров`,
				Required: `yes`,
				Default:  `/`,
				FormType: `text`,
			}, {
				Name:     `Name`,
				Title:    `Название раздела`,
				Required: `yes`,
				FormType: `text`,
			}, {
				Name:     `Redirect`,
				Title:    `Если не пусто, то содержит адрес безусловной переадресации`,
				FormType: `text`,
			}, {
				Name:     `Layout`,
				Title:    `Макет сайта`,
				FormType: `text`,
			}, {
				Name:     `IsAuthorized`,
				Title:    `=1 - доступ к разделу разрешен только авторизованным пользователям`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `IsMenuVisible`,
				Title:    `=1 - раздел отображается в стандартном меню`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `IsDisable`,
				Title:    `=1 - Раздел отключен, при попытке доступа выдается 404`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `Content`,
				Title:    `Контент`,
				FormType: `file`,
			}, {
				Name:     `ContentTime`,
				Title:    `Дата и время обновления контента`,
				FormType: `datetime`,
			}, {
				Name:     `ContentType`,
				Title:    `Mime type контента раздела, может быть переопределён контроллером`,
				Required: `yes`,
				Default:  `application/json`,
				FormType: `text`,
			}, {
				Name:     `ContentEncode`,
				Title:    `Кодировка контента (по умолчанию utf-8) для заголовка`,
				Required: `yes`,
				Default:  `utf-8`,
				FormType: `text`,
			}, {
				Name:     `Position`,
				Title:    `Сортировка, приоритет в роутинге`,
				Required: `yes`,
				Default:  `1`,
				FormType: `hidden`,
			}, {
				Name:     `Title`,
				Title:    `Заголовок раздела - title`,
				FormType: `text`,
			}, {
				Name:     `KeyWords`,
				Title:    `Ключевые слова - keywords`,
				FormType: `text`,
			}, {
				Name:     `Description`,
				Title:    `Описание раздела - description`,
				FormType: `textarea`,
			}, {
				Name:     `Controllers`,
				Title:    `Контроллеры урла (Path)`,
				FormType: `linkcross`,
				Uri:      `controllers`,
			}, {
				Name:     `Groups`,
				Title:    `Группы урла (права)`,
				FormType: `relation`,
				Uri:      `groups`,
			}},
		},
		`All`: types.Scenario{
			Description: `Все свойства`,
			Property: []types.Property{{
				Name: `Method`,
			}, {
				Name: `Domain`,
			}, {
				Name: `Uri`,
			}, {
				Name: `Name`,
			}, {
				Name: `Redirect`,
			}, {
				Name: `Layout`,
			}, {
				Name: `IsAuthorized`,
			}, {
				Name: `IsMenuVisible`,
			}, {
				Name: `IsDisable`,
			}, {
				Name: `Content`,
			}, {
				Name: `ContentTime`,
			}, {
				Name: `ContentType`,
			}, {
				Name: `ContentEncode`,
			}, {
				Name: `Position`,
			}, {
				Name: `Title`,
			}, {
				Name: `KeyWords`,
			}, {
				Name: `Description`,
			}, {
				Name: `Controllers`,
			}, {
				Name: `Groups`,
			}},
		},
		`GridAdmin`: types.Scenario{
			Description: `Список разделов в админке`,
			Property: []types.Property{{
				Name: `Id`,
			}, {
				Name: `Uri`,
			}, {
				Name: `Name`,
			}, {
				Name: `Layout`,
			}, {
				Name: `ContentType`,
			}, {
				Name: `ContentEncode`,
			}},
		},
	})
}
