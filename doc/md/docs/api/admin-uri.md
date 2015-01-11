[Главная](/) ::
[Начало](/docs/start.html) ::
[Справочник](/docs/reference.html) ::
[Учебник](/docs/tutorial.html) ::
[Системное api](/docs/api.html) ::
[Скачать](https://github.com/kshamiev/sungora)

## Разделы (Uri)

##### Опции списка разделов
***
`OPTIONS : http://localhost/api/v1.0/admin/uri/{token}/{type}`

> Controller: `base.Uri.ApiGrid(apiGridOptions)`

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
				"Description": "Список разделов в админке",
				"Property":
				[
					{
						"Name": "Uri",
						"Title": "URI без указания домена и протокола до необязательных параметров",
						"AliasDb": "`Uri`",
						"Required": "yes",
						"Readonly": "",
						"Default": "/",
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

##### Список разделов
***
`GET : http://localhost/api/v1.0/admin/uri/{token}/[page]`

> Controller: `base.Uri.ApiGrid(apiGridGet)`

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
					"Id": 3,
					"Uri": "/api/v1.0/session/[token]",
					"Name": "Восстановление пароля пользователя",
					"Layout": "",
					"ContentType": "application/json",
					"ContentEncode": "utf-8",
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

##### Сортировка разделов
***
`PUT : http://localhost/api/v1.0/admin/uri/{token}/position`

> Controller: `base.Uri.ApiGrid(apiGridPut)`

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

##### Реинициализация роутинга
***
`PUT : http://localhost/api/v1.0/admin/uri/{token}/route`

> Controller: `base.Uri.ApiGrid(apiGridPut)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя

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

##### Добавление раздела
***
`POST : http://localhost/api/v1.0/admin/uri/{token}`

> Controller: `base.Uri.ApiGrid(apiGridPost)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

		{
			"Method":
			[
				"GET",
				"OPTIONS"
			],
			"Domain": "",
			"Uri": "/page/page",
			"Name": "Раздел",
			"Redirect": "",
			"Layout": "",
			"IsAuthorized": false,
			"IsMenuVisible": false,
			"IsDisable": false,
			"Content": "",
			"ContentTime": "0001-01-01T00:00:00Z",
			"ContentType": "application/json",
			"ContentEncode": "utf-8",
			"Position": 6,
			"Title": "",
			"KeyWords": "",
			"Description": "",
			"Controllers": null,
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
			"errorMessage": "Ошибка проверки по сценарию [Uri.All]"
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

##### Опции раздела подробно
***
`OPTIONS : http://localhost/api/v1.0/admin/uri/obj/{token}/{id}`

> Controller: `base.Uri.ApiObj(apiObjOptions)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (раздела)

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
			    "Name": "Разделы",
			    "Description": "Все свойства",
			    "Property":
			    [
					{
						"Name": "Method",
						"Title": "Метод запроса",
						"AliasDb": "`Method`",
						"Required": "",
						"Readonly": "",
						"Default": "GET",
						"FormType": "checkbox",
						"FormMask": "",
						"Visible": "",
						"EnumSet":
						{
						    "CONNECT": "CONNECT",
						    "DELETE": "DELETE",
						    "GET": "GET",
						    "HEAD": "HEAD",
						    "LINK": "LINK",
						    "OPTIONS": "OPTIONS",
						    "PATCH": "PATCH",
						    "POST": "POST",
						    "PUT": "PUT",
						    "TRACE": "TRACE",
						    "UNLINK": "UNLINK",
						    "WS": "WS"
						},
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

##### Получение раздела
***
`GET : http://localhost/api/v1.0/admin/uri/obj/{token}/{id}`

> Controller: `base.Uri.ApiObj(apiObjGet)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (раздела)

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
				"Method":
				[
					"GET",
					"OPTIONS"
				],
				"Domain": "",
				"Uri": "/api/v1.0/types/scenario/{typ}/[scenario]",
				"Name": "Проверка доступности сервера",
				"Redirect": "",
				"Layout": "",
				"IsAuthorized": false,
				"IsMenuVisible": false,
				"IsDisable": false,
				"Content": "",
				"ContentTime": "0001-01-01T00:00:00Z",
				"ContentType": "application/json",
				"ContentEncode": "utf-8",
				"Position": 6,
				"Title": "",
				"KeyWords": "",
				"Description": "",
				"Controllers": null,
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

##### Изменение раздела
***
`PUT : http://localhost/api/v1.0/admin/uri/obj/{token}/{id}`

> Controller: `base.Uri.ApiObj(apiObjPut)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (раздела)

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

        {
            "Method":
            [
               "GET",
               "OPTIONS"
            ],
            "Domain": "",
            "Uri": "/page/page",
            "Name": "Раздел",
            "Redirect": "",
            "Layout": "",
            "IsAuthorized": false,
            "IsMenuVisible": false,
            "IsDisable": false,
            "Content": "",
            "ContentTime": "0001-01-01T00:00:00Z",
            "ContentType": "application/json",
            "ContentEncode": "utf-8",
            "Position": 6,
            "Title": "",
            "KeyWords": "",
            "Description": "",
            "Controllers": null,
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
			"errorMessage": "Ошибка проверки по сценарию [Uri.All]"
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

##### Удаление раздела
***
`DELETE : http://localhost/api/v1.0/admin/uri/obj/{token}/{id}`

> Controller: `base.Uri.ApiObj(apiObjDelete)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (раздела)

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

##### Опции контроллеров раздела
***
`OPTIONS : http://localhost/api/v1.0/admin/uri/obj/{token}/{id}/controllers`

> Controller: `base.Uri.ApiObjControllers(apiObjControllersOptions)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (раздела)

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
						"Name": "Uri_Id",
						"Title": "Раздел",
						"AliasDb": "`Uri_Id`",
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

##### Контроллеры раздела (получение)
***
`GET : http://localhost/api/v1.0/admin/uri/obj/{token}/{id}/controllers`

> Controller: `base.Uri.ApiObjControllers(apiObjControllersGet)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (раздела)

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
						"Uri_Id": 999,
						"Controllers_Id": 1,
						"Name": "Управление разделами (роутинг)"
					},
					...
				],
				"unbound":
				[
					{
						"Uri_Id": 0,
						"Controllers_Id": 2,
						"Name": "Восстановление пароля пользователя"
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

##### Контроллеры раздела (изменение)
***

`PUT : http://localhost/api/v1.0/admin/uri/obj/{token}/{id}/controllers`

> Controller: `base.Uri.ApiObjControllers(apiObjControllersPut)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (раздела)

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

        [
            {
	            "Uri_Id": 999,
	            "Controllers_Id": 1,
	            "Name": "Управление разделами (роутинг)"
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

##### Опции групп раздела
***
`OPTIONS : http://localhost/api/v1.0/admin/uri/obj/{token}/{id}/groups`

> Controller: `base.Uri.ApiObjGroups(apiObjGroupsOptions)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (раздела)

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

##### Группы раздела (получение)
***
`GET : http://localhost/api/v1.0/admin/uri/obj/{token}/{id}/groups`

> Controller: `base.Uri.ApiObjGroups(apiObjGroupsGet)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (раздела)

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
						"Uri_Id": 999,
						"Get": false,
						"Post": false,
						"Put": false,
						"Delete": false,
						"Options": false,
						"Name": "dev"
					},
					...
				],
				"unbound":
				[
					{
						"Groups_Id": 2,
						"Uri_Id": 0,
						"Get": false,
						"Post": false,
						"Put": false,
						"Delete": false,
						"Options": false,
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

##### Группы раздела (изменение)
***
`PUT : http://localhost/api/v1.0/admin/uri/obj/{token}/{id}/groups`

> Controller: `base.Uri.ApiObjGroups(apiObjGroupsPut)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (раздела)

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

        [
            {
               "Groups_Id": 2,
               "Uri_Id": 999,
               "Get": false,
               "Post": false,
               "Put": false,
               "Delete": false,
               "Options": false,
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

