// app Область данных приложения, Модули приложения, Роутинг.
package app

import (
	"sort"

	typDb "types/db"
)

// Данные приложения (БД) представленные в памяти
var Data = new(data)

// Данные приложения (БД) представленные в памяти
// Именя свойств структуры должны соответствовать именам моделей (если таковые имеются)
// Тип свойства обычно является типом описывающим структуру данных в БД.
// Обычно имена типов и моделей совпадают (когда всего одна модель).
// db:"cross" - для стуктур не имеющий свойства - поля Id
// db:"-" - игнорирование (централизовано не обрабатывать)
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

// Роутинг
var Routes = RouteList{}

// Структура роутинга (для сортировки)
type RouteList []*Route

// Структура роутинга
type Route struct {
	Id     uint64 // Идентификатор Uri
	Uri    string // адрес запроса без домена (/page/page == Uri.Uri)
	Domain string // Строка описывающий домен или его часть (shop.funtik.ru, shop, funtik.ru)
}

// Len() int Сортировка
func (self RouteList) Len() int {
	return len(self)
}

// Less(int, int) bool Сортировка
func (self RouteList) Less(i, j int) bool {
	if len(self[i].Uri) > len(self[j].Uri) {
		return true
	}
	if len(self[i].Uri) == len(self[j].Uri) {
		return len(self[i].Domain) > len(self[j].Domain)
	}
	return false
}

// Swap(int, int) Сортировка
func (self RouteList) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

// ReInitRoute() Инициализация роутинга (после изменения разделов Uri и в момент запуска приложения)
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
