[main](/) ::
[quick start](/docs/start.html) ::
[reference](/docs/reference.html) ::
[tutorial](/docs/tutorial.html) ::
[api](/docs/api.html) ::
[sample](/sample) ::
[download](https://github.com/kshamiev/sungora)

<a name="anchor_home"></a>
## Сессия

1. [Регистрация нового пользователя](#anchor_registration)
1. [Выход](#anchor_exit)
1. [Авторизация](#anchor_auth)
1. [Проверка и продление токена](#anchor_online)
1. [Восстановление пароля пользователя](#anchor_recovery)
1. [Получение текущего пользователя](#anchor_usercurrent)
1. [Получение капчи](#anchor_capcha)

<a name="anchor_registration"></a>
### [Регистрация нового пользователя](#anchor_home)
***
1. [POST](#anchor_registration) /api/v1.0/session/registration/[token]

	> Sample: `http://localhost/api/v1.0/session/registration`
    Controller: `base.Session.ApiRegistration`

	"CaptchaHash", "CaptchaValue" указываются в случае ошибочной регистрации с первой попытки.<br>
	Для этого необходимо выводить капчу.<br>
	Если указан необязательный параметр [token] то регистрируемый пользователь становиться рефералом
    После регистрации пользователю отправляется регистрационное писмьо с логином и паролем.
	* **Parameters**
		* token `string` - токен идентификации пользователя
	* **Request**

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
			    "Email": "konstantin@shamiev.ru",
			    "CaptchaHash": "TSkLJxxcxY5OUbaV6avs",
			    "CaptchaValue": "keystring"
			}
	* **Response** 200

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
                "errorCode": 0,
                "errorMessage": "",
			}
	* **Response** 400

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100160,
				"errorMessage": "Требуется ввод капчи. Капча не указана"
			}
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100161,
				"errorMessage": "Капча не верна [keystring]"
			}
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100139,
				"errorMessage": "Токен [token] не верный"
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
				"errorCode": 100162,
				"errorMessage": "Email [konstantin@shamiev.ru] уже занят"
			}
	* **Response** 409

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100163,
				"errorMessage": "Логин [konstantin@shamiev.ru] уже занят"
			}
	* **Response** 500

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100560,
				"errorMessage": "Ошибка добавления [User.Registration]"
			}
	* **Response** 500

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100142,
				"errorMessage": "Ошибка отправки регистрационного письма [konstantin@shamiev.ru]"
			}

<a name="anchor_exit"></a>
### [Выход](#anchor_home)
***
1. [DELETE](#anchor_exit) /api/v1.0/session/authorization/{token}

	> Sample: `http://localhost/api/v1.0/session/authorization/{token}`
    Controller: `base.Session.ApiMain(apiMainDelete)`
	* **Parameters**
		* token `string` - токен идентификации пользователя
	* **Request**

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			-
	* **Response** 200

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 0,
				"errorMessage": "",
			}
	* **Response** 409

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100171,
				"errorMessage": "Токен не указан"
			}

<a name="anchor_auth"></a>
### [Авторизация](#anchor_home)
***
1. [PUT](#anchor_auth) /api/v1.0/session/authorization

	> Sample: `http://localhost/api/v1.0/session/authorization`
    Controller: `base.Session.ApiMain(apiMainPut)`

	"CaptchaHash", "CaptchaValue" указываются в случае ошибочной авторизации с первой попытки.<br>
	Для этого необходимо выводить капчу.<br>
	"Language" - необязательный
	* **Parameters**

			-
	* **Request**

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"Login": "konstantin@shamiev.ru",
				"Password": "543bf54bf8984abeeb5e",
				"Language": "language (ru-ru|en-en)",
			   "CaptchaHash": "hashString",
			   "CaptchaValue": "keystring"
			}
	* **Response** 200

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 0,
				"errorMessage": "",
				"content": "hashString",
			}
	* **Response** 400

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100160,
				"errorMessage": "Требуется ввод капчи. Капча не указана"
			}
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100161,
				"errorMessage": "Капча не верна [keystring]"
			}
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100169,
				"errorMessage": "Пользователь удален"
			}
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100168,
				"errorMessage": "Пользователь заблокирован"
			}
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100167,
				"errorMessage": "Пароль не верен"
			}
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100166,
				"errorMessage": "Пользователь не найден"
			}
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100165,
				"errorMessage": "Логин не указан"
			}
	* **Response** 409

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100510,
				"errorMessage": "Ошибка получение входных данных запроса"
			}
	* **Response** 500

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100170,
				"errorMessage": "Ошибка сохранения пользователя при авторизации"
			}

<a name="anchor_online"></a>
### [Проверка и продление токена](#anchor_home)
***
1. [GET](#anchor_online) /api/v1.0/session/authorization/{token}

	> Sample: `http://localhost/api/v1.0/session/authorization/{token}`
    Controller: `base.Session.ApiMain(apiMainGet)`
	* **Parameters**
		* token `string` - токен идентификации пользователя
	* **Request**

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			-
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

	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100166,
				"errorMessage": "Пользователь не найден"
			}
	* **Response** 409

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100171,
				"errorMessage": "Токен не указан"
			}

<a name="anchor_recovery"></a>
### [Восстановление пароля пользователя](#anchor_home)
***
1. [PUT](#anchor_recovery) /api/v1.0/session/recovery

	> Sample: `http://localhost/api/v1.0/session/recovery`
    Controller: `base.Session.ApiRecovery`

	"CaptchaHash", "CaptchaValue" указываются в случае ошибочной регистрации с первой попытки.<br>
	Для этого необходимо выводить капчу.<br>
	* **Parameters**

			-
	* **Request**

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
			    "Email": "konstantin@shamiev.ru",
			    "CaptchaHash": "TSkLJxxcxY5OUbaV6avs",
			    "CaptchaValue": "keystring"
			}
	* **Response** 200

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 0,
				"errorMessage": "",
			}
	* **Response** 400

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100160,
				"errorMessage": "Требуется ввод капчи. Капча не указана"
			}
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100161,
				"errorMessage": "Капча не верна [keystring]"
			}
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100176,
				"errorMessage": "Email не указан"
			}
	* **Response** 404

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100166,
				"errorMessage": "Пользователь не найден"
			}
	* **Response** 409

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100178,
				"errorMessage": "Ошибка изменение пароля для пользователя"
			}
	* **Response** 409

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100180,
				"errorMessage": "Ошибка отправки письма с восстановлением [%s]"
			}
	* **Response** 409

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 100510,
				"errorMessage": "Ошибка получение входных данных запроса"
			}

<a name="anchor_usercurrent"></a>
### [Получение текущего пользователя](#anchor_home)
***
1. [GET](#anchor_usercurrent) /api/v1.0/session/[token]

	> Sample: `http://localhost/api/v1.0/session/[token]`
    Controller: `base.Session.ApiUserCurrent`
	* **Parameters**
		* token `string` - токен идентификации пользователя
	* **Request**

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			-
	* **Response** 200

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			{
				"errorCode": 0,
				"errorMessage": "",
				"content":
				{
					"Id": 1,
					"Users_Id": 0,
					"Login": "guestguestguest",
					"Password": "",
					"PasswordR": "",
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

<a name="anchor_capcha"></a>
### [Получение капчи](#anchor_home)
***
1. [GET](#anchor_capcha) /api/v1.0/session/captcha/native

	> Sample: `http://localhost/api/v1.0/session/captcha/native`
    Controller: `base.Session.ApiCaptcha`
	* **Parameters**

		-
	* **Request**

		Headers

			Content-Type: application/json; charset=utf-8
		Body

			-
	* **Response** 200

		Headers

			Content-Type: application/json; charset=utf-8
		Body

		    {
		       "errorCode": 0,
		       "errorMessage": "",
			   "content":
			   {
				   "CaptchaImage": "hashDataString",
				   "CaptchaHash": "TSkLJxxcxY5OUbaV6avs",
			   }
		    }
   	* **Response** 409

		Headers

			Content-Type: application/json; charset=utf-8
		Body

		    {
		       "errorCode": 100159,
		       "errorMessage": "Слишком частые запросы капчи",
		    }
	* **Response** 409

		Headers

			Content-Type: application/json; charset=utf-8
		Body

		    {
		       "errorCode": 100140,
		       "errorMessage": "Ошибка получения капчи %v",
		    }
