### Документирование api

Для работы со свагером мы используем библиотеку: [swaggo](https://github.com/swaggo/swag#api-operation)

Описание документирования api:
<pre>
//+funcName godoc
//+@Summary Авторизация пользователя по логину и паролю (ldap).     пишем кратко о чем речь и что принимает на входе
// @Description Возвращается токен авторизации и пользователья      пишем что возвращает и возможно подробности
//+@Tags tagName                                                    группировка api запросов
//+@Router /page/page [post]                                        относительный роутинг от базового и метод
//+@Param name TARGET PARAM true "com"                              входящие параметры
//+@Success 200 {TYPE} string "com"                                 положительный ответ
//+@Failure 400 {TYPE} request.Error "com"                          отрицательный ответ
//+@Failure 401 {TYPE} request.Error "user unauthorized"            пользователь не авторизован
//+@Accept json                                                     тип принимаемых данных
//+@Produce json                                                    тип возвращаемых данных
//+@Security ApiKeyAuth                                             запрос авторизованный по ключу или токену
</pre>

<pre>
+ Обязательные теги и теги по контексту (параметров может и не быть...)
TARGET      = header | path | query  | body | formData
PARAM       = string | int  | number | bool | file | userGolangStruct
TYPE        = string | int  | number | bool | file | object | array
</pre>

Пример:
<pre>
// Login авторизация пользователя по логину и паролю ldap
// @Summary авторизация пользователя по логину и паролю (ldap).
// @Description возвращается токен авторизации
// @Tags Auth
// @Router /auth/login [post]
// @Param credentials body models.Credentials true "реквизиты доступа"
// @Success 200 {string} string "успешная авторизация"
// @Failure 400 {object} request.Error "operation error"
// @Failure 401 {object} request.Error "unauthorized"
// @Failure 403 {object} request.Error "forbidden"
// @Failure 404 {object} request.Error "not found"
// @Accept json                                                    
// @Produce json                                                   
// @Security ApiKeyAuth
</pre>

Формирование документации в "Makefile.Sample"

Далее документация доступа на сервере после выкладки по адресу `/api/docs/index.html`

Проблемы:

* Не умеет работать с алиасами в импортах.
* Типы slice, map не поддерживаются для входных параметров (нужно оборачивать в отдельные типы) 

Принятые коды ответов:

- 200 Любой положительный ответ
- 301 Редирект (перманентный). Переход на другой запрос
- 302 Редирект (от логики). Переход на другой запрос
- 400 Ошибка работы с данными приложения
- 401 Пользователь не авторизован
- 403 Отказано в операции за отсутствием прав 
- 404 Данные по запросу не найдены
- 500 Ошибка работы сервера

Формат принимаемых и отлаваемых данных для API:

- Данные передаются в формате JSON
