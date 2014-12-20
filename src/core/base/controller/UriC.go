package controller

import (
	"app"
	"core"
	"core/base/model"
	"types"
	typDb "types/db"
	typReq "types/request"
)

type Uri struct {
	Session     *core.Session      // Сессия
	RW          *core.RW           // Управление вводом и выводом
	Controllers *typDb.Controllers // Соответсвующий контроллер из области данных по строковому ид.
}

// Создание контроллера
func NewUri(rw *core.RW, s *core.Session, c *typDb.Controllers) interface{} {
	var self = new(Uri)
	self.RW = rw
	self.Session = s
	self.Controllers = c
	return self
}

////

// ApiGrid Управление разделами. Список
func (self *Uri) ApiGrid() (err error) {
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
	return self.RW.ResponseJson(nil, 404, 520, `base.Uri.ApiGrid`)
}

// apiGridGet Список разделов
func (self *Uri) apiGridGet() (err error) {
	var data []*typDb.Uri
	// постраничный вывод
	var page, ok = self.RW.GetSegmentUriInt(`param`)
	if ok == false {
		page = 1
	}
	if data, err = model.GetUriGrid(int(page)); err != nil {
		return self.RW.ResponseJson(nil, 500, 500)
	}
	return self.RW.ResponseJson(data, 200, 0)
}

// apiGridPost Добавление раздела
func (self *Uri) apiGridPost() (err error) {
	// входящие данные
	var reqObj = new(typDb.Uri)
	if err = self.RW.RequestJsonParse(reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	// проверка
	var obj = model.NewUri(0)
	var validMessage map[string]string
	if validMessage, err = obj.Model.Set(reqObj, `All`); err != nil {
		return self.RW.ResponseJson(validMessage, 409, 530, `Uri.All`)

	}
	// сохранение
	if err = obj.Save(`All`, `Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 560, `Uri`)
	}

	return self.RW.ResponseJson(obj.Type.Id, 200, 0)
}

// apiGridPut Сортировка разделов
func (self *Uri) apiGridPut() (err error) {
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
		if err = model.SortingUri(reqPosition); err != nil {
			return self.RW.ResponseJson(nil, 409, 580, `Uri`)
		}
		return self.RW.ResponseJson(nil, 200, 0)
	case `route`:
		// реинициализация ротуинга
		app.ReInitRoute()
		return self.RW.ResponseJson(nil, 200, 0)
	}
	return self.RW.ResponseJson(nil, 409, 510)
}

// apiGridOptions Опции списка разделов
func (self *Uri) apiGridOptions() (err error) {
	var typ, ok = self.RW.GetSegmentUriString(`param`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 202)
	}
	switch typ {
	case `scenario`:
		var options *types.Scenario
		if options, err = types.GetScenario(`Uri`, `GridAdmin`); err != nil {
			return self.RW.ResponseJson(nil, 404, 570, `GridAdmin`, `Uri`)
		}
		return self.RW.ResponseJson(options, 200, 0)
	}
	return self.RW.ResponseJson(nil, 409, 203, typ)
}

////

// ApiObj Управление разделами. Детально
func (self *Uri) ApiObj() (err error) {
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
	return self.RW.ResponseJson(nil, 404, 520, `base.Uri.ApiObj`)
}

// apiObjGet Получение раздела
func (self *Uri) apiObjGet() (err error) {
	// входящие данные
	var id, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	//switch self.RW.Uri.Segment.Child {
	//case `content`:
	//	for i := range app.Data.Uri {
	//		if app.Data.Uri[i].Id == self.RW.Uri.Segment.Sid {
	//			self.RW.Content.File = app.Data.Uri[i].ContentFile
	//			self.RW.Content.Content = app.Data.Uri[i].Content
	//		}
	//	}
	//	return self.RW.Response()
	//}
	var obj = model.NewUri(id)
	if err = obj.Model.Load(`Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 500)
	}
	return self.RW.ResponseJson(obj.Type, 200, 0)
}

// apiObjPut Изменение раздела
func (self *Uri) apiObjPut() (err error) {
	// входящие данные
	var id, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var reqObj = new(typDb.Uri)
	if err = self.RW.RequestJsonParse(reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	// загрузка
	var obj = model.NewUri(id)
	if err = obj.Model.Load(`Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 500)
	}
	// проверка
	var validMessage map[string]string
	if validMessage, err = obj.Model.Set(reqObj, `All`); err != nil {
		return self.RW.ResponseJson(validMessage, 409, 530, `Uri.All`)

	}
	// сохранение
	if err = obj.Save(`All`, `Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 540, `Uri.All`, obj.Type.Id)
	}
	return self.RW.ResponseJson(nil, 200, 0)
}

// apiObjDelete Удаление раздела
func (self *Uri) apiObjDelete() (err error) {
	// входящие данные
	var id, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}

	var obj = model.NewUri(id)
	if err = obj.Remove(`Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 550, `Uri`, id)
	}
	return self.RW.ResponseJson(nil, 200, 0)
}

// apiObjOptions Опции редактирования раздела
func (self *Uri) apiObjOptions() (err error) {
	var options *types.Scenario
	if options, err = types.GetScenario(`Uri`, `All`); err != nil {
		return self.RW.ResponseJson(nil, 404, 570, `All`, `Uri`)
	}
	options.Sample = &typDb.Uri{
		Uri:           `/`,
		ContentType:   `text/html`,
		ContentEncode: `utf-8`,
	}
	return self.RW.ResponseJson(options, 200, 0)
}

////

// ApiObjControllers Управление контроллерами раздела
func (self *Uri) ApiObjControllers() (err error) {
	// диспетчер методов
	switch self.RW.Request.Method {
	case "GET":
		return self.apiObjControllersGet()
	case "PUT":
		return self.apiObjControllersPut()
	case "OPTIONS":
		return self.apiObjControllersOptions()
	}
	return self.RW.ResponseJson(nil, 404, 520, `base.Uri.ApiObjControllers`)
}

// apiObjControllersGet Контроллеры раздела (получение)
func (self *Uri) apiObjControllersGet() (err error) {
	// входящие данные
	var pid, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var obj = model.NewUri(pid)
	var tied, unbound = obj.LoadControllers()
	return self.RW.ResponseJson(map[string]interface{}{`tied`: tied, `unbound`: unbound}, 200, 0)
}

// apiObjControllersPut Контроллеры раздела (изменение)
func (self *Uri) apiObjControllersPut() (err error) {
	// входящие данные
	var pid, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var reqObj []*typDb.GroupsUri
	if err = self.RW.RequestJsonParse(&reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var obj = model.NewUri(pid)
	if err = obj.RemoveControllers(); err != nil {
		return self.RW.ResponseJson(nil, 409, 550, `base.Uri.apiObjControllersPut`, obj.Type.Id)
	}
	if err = obj.SaveControllers(reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 540, `base.Uri.apiObjControllersPut`, obj.Type.Id)
	}
	return self.RW.ResponseJson(nil, 200, 0)
}

// apiObjControllersOptions Опции контроллеров раздела
func (self *Uri) apiObjControllersOptions() (err error) {
	var options *types.Scenario
	if options, err = types.GetScenario(`GroupsUri`, `All`); err != nil {
		return self.RW.ResponseJson(nil, 404, 570, `All`, `GroupsUri`)
	}
	return self.RW.ResponseJson(options, 200, 0)
}

////

// ApiObjGroups Управление группами раздела
func (self *Uri) ApiObjGroups() (err error) {
	// диспетчер методов
	switch self.RW.Request.Method {
	case "GET":
		return self.apiObjGroupsGet()
	case "PUT":
		return self.apiObjGroupsPut()
	case "OPTIONS":
		return self.apiObjGroupsOptions()
	}
	return self.RW.ResponseJson(nil, 404, 520, `base.Uri.ApiObjGroups`)
}

// apiObjGroupsGet Группы раздела (получение)
func (self *Uri) apiObjGroupsGet() (err error) {
	// входящие данные
	var pid, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var obj = model.NewUri(pid)
	var tied, unbound = obj.LoadGroups()
	return self.RW.ResponseJson(map[string]interface{}{`tied`: tied, `unbound`: unbound}, 200, 0)
}

// apiObjGroupsPut Группы раздела (изменение)
func (self *Uri) apiObjGroupsPut() (err error) {
	// входящие данные
	var pid, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var reqObj []*typDb.GroupsUri
	if err = self.RW.RequestJsonParse(&reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var obj = model.NewUri(pid)
	if err = obj.RemoveGroups(); err != nil {
		return self.RW.ResponseJson(nil, 409, 550, `base.Uri.apiObjGroupsPut`, obj.Type.Id)
	}
	if err = obj.SaveGroups(reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 540, `base.Uri.apiObjGroupsPut`, obj.Type.Id)
	}
	return self.RW.ResponseJson(nil, 200, 0)
}

// apiObjGroupsOptions Опции групп раздела
func (self *Uri) apiObjGroupsOptions() (err error) {
	var options *types.Scenario
	if options, err = types.GetScenario(`GroupsUri`, `All`); err != nil {
		return self.RW.ResponseJson(nil, 404, 570, `All`, `GroupsUri`)
	}
	return self.RW.ResponseJson(options, 200, 0)
}
