[Главная](/) ::
[Начало](/docs/start.html) ::
[Справочник](/docs/reference.html) ::
[Учебник](/docs/tutorial.html) ::
[Системное api](/docs/api.html) ::
[Скачать](https://github.com/kshamiev/sungora)

## Контроллеры (Controllers)

##### Опции списка контроллеров
***
`OPTIONS : http://localhost/api/v1.0/admin/controllers/{token}/{type}`

> Controller: `base.Controllers.ApiGrid(apiGridOptions)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `type`:`string` - Тип опций
		- `scenario` - сценарий для данной страницы

* **Request**

    Headers

        Content-Type: application/json; charset=utf-8
    Body

* **Response** 200

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 0,
			"errorMessage": "",
			"content":
			{
				"Name": "Роутинг",
				"Description": "Список контроллеров в админке",
				"Property":
				[
					{
						"Name": "Name",
						"Title": "Человеческое название контроллера",
						"AliasDb": "`Name`",
						"Required": "",
						"Readonly": "",
						"Default": "",
						"FormType": "text",
						"FormMask": "",
						"Visible": "",
						"EnumSet": null,
						"Hint": "",
						"Help": "",
						"Placeholder": "",
						"Tab": 0,
						"Column": 0,
						"Uri": ""
					},
					...
				],
				"Sample": null
			}
		}

* **Response** 403

	Headers

		Content-Type: text/html; charset=utf-8
	Body

* **Response** 404

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100570,
			"errorMessage": "Сценарй [%s] не найден для [%s]",
		}

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
		 "errorCode": 100202,
		 "errorMessage": "Тип получаемых опций не указан",
		}

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
		 "errorCode": 100203,
		 "errorMessage": "Тип получаемых опций не реализован [%s]",
		}

##### Список контроллеров
***
`GET : http://localhost/api/v1.0/admin/controllers/{token}/[page]`

> Controller: `base.Controllers.ApiGrid(apiGridGet)`

* **Parameters**
	* `token`:`string` - токен идентификации пользователя
	* `page`:`int` - номер страницы

* **Request**

    Headers

        Content-Type: application/json; charset=utf-8
    Body

* **Response** 200

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 0,
			"errorMessage": "",
			"content":
			[
				{
					"Id": 15,
					"Name": "Капча",
					"Path": "zero/Session/ApiCaptcha",
					"Date": "2014-11-25T11:46:05+04:00",
				},
				...
			]
		}

* **Response** 403

	Headers

		Content-Type: text/html; charset=utf-8
	Body

* **Response** 500

	Headers

		Content-Type: application/json; charset=utf-8
	Body

        {
           "errorCode": 500,
           "errorMessage": "Ошибка формирование данных на сервере",
        }

##### Сортировка контроллеров
***
`PUT : http://localhost/api/v1.0/admin/controllers/{token}/position`

> Controller: `base.Controllers.ApiGrid(apiGridPut)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

		{
			"Id":23,			// id сортируемого элемента
			"TargetId":3	// id элемента после которого встанет сортируемый (0 – если в начало)
		}

* **Response** 200

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
		   "errorCode": 0,
		   "errorMessage": "",
		}

* **Response** 403

	Headers

		Content-Type: text/html; charset=utf-8
	Body

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100510,
			"errorMessage": "Ошибка получение входных данных запроса"
		}


* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100580,
			"errorMessage": "Ошибка сортировки [%s]"
		}

##### Добавление контроллера
***
`POST : http://localhost/api/v1.0/admin/controllers/{token}`

> Controller: `base.Controllers.ApiGrid(apiGridPost)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

		{
			"Name": "Такой-то контроллер",
			"Path": "zero/Uri/ApiObj",
			"IsBefore": false,
			"IsInternal": false,
			"IsDefault": false,
			"IsHidden": false,
			"Position": 12,
			"Date": "2014-11-25T14:55:19.1238584+04:00",
			"Domain": "",
			"Content": "",
			"ContentTime": "0001-01-01T00:00:00Z",
			"Groups": null
		}

* **Response** 200

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 0,
			"errorMessage": "",
			"content": 27
		}

* **Response** 403

	Headers

		Content-Type: text/html; charset=utf-8
	Body

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100530,
			"errorMessage": "Ошибка проверки по сценарию [Controllers.All]"
         	"content":
				...
		}

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100560,
			"errorMessage": "Ошибка добавления [%s]"
		}

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100510,
			"errorMessage": "Ошибка получение входных данных запроса"
		}

##### Опции контроллера подробно
***
`OPTIONS : http://localhost/api/v1.0/admin/controllers/obj/{token}/{id}`

> Controller: `base.Controllers.ApiObj(apiObjOptions)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (контроллера)

* **Request**

    Headers

        Content-Type: application/json; charset=utf-8
    Body

* **Response** 200

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
		   "errorCode": 0,
		   "errorMessage": "",
		   "content":
		   {
			   "Name": "Контроллеры",
			   "Description": "Все свойства",
			   "Property":
			   [
					{
		               "Name": "Name",
		               "Title": "Человеческое название контроллера",
		               "AliasDb": "`Name`",
		               "Required": "",
		               "Readonly": "",
		               "Default": "",
		               "FormType": "text",
		               "FormMask": "",
		               "Visible": "",
		               "EnumSet": null,
		               "Hint": "",
		               "Help": "",
		               "Placeholder": "",
		               "Tab": 0,
		               "Column": 0,
		               "Uri": ""
					},
					...
			   ],
			   "Sample": null
		   }
		}

* **Response** 403

	Headers

		Content-Type: text/html; charset=utf-8
	Body

* **Response** 404

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100570,
			"errorMessage": "Сценарй [%s] не найден для [%s]",
		}

##### Получение контроллера
***
`GET : http://localhost/api/v1.0/admin/controllers/obj/{token}/{id}`

> Controller: `base.Controllers.ApiObj(apiObjGet)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (контроллера)

* **Request**

	Headers

		Content-Type: application/json; charset=utf-8
	Body

* **Response** 200

	Headers

		Content-Type: application/json; charset=utf-8
	Body

        {
            "errorCode": 0,
            "errorMessage": "",
            "content":
            {
				"Name": "Управление разделами (роутинг) (редактирование)",
				"Path": "zero/Uri/ApiObj",
				"IsBefore": false,
				"IsInternal": false,
				"IsDefault": false,
				"IsHidden": false,
				"Position": 12,
				"Date": "2014-11-25T14:55:19.1238584+04:00",
				"Domain": "",
				"Content": "",
				"ContentTime": "0001-01-01T00:00:00Z",
				"Groups": null
			}
        }

* **Response** 403

	Headers

		Content-Type: text/html; charset=utf-8
	Body

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100500,
			"errorMessage": "Ошибка формирование данных на сервере"
		}

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100510,
			"errorMessage": "Ошибка получение входных данных запроса"
		}

##### Изменение контроллера
***
`PUT : http://localhost/api/v1.0/admin/controllers/obj/{token}/{id}`

> Controller: `base.Controllers.ApiObj(apiObjPut)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (контроллера)

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

		{
			"Name": "Такой-то контроллер",
			"Path": "zero/Uri/ApiObj",
			"IsBefore": false,
			"IsInternal": false,
			"IsDefault": false,
			"IsHidden": false,
			"Position": 12,
			"Date": "2014-11-25T14:55:19.1238584+04:00",
			"Domain": "",
			"Content": "",
			"ContentTime": "0001-01-01T00:00:00Z",
			"Groups": null
		}

* **Response** 200

	Headers

		Content-Type: application/json; charset=utf-8
	Body

        {
            "errorCode": 0,
            "errorMessage": "",
        }

* **Response** 403

	Headers

		Content-Type: text/html; charset=utf-8
	Body

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100500,
			"errorMessage": "Ошибка формирование данных на сервере"
		}

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100510,
			"errorMessage": "Ошибка получение входных данных запроса"
		}

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100530,
			"errorMessage": "Ошибка проверки по сценарию [Controllers.All]"
         	"content":
				...
		}

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100540,
			"errorMessage": "Ошибка сохранения [%s] [%d]"
		}

##### Удаление контроллера
***
`DELETE : http://localhost/api/v1.0/admin/controllers/obj/{token}/{id}`

> Controller: `base.Controllers.ApiObj(apiObjDelete)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (контроллера)

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

* **Response** 200

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 0,
			"errorMessage": "",
		}

* **Response** 403

	Headers

		Content-Type: text/html; charset=utf-8
	Body

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100510,
			"errorMessage": "Ошибка получение входных данных запроса"
		}

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100550,
			"errorMessage": "Ошибка удаления [%s] [%d]"
		}

##### Опции групп контроллера
***
`OPTIONS : http://localhost/api/v1.0/admin/controllers/obj/{token}/{id}/groups`

> Controller: `base.Controllers.ApiObjGroups(apiObjGroupsOptions)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (контроллера)

* **Request**

    Headers

        Content-Type: application/json; charset=utf-8
    Body

* **Response** 200

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
		    "errorCode": 0,
		    "errorMessage": "",
		    "content":
		    {
			    "Name": "Права",
			    "Description": "Все свойства",
			    "Property":
			    [
					{
						"Name": "Groups_Id",
						"Title": "Группа",
						"AliasDb": "`Groups_Id`",
						"Required": "",
						"Readonly": "",
						"Default": "",
						"FormType": "link",
						"FormMask": "",
						"Visible": "",
						"EnumSet": null,
						"Hint": "",
						"Help": "",
						"Placeholder": "",
						"Tab": 0,
						"Column": 0,
						"Uri": ""
					},
					...
			    ],
			    "Sample": null
		    }
		}

* **Response** 403

	Headers

		Content-Type: text/html; charset=utf-8
	Body

* **Response** 404

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100570,
			"errorMessage": "Сценарй [%s] не найден для [%s]",
		}

##### Группы контроллера (получение)
***
`GET : http://localhost/api/v1.0/admin/controllers/obj/{token}/{id}/groups`

> Controller: `base.Controllers.ApiObjGroups(apiObjGroupsGet)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (контроллера)

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

* **Response** 200

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 0,
			"errorMessage": "",
			"content":
			{
				"tied":
				[
					{
						"Groups_Id": 1,
						"Controllers_Id": 999,
						"Disable": false,
						"Name": "dev"
					},
					...
				],
				"unbound":
				[
					{
						"Groups_Id": 2,
						"Controllers_Id": 0,
						"Disable": false,
						"Name": "guest"
					},
					...
				]
			}
		}

* **Response** 403

	Headers

		Content-Type: text/html; charset=utf-8
	Body

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100510,
			"errorMessage": "Ошибка получение входных данных запроса"
		}

##### Группы контроллера (изменение)
***
`PUT : http://localhost/api/v1.0/admin/controllers/obj/{token}/{id}/groups`

> Controller: `base.Controllers.ApiObjGroups(apiObjGroupsPut)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (контроллера)

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

		[
			{
				"Groups_Id": 2,
				"Controllers_Id": 999,
				"Disable": false,
				"Name": "guest"
			},
			...
		],

* **Response** 200

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 0,
			"errorMessage": "",
		}

* **Response** 403

	Headers

		Content-Type: text/html; charset=utf-8
	Body

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100510,
			"errorMessage": "Ошибка получение входных данных запроса"
		}

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100550,
			"errorMessage": "Ошибка удаления [%s] [%d]"
		}

* **Response** 409

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 100540,
			"errorMessage": "Ошибка сохранения [%s] [%d]"
		}

##### Список проблемных контроллеров
***
`GET : http://localhost/api/v1.0/admin/controllers/problems/{token}`
Неверные пути, несуществующие контроллеры или их методы
> Controller: `base.Controllers.ApiProblem`

* **Parameters**
	* `token`:`string` - токен идентификации пользователя

* **Request**

    Headers

        Content-Type: application/json; charset=utf-8
    Body

* **Response** 200

	Headers

		Content-Type: application/json; charset=utf-8
	Body

		{
			"errorCode": 0,
			"errorMessage": "",
			"content":
			[
				{
					"Id": 15,
					"Name": "ConytollerName",
					"Path": "moduleName/ControllerName/MethosName",
					"Date": "2010-05-25T11:46:05+04:00",
				},
				...
			]
		}

* **Response** 403

	Headers

		Content-Type: text/html; charset=utf-8
	Body
