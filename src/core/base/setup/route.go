package setup

import (
	coreConfig "core/config"
	typDb "types/db"
)

// initUri Регистрация и инициализация роутинга
func initUri() {
	// Типы
	coreConfig.ConfigUri[`/api/v1.0/types/scenario/{typ}/[scenario]`] = typDb.Uri{
		Method:      []string{`OPTIONS`},
		Name:        `Получение сценариев для всех типов`,
		ContentType: `application/json`,
		Controllers: []string{`base/Types/ApiScenario`},
	}

	// Сервер
	coreConfig.ConfigUri[`/api/v1.0/server/ping`] = typDb.Uri{
		Method:      []string{`GET`},
		Name:        `Проверка доступности сервера`,
		ContentType: `application/json`,
		Controllers: []string{`base/Server/ApiPing`},
	}
	coreConfig.ConfigUri[`/api/v1.0/server/upload/{token}`] = typDb.Uri{
		Method:      []string{`POST`},
		Name:        `Загрузка бинарных данных`,
		ContentType: `application/json`,
		Controllers: []string{`base/Server/ApiUpload`},
	}

	// Сессия
	coreConfig.ConfigUri[`/api/v1.0/session/recovery`] = typDb.Uri{
		Method:      []string{`PUT`},
		Name:        `Восстановление пароля пользователя`,
		ContentType: `application/json`,
		Controllers: []string{`base/Session/ApiRecovery`},
	}
	coreConfig.ConfigUri[`/api/v1.0/session/[token]`] = typDb.Uri{
		Method:      []string{`GET`},
		Name:        `Получение текущего пользователя`,
		ContentType: `application/json`,
		Controllers: []string{`base/Session/ApiUserCurrent`},
	}
	coreConfig.ConfigUri[`/api/v1.0/session/authorization/[token]`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `DELETE`},
		Name:        `Авторизация, выход, проверка токена с его пролонгацией`,
		ContentType: `application/json`,
		Controllers: []string{`base/Session/ApiMain`},
	}
	coreConfig.ConfigUri[`/api/v1.0/session/registration/[token]`] = typDb.Uri{
		Method:      []string{`POST`},
		Name:        `Регистрация нового пользователя`,
		ContentType: `application/json`,
		Controllers: []string{`base/Session/ApiRegistration`},
	}
	coreConfig.ConfigUri[`/api/v1.0/session/captcha/native`] = typDb.Uri{
		Method:      []string{`GET`},
		Name:        `Получение капчи`,
		ContentType: `application/json`,
		Controllers: []string{`base/Session/ApiCaptcha`},
	}

	// Uri
	coreConfig.ConfigUri[`/api/v1.0/admin/uri/{token}/[param]`] = typDb.Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `OPTIONS`},
		Name:        `Управление разделами. Список`,
		ContentType: `application/json`,
		Controllers: []string{`base/Uri/ApiGrid`},
	}
	coreConfig.ConfigUri[`/api/v1.0/admin/uri/obj/{token}/{id}`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `DELETE`, `OPTIONS`},
		Name:        `Управление разделами. Детально`,
		ContentType: `application/json`,
		Controllers: []string{`base/Uri/ApiObj`},
	}
	coreConfig.ConfigUri[`/api/v1.0/admin/uri/obj/{token}/{id}/groups`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `OPTIONS`},
		Name:        `Управление группами раздела`,
		ContentType: `application/json`,
		Controllers: []string{`base/Uri/ApiObjGroups`},
	}
	coreConfig.ConfigUri[`/api/v1.0/admin/uri/obj/{token}/{id}/controllers`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `OPTIONS`},
		Name:        `Управление контроллерами раздела`,
		ContentType: `application/json`,
		Controllers: []string{`base/Uri/ApiObjControllers`},
	}

	// Controllers
	coreConfig.ConfigUri[`/api/v1.0/admin/controllers/{token}/[param]`] = typDb.Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `OPTIONS`},
		Name:        `Управление контроллерами. Список`,
		ContentType: `application/json`,
		Controllers: []string{`base/Controllers/ApiGrid`},
	}
	coreConfig.ConfigUri[`/api/v1.0/admin/controllers/obj/{token}/{id}`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `DELETE`, `OPTIONS`},
		Name:        `Управление контроллерами. Детально`,
		ContentType: `application/json`,
		Controllers: []string{`base/Controllers/ApiObj`},
	}
	coreConfig.ConfigUri[`/api/v1.0/admin/controllers/obj/{token}/{id}/groups`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `OPTIONS`},
		Name:        `Управление группами контроллера`,
		ContentType: `application/json`,
		Controllers: []string{`base/Controllers/ApiObjGroups`},
	}
	coreConfig.ConfigUri[`/api/v1.0/admin/controllers/problems/{token}`] = typDb.Uri{
		Method:      []string{`GET`},
		Name:        `Список проблемных контроллеров`,
		ContentType: `application/json`,
		Controllers: []string{`base/Controllers/ApiProblem`},
	}

	// Users
	coreConfig.ConfigUri[`/api/v1.0/admin/users/{token}/[param]`] = typDb.Uri{
		Method:      []string{`GET`, `POST`, `OPTIONS`},
		Name:        `Управление пользователями. Список`,
		ContentType: `application/json`,
		Controllers: []string{`base/Users/ApiGrid`},
	}
	coreConfig.ConfigUri[`/api/v1.0/admin/users/obj/{token}/{id}`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `DELETE`, `OPTIONS`},
		Name:        `Управление пользователями. Подробно`,
		ContentType: `application/json`,
		Controllers: []string{`base/Users/ApiObj`},
	}
	coreConfig.ConfigUri[`/api/v1.0/admin/users/obj/{token}/{id}/groups`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `OPTIONS`},
		Name:        `Управление группами пользователя`,
		ContentType: `application/json`,
		Controllers: []string{`base/Users/ApiObjGroups`},
	}

	// Groups
	coreConfig.ConfigUri[`/api/v1.0/admin/groups/{token}/[param]`] = typDb.Uri{
		Method:      []string{`GET`, `POST`, `OPTIONS`},
		Name:        `Управление группами. Список`,
		ContentType: `application/json`,
		Controllers: []string{`base/Groups/ApiGrid`},
	}
	coreConfig.ConfigUri[`/api/v1.0/admin/groups/obj/{token}/{id}`] = typDb.Uri{
		Method:      []string{`GET`, `PUT`, `DELETE`, `OPTIONS`},
		Name:        `Управление группами. Подробно`,
		ContentType: `application/json`,
		Controllers: []string{`base/Groups/ApiObj`},
	}

	coreConfig.ConfigUri[`/`] = typDb.Uri{
		Method:      []string{`GET`},
		Name:        `Главная страница`,
		Layout:      `/default.html`,
		Controllers: []string{`base/Page/Content`},
	}
	coreConfig.ConfigUri[`/project`] = typDb.Uri{
		Method:      []string{`GET`},
		Name:        `Информационная страница о проекте`,
		Layout:      `/default.html`,
		Controllers: []string{`base/Page/Content`},
	}
}
