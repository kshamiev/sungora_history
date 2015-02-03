package controller

import (
	"app"
	"core"
	"core/base/model"
	"types"
	typDb "types/db"
	typReq "types/request"
)

type Controllers struct {
	Session     *core.Session      // Сессия
	RW          *core.RW           // Управление вводом и выводом
	Controllers *typDb.Controllers // Соответсвующий контроллер из области данных по строковому ид.
}

// Создание контроллера
func NewControllers(rw *core.RW, s *core.Session, c *typDb.Controllers) interface{} {
	var self = new(Controllers)
	self.RW = rw
	self.Session = s
	self.Controllers = c
	return self
}

////

// ApiGrid Управление контроллерами. Список
func (self *Controllers) ApiGrid() (err error) {
	// диспетчер методов
	switch self.RW.Request.Method {
	case "GET":
		return self.apiGridGet()
	case "POST":
		return self.apiGridPost()
	case "PUT":
		return self.apiGridPut()
	case "OPTIONS":
		return self.apiGridOptions()
	}
	return self.RW.ResponseJson(nil, 404, 520, `base.Controllers.ApiGrid`)
}

// apiGridGet Список контроллеров
func (self *Controllers) apiGridGet() (err error) {
	var data []*typDb.Controllers
	// постраничный вывод
	var page, ok = self.RW.GetSegmentUriInt(`param`)
	if ok == false {
		page = 1
	}
	if data, err = model.GetControllersGrid(int(page)); err != nil {
		return self.RW.ResponseJson(nil, 500, 500)
	}
	return self.RW.ResponseJson(data, 200, 0)
}

// apiGridPost Добавление контроллера
func (self *Controllers) apiGridPost() (err error) {
	// входящие данные
	var reqObj = new(typDb.Controllers)
	if err = self.RW.RequestJsonParse(reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	// проверка
	var obj = model.NewControllers(0)
	var validMessage map[string]string
	if validMessage, err = obj.Model.Set(reqObj, `All`); err != nil {
		return self.RW.ResponseJson(validMessage, 409, 530, `Controllers.All`)

	}
	// сохранение
	if err = obj.Save(`All`, `Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 560, `Controllers`)
	}

	return self.RW.ResponseJson(obj.Type.Id, 200, 0)
}

// apiGridPut Сортировка контроллеров
func (self *Controllers) apiGridPut() (err error) {
	var typ, ok = self.RW.GetSegmentUriString(`param`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	switch typ {
	case `position`:
		// входящие данные
		var reqPosition = new(typReq.Position)
		if err = self.RW.RequestJsonParse(reqPosition); err != nil {
			return self.RW.ResponseJson(nil, 409, 510)
		}
		// сортировка
		if err = model.SortingControllers(reqPosition); err != nil {
			return self.RW.ResponseJson(nil, 409, 580, `Controllers`)
		}
		return self.RW.ResponseJson(nil, 200, 0)
	}
	return self.RW.ResponseJson(nil, 409, 510)
}

// apiGridOptions Опции списка контроллеров
func (self *Controllers) apiGridOptions() (err error) {
	var typ, ok = self.RW.GetSegmentUriString(`param`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 575)
	}
	switch typ {
	case `scenario`:
		var options *types.Scenario
		if options, err = types.GetScenario(`Controllers`, `GridAdmin`); err != nil {
			return self.RW.ResponseJson(nil, 404, 104, `Controllers`, `GridAdmin`)
		}
		return self.RW.ResponseJson(options, 200, 0)
	}
	return self.RW.ResponseJson(nil, 409, 570, typ)
}

////

// ApiObj Управление контроллерами. Детально
func (self *Controllers) ApiObj() (err error) {
	// диспетчер методов
	switch self.RW.Request.Method {
	case "GET":
		return self.apiObjGet()
	case "PUT":
		return self.apiObjPut()
	case "DELETE":
		return self.apiObjDelete()
	case "OPTIONS":
		return self.apiObjOptions()
	}
	return self.RW.ResponseJson(nil, 404, 520, `base.Controllers.ApiObj`)
}

// apiObjGet Получение контроллера
func (self *Controllers) apiObjGet() (err error) {
	// входящие данные
	var id, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	//switch self.RW.Uri.Segment.Child {
	//case `content`:
	//	for i := range app.Data.Controllers {
	//		if app.Data.Controllers[i].Id == self.RW.Uri.Segment.Sid {
	//			self.RW.Content.File = app.Data.Controllers[i].Path + `.html`
	//			self.RW.Content.Content = app.Data.Controllers[i].Content
	//		}
	//	}
	//	return self.RW.Response()
	//}
	var obj = model.NewControllers(id)
	if err = obj.Model.Load(`Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 500)
	}
	return self.RW.ResponseJson(obj.Type, 200, 0)
}

// apiObjPut Изменение контроллера
func (self *Controllers) apiObjPut() (err error) {
	// входящие данные
	var id, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var reqObj = new(typDb.Controllers)
	if err = self.RW.RequestJsonParse(reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	// загрузка
	var obj = model.NewControllers(id)
	if err = obj.Model.Load(`Id`); err != nil {
		return self.RW.ResponseJson(nil, 404, 500)
	}
	// проверка
	var validMessage map[string]string
	if validMessage, err = obj.Model.Set(reqObj, `All`); err != nil {
		return self.RW.ResponseJson(validMessage, 409, 530, `Controllers.All`)

	}
	// сохранение
	if err = obj.Save(`All`, `Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 540, `Controllers.All`, obj.Type.Id)
	}
	return self.RW.ResponseJson(nil, 200, 0)
}

// apiObjDelete Удаление контроллера
func (self *Controllers) apiObjDelete() (err error) {
	// входящие данные
	var id, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}

	var obj = model.NewControllers(id)
	if err = obj.Remove(`Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 550, `Controllers`, id)
	}
	return self.RW.ResponseJson(nil, 200, 0)
}

// apiObjOptions Опции редактирования контроллера
func (self *Controllers) apiObjOptions() (err error) {
	var options *types.Scenario
	if options, err = types.GetScenario(`Controllers`, `All`); err != nil {
		return self.RW.ResponseJson(nil, 404, 570, `All`, `Controllers`)
	}
	options.Sample = &typDb.Controllers{
		Name: `SampleName`,
	}
	return self.RW.ResponseJson(options, 200, 0)
}

////

// ApiObjGroups Управление группами контроллера
func (self *Controllers) ApiObjGroups() (err error) {
	// диспетчер методов
	switch self.RW.Request.Method {
	case "GET":
		return self.apiObjGroupsGet()
	case "PUT":
		return self.apiObjGroupsPut()
	case "OPTIONS":
		return self.apiObjGroupsOptions()
	}
	return self.RW.ResponseJson(nil, 404, 520, `base.Controllers.ApiObjGroups`)
}

// apiObjGroupsGet Группы контроллера (получение)
func (self *Controllers) apiObjGroupsGet() (err error) {
	// входящие данные
	var pid, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var obj = model.NewControllers(pid)
	var tied, unbound = obj.LoadGroups()
	return self.RW.ResponseJson(map[string]interface{}{`tied`: tied, `unbound`: unbound}, 200, 0)
}

// apiObjGroupsPut Группы контроллера (изменение)
func (self *Controllers) apiObjGroupsPut() (err error) {
	// входящие данные
	var pid, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var reqObj []*typDb.GroupsUri
	if err = self.RW.RequestJsonParse(&reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var obj = model.NewControllers(pid)
	if err = obj.RemoveGroups(); err != nil {
		return self.RW.ResponseJson(nil, 409, 550, `base.Controllers.apiObjGroupsPut`, obj.Type.Id)
	}
	if err = obj.SaveGroups(reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 540, `base.Controllers.apiObjGroupsPut`, obj.Type.Id)
	}
	return self.RW.ResponseJson(nil, 200, 0)
}

// apiObjGroupsOptions Опции групп контроллера
func (self *Controllers) apiObjGroupsOptions() (err error) {
	var options *types.Scenario
	if options, err = types.GetScenario(`GroupsUri`, `All`); err != nil {
		return self.RW.ResponseJson(nil, 404, 570, `All`, `GroupsUri`)
	}
	return self.RW.ResponseJson(options, 200, 0)
}

////

// ApiProblem Список проблемных контроллеров
func (self *Controllers) ApiProblem() (err error) {
	var data = app.CheckControllers()
	return self.RW.ResponseJson(data, 200, 0)
}
