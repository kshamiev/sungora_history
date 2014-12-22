[main](/) ::
[quick start](/docs/start.html) ::
[reference](/docs/reference.html) ::
[tutorial](/docs/tutorial.html) ::
[api](/docs/api.html) ::
[sample](/sample) ::
[download](https://github.com/kshamiev/sungora)

# Компонент: Роутинг
Этот компонент одновременно является и понятием.

#### Реализация:
При старте программы происходит инициализация модулей: `src/modules.go`.
В модулях вызываются функции иницализации разделов, контроллеров: `src/core/base/setup/setup.go`.
Информация обо всех контроллерах и разделах описанных в модулях заносится во временную конфигурацию `src/core/config/app.go`.
Также инициализируется массив функций конструторов контроллеров `src/app.Controller`

Далее в момент загрузки данных приложения с учетом БД, происходит сверка с конфигруацией и занесение ее также как и остальных данных в область приложения `src/app.Data`

После этого происходит инициализация собственно самого роутинга в прямом его смысле `src/app.Routes`.<br>
В процессе работы сервера при получении запроса происходит поиск роутинга именно по нему.<br>
После этого по найденому роутингу (uri) производиться поиск соответсвующего ему раздела в области данных по Id.

***
### Контроллеры
`type Controllers`

Адресация контроллеров задается строковым индетификатором.
Котороый предстваляет из себя три словоформы соединенные знаком '/'. Где:
- ModuleName - имя модуля
- ControllerName - имя контроллера
- MetodName - имя метода

Пример: `base/Users/ApiObj`

Сами контрллеры создаются под строковыми идентификаторами состоящими из двух словоформ соединенные знаком '/'. Где:
- ModuleName - имя модуля
- ControllerName - имя контроллера

Пример: `base/Users`

Конфигурация контроллеров задается в модулях их содержащих.
По пути: `src/app/ModuleName/setup/controllers.go`

Пример:

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
	...

### Разделы
`type Uri`

Роутинг задается с помощью свойства `Uri` типа `Uri`.
В нем указывается относительный путь страницы.
Также в нем могут быть указаны параметры в любой количестве:
- {} - обязательные (могут быть указаны в любой части uri)
- [] - не обязательные (могут быть указаны только в конце uri)

Примеры:

	`/page/page/page/{token}/[id]`
	`/{lang}/page/page/obj/{token}/[id]/[relation]`

К каждому роутингу (Uri) можно привязать опциональное количество контроллеров.
Указав их строковые идетификаторы в виде списка в свойстве `Controllers`

Конфигурация роутинга задается в модулях реализующих их работу.
По пути: `src/app/ModuleName/setup/route.go`

Пример:

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
	...
