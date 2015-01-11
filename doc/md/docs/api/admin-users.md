[Главная](/) ::
[Начало](/docs/start.html) ::
[Справочник](/docs/reference.html) ::
[Учебник](/docs/tutorial.html) ::
[Системное api](/docs/api.html) ::
[Скачать](https://github.com/kshamiev/sungora)

## Пользователи (Users)

##### Опции списка пользователей
***
`OPTIONS : http://localhost/api/v1.0/admin/users/{token}/{type}`

> Controller: `base.Users.ApiGrid(apiGridOptions)`

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
				"Name": "Пользователи",
				"Description": "Список пользователей в админке",
				"Property":
				[
					{
						"Name": "Email",
						"Title": "Email",
						"AliasDb": "`Email`",
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

##### Список пользователей
***
`GET : http://localhost/api/v1.0/admin/users/{token}/[page]`

> Controller: `base.Users.ApiGrid(apiGridGet)`

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


					"Id": 2,
					"Email": "guest@guest.guest",
					"LastName": "",
					"Name": "",
					"MiddleName": "",
					"Date": "0001-01-01T00:00:00Z",
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

##### Добавление пользователя
***
`POST : http://localhost/api/v1.0/admin/users/{token}`

> Controller: `base.Users.ApiGrid(apiGridPost)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

		{
			"Users_Id": 0,
			"Login": "LoginName",
			"Password": "",
			"Email": "name@name.zone",
			"LastName": "",
			"Name": "",
			"MiddleName": "",
			"IsAccess": false,
			"IsCondition": false,
			"IsActivated": false,
			"DateOnline": "0001-01-01T00:00:00Z",
			"Date": "0001-01-01T00:00:00Z",
			"Del": false,
			"Hash": "",
			"Token": "",
			"Language": "",
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

##### Опции пользователя подробно
***
`OPTIONS : http://localhost/api/v1.0/admin/users/obj/{token}/{id}`

> Controller: `base.Users.ApiObj(apiObjOptions)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (пользователя)

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
			   "Name": "Пользователи",
			   "Description": "Все свойства",
			   "Property":
			   [
					{
						"Name": "Users_Id",
						"Title": "Пользователь",
						"AliasDb": "`Users_Id`",
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

##### Получение пользователя
***
`GET : http://localhost/api/v1.0/admin/users/obj/{token}/{id}`

> Controller: `base.Users.ApiObj(apiObjGet)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (пользователя)

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
				"Users_Id": 0,
				"Login": "guestguestguest",
				"Password": "",
				"Email": "guest@guest.guest",
				"LastName": "",
				"Name": "",
				"MiddleName": "",
				"IsAccess": false,
				"IsCondition": false,
				"IsActivated": false,
				"DateOnline": "0001-01-01T00:00:00Z",
				"Date": "0001-01-01T00:00:00Z",
				"Del": false,
				"Hash": "",
				"Token": "",
				"Language": "",
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

##### Изменение пользователя
***
`PUT : http://localhost/api/v1.0/admin/users/obj/{token}/{id}`

> Controller: `base.Users.ApiObj(apiObjPut)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (пользователя)

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

		{
			"Users_Id": 0,
			"Login": "LoginName",
			"Password": "",
			"Email": "name@name.zone",
			"LastName": "",
			"Name": "",
			"MiddleName": "",
			"IsAccess": false,
			"IsCondition": false,
			"IsActivated": false,
			"DateOnline": "0001-01-01T00:00:00Z",
			"Date": "0001-01-01T00:00:00Z",
			"Del": false,
			"Hash": "",
			"Token": "",
			"Language": "",
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

##### Удаление пользователя
***
`DELETE : http://localhost/api/v1.0/admin/users/obj/{token}/{id}`

> Controller: `base.Users.ApiObj(apiObjDelete)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (пользователя)

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

##### Опции групп пользователя
***
`OPTIONS : http://localhost/api/v1.0/admin/users/obj/{token}/{id}/groups`

> Controller: `base.Users.ApiObjGroups(apiObjGroupsOptions)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (пользователя)

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

##### Группы пользователя (получение)
***
`GET : http://localhost/api/v1.0/admin/users/obj/{token}/{id}/groups`

> Controller: `base.Users.ApiObjGroups(apiObjGroupsGet)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (пользователя)

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
						"Users_Id": 2,
						"Groups_Id": 2,
						"Name": "guest"
					},
					...
				],
				"unbound":
				[
					{
						"Users_Id": 0,
						"Groups_Id": 1,
						"Name": "developer"
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

##### Группы пользователя (изменение)
***
`PUT : http://localhost/api/v1.0/admin/users/obj/{token}/{id}/groups`

> Controller: `base.Users.ApiObjGroups(apiObjGroupsPut)`

- **Parameters**
	- `token`:`string` - токен идентификации пользователя
	- `id`:`uint64` - id uri (пользователя)

* **Request**

    Headers

		Content-Type: application/json; charset=utf-8
    Body

		[
			{
				"Users_Id": 0,
				"Groups_Id": 1,
				"Name": "developer"
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

