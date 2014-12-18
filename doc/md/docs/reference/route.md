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
При старте программы происходит инициализация модулей: `src/modules.go`  
В модулях вызываются функции иницализации разделов, контроллеров и другие: `src/core/base/setup/setup.go`
Информация обо всех контроллерах и разделах описанных в модулях заносится во временную конфигурацию в базовом контроллере `core/controller/config.go`
Также создаются реальные контроллеры в `core/controller.Controllers`

Далее в момент загрузки данных приложения с учетом БД, происходит сверка с конфигруацией и занесение ее также как и остальных данных в область приложения `src/app.Data`

После этого происходит инициализация собственно самого роутинга в прямом его смысле `src/app.Routes`.

В процессе работы сервера при получении запроса происходит поиск роутинга именно по нему. И только после этого найденый роутинг (uri) ищеться в области данных по Id.

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
	...
