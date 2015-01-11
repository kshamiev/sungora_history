package types

import (
	"reflect"

	"lib/logs"
)

// Копирование объектов. Копируются данные свойств оригинальные имена которых совпадают в исходном и целевых объектах.
// Копирвоание срезов, хешей, единичных объектов (!по указателям!).
func Copy(typInput, typTarget interface{}) (err error) {
	return CopyFilter(typInput, typTarget, func(src, dst interface{}) bool {
		return true
	})
}

// Копирование объектов. Копируются данные свойств оригинальные имена которых совпадают в исходном и целевых объектах.
// Копирвоание срезов, хешей, единичных объектов (!по указателям!).
// Фильтрующая функция получает исходный и результирующий объект. В случае если функция вернет true объект копируется, если false то объект пропускается.
// Фильтрующая функция может модифицировать результирующий объект.
func CopyFilter(typInput, typTarget interface{}, f func(src, dst interface{}) bool) (err error) {
	var resp bool

	// инициализация рефлексии исходного объекта
	objType := reflect.TypeOf(typInput)
	if objType.Kind() != reflect.Ptr {
		return logs.Base.Error(100, objType.String()).Err
	}
	refModel := reflect.ValueOf(typInput)

	// инициализация рефлексии целевого объекта
	objTypeTarget := reflect.TypeOf(typTarget)
	if objTypeTarget.Kind() != reflect.Ptr {
		return logs.Base.Error(101, objTypeTarget.String()).Err
	}
	refModelTarget := reflect.ValueOf(typTarget)

	var objValue reflect.Value
	if objTypeTarget.Elem().Kind() == reflect.Map {
		objValue = reflect.MakeMap(objTypeTarget.Elem())
	} else if objTypeTarget.Elem().Kind() == reflect.Slice {
		objValue = reflect.MakeSlice(objTypeTarget.Elem(), 0, 0)
	} else if objTypeTarget.Elem().Kind() == reflect.Struct {
		return copyStructure(refModel, refModelTarget)
	}

	// Наполнение
	refModel = refModel.Elem()
	for _, index := range refModel.MapKeys() {
		// исходный объект
		elm := refModel.MapIndex(index)
		// целевой объект
		obj := reflect.New(objTypeTarget.Elem().Elem().Elem())
		// Функция фильтрации при копировании
		resp = f(elm.Interface(), obj.Interface())
		if resp == false {
			continue
		}
		// копирование
		if err = copyStructure(elm, obj); err != nil {
			return
		}
		if objValue.Type().Kind() == reflect.Map {
			objValue.SetMapIndex(index, obj)
		} else if objValue.Type().Kind() == reflect.Slice {
			objValue = reflect.Append(objValue, obj)
		}
	}
	refModelTarget.Elem().Set(objValue)
	return
}

// Непосредственное копирование данных свойств исходного объекта в целевой
func copyStructure(typInput, typTarget reflect.Value) (err error) {
	if typInput.Kind() == reflect.Ptr {
		typInput = typInput.Elem()
	}
	if typTarget.Kind() == reflect.Ptr {
		typTarget = typTarget.Elem()
	}

	//typInput = typInput.Elem()
	if typInput.CanAddr() == false {
		return logs.Base.Error(102, typInput.Type().String()).Err
	}
	//typTarget = typTarget.Elem()
	if typTarget.CanAddr() == false {
		return logs.Base.Error(103, typTarget.Type().String()).Err
	}

	// наполнение
	num := typInput.NumField()
	for i := 0; i < num; i++ {
		// исходный объект - свойство
		field := typInput.Field(i)
		fieldName := typInput.Type().Field(i).Name
		t := field.Type().String()
		// целевой объект - свойство
		prop := typTarget.FieldByName(fieldName)
		if false == prop.CanSet() { // пропускаем не совпадающие свойства по оригинальному названию. Что не явялется ошибкой в данном случае
			continue
		}
		// присвоение по умолчанию
		switch t {
		case "bool":
			prop.SetBool(field.Interface().(bool))
		case "int8":
			prop.SetInt(int64(field.Interface().(int8)))
		case "int16":
			prop.SetInt(int64(field.Interface().(int16)))
		case "int32":
			prop.SetInt(int64(field.Interface().(int32)))
		case "int64":
			prop.SetInt(int64(field.Interface().(int64)))
		case "uint8":
			prop.SetUint(uint64(field.Interface().(uint8)))
		case "uint16":
			prop.SetUint(uint64(field.Interface().(uint16)))
		case "uint32":
			prop.SetUint(uint64(field.Interface().(uint32)))
		case "uint64":
			prop.SetUint(uint64(field.Interface().(uint64)))
		case "float32", "float64":
			prop.SetFloat(field.Interface().(float64))
		case "string":
			prop.SetString(field.Interface().(string))
		case "[]uint8":
			prop.SetBytes(field.Interface().([]byte))
			//case "[]string", "time.Time", "time.Duration":
		default:
			prop.Set(field)
		}
	}
	return
}
