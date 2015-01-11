[Главная](/) ::
[Начало](/docs/start.html) ::
[Справочник](/docs/reference.html) ::
[Учебник](/docs/tutorial.html) ::
[Системное api](/docs/api.html) ::
[Скачать](https://github.com/kshamiev/sungora)

## Группы (Groups)

##### Опции списка групп
***
`OPTIONS : http://localhost/api/v1.0/admin/groups/{token}/{type}`

> Controller: `base.Groups.ApiGrid(apiGridOptions)`

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
				"Name": "Группы",
				"Description": "Список групп в админке",
				"Property":
				[
					{
						"Name": "Name",
						"Title": "Наименование",
						"AliasDb": "`Name`",
						"Required": "yes",
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

##### Список групп
***
`GET : http://localhost/api/v1.0/admin/groups/{token}/[page]`

> Controller: `base.Groups.ApiGrid(apiGridGet)`

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
					"Id": 1,
					"Name": "devdevdev",
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

##### Добавление группы
***
`POST : http://localhost/api/v1.0/admin/groups/{token}`

> Controller: `base.Groups.ApiGrid(apiGridPost)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

		{
			"Name": "Название группы",
			"Description": "Описание",
			"IsDefault": false,
			"Del": false,
			"Hash": ""
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

##### Опции группы подробно
***
`OPTIONS : http://localhost/api/v1.0/admin/groups/obj/{token}/{id}`

> Controller: `base.Groups.ApiObj(apiObjOptions)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (группы)

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
			   "Name": "Группы",
			   "Description": "Все свойства",
			   "Property":
			   [
					{
						"Name": "Name",
						"Title": "Наименование",
						"AliasDb": "`Name`",
						"Required": "yes",
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

##### Получение группы
***
`GET : http://localhost/api/v1.0/admin/groups/obj/{token}/{id}`

> Controller: `base.Groups.ApiObj(apiObjGet)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (группы)

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
				"Name": "devdevdev",
				"Description": "",
				"IsDefault": false,
				"Del": false,
				"Hash": ""
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

##### Изменение группы
***
`PUT : http://localhost/api/v1.0/admin/groups/obj/{token}/{id}`

> Controller: `base.Groups.ApiObj(apiObjPut)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (группы)

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

		{
			"Name": "Название группы",
			"Description": "Описание",
			"IsDefault": false,
			"Del": false,
			"Hash": ""
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

##### Удаление группы
***
`DELETE : http://localhost/api/v1.0/admin/groups/obj/{token}/{id}`

> Controller: `base.Groups.ApiObj(apiObjDelete)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (группы)

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

