// Тип Пользователи
package db

import (
	"time"
	"types"
)

// Пользователи
type Users struct {
	Id          uint64    // Id
	Users_Id    uint64    // Пользователь
	Login       string    // Логин пользователя
	Password    string    // Пароль пользователя (SHA256)
	PasswordR   string    `db:"-"` // Пароль пользователя (SHA256) (повторно)
	Email       string    // Email
	LastName    string    // Фамилия
	Name        string    // Имя
	MiddleName  string    // Отчество
	IsAccess    bool      // Доступ разрешен
	IsCondition bool      // Условия пользователя
	IsActivated bool      // Активированный Email
	DateOnline  time.Time // Дата последнего посещения
	Date        time.Time // Дата регистрации
	Del         bool      // Запись удалена
	Hash        string    // Контрольная сумма для синхронизации (SHA256)
	Token       string    // Кука активации и идентификации
	Language    string    `db:"-"` // Язык интерфейса
	Groups      []uint64  `db:"-"` // Группы пользователя
}

func init() {
	// Набор сценариев для типа
	types.SetScenario(`Users`, map[string]types.Scenario{
		`root`: types.Scenario{
			Name:        `Пользователи`,
			Description: `Базовая конфигурация всех свойств для всех сценарией указанного типа`,
			Sample:      new(Users),
			Property: []types.Property{{
				Name:     `Id`,
				Title:    `Id`,
				Readonly: `yes`,
				Required: `yes`,
				FormType: `hidden`,
			}, {
				Name:     `Users_Id`,
				Title:    `Пользователь`,
				FormType: `link`,
			}, {
				Name:     `Login`,
				Title:    `Логин пользователя`,
				Required: `yes`,
				FormType: `text`,
			}, {
				Name:     `Password`,
				Title:    `Пароль пользователя (SHA256)`,
				FormType: `text`,
			}, {
				Name:     `Email`,
				Title:    `Email`,
				Required: `yes`,
				FormType: `text`,
			}, {
				Name:     `LastName`,
				Title:    `Фамилия`,
				FormType: `text`,
			}, {
				Name:     `Name`,
				Title:    `Имя`,
				Required: `yes`,
				FormType: `text`,
			}, {
				Name:     `MiddleName`,
				Title:    `Отчество`,
				FormType: `text`,
			}, {
				Name:     `IsAccess`,
				Title:    `Доступ разрешен`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `IsCondition`,
				Title:    `Условия пользователя`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `IsActivated`,
				Title:    `Активированный`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `DateOnline`,
				Title:    `Дата последнего посещения`,
				FormType: `datetime`,
			}, {
				Name:     `Date`,
				Title:    `Дата регистрации`,
				FormType: `datetime`,
			}, {
				Name:     `Del`,
				Title:    `Запись удалена`,
				Required: `yes`,
				FormType: `bool`,
			}, {
				Name:     `Hash`,
				Title:    `Контрольная сумма для синхронизации (SHA256)`,
				FormType: `hidden`,
			}, {
				Name:     `Token`,
				Title:    `Кука активации и идентификации`,
				FormType: `text`,
			}, {
				Name:     `Language`,
				Title:    `Используемый язык на сайте (в интерфейсе)`,
				FormType: `text`,
			}, {
				Name:     `Groups`,
				Title:    `Группы пользователя`,
				FormType: `linkcross`,
				Uri:      `groups`,
			}},
		},
		`All`: types.Scenario{
			Description: `Все свойства`,
			Property: []types.Property{{
				Name: `Users_Id`,
			}, {
				Name: `Login`,
			}, {
				Name: `Password`,
			}, {
				Name: `Email`,
			}, {
				Name: `LastName`,
			}, {
				Name: `Name`,
			}, {
				Name: `MiddleName`,
			}, {
				Name: `IsAccess`,
			}, {
				Name: `IsCondition`,
			}, {
				Name: `IsActivated`,
			}, {
				Name: `DateOnline`,
			}, {
				Name: `Date`,
			}, {
				Name: `Del`,
			}, {
				Name: `Hash`,
			}, {
				Name: `Token`,
			}, {
				Name: `Language`,
			}, {
				Name: `Groups`,
			}},
		},
		`Registration`: types.Scenario{
			Description: `Регистрация`,
			Property: []types.Property{{
				Name: `Users_Id`,
			}, {
				Name: `Login`,
			}, {
				Name: `Password`,
			}, {
				Name: `Email`,
			}, {
				Name: `LastName`,
			}, {
				Name: `Name`,
			}, {
				Name: `MiddleName`,
			}, {
				Name: `IsAccess`,
			}, {
				Name: `IsCondition`,
			}, {
				Name: `Date`,
			}},
		},
		`Recovery`: types.Scenario{
			Description: `Восстановление`,
			Property: []types.Property{{
				Name: `Password`,
			}},
		},
		`Profile`: types.Scenario{
			Description: `Профиль`,
			Property: []types.Property{{
				Name: `LastName`,
			}, {
				Name: `Name`,
			}, {
				Name: `MiddleName`,
			}, {
				Name: `Password`,
			}, {
				Name: `PasswordR`,
			}},
		},
		`Online`: types.Scenario{
			Description: `Статус онлайн`,
			Property: []types.Property{{
				Name: `DateOnline`,
			}, {
				Name: `Token`,
			}},
		},
		`GridAdmin`: types.Scenario{
			Description: `Список пользователей в админке`,
			Property: []types.Property{{
				Name: `Id`,
			}, {
				Name: `Email`,
			}, {
				Name: `LastName`,
			}, {
				Name: `Name`,
			}, {
				Name: `MiddleName`,
			}, {
				Name: `Date`,
			}},
		},
	})
}
