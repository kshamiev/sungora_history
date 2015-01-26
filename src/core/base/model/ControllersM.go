// Модель Контроллеры
package model

import (
	"app"
	"core"
	"core/model"
	"lib/database"
	typDb "types/db"
)

type Controllers struct {
	Model *model.Model
	Type  *typDb.Controllers
	Db    *model.Db
}

// NewControllers Создание объекта модели
func NewControllers(id uint64) *Controllers {
	var self = new(Controllers)
	self.Type = new(typDb.Controllers)
	self.Type.Id = id
	self.Model = model.NewModel(self.Type, self)
	self.Db = model.NewDb(self.Type, core.Config.Main.UseDb, 0)
	return self
}

// NewControllers Создание объекта модели
func NewControllersType(typ typDb.Controllers) *Controllers {
	var self = new(Controllers)
	self.Type = &typ
	self.Model = model.NewModel(self.Type, self)
	self.Db = model.NewDb(self.Type, core.Config.Main.UseDb, 0)
	return self
}

// VScenarioSql Пример общей валидации для сценария
func (self *Controllers) VScenarioAll(typ typDb.Controllers) (err error) {
	return
}

// VProperty Пример валидации свойства
func (self *Controllers) VPropertySample(scenario string, value uint64) (err error) {
	self.Type.Id = value
	return
}

func (self *Controllers) Save(scenario, key string) (err error) {
	if err = self.Db.Save(scenario, key); err != nil {
		return
	}
	if err = self.Model.Save(key); err != nil {
		return
	}
	return
}

func (self *Controllers) Remove(key string) (err error) {
	if err = self.Db.Remove(key); err != nil {
		return
	}
	if err = self.Model.Remove(key); err != nil {
		return
	}
	if err = self.RemoveGroups(); err != nil {
		return
	}
	return
}

////

func (self *Controllers) LoadGroups() (tied, unbound []*typDb.GroupsUri) {
	if self.Type.Id == 0 {
		return
	}
	// находим связанные группы
	var dataMap = make(map[uint64]typDb.GroupsUri)
	for i := range app.Data.GroupsUri {
		if app.Data.GroupsUri[i].Controllers_Id == self.Type.Id && app.Data.GroupsUri[i].Groups_Id > 0 {
			dataMap[app.Data.GroupsUri[i].Groups_Id] = *app.Data.GroupsUri[i]
		}
	}
	// формируем данные
	for i := range app.Data.Groups {
		if elm, ok := dataMap[app.Data.Groups[i].Id]; ok == true {
			elm.Name = app.Data.Groups[i].Name
			tied = append(tied, &elm)
		} else {
			unbound = append(unbound, &typDb.GroupsUri{
				Groups_Id: app.Data.Groups[i].Id,
				Name:      app.Data.Groups[i].Name,
			})
		}
	}
	return
}

func (self *Controllers) SaveGroups(tied []*typDb.GroupsUri) (err error) {
	if self.Type.Id == 0 || len(tied) == 0 {
		return
	}
	app.Data.GroupsUri = append(app.Data.GroupsUri, tied...)
	// сохраняем в БД
	var db database.DbFace
	if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
		return
	}
	defer db.Free()
	for i := range tied {
		db.Insert(tied[i], `GroupsUri`)
	}
	return
}

func (self *Controllers) RemoveGroups() (err error) {
	if self.Type.Id == 0 {
		return
	}
	var slice []*typDb.GroupsUri
	var sliceId []uint64
	for i := range app.Data.GroupsUri {
		if app.Data.GroupsUri[i].Controllers_Id == self.Type.Id && app.Data.GroupsUri[i].Groups_Id > 0 {
			sliceId = append(sliceId, app.Data.GroupsUri[i].Groups_Id)
			//app.Data.GroupsUri = append(app.Data.GroupsUri[:i], app.Data.GroupsUri[i+1:]...)
		} else {
			slice = append(slice, app.Data.GroupsUri[i])
		}
	}
	app.Data.GroupsUri = slice
	// удаление из БД
	var db database.DbFace
	if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
		return
	}
	defer db.Free()
	for i := range sliceId {
		db.Query(`base/controllers/3`, self.Type.Id, sliceId[i])
	}
	return
}

//
