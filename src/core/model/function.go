package model

import (
	"lib/logs"
	"reflect"
	"sync"
	"time"
)

// getProperty получение значения свойства объекта
func getProperty(typ interface{}, property string) (value interface{}, err error) {
	objValue := reflect.ValueOf(typ)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	prop := objValue.FieldByName(property)
	if prop.IsValid() == false {
		return value, logs.Error(1, "у объекта "+objValue.Type().String()+" остутсвует свойство "+property).Error
	}
	return prop.Interface(), err
}

// setProperty уставнока свойства объекту модели
func setProperty(typ interface{}, property string, value interface{}, required bool) (err error) {
	objValue := reflect.ValueOf(typ)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	prop := objValue.FieldByName(property)
	if prop.IsValid() == false {
		return logs.Error(1, "у объекта "+objValue.Type().String()+" остутсвует свойство "+property).Error
	}
	t := prop.Type().String()
	switch t {
	case "bool":
		prop.SetBool(value.(bool))
	case "int8":
		if value.(int8) == 0 {
			if required == true {
				return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
			}
			prop.SetInt(0)
		} else {
			prop.SetInt(int64(value.(int8)))
		}
	case "int16":
		if value.(int16) == 0 {
			if required == true {
				return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
			}
			prop.SetInt(0)
		} else {
			prop.SetInt(int64(value.(int16)))
		}
	case "int32":
		if value.(int32) == 0 {
			if required == true {
				return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
			}
			prop.SetInt(0)
		} else {
			prop.SetInt(int64(value.(int32)))
		}
	case "int64":
		if value.(int64) == 0 {
			if required == true {
				return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
			}
			prop.SetInt(0)
		} else {
			prop.SetInt(value.(int64))
		}
	case "uint8":
		if value.(uint8) == 0 {
			if required == true {
				return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
			}
			prop.SetUint(0)
		} else {
			prop.SetUint(uint64(value.(uint8)))
		}
	case "uint16":
		if value.(uint16) == 0 {
			if required == true {
				return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
			}
			prop.SetUint(0)
		} else {
			prop.SetUint(uint64(value.(uint16)))
		}
	case "uint32":
		if value.(uint32) == 0 {
			if required == true {
				return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
			}
			prop.SetUint(0)
		} else {
			prop.SetUint(uint64(value.(uint32)))
		}
	case "uint64":
		if value.(uint64) == 0 {
			if required == true {
				return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
			}
			prop.SetUint(0)
		} else {
			prop.SetUint(value.(uint64))
		}
	case "float64":
		if required == true && 0 == value.(float64) {
			return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
		}
		if value == 0 {
			prop.SetFloat(0)
		} else {
			prop.SetFloat(value.(float64))
		}
	case "string":
		if required == true && "" == value.(string) {
			return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
		}
		prop.SetString(value.(string))
	case "[]string":
		if required == true && 0 == len(value.([]string)) {
			return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
		}
		prop.Set(reflect.ValueOf(value))
	case "time.Time":
		if required == true && "0001-01-01 00:00:00 +0000 UTC" == value.(time.Time).String() {
			return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
		}
		prop.Set(reflect.ValueOf(value))
	case "time.Duration":
		if required == true && "0" == value.(time.Duration).String() {
			return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
		}
		prop.Set(reflect.ValueOf(value))
	case "[]uint8":
		if required == true && 0 == len(value.([]byte)) {
			return logs.Error(1, "своство '"+property+"' обязательно для заполнения").Error
		}
		prop.SetBytes(value.([]byte))
	default:
		prop.Set(reflect.ValueOf(value))
	}
	return
}

var mutexId sync.Mutex
var storageId = make(map[string]uint64)

// GenerateId Генерация идентификаторов объектов (если БД не используется)
func GenerateId(source string) (id uint64) {
	// Блокирвка
	mutexId.Lock()
	defer mutexId.Unlock()
	if _, ok := storageId[source]; ok == false {
		storageId[source] = 1
	} else {
		storageId[source]++
	}
	return storageId[source]
}
