// Модель Пользователи
package model

import (
	"app"
	"core"
	"core/model"
	"lib"
	"lib/database"
	typDb "types/db"
)

type Users struct {
	Model *model.Model
	Type  *typDb.Users
	Db    *model.Db
}

// NewUsers Создание объекта модели
func NewUsers(id uint64) *Users {
	var self = new(Users)
	self.Type = new(typDb.Users)
	self.Type.Id = id
	self.Model = model.NewModel(self.Type, self)
	self.Db = model.NewDb(self.Type, core.Config.Main.UseDb, 0)
	return self
}

// NewUsers Создание объекта модели
func NewUsersType(typ typDb.Users) *Users {
	var self = new(Users)
	self.Type = &typ
	self.Model = model.NewModel(self.Type, self)
	self.Db = model.NewDb(self.Type, core.Config.Main.UseDb, 0)
	return self
}

func (self *Users) Save(scenario, key string) (err error) {
	if err = self.Db.Save(scenario, key); err != nil {
		return
	}
	if err = self.Model.Save(key); err != nil {
		return
	}
	return
}

// VScenarioSql Пример общей валидации для сценария
func (self *Users) VScenarioAll(typ typDb.Users) (err error) {
	return
}

// VPropertyLogin Проверка логина
func (self *Users) VPropertyLogin(scenario string, value string) bool {
	switch scenario {
	case `Registration`:
		var flag = false
		for _, val := range app.Data.Users {
			if val.Login == value {
				flag = true
				break
			}
		}
		if flag == false {
			self.Type.Login = value
		} else {
			return false
		}
	default:
		self.Type.Login = value
	}
	return true
}

// VPropertyEmail Проверка Email
func (self *Users) VPropertyEmail(scenario string, value string) bool {
	switch scenario {
	case `Registration`:
		var flag = false
		for i := range app.Data.Users {
			if app.Data.Users[i].Email == value {
				flag = true
				break
			}
		}
		if flag == false {
			self.Type.Email = value
		} else {
			return false
		}
	default:
		self.Type.Email = value
	}
	return true
}

// VScenarioProfile Пример общей валидации для сценария
func (self *Users) VScenarioProfile(typ typDb.Users) bool {
	if typ.Password != typ.PasswordR {
		return false
	}
	return true
}

// VPropertyPassword Формирование пароля
func (self *Users) VPropertyPassword(scenario string, value string) (err error) {
	self.Type.Password = lib.String.CreatePasswordHash(value)
	return
}

func (self *Users) Remove(key string) (err error) {
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

func (self *Users) LoadGroups() (tied, unbound []*typDb.UsersGroups) {
	if self.Type.Id == 0 {
		return
	}
	// находим связанные группы
	var dataMap = make(map[uint64]typDb.UsersGroups)
	for i := range app.Data.UsersGroups {
		if app.Data.UsersGroups[i].Users_Id == self.Type.Id && app.Data.UsersGroups[i].Groups_Id > 0 {
			dataMap[app.Data.UsersGroups[i].Groups_Id] = *app.Data.UsersGroups[i]
		}
	}
	// формируем данные
	for i := range app.Data.Groups {
		if elm, ok := dataMap[app.Data.Groups[i].Id]; ok == true {
			elm.Name = app.Data.Groups[i].Name
			tied = append(tied, &elm)
		} else {
			unbound = append(unbound, &typDb.UsersGroups{
				Groups_Id: app.Data.Groups[i].Id,
				Name:      app.Data.Groups[i].Name,
			})
		}
	}
	return
}

func (self *Users) SaveGroups(tied []*typDb.UsersGroups) (err error) {
	if self.Type.Id == 0 || len(tied) == 0 {
		return
	}
	app.Data.UsersGroups = append(app.Data.UsersGroups, tied...)
	// сохраняем в БД
	var db database.DbFace
	if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
		return
	}
	defer db.Free()
	for i := range tied {
		db.Query(`base/users/0`, self.Type.Id, tied[i].Groups_Id)
	}
	return
}

func (self *Users) RemoveGroups() (err error) {
	if self.Type.Id == 0 {
		return
	}
	var slice []*typDb.UsersGroups
	var sliceId []uint64
	for i := range app.Data.UsersGroups {
		if app.Data.UsersGroups[i].Users_Id == self.Type.Id {
			sliceId = append(sliceId, app.Data.UsersGroups[i].Groups_Id)
			//app.Data.GroupsUri = append(app.Data.GroupsUri[:i], app.Data.GroupsUri[i+1:]...)
		} else {
			slice = append(slice, app.Data.UsersGroups[i])
		}
	}
	app.Data.UsersGroups = slice
	// удаление из БД
	var db database.DbFace
	if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
		return
	}
	defer db.Free()
	for i := range sliceId {
		db.Query(`base/users/1`, self.Type.Id, sliceId[i])
	}
	return
}

//
