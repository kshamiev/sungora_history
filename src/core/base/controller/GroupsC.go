package controller

import (
	"core"
	"core/base/model"
	"types"
	typDb "types/db"
)

type Groups struct {
	Session     *core.Session      // Сессия
	RW          *core.RW           // Управление вводом и выводом
	Controllers *typDb.Controllers // Соответсвующий контроллер из области данных по строковому ид.
}

// Создание контроллера
func NewGroups(rw *core.RW, s *core.Session, c *typDb.Controllers) interface{} {
	var self = new(Groups)
	self.RW = rw
	self.Session = s
	self.Controllers = c
	return self
}

////

// ApiGrid Управление групами. Список
func (self *Groups) ApiGrid() (err error) {
	// диспетчер методов
	switch self.RW.Request.Method {
	case "GET":
		return self.apiGridGet()
	case "POST":
		return self.apiGridPost()
	case "OPTIONS":
		return self.apiGridOptions()
	}
	return self.RW.ResponseJson(nil, 404, 520, `base.Groups.ApiGrid`)
}

// apiGridGet Список групп
func (self *Groups) apiGridGet() (err error) {
	var data []*typDb.Groups
	// постраничный вывод
	var page, ok = self.RW.GetSegmentUriInt(`param`)
	if ok == false {
		page = 1
	}
	if data, err = model.GetGroupsGrid(int(page)); err != nil {
		return self.RW.ResponseJson(nil, 500, 500)
	}
	return self.RW.ResponseJson(data, 200, 0)
}

// apiGridPost Добавление группы
func (self *Groups) apiGridPost() (err error) {
	// входящие данные
	var reqObj = new(typDb.Groups)
	if err = self.RW.RequestJsonParse(reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	// проверка
	var obj = model.NewGroups(0)
	var validMessage map[string]string
	if validMessage, err = obj.Model.Set(reqObj, `All`); err != nil {
		return self.RW.ResponseJson(validMessage, 409, 530, `Groups.All`)

	}
	// сохранение
	if err = obj.Save(`All`, `Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 560, `Groups`)
	}

	return self.RW.ResponseJson(obj.Type.Id, 200, 0)
}

// apiGridOptions Опции списка групп
func (self *Groups) apiGridOptions() (err error) {
	var typ, ok = self.RW.GetSegmentUriString(`param`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 202)
	}
	switch typ {
	case `scenario`:
		var options *types.Scenario
		if options, err = types.GetScenario(`Groups`, `GridAdmin`); err != nil {
			return self.RW.ResponseJson(nil, 404, 570, `GridAdmin`, `Groups`)
		}
		return self.RW.ResponseJson(options, 200, 0)
	}
	return self.RW.ResponseJson(nil, 409, 203, typ)
}

////

// ApiObj Управление группами. Детально
func (self *Groups) ApiObj() (err error) {
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
	return self.RW.ResponseJson(nil, 404, 520, `base.Groups.ApiObj`)
}

// apiObjGet Получение группы
func (self *Groups) apiObjGet() (err error) {
	// входящие данные
	var id, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var obj = model.NewGroups(id)
	if err = obj.Model.Load(`Id`); err != nil {
		return self.RW.ResponseJson(nil, 404, 500)
	}
	return self.RW.ResponseJson(obj.Type, 200, 0)
}

// apiObjPut Изменение группы
func (self *Groups) apiObjPut() (err error) {
	// входящие данные
	var id, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var reqObj = new(typDb.Groups)
	if err = self.RW.RequestJsonParse(reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	// загрузка
	var obj = model.NewGroups(id)
	if err = obj.Model.Load(`Id`); err != nil {
		return self.RW.ResponseJson(nil, 404, 500)
	}
	// проверка
	var validMessage map[string]string
	if validMessage, err = obj.Model.Set(reqObj, `All`); err != nil {
		return self.RW.ResponseJson(validMessage, 409, 530, `Groups.All`)

	}
	// сохранение
	if err = obj.Save(`All`, `Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 540, `Groups.All`, obj.Type.Id)
	}
	return self.RW.ResponseJson(nil, 200, 0)
}

// apiObjDelete Удаление группы
func (self *Groups) apiObjDelete() (err error) {
	// входящие данные
	var id, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}

	var obj = model.NewGroups(id)
	if err = obj.Remove(`Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 550, `Groups`, id)
	}
	return self.RW.ResponseJson(nil, 200, 0)
}

// apiObjOptions Опции редактирования группы
func (self *Groups) apiObjOptions() (err error) {
	var options *types.Scenario
	if options, err = types.GetScenario(`Groups`, `All`); err != nil {
		return self.RW.ResponseJson(nil, 404, 570, `All`, `Groups`)
	}
	options.Sample = &typDb.Groups{
		Name: `SampleName`,
	}
	return self.RW.ResponseJson(options, 200, 0)
}
