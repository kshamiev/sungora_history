// запуск теста
// SET GOPATH=C:\Work\projectName
// go test -v lib
// go test -v -bench . lib
package lib_test

import (
	"testing"
	"time"

	"lib"
	"lib/logs"
)

type Uri struct {
	Id            uint64    // Id
	Method        []string  // Метод доступа, разрешается перечисление через запятую без пробелов. ALL,POST,GET,DELETE,UPDATE и т.п.
	Domain        string    // Домен или regexp описывающий домен
	Uri           string    // URI без указания домена и протокола до необязательных параметров
	UriSegment    string    // Сегменированные параметры URI
	Name          string    // Название раздела
	Redirect      string    // Если не пусто, то содержит адрес безусловной переадресации
	Layout        string    // Макет сайта
	IsAuthorized  bool      // =1 - доступ к разделу разрешен только авторизованным пользователям
	IsMenuVisible bool      // =1 - раздел отображается в стандартном меню
	IsDisable     bool      // =1 - Раздел отключен, при попытке доступа выдается 404
	Content       []byte    // Шаблон или контент раздела (текстовые или бинарные данные)
	ContentFile   string    // Имя загруженного файла
	ContentTime   time.Time // Дата и время копии файла контента в файловой системе (синхронизация БД с файловой системой)
	ContentType   string    // Mime type контента раздела, может быть переопределён контроллером
	ContentEncode string    // Кодировка контента (по умолчанию utf-8) для заголовка
	Position      int32     // Сортировка, приоритет в роутинге
	Title         string    // Заголовок раздела - title
	KeyWords      string    // Ключевые слова - keywords
	Description   string    // Описание раздела - description
	Controllers   []string  `db:"-"` // Контроллеры урла (Path)
}

func TestSliceSorting(t *testing.T) {

	var data []*Uri

	// Groups
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/admin/groups/obj/{token}`,
		UriSegment:  `[sid]/[child]/[cid]`,
		Name:        `Управление разделами (роутинг)`,
		ContentType: `application/json`,
		Controllers: []string{`base/Groups/ApiObj`},
	})
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/admin/groups/{token}`,
		UriSegment:  `[self]/[sid]`,
		Name:        `Управление разделами (роутинг)`,
		ContentType: `application/json`,
		Controllers: []string{`base/Groups/ApiGrid`},
	})
	// Users
	data = append(data, &Uri{
		Method:      []string{`GET`, `PUT`},
		Uri:         `/api/v1.0/admin/users/obj/groups/{token}/{pid}`,
		Name:        `Управление контроллерами разделов`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Uri/ApiObjGroups`},
	})
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/admin/users/obj/{token}`,
		UriSegment:  `[sid]/[child]/[cid]`,
		Name:        `Управление разделами (роутинг)`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Users/ApiObj`},
	})
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/admin/users/{token}`,
		UriSegment:  `[self]/[sid]`,
		Name:        `Управление разделами (роутинг)`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Users/ApiGrid`},
	})
	// Controllers
	data = append(data, &Uri{
		Method:      []string{`GET`, `PUT`},
		Uri:         `/api/v1.0/admin/controllers/obj/groups/{token}/{pid}`,
		Name:        `Управление контроллерами разделов`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Uri/ApiObjGroups`},
	})
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/admin/controllers/obj/{token}`,
		UriSegment:  `[sid]/[child]/[cid]`,
		Name:        `Управление разделами (роутинг)`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Controllers/ApiObj`},
	})
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/admin/controllers/{token}`,
		UriSegment:  `[self]/[sid]`,
		Name:        `Управление разделами (роутинг)`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Controllers/ApiGrid`},
	})
	// uri
	data = append(data, &Uri{
		Method:      []string{`GET`, `PUT`},
		Uri:         `/api/v1.0/admin/uri/obj/groups/{token}/{pid}`,
		Name:        `Управление контроллерами разделов`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Uri/ApiObjGroups`},
	})
	data = append(data, &Uri{
		Method:      []string{`GET`, `PUT`},
		Uri:         `/api/v1.0/admin/uri/obj/controllers/{token}/{pid}`,
		UriSegment:  `[sid]/[child]/[cid]`,
		Name:        `Управление контроллерами разделов`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Uri/ApiObjControllers`},
	})
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/admin/uri/obj/{token}`,
		UriSegment:  `[sid]/[child]/[cid]`,
		Name:        `Управление разделами (роутинг)`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Uri/ApiObj`},
	})
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/admin/uri/{token}`,
		UriSegment:  `[self]/[sid]`,
		Name:        `Управление разделами (роутинг)`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Uri/ApiGrid`},
	})
	// сессия
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/session/recovery`,
		UriSegment:  `[token]`,
		Name:        `Восстановление пароля пользователя`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Session/ApiRecovery`},
	})
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/session/authorization`,
		Name:        `Аутентификация`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Session/ApiMain`},
	})
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/session/registration`,
		UriSegment:  `[token]`,
		Name:        `Регистрация нового пользователя`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Session/ApiRegistration`},
	})
	data = append(data, &Uri{
		Method:      []string{`GET`},
		Uri:         `/api/v1.0/session/captcha/native`,
		Name:        `Капча`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Session/ApiCaptcha`},
	})
	// сервер
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/server/ping`,
		UriSegment:  `[token]`,
		Name:        `Проверка доступности сервера`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Server/ApiPing`},
	})
	data = append(data, &Uri{
		Method:      []string{`GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`},
		Uri:         `/api/v1.0/server/upload`,
		UriSegment:  `[token]`,
		Name:        `Проверка доступности сервера`,
		ContentType: `application/json`,
		Controllers: []string{`zero/Server/ApiUpload`},
	})

	t.Log("\nСортирвока по возрастанию")
	lib.Slice.SortingSliceAsc(data, `Uri`)
	for i := range data {
		t.Log(data[i].Uri)
	}
	t.Log("\nСортирвока по убыванию")
	lib.Slice.SortingSliceDesc(data, `Uri`)
	for i := range data {
		t.Log(data[i].Uri)
	}
	t.Log("\nПользовательская сортирвока")
	lib.Slice.SortingSliceFunc(data, func(left, right interface{}) bool {
		return len(left.(*Uri).Uri) > len(right.(*Uri).Uri)
	})
	for i := range data {
		t.Log(data[i].Uri)
	}
	logs.Dumper()
}
