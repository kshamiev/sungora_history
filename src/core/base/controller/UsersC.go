package controller

import (
	"core"
	"core/base/model"
	"types"
	typDb "types/db"
)

type Users struct {
	Session     *core.Session      // Сессия
	RW          *core.RW           // Управление вводом и выводом
	Controllers *typDb.Controllers // Соответсвующий контроллер из области данных по строковому ид.
}

// Создание контроллера
func NewUsers(rw *core.RW, s *core.Session, c *typDb.Controllers) interface{} {
	var self = new(Users)
	self.RW = rw
	self.Session = s
	self.Controllers = c
	return self
}

////

// ApiGrid Управление пользователями. Список
func (self *Users) ApiGrid() (err error) {
	// диспетчер методов
	switch self.RW.Request.Method {
	case "GET":
		return self.apiGridGet()
	case "POST":
		return self.apiGridPost()
	case "OPTIONS":
		return self.apiGridOptions()
	}
	return self.RW.ResponseJson(nil, 404, 520, `base.Users.ApiGrid`)
}

// apiGridGet Список пользователей
func (self *Users) apiGridGet() (err error) {
	var data []*typDb.Users
	// постраничный вывод
	var page, ok = self.RW.GetSegmentUriInt(`param`)
	if ok == false {
		page = 1
	}
	if data, err = model.GetUsersGrid(int(page)); err != nil {
		return self.RW.ResponseJson(nil, 500, 500)
	}
	return self.RW.ResponseJson(data, 200, 0)
}

// apiGridPost Добавление пользователя
func (self *Users) apiGridPost() (err error) {
	// входящие данные
	var reqObj = new(typDb.Users)
	if err = self.RW.RequestJsonParse(reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	// проверка
	var obj = model.NewUsers(0)
	var validMessage map[string]string
	if validMessage, err = obj.Model.Set(reqObj, `All`); err != nil {
		return self.RW.ResponseJson(validMessage, 409, 530, `Users.All`)

	}
	// сохранение
	if err = obj.Save(`All`, `Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 560, `Users`)
	}

	return self.RW.ResponseJson(obj.Type.Id, 200, 0)
}

// apiGridOptions Опции списка пользователей
func (self *Users) apiGridOptions() (err error) {
	var typ, ok = self.RW.GetSegmentUriString(`param`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 575)
	}
	switch typ {
	case `scenario`:
		var options *types.Scenario
		if options, err = types.GetScenario(`Users`, `GridAdmin`); err != nil {
			return self.RW.ResponseJson(nil, 404, 104, `Users`, `GridAdmin`)
		}
		return self.RW.ResponseJson(options, 200, 0)
	}
	return self.RW.ResponseJson(nil, 409, 570, typ)
}

////

// ApiObj Управление пользователями. Детально
func (self *Users) ApiObj() (err error) {
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
	return self.RW.ResponseJson(nil, 404, 520, `base.Users.ApiObj`)
}

// apiObjGet Получение пользователя
func (self *Users) apiObjGet() (err error) {
	// входящие данные
	var id, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var obj = model.NewUsers(id)
	if err = obj.Model.Load(`Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 500)
	}
	return self.RW.ResponseJson(obj.Type, 200, 0)
}

// apiObjPut Изменение пользователя
func (self *Users) apiObjPut() (err error) {
	// входящие данные
	var id, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var reqObj = new(typDb.Users)
	if err = self.RW.RequestJsonParse(reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	// загрузка
	var obj = model.NewUsers(id)
	if err = obj.Model.Load(`Id`); err != nil {
		return self.RW.ResponseJson(nil, 404, 500)
	}
	// проверка
	var validMessage map[string]string
	if validMessage, err = obj.Model.Set(reqObj, `All`); err != nil {
		return self.RW.ResponseJson(validMessage, 409, 530, `Users.All`)

	}
	// сохранение
	if err = obj.Save(`All`, `Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 540, `Users.All`, obj.Type.Id)
	}
	return self.RW.ResponseJson(nil, 200, 0)
}

// apiObjDelete Удаление пользователя
func (self *Users) apiObjDelete() (err error) {
	// входящие данные
	var id, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}

	var obj = model.NewUsers(id)
	if err = obj.Remove(`Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 550, `Users`, id)
	}
	return self.RW.ResponseJson(nil, 200, 0)
}

// apiObjOptions Опции редактирования пользователя
func (self *Users) apiObjOptions() (err error) {
	var options *types.Scenario
	if options, err = types.GetScenario(`Users`, `All`); err != nil {
		return self.RW.ResponseJson(nil, 404, 570, `All`, `Users`)
	}
	options.Sample = &typDb.Users{
		Login: `SampleLogin`,
		Email: `samplename@domain.ru`,
	}
	return self.RW.ResponseJson(options, 200, 0)
}

////

// ApiObjGroups Управление группами пользователя
func (self *Users) ApiObjGroups() (err error) {
	// диспетчер методов
	switch self.RW.Request.Method {
	case "GET":
		return self.apiObjGroupsGet()
	case "PUT":
		return self.apiObjGroupsPut()
	case "OPTIONS":
		return self.apiObjGroupsOptions()
	}
	return self.RW.ResponseJson(nil, 404, 520, `base.Users.ApiObjGroups`)
}

// apiObjGroupsGet Группы пользователя (получение)
func (self *Users) apiObjGroupsGet() (err error) {
	// входящие данные
	var pid, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var obj = model.NewUsers(pid)
	var tied, unbound = obj.LoadGroups()
	return self.RW.ResponseJson(map[string]interface{}{`tied`: tied, `unbound`: unbound}, 200, 0)
}

// apiObjGroupsPut Группы пользователя (изменение)
func (self *Users) apiObjGroupsPut() (err error) {
	// входящие данные
	var pid, ok = self.RW.GetSegmentUriUint(`id`)
	if ok == false {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var reqObj []*typDb.UsersGroups
	if err = self.RW.RequestJsonParse(&reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}
	var obj = model.NewUsers(pid)
	if err = obj.RemoveGroups(); err != nil {
		return self.RW.ResponseJson(nil, 409, 550, `base.Users.apiObjGroupsPut`, obj.Type.Id)
	}
	if err = obj.SaveGroups(reqObj); err != nil {
		return self.RW.ResponseJson(nil, 409, 540, `base.Users.apiObjGroupsPut`, obj.Type.Id)
	}
	return self.RW.ResponseJson(nil, 200, 0)
}

// apiObjGroupsOptions Опции групп пользователя
func (self *Users) apiObjGroupsOptions() (err error) {
	var options *types.Scenario
	if options, err = types.GetScenario(`UsersGroups`, `All`); err != nil {
		return self.RW.ResponseJson(nil, 404, 570, `All`, `UsersGroups`)
	}
	return self.RW.ResponseJson(options, 200, 0)
}
