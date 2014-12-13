// Модель Роутинг
package model

import (
	"app"
	"core"
	"core/model"
	"io/ioutil"
	"lib/database"
	"lib/logs"
	"os"
	typDb "types/db"
)

type Uri struct {
	Model *model.Model
	Type  *typDb.Uri
	Db    *model.Db
}

// NewUri Создание объекта модели
func NewUri(id uint64) *Uri {
	var self = new(Uri)
	self.Type = new(typDb.Uri)
	self.Type.Id = id
	self.Model = model.NewModel(self.Type, self)
	self.Db = model.NewDb(self.Type, core.Config.Main.UseDb, 0)
	return self
}

// NewUri Создание объекта модели
func NewUriType(typ typDb.Uri) *Uri {
	var self = new(Uri)
	self.Type = &typ
	self.Model = model.NewModel(self.Type, self)
	self.Db = model.NewDb(self.Type, core.Config.Main.UseDb, 0)
	return self
}

// VScenarioSql Пример общей валидации для сценария
func (self *Uri) VScenarioAll(typ typDb.Uri) (err error) {
	return
}

// VPropertySample Пример валидации свойства
func (self *Uri) VPropertySample(scenario string, value uint64) (err error) {
	self.Type.Id = value
	return
}

func (self *Uri) VPropertyUri(scenario string, value string) (err error) {
	if value == `` {
		return logs.Error(201).Error
	}
	for i := range app.Data.Uri {
		if app.Data.Uri[i].Uri == value && (self.Type.Id == 0 || self.Type.Id != app.Data.Uri[i].Id) {
			return logs.Error(1, value).Error
		}
	}
	self.Type.Uri = value
	return
}

func (self *Uri) VPropertyContentType(scenario string, value string) (err error) {
	if value == `` {
		value = `text/html`
	}
	self.Type.ContentType = value
	return
}

func (self *Uri) VPropertyContentEncode(scenario string, value string) (err error) {
	if value == `` {
		value = `utf-8`
	}
	self.Type.ContentEncode = value
	return
}

func (self *Uri) VPropertyPosition(scenario string, value int32) (err error) {
	if self.Type.Id == 0 {
		self.Type.Position = app.Data.MaxPostion[`Uri`]
		app.Data.MaxPostion[`Uri`]++
	}
	return
}

func (self *Uri) Save(scenario, key string) (err error) {
	if err = self.Db.Save(scenario, key); err != nil {
		return
	}
	if err = self.Model.Save(key); err != nil {
		return
	}
	//if err = self.RemoveControllers(); err != nil {
	//	return
	//}
	//if err = self.SaveControllers(); err != nil {
	//	return
	//}
	return
}

func (self *Uri) Remove(key string) (err error) {
	if err = self.Db.Remove(key); err != nil {
		return
	}
	if err = self.Model.Remove(key); err != nil {
		return
	}
	if err = self.RemoveControllers(); err != nil {
		return
	}
	if err = self.RemoveGroups(); err != nil {
		return
	}
	return
}

////

// UriContentUpdate обновление контента uri по отношению к файловой системе
func UriContentUpdate(rw *core.RW, u *typDb.Uri) (err error) {
	if u.Id == 0 {
		return
	}
	var fi os.FileInfo
	var path = rw.DocumentRoot + u.Uri
	if fi, err = os.Stat(path); err != nil {
		// return logs.Error(146, path, err).Error
		return
	}
	var con []byte
	if fi.ModTime().Sub(u.ContentTime) > 0 {
		if con, err = ioutil.ReadFile(path); nil != err {
			return logs.Error(146, path, err).Error
		} else {
			u.Content = con
			u.ContentTime = fi.ModTime()
			// сохраняем в БД
			if core.Config.Main.UseDb > 0 {
				var db database.DbFace
				if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
					return
				}
				// сохранение
				_, err = db.Update(u, `Uri`, `Id`)
				db.Free()
				//self.Db.Save(`All`, `Id`)
			}
		}
	}
	return
}

////

func (self *Uri) LoadControllers() (tied, unbound []*typDb.GroupsUri) {
	if self.Type.Id == 0 {
		return
	}
	// находим связанные контроллеры
	var dataMap = make(map[uint64]typDb.GroupsUri)
	for i := range app.Data.GroupsUri {
		if app.Data.GroupsUri[i].Uri_Id == self.Type.Id && app.Data.GroupsUri[i].Controllers_Id > 0 {
			dataMap[app.Data.GroupsUri[i].Controllers_Id] = *app.Data.GroupsUri[i]
		}
	}
	// формируем данные
	for i := range app.Data.Controllers {
		if elm, ok := dataMap[app.Data.Controllers[i].Id]; ok == true {
			elm.Name = app.Data.Controllers[i].Name
			tied = append(tied, &elm)
		} else {
			unbound = append(unbound, &typDb.GroupsUri{
				Controllers_Id: app.Data.Controllers[i].Id,
				Name:           app.Data.Controllers[i].Name,
			})
		}
	}
	return
}

func (self *Uri) SaveControllers(tied []*typDb.GroupsUri) (err error) {
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

func (self *Uri) RemoveControllers() (err error) {
	if self.Type.Id == 0 {
		return
	}
	var slice []*typDb.GroupsUri
	var sliceId []uint64
	for i := range app.Data.GroupsUri {
		if app.Data.GroupsUri[i].Uri_Id == self.Type.Id && app.Data.GroupsUri[i].Controllers_Id > 0 {
			sliceId = append(sliceId, app.Data.GroupsUri[i].Controllers_Id)
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
		db.Query(`base/uri/3`, self.Type.Id, sliceId[i])
	}
	return
}

////

func (self *Uri) LoadGroups() (tied, unbound []*typDb.GroupsUri) {
	if self.Type.Id == 0 {
		return
	}
	// находим связанные группы
	var dataMap = make(map[uint64]typDb.GroupsUri)
	for i := range app.Data.GroupsUri {
		if app.Data.GroupsUri[i].Uri_Id == self.Type.Id && app.Data.GroupsUri[i].Groups_Id > 0 {
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

func (self *Uri) SaveGroups(tied []*typDb.GroupsUri) (err error) {
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

func (self *Uri) RemoveGroups() (err error) {
	if self.Type.Id == 0 {
		return
	}
	var slice []*typDb.GroupsUri
	var sliceId []uint64
	for i := range app.Data.GroupsUri {
		if app.Data.GroupsUri[i].Uri_Id == self.Type.Id && app.Data.GroupsUri[i].Groups_Id > 0 {
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
		db.Query(`base/uri/4`, self.Type.Id, sliceId[i])
	}
	return
}

//
