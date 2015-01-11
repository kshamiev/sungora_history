// Интернационализация.
//
// Реализация интерактивного перевода текстовых данных на нужный язык.
package i18n

import (
	"fmt"
)

// Коды модулей. Для формирования ункиального глобального кода сообщения
// Заполняется в момент инициализации модулей
var ModuleCode = make(map[string]int)

// Сообщения всех уровней на разных языках
var Messages = make(map[string]map[int]string)

// Перевод сообщений
//    + moduleName string имя модуля
//    + lang string язык
//    + codeLocal int локальный код сообщения в рамках модуля
//    + params ...interface{} параметры вставляемые в переводимое сообщение
//    - int глобальный неизменяемый код сообщения для вывода в логи или клиенту
//    - string свормированное сообщение на нужном языке
func Message(moduleName, lang string, codeLocal int, params ...interface{}) (code int, message string) {
	var ok bool
	code = ModuleCode[moduleName]*1000 + codeLocal
	if message, ok = Messages[lang][code]; ok == true {
		message = fmt.Sprintf(message, params...)
	} else if 0 < len(params) {
		if s, ok := params[0].(string); ok == true {
			message = fmt.Sprintf(s, params[1:]...)
		}
	}
	return
}

// Текстовые данные на разных языках под текстовыми ключами
var Data = make(map[string]map[string]string)

// Перевод по ключевому слову
//    + lang string язык
//    + key string текстовой ключ
//    + params ...interface{} параметры вставляемые в перевод
//    - string свормированный перевод на нужном языке
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
