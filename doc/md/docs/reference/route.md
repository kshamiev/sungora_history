[main](/) ::
[quick start](/docs/start.html) ::
[reference](/docs/reference.html) ::
[tutorial](/docs/tutorial.html) ::
[api](/docs/api.html) ::
[sample](/sample) ::
[download](/https://github.com/kshamiev/sungora)

# Компонент: Роутинг
***
### Контроллеры
`type Controllers`

Адресация контроллеров задается строковым индетификатором.
Котороый предстваляет из себя три словоформы соединенные знаком '/'. Где:
- ModuleName - имя модуля
- ControllerName - имя контроллера
- MetodName - имя метода

Пример: `zero/Users/ApiObj`

Сами контрллеры создаются под строковыми идентификаторами состоящими из двух словоформ соединенные знаком '/'. Где:
- ModuleName - имя модуля
- ControllerName - имя контроллера

Пример: `zero/Users`

Конфигурация контроллеров задается в модулях их содержащих.
По пути: `src/app/ModuleName/setup/controllers.go`

Пример:

	controller.Controllers[`zero/Users`] = new(zeroController.Users)
	controller.ConfigControllers[`zero/Users/ApiObjGroups`] = typDb.Controllers{
		Name: `Управление контроллерами разделов`,
	}
	controller.ConfigControllers[`zero/Users/ApiGrid`] = typDb.Controllers{
		Name: `Управление разделами (роутинг)`,
	}
	controller.ConfigControllers[`zero/Users/ApiObj`] = typDb.Controllers{
		Name: `Управление разделами (роутинг) (редактирование)`,
	}

### Разделы
`type Uri`

Роутинг задается с помощью свойства Uri типа Uri.
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

	// Groups
	controller.ConfigUri[`/api/v1.0/admin/groups/obj/{token}/[sid]/[child]/[cid]`] = typDb.Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Name:        `Управление разделами (роутинг)`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Groups/ApiObj`},
	}
	controller.ConfigUri[`/api/v1.0/admin/groups/{token}/[self]/[sid]`] = typDb.Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Name:        `Управление разделами (роутинг)`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Groups/ApiGrid`},
	}










