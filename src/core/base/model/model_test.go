// запуск теста
// SET GOPATH=C:\Work\zegota
// go test -v app/base/model | go test -v -bench . app/base/model
package model_test

import (
	"sort"
	"testing"

	"app"
	"app/baseА/model"
	"core"
	"core/config/typ"
	"lib"
	typDb "types/db"
)

func BenchmarkRange(b *testing.B) {
	b.StopTimer()
	var dataSlice []*typDb.Uri
	for i := uint64(0); i < 100000; i++ {
		var u = &typDb.Uri{
			Id:      i + 1,
			Content: dataBig,
		}
		dataSlice = append(dataSlice, u)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for i := range dataSlice {
			if dataSlice[i].Uri == `test` {
			}
		}
	}
}

func BenchmarkSearchBlank(b *testing.B) {
	b.StopTimer()
	var dataSlice []*typDb.Uri
	for i := uint64(0); i < 100000; i++ {
		var u = &typDb.Uri{
			Id:      i + 1,
			Content: dataBig,
		}
		dataSlice = append(dataSlice, u)
	}
	b.StartTimer()

	var index int
	for i := 0; i < b.N; i++ {
		index = sort.Search(len(dataSlice), func(i int) bool { return dataSlice[i].Id >= 51000 })
		if index < len(dataSlice) && dataSlice[index].Id == 51000 {
			// b.Log(dataSlice[index].Id)
		}
	}
}

// Search
func BenchmarkLoadId(b *testing.B) {
	b.StopTimer()

	core.Config = new(typ.Configuration)
	core.Config.Main = new(typ.Main)
	core.Config.Main.Db = 1

	app.Data.Uri = nil
	//var dataSlice []*typDb.Uri
	for i := uint64(0); i < 1000000; i++ {
		var u = &typDb.Uri{
			Id:      i + 1,
			Content: dataBig,
		}
		app.Data.Uri = append(app.Data.Uri, u)
	}
	app.Data.Uri[551110].Uri = `popcorn`
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var uri = model.NewUri(51111)
		uri.Model.Load(`Id`)
		//b.Log(uri.Type.Uri)
	}
}

// Search
func BenchmarkLoadProp(b *testing.B) {
	b.StopTimer()

	core.Config = new(typ.Configuration)
	core.Config.Main = new(typ.Main)
	core.Config.Main.Db = 1

	app.Data.Uri = nil
	//var dataSlice []*typDb.Uri
	for i := uint64(0); i < 1000000; i++ {
		var u = &typDb.Uri{
			Id:      i + 1,
			Content: dataBig,
		}
		app.Data.Uri = append(app.Data.Uri, u)
	}
	app.Data.Uri[551110].Uri = `popcorn`
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var uri = model.NewUri(0)
		uri.Type.Uri = `popcorn`
		uri.Model.Load(`Uri`)
		//b.Log(uri.Type.Id)
	}
}

func BenchmarkUriGrid(b *testing.B) {
	b.StopTimer()

	core.Config = new(typ.Configuration)
	core.Config.Main = new(typ.Main)
	core.Config.Main.Db = 1

	app.Data.Uri = nil
	//var dataSlice []*typDb.Uri
	for i := uint64(0); i < 1000000; i++ {
		var u = &typDb.Uri{
			Id:      i + 1,
			Content: dataBig,
		}
		app.Data.Uri = append(app.Data.Uri, u)
	}
	app.Data.Uri[551110].Uri = `popcorn`
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		model.GetUriGrid(30, 20, `Position`)
	}
}

func BenchmarkSort1(b *testing.B) {
	b.StopTimer()
	var dataSlice []*typDb.Uri
	for i := uint64(0); i < 100000; i++ {
		var u = &typDb.Uri{
			Id:      i + 1,
			Content: dataBig,
		}
		dataSlice = append(dataSlice, u)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		lib.Slice.SortingSliceDesc(dataSlice, `Id`)
	}
}

func BenchmarkSort2(b *testing.B) {
	b.StopTimer()
	var dataSlice []*typDb.Uri
	for i := uint64(0); i < 100000; i++ {
		var u = &typDb.Uri{
			Id:      i + 1,
			Content: dataBig,
		}
		dataSlice = append(dataSlice, u)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Sor(dataSlice, func(left, right *typDb.Uri) bool {
			return left.Id > right.Id
		})
	}
}

func BenchmarkSort3(b *testing.B) {
	b.StopTimer()
	var dataSlice []*typDb.Uri
	for i := uint64(0); i < 100000; i++ {
		var u = &typDb.Uri{
			Id:      i + 1,
			Content: dataBig,
		}
		dataSlice = append(dataSlice, u)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Sorting(dataSlice)
	}
}

var dataSmall = `Знаменитый`
var dataBig = `Знаменитый российский писатель Борис Акунин (Григорий Чхартишвили), автор романов о сыщике Фандорине, сообщил в своем блоге на сайте радиостанции "Эхо Москвы" о том, что покидает Россию.`

////

func Sor(slice []*typDb.Uri, f func(left, right *typDb.Uri) bool) {
	ps := &sor{
		data:     slice,
		function: f,
	}
	sort.Sort(ps)
}

type sor struct {
	data     []*typDb.Uri
	function func(left, right *typDb.Uri) bool
}

func (self sor) Len() int {
	return len(self.data)
}

func (self sor) Less(i, j int) bool {
	//return self[i].Id > self[j].Id
	return self.function(self.data[i], self.data[j])
}

func (self sor) Swap(i, j int) {
	self.data[i], self.data[j] = self.data[j], self.data[i]
}

////

func Sorting(s []*typDb.Uri) {
	sort.Sort(sorting(s))
}

type sorting []*typDb.Uri

func (self sorting) Len() int {
	return len(self)
}

func (self sorting) Less(i, j int) bool {
	return self[i].Id > self[j].Id
}

func (self sorting) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}
