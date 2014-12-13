package setup

import (
	"core/base/config"
	_ "core/base/i18n/enen"
	_ "core/base/i18n/ruru"
	moduleServer "core/base/server"
	"core/i18n"
	"core/server"
)

func init() {
	// Контроллеры
	initControllers()

	// Роутинг
	initUri()

	// Интернационализация. Код модуля
	i18n.ModuleCode[config.MODULE_NAME] = config.MODULE_CODE

	// Переопределения ядра
	server.FactorNewUri = moduleServer.FactorNewUri
	server.FactorNewUsers = moduleServer.FactorNewUsers
	server.FactorNewSession = moduleServer.FactorNewSession
	server.FactorAccess = moduleServer.FactorAccess
}
