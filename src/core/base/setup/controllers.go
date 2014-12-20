package setup

import (
	"app"
	moduleController "core/base/controller"
	coreConfig "core/config"
	typDb "types/db"
)

// initControllers Регистрация и инициализация контроллеров
func initControllers() {
	// Контроллер Groups
	app.Controller[`base/Groups`] = moduleController.NewGroups
	// методы
	coreConfig.ConfigControllers[`base/Groups/ApiGrid`] = typDb.Controllers{
		Name: `Управление групами. Список`,
	}
	coreConfig.ConfigControllers[`base/Groups/ApiObj`] = typDb.Controllers{
		Name: `Управление группами. Детально`,
	}

	// Контроллер Users
	app.Controller[`base/Users`] = moduleController.NewUsers
	// методы
	coreConfig.ConfigControllers[`base/Users/ApiGrid`] = typDb.Controllers{
		Name: `Управление пользователями. Список`,
	}
	coreConfig.ConfigControllers[`base/Users/ApiObj`] = typDb.Controllers{
		Name: `Управление пользователями. Детально`,
	}
	coreConfig.ConfigControllers[`base/Users/ApiObjGroups`] = typDb.Controllers{
		Name: `Управление группами пользователя`,
	}

	// Контроллер Controllers
	app.Controller[`base/Controllers`] = moduleController.NewControllers
	// методы
	coreConfig.ConfigControllers[`base/Controllers/ApiGrid`] = typDb.Controllers{
		Name: `Управление контроллерами. Список`,
	}
	coreConfig.ConfigControllers[`base/Controllers/ApiObj`] = typDb.Controllers{
		Name: `Управление контроллерами. Детально`,
	}
	coreConfig.ConfigControllers[`base/Controllers/ApiObjGroups`] = typDb.Controllers{
		Name: `Управление группами контроллера`,
	}
	coreConfig.ConfigControllers[`base/Controllers/ApiProblem`] = typDb.Controllers{
		Name: `Список проблемных контроллеров`,
	}

	// Контроллер Uri
	app.Controller[`base/Uri`] = moduleController.NewUri
	// методы
	coreConfig.ConfigControllers[`base/Uri/ApiGrid`] = typDb.Controllers{
		Name: `Управление разделами (роутинг). Список`,
	}
	coreConfig.ConfigControllers[`base/Uri/ApiObj`] = typDb.Controllers{
		Name: `Управление разделами (роутинг). Детально`,
	}
	coreConfig.ConfigControllers[`base/Uri/ApiObjGroups`] = typDb.Controllers{
		Name: `Управление группами разделов`,
	}
	coreConfig.ConfigControllers[`base/Uri/ApiObjControllers`] = typDb.Controllers{
		Name: `Управление контроллерами разделов`,
	}

	// Контроллер Сессия
	app.Controller[`base/Session`] = moduleController.NewSession
	// методы
	coreConfig.ConfigControllers[`base/Session/ApiRecovery`] = typDb.Controllers{
		Name: `Восстановление пароля пользователя`,
	}
	coreConfig.ConfigControllers[`base/Session/ApiMain`] = typDb.Controllers{
		Name: `Авторизация, выход, проверка токена с его пролонгацией`,
	}
	coreConfig.ConfigControllers[`base/Session/ApiRegistration`] = typDb.Controllers{
		Name: `Регистрация нового пользователя`,
	}
	coreConfig.ConfigControllers[`base/Session/ApiUserCurrent`] = typDb.Controllers{
		Name: `Получение текущего пользователя`,
	}
	coreConfig.ConfigControllers[`base/Session/ApiCaptcha`] = typDb.Controllers{
		Name: `Получение капчи`,
	}

	// Контроллер Сервер
	app.Controller[`base/Server`] = moduleController.NewServer
	// методы
	coreConfig.ConfigControllers[`base/Server/ApiPing`] = typDb.Controllers{
		Name: `Проверка доступности сервера`,
	}
	coreConfig.ConfigControllers[`base/Server/ApiUpload`] = typDb.Controllers{
		Name: `Загрузка бинарных данных`,
	}

	// Контроллер Типы
	app.Controller[`base/Types`] = moduleController.NewTypes
	// методы
	coreConfig.ConfigControllers[`base/Types/ApiScenario`] = typDb.Controllers{
		Name: `Получение сценариев для всех типов`,
	}

	// Контроллер информационная страница
	app.Controller[`base/Page`] = moduleController.NewPage
	// методы
	coreConfig.ConfigControllers[`base/Page/Block`] = typDb.Controllers{
		Name: `Информационный блок`,
	}
	coreConfig.ConfigControllers[`base/Page/Content`] = typDb.Controllers{
		Name: `Информационная страница`,
	}

}
