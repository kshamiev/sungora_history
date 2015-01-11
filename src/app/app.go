// Область приложения.
//
// Модули, данные, роутинг приложения.
package app

import (
	"reflect"
	"sort"
	"strings"

	"core"
	"lib/logs"
	typDb "types/db"
)

// Данные приложения, (БД) представленные в памяти.
// Именя свойств структуры должны соответствовать именам моделей (если таковые имеются).
// Тип свойства обычно является типом описывающим структуру данных в БД.
// Обычно имена типов и моделей совпадают (когда всего одна модель).
// db:"cross" - для стуктур не имеющий свойства - поля Id.
// db:"-" - игнорирование при загрузке данных.
var Data = new(data)

type data struct {
	Controllers []*typDb.Controllers  ``
	GroupsUri   []*typDb.GroupsUri    `db:"cross"`
	Groups      []*typDb.Groups       ``
	Users       []*typDb.Users        ``
	UsersGroups []*typDb.UsersGroups  `db:"cross"`
	Uri         []*typDb.Uri          ``
	Content     []*typDb.Content      `db:"cross"`
	UriDefault  map[string]*typDb.Uri `db:"-"`
	MaxPostion  map[string]int32      `db:"-"`
}

//// Роутинг

// Отсортированные данные роутинга для поиска uri
var Routes = RouteList{}

// Структура роутинга (для сортировки)
type RouteList []*Route

// Структура роутинга
type Route struct {
	Id     uint64 // Идентификатор Uri
	Uri    string // адрес запроса без домена (/page/page == Uri.Uri)
	Domain string // Строка описывающий домен или его часть (shop.funtik.ru, shop, funtik.ru)
}

// Сортировка
func (self RouteList) Len() int {
	return len(self)
}

// Сортировка
func (self RouteList) Less(i, j int) bool {
	if len(self[i].Uri) > len(self[j].Uri) {
		return true
	}
	if len(self[i].Uri) == len(self[j].Uri) {
		return len(self[i].Domain) > len(self[j].Domain)
	}
	return false
}

// Сортировка
func (self RouteList) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

// Инициализация роутинга (после изменения разделов Uri и в момент запуска приложения)
func ReInitRoute() {
	var data RouteList
	for i := range Data.Uri {
		data = append(data, &Route{
			Id:     Data.Uri[i].Id,
			Uri:    Data.Uri[i].Uri,
			Domain: Data.Uri[i].Domain,
		})
	}
	sort.Sort(data)
	Routes = data
}

//// Контроллеры

// Хеш функций-конструкторов контроллеров
var Controller = make(map[string]func(rw *core.RW, s *core.Session, c *typDb.Controllers) interface{})

// Проверка корректности и работспособности контроллеров.
//    + []*typDb.Controllers - срез ошибочных контроллеров
func CheckControllers() (data []*typDb.Controllers) {
	for _, c := range Data.Controllers {
		// путь до контроллера и его метода в неправильном формате
		l := strings.Split(c.Path, `/`)
		if len(l) != 3 {
			logs.Critical(172, c.Path)
			data = append(data, c)
			continue
		}
		// нет такого контроллера
		ctrF, ok := Controller[l[0]+`/`+l[1]]
		if false == ok {
			logs.Critical(173, l[0], l[1])
			data = append(data, c)
			continue
		}
		// нет такого метода
		var ctr = ctrF(nil, nil, nil)
		objValue := reflect.ValueOf(ctr)
		met := objValue.MethodByName(l[2])
		if met.IsValid() == false {
			logs.Critical(174, l[0], l[1], l[2])
			data = append(data, c)
			continue
		}
	}
	return data
}
