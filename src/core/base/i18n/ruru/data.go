package ruru

import (
	"core/i18n"
)

func init() {
	if _, ok := i18n.Data[`ru-ru`]; ok == false {
		i18n.Data[`ru-ru`] = make(map[string]string)
	}
	for key, message := range Data {
		i18n.Data[`ru-ru`][key] = message
	}
}

var Data = map[string]string{
	`hello`: `Привет Мир!`,
}
