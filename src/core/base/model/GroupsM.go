// Модель Группы
package model

import (
	"core"
	"core/model"
	typDb "types/db"
)

type Groups struct {
	Model *model.Model
	Type  *typDb.Groups
	Db    *model.Db
}

// NewGroups Создание объекта модели
func NewGroups(id uint64) *Groups {
	var self = new(Groups)
	self.Type = new(typDb.Groups)
	self.Type.Id = id
	self.Model = model.NewModel(self.Type, self)
	self.Db = model.NewDb(self.Type, core.Config.Main.UseDb, 0)
	return self
}

// NewGroups Создание объекта модели
func NewGroupsType(typ typDb.Groups) *Groups {
	var self = new(Groups)
	self.Type = &typ
	self.Model = model.NewModel(self.Type, self)
	self.Db = model.NewDb(self.Type, core.Config.Main.UseDb, 0)
	return self
}

// VScenarioAll Пример общей валидации для сценария
func (self *Groups) VScenarioAll(typ typDb.Groups) (err error) {
	return
}

// VPropertySample Пример валидации свойства
func (self *Groups) VPropertySample(scenario string, value uint64) (err error) {
	self.Type.Id = value
	return
}

func (self *Groups) Save(scenario, key string) (err error) {
	if err = self.Db.Save(scenario, key); err != nil {
		return
	}
	if err = self.Model.Save(key); err != nil {
		return
	}
	return
}

func (self *Groups) Remove(key string) (err error) {
	if err = self.Db.Remove(key); err != nil {
		return
	}
	if err = self.Model.Remove(key); err != nil {
		return
	}
	return
}
