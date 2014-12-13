package model

import (
	"errors"
	"reflect"
	"strings"

	"lib/database"
	"types"
)

// Работа модели с БД
type Db struct {
	typ        interface{} // Тип модели
	typName    string      // Имя обрабатываемого типа модели
	useDb      int64       // Используемая БД
	numConnect int8        // Номер конфига к БД
}

func NewDb(typ interface{}, useDb int64, numConnect int8) *Db {
	var self = new(Db)
	self.typ = typ
	Value := reflect.ValueOf(self.typ)
	if Value.Kind() == reflect.Ptr {
		Value = Value.Elem()
	}
	self.typName = Value.Type().Name()
	self.useDb = useDb
	self.numConnect = numConnect
	return self
}

func (self *Db) Load(scenarioName, key string) (err error) {
	// получение свойств
	var scenario *types.Scenario
	if scenario, err = types.GetScenario(self.typName, scenarioName); err != nil {
		return
	}
	var fields = make([]string, 0)
	for i := range scenario.Property {
		fields = append(fields, scenario.Property[i].Name)
	}
	// получение ключевого поля (Id)
	param, err := getProperty(self.typ, key)
	if err != nil {
		return
	}
	var sql = "SELECT `" + strings.Join(fields, "`, `") + "` FROM " + self.typName + " WHERE `" + key + "` = ?"
	var db database.DbFace
	if db, err = database.NewDb(self.useDb, self.numConnect); err != nil {
		return
	}
	err = db.Select(self.typ, sql, param)
	db.Free()
	return
}

func (self *Db) Save(scenarioName, key string) (err error) {
	// получение свойств
	var scenario *types.Scenario
	if scenario, err = types.GetScenario(self.typName, scenarioName); err != nil {
		return
	}
	var fields = make(map[string]string)
	for i := range scenario.Property {
		fields[scenario.Property[i].Name] = scenario.Property[i].Name
	}
	// Тип
	typValue := reflect.ValueOf(self.typ)
	if typValue.Kind() == reflect.Ptr {
		typValue = typValue.Elem()
	}
	typField := typValue.FieldByName(key)
	if typField.IsValid() == false {
		return errors.New(`Объекты типа ` + typValue.Type().Name() + ` не имеет свойства ` + key)
	}
	// коннект
	var db database.DbFace
	if db, err = database.NewDb(self.useDb, self.numConnect); err != nil {
		return
	}
	// сохранение
	if key == `Id` && typField.Interface().(uint64) == 0 {
		var id uint64
		if id, err = db.Insert(self.typ, self.typName, fields); id > 0 {
			typField.SetUint(id)
		}
	} else {
		_, err = db.Update(self.typ, self.typName, key, fields)

	}
	db.Free()
	return
}

func (self *Db) Remove(key string) (err error) {
	var sql = "DELETE FROM `" + self.typName + "` WHERE `" + key + "` = ?"
	param, err := getProperty(self.typ, key)
	if err != nil {
		return err
	}
	var db database.DbFace
	if db, err = database.NewDb(self.useDb, self.numConnect); err != nil {
		return
	}
	err = db.Query(sql, param)
	db.Free()
	return
}
