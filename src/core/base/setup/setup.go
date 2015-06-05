package setup

import (
	"core/base/config"
	_ "core/base/i18n/enen"
	_ "core/base/i18n/ruru"
	moduleServer "core/base/server"
	"core/server"
	"lib/i18n"
)

func init() {
	// Контроллеры
	initControllers()

	// Роутинг
	initUri()

	// Интернационализация
	i18n.ModuleCode[config.MODULE_NAME] = config.MODULE_CODE

	// Переопределения ядра
	server.FactorNewUri = moduleServer.FactorNewUri
	server.FactorNewUsers = moduleServer.FactorNewUsers
	server.FactorNewSession = moduleServer.FactorNewSession
	server.FactorAccess = moduleServer.FactorAccess
}
