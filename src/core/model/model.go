// model
// @apidoc Базовая модель
// Базовый инструментарий по работе с моделями
// Сохранение, удаление и загрузка. Приоритет выполнениия память, кеш, БД.
// Валидация данных
package model

import (
	"errors"
	"io/ioutil"
	"reflect"
	"sort"

	"app"
	"lib/logs"
	"lib/uploader"
	"types"
)

// Базовый и вспоготальный функционал для всех моделей
type Model struct {
	model   interface{} // Делегируемая модель
	typ     interface{} // Тип модели
	typName string      // Имя обрабатываемого типа модели
}

func NewModel(typ, model interface{}) *Model {
	var self = new(Model)
	self.model = model
	self.typ = typ
	Value := reflect.ValueOf(self.typ)
	if Value.Kind() == reflect.Ptr {
		Value = Value.Elem()
	}
	self.typName = Value.Type().Name()
	return self
}

// VPropertiesFile Валидация бинарных данных
func (self *Model) VPropertiesFile(propertyName string, value string) (err error) {
	if value != "" {
		//var propertyFile = strings.Replace(propertyName, "File", "", 1)
		resp, err := uploader.Get(value)
		if err == nil {
			logs.Info(123, resp.Data.PathSys)
			data, _ := ioutil.ReadFile(resp.Data.PathSys)
			setProperty(self.typ, propertyName, data, false)
			//setProperty(self.typ, propertyFile, data, false)
			//setProperty(self.typ, propertyName, resp.Data.NameOriginal, false)
			// удаление временного файла
			err = uploader.Delete(value)
			if err != nil {
				logs.Warning(118, value, err)
			}
		}
	}
	return err
}

//func (self *Model) VPropertiesFile(propertyName string, value string) (err error) {
//	if value != "" {
//		var propertyFile = strings.Replace(propertyName, "File", "", 1)
//		resp, err := uploader.Get(value)
//		if err == nil {
//			logs.Info(123, resp.Data.PathSys)
//			data, _ := ioutil.ReadFile(resp.Data.PathSys)
//			setProperty(self.typ, propertyFile, data, false)
//			setProperty(self.typ, propertyName, resp.Data.NameOriginal, false)
//			// удаление временного файла
//			err = uploader.Delete(value)
//			if err != nil {
//				logs.Warning(118, value, err)
//			}
//		}
//	}
//	return err
//}

//0  *uploader.Response {
//1  .  FilePrefix: ""
//2  .  Field: "filename"
//3  .  Data: uploader.data {
//4  .  .  Name: "filename"
//5  .  .  NameOriginal: "Hydrangeas.jpg"
//6  .  .  NameFilePrefix: "1395074891013683800-"
//7  .  .  NameFile: "1395074891013683800-VKUjuLwt3e7HMMbO.jpg"
//8  .  .  Extension: ".jpg"
//9  .  .  PathSys: "C:\\Work\\spl\\upload\\uploader\\1395074891013683800-VKUjuLwt3e7HMMbO.jpg"
//10  .  .  PathWeb: "/1395074891013683800-VKUjuLwt3e7HMMbO.jpg"
//11  .  .  IsPicture: true
//12  .  .  PictureWidth: 1024
//13  .  .  PictureHeight: 768
//14  .  .  Size: 595284
//15  .  .  ContentType: "image/jpeg"
//16  .  }
//17  }

// Set Установка (копирование) передаваемого объекта в исходный
// scenarioName - сценарий опций типа
func (self *Model) Set(modelType interface{}, scenarioName string) (propertyMessage map[string]string, err error) {
	propertyMessage = make(map[string]string)
	// Прямое копирование
	if scenarioName == "" {
		err = types.Copy(modelType, self.typ)
		return propertyMessage, logs.Error(120, err).Error
	}
	objValue := reflect.ValueOf(self.model)
	objValueSet := reflect.ValueOf(modelType)
	if objValueSet.Kind() == reflect.Ptr {
		objValueSet = objValueSet.Elem()
	}

	// Проверка и копирование по сценарию
	var options *types.Scenario
	if options, err = types.GetScenario(self.typName, scenarioName); err != nil {
		return propertyMessage, logs.Error(121, scenarioName, self.typName).Error
	}
	// Общая проверка для сценария
	var method = objValue.MethodByName("VScenario" + scenarioName)
	if method.IsValid() == true {
		var in = []reflect.Value{}
		in = append(in, objValueSet)
		out := method.Call(in)
		if nil != out[0].Interface() {
			err = out[0].Interface().(error)
			return propertyMessage, logs.Error(119, "VScenario"+scenarioName, err).Error
		}
	}
	// Проверка по свойствам
	for i := range options.Property {
		var propertyName = options.Property[i].Name
		var fieldSet = objValueSet.FieldByName(propertyName)
		if fieldSet.IsValid() == false {
			propertyMessage[propertyName] = "property '" + propertyName + "' not exists in type"
			continue
		}
		var validMethod = "VProperty" + propertyName
		// Персональная проверка
		var method = objValue.MethodByName(validMethod)
		if method.IsValid() == true {
			var in = []reflect.Value{}
			in = append(in, reflect.ValueOf(scenarioName))
			in = append(in, fieldSet)
			out := method.Call(in)
			if nil != out[0].Interface() {
				err = out[0].Interface().(error)
				propertyMessage[propertyName] = err.Error()
			}
			continue
		}
		// Присвоение по умолчанию
		if options.Property[i].Readonly == `yes` { // пропускаем только для чтения
			continue
		}
		switch options.Property[i].FormType {
		case "relation", "linkcross":
			// пропускаем работу по кросс связям и отношениям
		case "file":
			if err = self.VPropertiesFile(propertyName, string(fieldSet.Interface().([]byte))); err != nil {
				propertyMessage[propertyName] = err.Error()
			}
		case "img":
			if err = self.VPropertiesFile(propertyName, string(fieldSet.Interface().([]byte))); err != nil {
				propertyMessage[propertyName] = err.Error()
			}
		default:
			if options.Property[i].Required == `yes` {
				err = setProperty(self.typ, propertyName, fieldSet.Interface(), true)
			} else {
				err = setProperty(self.typ, propertyName, fieldSet.Interface(), false)
			}
			if err != nil {
				propertyMessage[propertyName] = err.Error()
			}
		}
	}
	if 0 < len(propertyMessage) {
		err = errors.New(`Error set model`)
	}
	return
}

// Save Сохранение и обновление объекта в памяти
func (self *Model) Save(key string) (err error) {
	// Тип
	typValue := reflect.ValueOf(self.typ)
	if typValue.Kind() == reflect.Ptr {
		typValue = typValue.Elem()
	}
	typField := typValue.FieldByName(key)
	if typField.IsValid() == false {
		return logs.Error(113, key, typValue.Type().Name()).Error
	}
	// Данные
	dataValue := reflect.ValueOf(app.Data)
	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}
	dataField := dataValue.FieldByName(self.typName)
	if dataField.IsValid() == false {
		return logs.Error(124, typValue.Type().Name()).Error
	}
	if dataField.Type().Kind() != reflect.Slice {
		return logs.Error(125, typValue.Type().Name()).Error
	}
	// Сохранение
	if key == `Id` {
		if typField.Interface().(uint64) == 0 { // новый объект (полностью), добавление
			typField.SetUint(GenerateId(self.typName))
			dataField.Set(reflect.Append(dataField, typValue.Addr()))
		} else {
			index := sort.Search(dataField.Len(), func(i int) bool {
				f := dataField.Index(i)
				if f.Kind() == reflect.Ptr {
					f = f.Elem()
				}
				return f.FieldByName(key).Interface().(uint64) >= typField.Interface().(uint64)
			})
			if index < dataField.Len() { // существующий объект, обновление
				f := dataField.Index(index)
				if f.Kind() == reflect.Ptr {
					f = f.Elem()
				}
				if typField.Interface() == f.FieldByName(key).Interface() {
					f.Set(typValue)
				}
			} else { // новый существующий объект, добавление
				// typField.SetUint(GenerateId(self.typName))
				dataField.Set(reflect.Append(dataField, typValue.Addr()))
			}
		}
	} else {
		var flagSave bool
		for i := 0; i < dataField.Len(); i++ {
			f := dataField.Index(i)
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			if typField.Interface() == f.FieldByName(key).Interface() {
				f.Set(typValue)
				flagSave = true
				break
			}
		}
		if flagSave == false {
			dataField.Set(reflect.Append(dataField, typValue.Addr()))
		}
	}
	return
}

// Add Добавление объекта в память (после сохранение в БД)
/*
func (self *Model) Add(key string) (err error) {
	// Тип
	typValue := reflect.ValueOf(self.typ)
	if typValue.Kind() == reflect.Ptr {
		typValue = typValue.Elem()
	}
	typField := typValue.FieldByName(key)
	if typField.IsValid() == false {
		return logs.Error(113, key, typValue.Type().Name()).Error
	}
	// Данные
	dataValue := reflect.ValueOf(self.data)
	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}
	dataField := dataValue.FieldByName(typValue.Type().Name())
	if dataField.IsValid() == false {
		return logs.Error(124, typValue.Type().Name()).Error
	}
	if dataField.Type().Kind() != reflect.Slice {
		return logs.Error(125, typValue.Type().Name()).Error
	}
	// Сохранение
	dataField.Set(reflect.Append(dataField, typValue.Addr()))
	return
}
*/

// Remove Удаление объекта из памяти и БД
func (self *Model) Remove(key string) (err error) {
	// Тип
	typValue := reflect.ValueOf(self.typ)
	if typValue.Kind() == reflect.Ptr {
		typValue = typValue.Elem()
	}
	typField := typValue.FieldByName(key)
	// Данные
	dataValue := reflect.ValueOf(app.Data)
	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}
	dataField := dataValue.FieldByName(self.typName)
	if dataField.IsValid() == false {
		return logs.Error(124, typValue.Type().Name()).Error
	}
	// Срез
	if dataField.Type().Kind() != reflect.Slice {
		return logs.Error(125, typValue.Type().Name()).Error
	}
	if key == `Id` {
		index := sort.Search(dataField.Len(), func(i int) bool {
			f := dataField.Index(i)
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			return f.FieldByName(key).Interface().(uint64) >= typField.Interface().(uint64)
		})
		if index < dataField.Len() {
			f := dataField.Index(index)
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			if typField.Interface() == f.FieldByName(key).Interface() {
				dataField.Set(reflect.AppendSlice(dataField.Slice(0, index), dataField.Slice(index+1, dataField.Len())))
			}
		}

	} else {
		for i := 0; i < dataField.Len(); i++ {
			f := dataField.Index(i)
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			if typField.Interface() == f.FieldByName(key).Interface() {
				//var z = reflect.Zero(typValue.Type())
				//f.Set(z)
				dataField.Set(reflect.AppendSlice(dataField.Slice(0, i), dataField.Slice(i+1, dataField.Len())))
				break
			}
		}
	}
	return
}

// Load Загрузка объекта из БД в память и в объект
func (self *Model) Load(key string) (err error) {
	// Тип
	typValue := reflect.ValueOf(self.typ)
	if typValue.Kind() == reflect.Ptr {
		typValue = typValue.Elem()
	}
	typField := typValue.FieldByName(key)
	// Данные
	dataValue := reflect.ValueOf(app.Data)
	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}
	dataField := dataValue.FieldByName(self.typName)
	if dataField.IsValid() == false {
		return logs.Error(124, typValue.Type().Name()).Error
	}
	// Срез
	if dataField.Type().Kind() != reflect.Slice {
		return logs.Error(125, typValue.Type().Name()).Error
	}
	var flag bool
	if key == `Id` {
		index := sort.Search(dataField.Len(), func(i int) bool {
			f := dataField.Index(i)
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			return f.FieldByName(key).Interface().(uint64) >= typField.Interface().(uint64)
		})
		if index < dataField.Len() {
			f := dataField.Index(index)
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			if typField.Interface() == f.FieldByName(key).Interface() {
				typValue.Set(f)
				flag = true
			}
		}
	} else {
		for i := 0; i < dataField.Len(); i++ {
			f := dataField.Index(i)
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			if typField.Interface() == f.FieldByName(key).Interface() {
				typValue.Set(f)
				flag = true
				break
			}
		}
	}
	if flag == false {
		return logs.Error(159, typValue.Type().Name(), typField.Interface()).Error
	}
	return
}
