[main](/) ::
[quick start](/docs/start.html) ::
[reference](/docs/reference.html) ::
[tutorial](/docs/tutorial.html) ::
[api](/docs/api.html) ::
[sample](/sample) ::
[download](/https://github.com/kshamiev/sungora)

# Api
***

Общее положение:
Используется концепция REST.
Передача данных осуществляется в формате JSON.

Спецсимволы в ури запроса:
+ {} - обязательный параметр (любое скалярное значение)
+ [] - необязательный параметр (любое скалярное значение)

Спецсимволы в теле запроса и ответа:
+ {} - объект
+ [] - список (масссив) значений (объектов)

В описании 'Controller' указывается обрабатывающий запрос контроллер Go `moduleName/ControllerName/metodName`.

##### [Сервер](/docs/api/server.html)

##### [Сессия](/docs/api/session.html)

##### Административная часть

* [Разделы (Uri)](/docs/api/admin-uri.html)
* [Контроллеры (Controllers)](/docs/api/admin-controllers.html)
* [Пользователи (Users)](/docs/api/admin-users.html)
* [Группы (Groups)](/docs/api/admin-groups.html)
