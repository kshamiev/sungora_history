package enen

import (
	"core/i18n"
)

func init() {
	if _, ok := i18n.Data[`en-en`]; ok == false {
		i18n.Data[`en-en`] = make(map[string]string)
	}
	for key, message := range Data {
		i18n.Data[`en-en`][key] = message
	}
}

var Data = map[string]string{
	`hello`: `Hello Word!`,
}
