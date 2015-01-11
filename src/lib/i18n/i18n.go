package i18n

import (
	"fmt"
)

var ModuleCode = make(map[string]int)

// ЛОГИ
var Messages = make(map[string]map[int]string)

// Error перевод сообщений для лога и формирование глобального кода
func Message(moduleName, lang string, codeLocal int, messages ...interface{}) (code int, message string) {
	var ok bool
	code = ModuleCode[moduleName]*1000 + codeLocal
	if message, ok = Messages[lang][code]; ok == true {
		message = fmt.Sprintf(message, messages...)
	} else if 0 < len(messages) {
		if s, ok := messages[0].(string); ok == true {
			message = fmt.Sprintf(s, messages[1:]...)
		}
	}
	return
}

// ПЕРЕВОДЫ
var Data = make(map[string]map[string]string)

// Translation Перевод по ключевому слову
func Translation(lang, key string, messages ...interface{}) (message string) {
	var ok bool
	if message, ok = Data[lang][key]; ok == true {
		message = fmt.Sprintf(message, messages...)
	} else if 0 < len(messages) {
		if s, ok := messages[0].(string); ok == true {
			message = fmt.Sprintf(s, messages[1:]...)
		}
	}
	return
}

//
