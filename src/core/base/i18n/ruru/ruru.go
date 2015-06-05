package ruru

import (
	"core/base/config"
	"lib/i18n"
)

func init() {
	// Message
	if _, ok := i18n.Messages[config.MODULE_NAME]; ok == false {
		i18n.Messages[config.MODULE_NAME] = make(map[string]map[int]string)
	}
	if _, ok := i18n.Messages[config.MODULE_NAME][`ru-ru`]; ok == false {
		i18n.Messages[config.MODULE_NAME][`ru-ru`] = make(map[int]string)
	}
	for code, message := range Messages {
		i18n.Messages[config.MODULE_NAME][`ru-ru`][code] = message
	}
	// Translation
	if _, ok := i18n.Data[`ru-ru`]; ok == false {
		i18n.Data[`ru-ru`] = make(map[string]string)
	}
	for key, message := range Data {
		i18n.Data[`ru-ru`][key] = message
	}
}
