package i18n

import (
	"fmt"
)

var ModuleCode = make(map[string]int16)

// ЛОГИ
var Messages = make(map[string]map[int]string)

// Error перевод сообщений для лога и формирование глобального кода
func Message(moduleName, lang string, codeLocal int16, messages ...interface{}) (code int, message string) {
	var ok bool
	code = int(ModuleCode[moduleName])*1000 + int(codeLocal)
	if message, ok = Messages[lang][code]; ok == true {
		message = fmt.Sprintf(message, messages...)
	} else if 0 < len(messages) {
		if s, ok := messages[0].(string); ok == true {
			message = fmt.Sprintf(s, messages[1:]...)
		}
		code = int(codeLocal)
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
