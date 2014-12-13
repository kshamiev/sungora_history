// app общеее (настраиваемое) пространство данных и функциоанала приложения
package app

import (
	"sort"
	typDb "types/db"
)

var Data = new(data)

// Данные (БД) представленные в памяти
// Именя свойств структуры должны соответствовать именам моделей
// Обычно имена типов и моделей совпадают.
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

var Routes = RouteList{}

type RouteList []*Route

type Route struct {
	Id     uint64 // Идентификатор Uri
	Uri    string // Uri.Uri
	Domain string // Домен или regexp описывающий домен
}

func (self RouteList) Len() int {
	return len(self)
}

func (self RouteList) Less(i, j int) bool {
	if len(self[i].Uri) > len(self[j].Uri) {
		return true
	}
	if len(self[i].Uri) == len(self[j].Uri) {
		return len(self[i].Domain) > len(self[j].Domain)
	}
	return false
}

func (self RouteList) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

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
