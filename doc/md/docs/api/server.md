[main](/) ::
[quick start](/docs/start.html) ::
[reference](/docs/reference.html) ::
[tutorial](/docs/tutorial.html) ::
[api](/docs/api.html) ::
[sample](/sample) ::
[download](/https://github.com/kshamiev/sungora)

<a name="anchor_home"></a>
## Сервер

1. [Проверка доступности сервера](#anchor_ping)
2. [Загрузка бинарных данных](#anchor_upload)
3. [Получение сценариев для всех типов](#anchor_scenario)

<a name="anchor_ping"></a>
### [Проверка доступности сервера](#anchor_home)
***
1. [GET](#anchor_ping) /api/v1.0/server/ping

	> Sample: `http://localhost/api/v1.0/server/ping`
    Controller: `base.Server.ApiPing`
	* **Parameters**

			-
	* **Request**

		Headers

			-
		Body

			-
	* **Response** 200

		Headers

			Content-Type: application/json; charset=utf-8
		Body

		    {
		       "errorCode": 0,
		       "errorMessage": ""
		    }

<a name="anchor_upload"></a>
### [Загрузка бинарных данных](#anchor_home)
***
1. [POST](#anchor_upload) /api/v1.0/server/upload/{token}

	> Sample: `http://localhost/api/v1.0/server/upload/{token}`
    Controller: `base.Server.ApiUpload`
	* **Parameters**
		* token `string` - токен идентификации пользователя
	* **Request**

		Headers

			Content-Type	multipart/form-data; boundary=-...
		Body

			бинарные данные. myFile - имя переменной формы "контейнера"
	* **Response** 200

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 0,
				"errorMessage": "",
				"content": "1415123595698052500" // идентификатор файла на сервере. передается далее с формой.
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
		       "errorCode": 100171,
		       "errorMessage": "Токен не указан"
		    }
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

		    {
		       "errorCode": 100139,
		       "errorMessage": "Токен [%s] не верный"
		    }
	* **Response** 409

		Headers

			Content-Type: application/json; charset=utf-8
		Body

		    {
		       "errorCode": 100590,
		       "errorMessage": "Ошибка загрузки файла ..."
		    }

<a name="anchor_scenario"></a>
### [Получение сценариев для всех типов](#anchor_home)
***
1. [OPTIONS](#anchor_scenario) /api/v1.0/types/scenario/{typ}/[scenario]

	> Sample: `http://localhost/api/v1.0/types/scenario/uri`
    Controller: `base.Server.ApiTypesScenario`

    Этот запрос сделан только для примера, ознакомления и тестирования.
    Не используйте его в продакшене.
	* **Parameters**
		* typ `string` - тип объектов
		* scenario `string` - сценарий типа
	* **Request**

		Headers

			Content-Type	multipart/form-data; boundary=-...
		Body

			бинарные данные. myFile - имя переменной формы "контейнера"
	* **Response** 200 Список сценариев

		Headers

			Content-Type: application/json; charset=utf-8
		Body

		    {
		       "errorCode": 0,
		       "errorMessage": "",
		       "content":
		       [
		           "all",
		           "registration",
		           "recovery",
		           "profile",
		           "online",
		           "root"
		       ]
		    }
	* **Response** 200 Выбранный сценарий

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
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

		    {
		       "errorCode": 100122,
		       "errorMessage": "Нет ни одного сценария [%s]"
		    }
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

		    {
		       "errorCode": 100570,
		       "errorMessage": "Сценарй [%s] не найден для [%s]"
		    }
	* **Response** 409

		Headers

			Content-Type: application/json; charset=utf-8
		Body

		    {
		       "errorCode": 100120,
		       "errorMessage": "Тип не указан"
		    }
