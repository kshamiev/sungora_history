package setup

import (
	"core/controller"
	typDb "types/db"
)

// initUri Регистрация и инициализация роутинга
func initUri() {
	// Типы
	controller.ConfigUri[`/api/v1.0/types/scenario/{typ}/[scenario]`] = typDb.Uri{
		Method:      []string{`OPTIONS`},
		Name:        `Получение сценариев для всех типов`,
		ContentType: `application/json`,
		Controllers: []string{`base/Types/ApiScenario`},
	}

	// Сервер
	controller.ConfigUri[`/api/v1.0/server/ping`] = typDb.Uri{
		Method:      []string{`GET`},
		Name:        `Проверка доступности сервера`,
		ContentType: `application/json`,
		Controllers: []string{`base/Server/ApiPing`},
	}
	controller.ConfigUri[`/api/v1.0/server/upload/{token}`] = typDb.Uri{
		Method:      []string{`POST`},
		Name:        `Загрузка бинарных данных`,
		ContentType: `application/json`,
		Controllers: []string{`base/Server/ApiUpload`},
	}

	// Сессия
	controller.ConfigUri[`/api/v1.0/session/recovery`] = typDb.Uri{
		Method:      []string{`PUT`},
		Name:        `Восстановление пароля пользователя`,
		ContentType: `application/json`,
		Controllers: []string{`base/Session/ApiRecovery`},
	}
	controller.ConfigUri[`/api/v1.0/session/[token]`] = typDb.Uri{
		Method:      []string{`GET`},
		Name:        `Получение текущего пользователя`,
		ContentType: `application/json`,
		Controllers: []string{`base/Session/ApiUserCurrent`},
	}
	controller.ConfigUri[`/api/v1.0/session/authorization/[token]`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `DELETE`},
		Name:        `Авторизация, выход, проверка токена с его пролонгацией`,
		ContentType: `application/json`,
		Controllers: []string{`base/Session/ApiMain`},
	}
	controller.ConfigUri[`/api/v1.0/session/registration/[token]`] = typDb.Uri{
		Method:      []string{`POST`},
		Name:        `Регистрация нового пользователя`,
		ContentType: `application/json`,
		Controllers: []string{`base/Session/ApiRegistration`},
	}
	controller.ConfigUri[`/api/v1.0/session/captcha/native`] = typDb.Uri{
		Method:      []string{`GET`},
		Name:        `Получение капчи`,
		ContentType: `application/json`,
		Controllers: []string{`base/Session/ApiCaptcha`},
	}

	// Uri
	controller.ConfigUri[`/api/v1.0/admin/uri/{token}/[param]`] = typDb.Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `OPTIONS`},
		Name:        `Управление разделами. Список`,
		ContentType: `application/json`,
		Controllers: []string{`base/Uri/ApiGrid`},
	}
	controller.ConfigUri[`/api/v1.0/admin/uri/obj/{token}/{id}`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `DELETE`, `OPTIONS`},
		Name:        `Управление разделами. Детально`,
		ContentType: `application/json`,
		Controllers: []string{`base/Uri/ApiObj`},
	}
	controller.ConfigUri[`/api/v1.0/admin/uri/obj/{token}/{id}/groups`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `OPTIONS`},
		Name:        `Управление группами раздела`,
		ContentType: `application/json`,
		Controllers: []string{`base/Uri/ApiObjGroups`},
	}
	controller.ConfigUri[`/api/v1.0/admin/uri/obj/{token}/{id}/controllers`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `OPTIONS`},
		Name:        `Управление контроллерами раздела`,
		ContentType: `application/json`,
		Controllers: []string{`base/Uri/ApiObjControllers`},
	}

	// Controllers
	controller.ConfigUri[`/api/v1.0/admin/controllers/{token}/[param]`] = typDb.Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `OPTIONS`},
		Name:        `Управление контроллерами. Список`,
		ContentType: `application/json`,
		Controllers: []string{`base/Controllers/ApiGrid`},
	}
	controller.ConfigUri[`/api/v1.0/admin/controllers/obj/{token}/{id}`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `DELETE`, `OPTIONS`},
		Name:        `Управление контроллерами. Детально`,
		ContentType: `application/json`,
		Controllers: []string{`base/Controllers/ApiObj`},
	}
	controller.ConfigUri[`/api/v1.0/admin/controllers/obj/{token}/{id}/groups`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `OPTIONS`},
		Name:        `Управление группами контроллера`,
		ContentType: `application/json`,
		Controllers: []string{`base/Controllers/ApiObjGroups`},
	}
	controller.ConfigUri[`/api/v1.0/admin/controllers/problems/{token}`] = typDb.Uri{
		Method:      []string{`GET`},
		Name:        `Список проблемных контроллеров`,
		ContentType: `application/json`,
		Controllers: []string{`base/Controllers/ApiProblem`},
	}

	// Users
	controller.ConfigUri[`/api/v1.0/admin/users/{token}/[param]`] = typDb.Uri{
		Method:      []string{`GET`, `POST`, `OPTIONS`},
		Name:        `Управление пользователями. Список`,
		ContentType: `application/json`,
		Controllers: []string{`base/Users/ApiGrid`},
	}
	controller.ConfigUri[`/api/v1.0/admin/users/obj/{token}/{id}`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `DELETE`, `OPTIONS`},
		Name:        `Управление пользователями. Подробно`,
		ContentType: `application/json`,
		Controllers: []string{`base/Users/ApiObj`},
	}
	controller.ConfigUri[`/api/v1.0/admin/users/obj/{token}/{id}/groups`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `OPTIONS`},
		Name:        `Управление группами пользователя`,
		ContentType: `application/json`,
		Controllers: []string{`base/Users/ApiObjGroups`},
	}

	// Groups
	controller.ConfigUri[`/api/v1.0/admin/groups/{token}/[param]`] = typDb.Uri{
		Method:      []string{`GET`, `POST`, `OPTIONS`},
		Name:        `Управление группами. Список`,
		ContentType: `application/json`,
		Controllers: []string{`base/Groups/ApiGrid`},
	}
	controller.ConfigUri[`/api/v1.0/admin/groups/obj/{token}/{id}`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `DELETE`, `OPTIONS`},
		Name:        `Управление группами. Подробно`,
		ContentType: `application/json`,
		Controllers: []string{`base/Groups/ApiObj`},
	}

	controller.ConfigUri[`/`] = typDb.Uri{
		Method:      []string{`GET`},
		Name:        `Главная страница`,
		Layout:      `/default.html`,
		Controllers: []string{`base/Page/Content`},
	}
	controller.ConfigUri[`/project`] = typDb.Uri{
		Method:      []string{`GET`},
		Name:        `Информационная страница о проекте`,
		Layout:      `/default.html`,
		Controllers: []string{`base/Page/Content`},
	}
}
