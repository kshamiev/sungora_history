// Библиотека: Интернационализация.
//
// Реализация интерактивного перевода текстовых данных на нужный язык.
package i18n

import (
	"fmt"
)

// Коды модулей. Для формирования ункиального глобального кода сообщения
// Заполняется в момент инициализации модулей при старте приложения
var ModuleCode = map[string]int{
	`base`: 100,
}

// Сообщения всех уровней на разных языках
// Заполняется в момент инициализации модулей при старте приложения
var Messages = map[string]map[string]map[int]string{
	`base`: {
		`ru-ru`: {
			1000: `Тестовое сообщение с параметром: [%s]`,
		},
		`en-en`: {
			1000: `Test message with a parameter: [%s]`,
		},
	},
}

// Перевод сообщений
// + moduleName string имя модуля
// + lang string язык
// + codeLocal int локальный код сообщения в рамках модуля
// + params ...interface{} параметры вставляемые в переводимое сообщение
// - int глобальный уникальный код сообщения для вывода в логи или клиенту
// - string сформированное сообщение на нужном языке
func Message(moduleName, lang string, codeLocal int, params ...interface{}) (code int, message string) {
	var ok bool
	code = codeLocal
	if message, ok = Messages[moduleName][lang][codeLocal]; ok == true {
		code = ModuleCode[moduleName]*10000 + codeLocal
		message = fmt.Sprintf(message, params...)

	} else if 0 < len(params) {
		if s, ok := params[0].(string); ok == true {
			message = fmt.Sprintf(s, params[1:]...)
		}
	}
	return
}

// Текстовые данные на разных языках под текстовыми ключами
var Data = map[string]map[string]string{
	`ru-ru`: {
		`key`: `value`,
	},
	`en-en`: {
		`key`: `value`,
	},
}

// Перевод по ключевому слову
// + lang string язык
// + key string текстовой ключ
// + params ...interface{} параметры вставляемые в перевод
// - string сформированный перевод на нужном языке
func Translation(lang, key string, params ...interface{}) (message string) {
	var ok bool
	if message, ok = Data[lang][key]; ok == true {
		message = fmt.Sprintf(message, params...)
	} else if 0 < len(params) {
		if s, ok := params[0].(string); ok == true {
			message = fmt.Sprintf(s, params[1:]...)
		}
	}
	return
}

//
