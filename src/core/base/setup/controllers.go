package setup

import (
	moduleController "core/base/controller"
	"core/controller"
	typDb "types/db"
)

// initControllers Регистрация и инициализация контроллеров
func initControllers() {
	// Контроллер Groups
	controller.Controllers[`base/Groups`] = new(moduleController.Groups)
	controller.ConfigControllers[`base/Groups/ApiGrid`] = typDb.Controllers{
		Name: `Управление групами. Список`,
	}
	controller.ConfigControllers[`base/Groups/ApiObj`] = typDb.Controllers{
		Name: `Управление группами. Детально`,
	}

	// Контроллер Users
	controller.Controllers[`base/Users`] = new(moduleController.Users)
	controller.ConfigControllers[`base/Users/ApiGrid`] = typDb.Controllers{
		Name: `Управление пользователями. Список`,
	}
	controller.ConfigControllers[`base/Users/ApiObj`] = typDb.Controllers{
		Name: `Управление пользователями. Детально`,
	}
	controller.ConfigControllers[`base/Users/ApiObjGroups`] = typDb.Controllers{
		Name: `Управление группами пользователя`,
	}

	// Контроллер Controllers
	controller.Controllers[`base/Controllers`] = new(moduleController.Controllers)
	controller.ConfigControllers[`base/Controllers/ApiGrid`] = typDb.Controllers{
		Name: `Управление контроллерами. Список`,
	}
	controller.ConfigControllers[`base/Controllers/ApiObj`] = typDb.Controllers{
		Name: `Управление контроллерами. Детально`,
	}
	controller.ConfigControllers[`base/Controllers/ApiObjGroups`] = typDb.Controllers{
		Name: `Управление группами контроллера`,
	}
	controller.ConfigControllers[`base/Controllers/ApiProblem`] = typDb.Controllers{
		Name: `Список проблемных контроллеров`,
	}

	// Контроллер Uri
	controller.Controllers[`base/Uri`] = new(moduleController.Uri)
	controller.ConfigControllers[`base/Uri/ApiGrid`] = typDb.Controllers{
		Name: `Управление разделами (роутинг). Список`,
	}
	controller.ConfigControllers[`base/Uri/ApiObj`] = typDb.Controllers{
		Name: `Управление разделами (роутинг). Детально`,
	}
	controller.ConfigControllers[`base/Uri/ApiObjGroups`] = typDb.Controllers{
		Name: `Управление группами разделов`,
	}
	controller.ConfigControllers[`base/Uri/ApiObjControllers`] = typDb.Controllers{
		Name: `Управление контроллерами разделов`,
	}

	// Контроллер Сессия
	controller.Controllers[`base/Session`] = new(moduleController.Session)
	controller.ConfigControllers[`base/Session/ApiRecovery`] = typDb.Controllers{
		Name: `Восстановление пароля пользователя`,
	}
	controller.ConfigControllers[`base/Session/ApiMain`] = typDb.Controllers{
		Name: `Авторизация, выход, проверка токена с его пролонгацией`,
	}
	controller.ConfigControllers[`base/Session/ApiRegistration`] = typDb.Controllers{
		Name: `Регистрация нового пользователя`,
	}
	controller.ConfigControllers[`base/Session/ApiUserCurrent`] = typDb.Controllers{
		Name: `Получение текущего пользователя`,
	}
	controller.ConfigControllers[`base/Session/ApiCaptcha`] = typDb.Controllers{
		Name: `Получение капчи`,
	}

	// Контроллер Сервер
	controller.Controllers[`base/Server`] = new(moduleController.Server)
	controller.ConfigControllers[`base/Server/ApiPing`] = typDb.Controllers{
		Name: `Проверка доступности сервера`,
	}
	controller.ConfigControllers[`base/Server/ApiUpload`] = typDb.Controllers{
		Name: `Загрузка бинарных данных`,
	}

	// Контроллер Типы
	controller.Controllers[`base/Types`] = new(moduleController.Types)
	controller.ConfigControllers[`base/Types/ApiScenario`] = typDb.Controllers{
		Name: `Получение сценариев для всех типов`,
	}

	// Контроллер информационная страница
	controller.Controllers[`base/Page`] = new(moduleController.Page)
	// методы
	controller.ConfigControllers[`base/Page/Block`] = typDb.Controllers{
		Name: `Информационный блок`,
	}
	controller.ConfigControllers[`base/Page/Content`] = typDb.Controllers{
		Name: `Информационная страница`,
	}

}
