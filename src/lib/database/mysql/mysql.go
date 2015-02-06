// Библиотека работы с БД Mysql.
//
// Непосредственная работа с БД. Выполнение запросов.
// Работа с БД Mysql в парадигме ORM.
// Отложенное выполнение запросов к БД.
// Генерация идентификаторов объектов к БД.
package mysql

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"

	"lib"
	"lib/logs"
)

// Стек соединений с БД
var conn = make(map[int8]map[string]*Db)

// служебная переменная для реализации блокировок
var mutexNew sync.Mutex

// Структура по работе с БД
type Db struct {
	Connect mysql.Conn // Конннект к БД
	free    bool       // Статус блокировки (использования)
	time    time.Time  // Дата и время последнего использования коннекта
	logs    bool       // Логирование запросов
}

// New Конструктор соединений с БД
// id выбранная конфигурация
func NewDb(id int8) (obj *Db, err error) {
	var ok bool
	var d *CfgMysql
	// блокировка
	mutexNew.Lock()
	defer mutexNew.Unlock()
	// проверка конфигурации
	if d, ok = cfgMysql[id]; false == ok {
		return nil, logs.Base.Error(800, id).Err
	}
	if _, ok = conn[id]; false == ok {
		conn[id] = make(map[string]*Db)
	}
	// поиск старого (свободного) коннекта
	//fmt.Println("SEARCH CONNECT")
	for len(conn[id]) >= int(cfgMysql[id].CntConn) && 0 < len(conn[id]) {
		for i := range conn[id] {
			if true == conn[id][i].free {
				conn[id][i].free = false
				conn[id][i].time = lib.Time.Now()
				//fmt.Println("FIND CONNECT OK")
				return conn[id][i], nil
			}
		}
	}
	//fmt.Println("NEW CONNECT")
	obj = new(Db)
	if d.Type == `tcp` {
		obj.Connect = mysql.New("tcp", "", fmt.Sprintf("%s:%d", d.Host, d.Port), d.Login, d.Password, d.Name)
	} else {
		obj.Connect = mysql.New("unix", "", d.Socket, d.Login, d.Password, d.Name)
	}
	if err := obj.Connect.Connect(); err != nil {
		return nil, logs.Base.Error(806, id, err).Err
	}
	obj.time = lib.Time.Now()
	obj.free = false
	obj.logs = d.Logs
	//conn[confId] = append(conn[confId], obj)
	conn[id][obj.time.String()] = obj
	// fmt.Println("\nnew coonnect. count connect:", len(conn[id]))
	return obj, nil
}

// Free Освобождение коннекта для других процессов (запросов)
func (self *Db) Free() {
	self.time = lib.Time.Now()
	self.free = true
}

// LoadData Загрузка БД в память
// object объект со свойствами одноименными с таблицами в БД
// Свойства object могут быть хешами либо срезами структур (рекомендуется ссылочные `[]*Users`)
// Свойста структур соответсвуют полям или столбцам в таблицах
// Все свойства должны быть публичны
func (self *Db) SelectData(object interface{}) (err error) {
	objValue := reflect.ValueOf(object)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	num := objValue.NumField()
	for i := 0; i < num; i++ {
		field := objValue.Field(i)
		// определяем имя источника
		var source = objValue.Type().Field(i).Name
		var db = objValue.Type().Field(i).Tag.Get(`db`)
		if db == `-` {
			continue
		}
		if false == field.CanSet() {
			return logs.Base.Error(803, source).Err
		}
		var sql string
		if db == `cross` {
			sql = "SELECT * FROM " + source
		} else {
			sql = "SELECT * FROM " + source + " ORDER BY Id ASC"
		}
		// читаем данные
		switch field.Type().Kind() {
		case reflect.Slice:
			if err = self.SelectSlice(field.Addr().Interface(), sql); err != nil {
				return
			}
		case reflect.Map:
			if err = self.SelectMap(field.Addr().Interface(), sql); err != nil {
				return
			}
		default:
			return logs.Base.Error(812, source).Err
		}
	}
	return err
}

// LoadArray Загрузка списка объектов из БД. Сопоставление свойств структур полям в БД
func (self *Db) SelectMap(ObjectList interface{}, sql string, params ...interface{}) (err error) {
	// рефлеския объекта
	objType := reflect.TypeOf(ObjectList)
	if objType.Kind() != reflect.Ptr {
		return logs.Base.Error(101, objType.String(), sql).Err
	}
	if objType.Elem().Kind() != reflect.Map {
		return logs.Base.Error(813, objType.String(), sql).Err
	}
	var objValue = reflect.MakeMap(objType.Elem())
	var fieldMap = make(map[string]string)
	var obj = reflect.New(objType.Elem().Elem().Elem())
	obj = obj.Elem()
	num := obj.NumField()
	for i := 0; i < num; i++ {
		field := obj.Field(i)
		fieldName := obj.Type().Field(i).Name
		fieldTag := obj.Type().Field(i).Tag.Get(`db`)
		if fieldTag == `-` {
			continue
		}
		if false == field.IsValid() || false == field.CanSet() {
			return logs.Base.Error(804, fieldName, sql).Err
		}
		if fieldTag == `` {
			fieldMap[fieldName] = fieldName
		} else {
			fieldMap[fieldTag] = fieldName
		}
	}
	// запрос
	if strings.LastIndex(sql, " ") == -1 {
		if sql, err = GetQuery(sql); err != nil {
			return
		}
	}
	//sql, params = sqlParse(sql, params)
	var res mysql.Result
	var rows []mysql.Row
	var row mysql.Row
	var stm mysql.Stmt
	defer func() {
		if stm != nil {
			stm.Delete()
		}
	}()
	stm, err = self.Connect.Prepare(sql)
	if err != nil {
		return logs.Base.Error(801, sql, err).Err
	}
	if len(params) > 0 {
		stm.Bind(params...)
	}
	rows, res, err = stm.Exec()
	if err != nil {
		return logs.Base.Error(802, sql, err).Err
	}
	if self.logs == true {
		logs.Base.Info(999, sql)
	}
	// пустой результат
	if 0 == len(rows) {
		return
	}
	// соответствие структуры типа и структуры запроса
	var fieldRes = make(map[string]string)
	for _, field := range res.Fields() {
		if _, ok := fieldMap[field.Name]; ok == false {
			return logs.Base.Error(803, field.Name, sql).Err
		}
		fieldRes[field.Name] = fieldMap[field.Name]
	}
	// наполнение результата
	for _, row = range rows {
		objectId := uint64(0)
		var obj = reflect.New(objType.Elem().Elem().Elem())
		obj = obj.Elem()
		for fAlias, fName := range fieldRes {
			// logs.Dumper(fAlias, res.Map(fAlias))
			prop := obj.FieldByName(fName)
			// заносим полученные значения согласно типам свойств
			switch prop.Type().String() {
			case "bool":
				prop.SetBool(row.Bool(res.Map(fAlias)))
			case "int8", "int64", "int32", "int16":
				prop.SetInt(row.Int64(res.Map(fAlias)))
			case "uint64":
				val := row.Uint64(res.Map(fAlias))
				prop.SetUint(val)
				if fAlias == `Id` {
					objectId = val
				}
			case "uint8", "uint32", "uint16":
				prop.SetUint(row.Uint64(res.Map(fAlias)))
			case "float32", "float64":
				prop.SetFloat(row.Float(res.Map(fAlias)))
			case "string":
				prop.SetString(row.Str(res.Map(fAlias)))
			case "[]uint8":
				prop.SetBytes(row.Bin(res.Map(fAlias)))
			case "[]string":
				val := row.Str(res.Map(fAlias))
				slc := reflect.MakeSlice(prop.Type(), 0, 0)
				if "" != val {
					list := strings.Split(val, ",")
					for i := range list {
						slc = reflect.Append(slc, reflect.ValueOf(list[i]))
					}
				}
				prop.Set(slc)
			case "time.Time":
				prop.Set(reflect.ValueOf(row.Time(res.Map(fAlias), lib.Time.Location)))
			case "time.Duration":
				prop.Set(reflect.ValueOf(row.Duration(res.Map(fAlias))))
			}
		}
		if objectId == 0 {
			objectId++
		}
		objValue.SetMapIndex(reflect.ValueOf(objectId), obj.Addr())
	}
	reflect.ValueOf(ObjectList).Elem().Set(objValue)
	return
}

// LoadArray Загрузка списка объектов из БД. Сопоставление свойств структур полям в БД
func (self *Db) SelectSlice(ObjectList interface{}, sql string, params ...interface{}) (err error) {
	// рефлеския объекта
	objType := reflect.TypeOf(ObjectList)
	if objType.Kind() != reflect.Ptr {
		return logs.Base.Error(101, objType.String(), sql).Err
	}
	if objType.Elem().Kind() != reflect.Slice {
		return logs.Base.Error(807, objType.String(), sql).Err
	}
	var objValue = reflect.MakeSlice(objType.Elem(), 0, 0)
	var fieldMap = make(map[string]string)
	var obj = reflect.New(objType.Elem().Elem().Elem())
	obj = obj.Elem()
	num := obj.NumField()
	for i := 0; i < num; i++ {
		field := obj.Field(i)
		fieldName := obj.Type().Field(i).Name
		fieldTag := obj.Type().Field(i).Tag.Get(`db`)
		if fieldTag == `-` {
			continue
		}
		if false == field.IsValid() || false == field.CanSet() {
			return logs.Base.Error(804, fieldName, sql).Err
		}
		if fieldTag == `` {
			fieldMap[fieldName] = fieldName
		} else {
			fieldMap[fieldTag] = fieldName
		}
	}
	// запрос
	if strings.LastIndex(sql, " ") == -1 {
		if sql, err = GetQuery(sql); err != nil {
			return
		}
	}
	//sql, params = sqlParse(sql, params)
	var res mysql.Result
	var rows []mysql.Row
	var row mysql.Row
	var stm mysql.Stmt
	defer func() {
		if stm != nil {
			stm.Delete()
		}
	}()
	stm, err = self.Connect.Prepare(sql)
	if err != nil {
		return logs.Base.Error(801, sql, err).Err
	}
	if len(params) > 0 {
		stm.Bind(params...)
	}
	rows, res, err = stm.Exec()
	if err != nil {
		return logs.Base.Error(802, sql, err).Err
	}
	if self.logs == true {
		logs.Base.Info(999, sql)
	}
	// пустой результат
	if 0 == len(rows) {
		return
	}
	// соответствие структуры типа и структуры запроса
	var fieldRes = make(map[string]string)
	for _, field := range res.Fields() {
		if _, ok := fieldMap[field.Name]; ok == false {
			return logs.Base.Error(803, field.Name, sql).Err
		}
		fieldRes[field.Name] = fieldMap[field.Name]
	}
	// наполнение результата
	for _, row = range rows {

		var obj = reflect.New(objType.Elem().Elem().Elem())
		obj = obj.Elem()
		for fAlias, fName := range fieldRes {
			// logs.Dumper(fAlias, res.Map(fAlias))
			prop := obj.FieldByName(fName)
			// заносим полученные значения согласно типам свойств
			switch prop.Type().String() {
			case "bool":
				prop.SetBool(row.Bool(res.Map(fAlias)))
			case "int8", "int64", "int32", "int16":
				prop.SetInt(row.Int64(res.Map(fAlias)))
			case "uint8", "uint64", "uint32", "uint16":
				prop.SetUint(row.Uint64(res.Map(fAlias)))
			case "float32", "float64":
				prop.SetFloat(row.Float(res.Map(fAlias)))
			case "string":
				prop.SetString(row.Str(res.Map(fAlias)))
			case "[]uint8":
				prop.SetBytes(row.Bin(res.Map(fAlias)))
			case "[]string":
				val := row.Str(res.Map(fAlias))
				slc := reflect.MakeSlice(prop.Type(), 0, 0)
				if "" != val {
					list := strings.Split(val, ",")
					for i := range list {
						slc = reflect.Append(slc, reflect.ValueOf(list[i]))
					}
				}
				prop.Set(slc)
			case "time.Time":
				prop.Set(reflect.ValueOf(row.Time(res.Map(fAlias), lib.Time.Location)))
			case "time.Duration":
				prop.Set(reflect.ValueOf(row.Duration(res.Map(fAlias))))
			}
		}
		objValue = reflect.Append(objValue, obj.Addr())
	}
	reflect.ValueOf(ObjectList).Elem().Set(objValue)
	return
}

// Load Загрузка объекта из БД. Сопоставление свойств структур полям в БД
// Object = *TypeName
func (self *Db) Select(Object interface{}, sql string, params ...interface{}) (err error) {
	// рефлеския объекта
	var objValue = reflect.ValueOf(Object)
	if objValue.Kind() != reflect.Ptr {
		return logs.Base.Error(101, objValue.Type().String()+`:`+sql).Err
	}
	if objValue.IsNil() == true {
		return logs.Base.Error(103, objValue.Type().String()+`:`+sql).Err
	}
	objValue = objValue.Elem()
	var fieldMap = make(map[string]string)
	num := objValue.NumField()
	for i := 0; i < num; i++ {
		field := objValue.Field(i)
		fieldName := objValue.Type().Field(i).Name
		fieldTag := objValue.Type().Field(i).Tag.Get(`db`)
		if fieldTag == `-` {
			continue
		}
		if false == field.IsValid() || false == field.CanSet() {
			return logs.Base.Error(804, fieldName, sql).Err
		}
		if fieldTag == `` {
			fieldMap[fieldName] = fieldName
		} else {
			fieldMap[fieldTag] = fieldName
		}
	}
	// запрос
	if strings.LastIndex(sql, " ") == -1 {
		if sql, err = GetQuery(sql); err != nil {
			return
		}
	}
	//sql, params = sqlParse(sql, params)
	var res mysql.Result
	var rows []mysql.Row
	var row mysql.Row
	var stm mysql.Stmt
	defer func() {
		if stm != nil {
			stm.Delete()
		}
	}()
	stm, err = self.Connect.Prepare(sql)
	if err != nil {
		return logs.Base.Error(801, sql, err).Err
	}
	if len(params) > 0 {
		stm.Bind(params...)
	}
	rows, res, err = stm.Exec()
	if err != nil {
		return logs.Base.Error(802, sql, err).Err
	}
	if self.logs == true {
		logs.Base.Info(999, sql)
	}
	// пустой результат
	if 0 == len(rows) {
		return logs.Base.Info(805, sql).Err
	}
	// соответствие структуры типа и структуры запроса
	var fieldRes = make(map[string]string)
	for _, field := range res.Fields() {
		if _, ok := fieldMap[field.Name]; ok == false {
			return logs.Base.Error(804, field.Name, sql).Err
		}
		fieldRes[field.Name] = fieldMap[field.Name]
	}
	// наполнение результата
	row = rows[0]
	for fAlias, fName := range fieldRes {
		// logs.Dumper(fAlias, res.Map(fAlias))
		prop := objValue.FieldByName(fName)
		// заносим полученные значения согласно типам свойств
		switch prop.Type().String() {
		case "bool":
			prop.SetBool(row.Bool(res.Map(fAlias)))
		case "int8", "int64", "int32", "int16":
			prop.SetInt(row.Int64(res.Map(fAlias)))
		case "uint8", "uint64", "uint32", "uint16":
			prop.SetUint(row.Uint64(res.Map(fAlias)))
		case "float32", "float64":
			prop.SetFloat(row.Float(res.Map(fAlias)))
		case "string":
			prop.SetString(row.Str(res.Map(fAlias)))
		case "[]uint8":
			prop.SetBytes(row.Bin(res.Map(fAlias)))
		case "[]string":
			val := row.Str(res.Map(fAlias))
			slc := reflect.MakeSlice(prop.Type(), 0, 0)
			if "" != val {
				list := strings.Split(val, ",")
				for i := range list {
					slc = reflect.Append(slc, reflect.ValueOf(list[i]))
				}
			}
			prop.Set(slc)
		case "time.Time":
			prop.Set(reflect.ValueOf(row.Time(res.Map(fAlias), lib.Time.Location)))
		case "time.Duration":
			prop.Set(reflect.ValueOf(row.Duration(res.Map(fAlias))))
		}
	}
	return
}

// SaveInsert Сохранение объекта в БД. Сопоставление свойств структур полям в БД
func (self *Db) Insert(Object interface{}, source string, properties ...map[string]string) (insertId uint64, err error) {
	// инициализация
	var stm mysql.Stmt
	defer func() {
		if stm != nil {
			stm.Delete()
		}
	}()
	var res mysql.Result
	var sql string
	var paramList []interface{}
	// запрос
	objValue := reflect.ValueOf(Object)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	num := objValue.NumField()
	for i := 0; i < num; i++ {
		prop := objValue.Field(i)
		propName := objValue.Type().Field(i).Name
		propTag := objValue.Type().Field(i).Tag.Get(`db`)
		// пропускаем защищенные от изменения и не указанные свойства
		if `-` == propTag || false == prop.CanSet() {
			continue
		}
		//	пропускаем свойства которые не надо сохранять
		if 0 < len(properties) {
			if _, ok := properties[0][propName]; ok == false {
				continue
			}
		}
		// определение алиаса для БД
		if propTag != `` {
			propName = propTag
		}
		switch prop.Type().String() {
		case "bool":
			val := prop.Interface().(bool)
			if false == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, 1)
			}
			sql += "`" + propName + "` = ?, "
		case "int64":
			val := prop.Interface().(int64)
			if 0 == val {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "int32":
			val := prop.Interface().(int32)
			if 0 == val {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "int16":
			val := prop.Interface().(int16)
			if 0 == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "int8":
			val := prop.Interface().(int8)
			if 0 == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "uint64":
			val := prop.Interface().(uint64)
			if 0 == val {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "uint32":
			val := prop.Interface().(uint32)
			if 0 == val {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "uint16":
			val := prop.Interface().(uint16)
			if 0 == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "uint8":
			val := prop.Interface().(uint8)
			if 0 == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "float32":
			val := prop.Interface().(float32)
			if 0 == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		case "float64":
			val := prop.Interface().(float64)
			if 0 == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		case "string":
			val := prop.Interface().(string)
			if "" == val {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		case "[]string":
			val := strings.Join(prop.Interface().([]string), ",")
			if "" == val {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		case "[]uint8":
			val := prop.Interface().([]uint8)
			if 0 == len(val) {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		case "time.Time":
			val := prop.Interface().(time.Time)
			if val.IsZero() == true {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		case "time.Duration":
			val := prop.Interface().(time.Duration)
			if val <= 0 {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		}
	}
	sql = "INSERT " + source + " SET " + sql[:len(sql)-2]
	// запрос и параметры
	if stm, err = self.Connect.Prepare(sql); err != nil {
		return insertId, logs.Base.Error(801, sql, err).Err
	}
	stm.Bind(paramList...)
	// выполнение запроса
	if _, res, err = stm.Exec(); err != nil {
		return insertId, logs.Base.Error(802, sql, err).Err
	}
	if self.logs == true {
		logs.Base.Info(999, sql)
	}
	insertId = res.InsertId()
	return
}

// SaveUpdate Изменение объекта в БД. Сопоставление свойств структур полям в БД
func (self *Db) Update(Object interface{}, source, key string, properties ...map[string]string) (affectedRow uint64, err error) {
	// инициализация
	var stm mysql.Stmt
	defer func() {
		if stm != nil {
			stm.Delete()
		}
	}()
	var res mysql.Result
	var sql string
	var paramList []interface{}
	// запрос
	objValue := reflect.ValueOf(Object)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	field := objValue.FieldByName(key)
	if field.IsValid() == false {
		return affectedRow, logs.Base.Error(808, key, source).Err
	}
	num := objValue.NumField()
	for i := 0; i < num; i++ {
		prop := objValue.Field(i)
		propName := objValue.Type().Field(i).Name
		propTag := objValue.Type().Field(i).Tag.Get(`db`)
		// пропускаем защищенные от изменения и не указанные свойства
		if `-` == propTag || false == prop.CanSet() {
			continue
		}
		//	пропускаем свойства которые не надо сохранять
		if 0 < len(properties) {
			if _, ok := properties[0][propName]; ok == false {
				continue
			}
		}
		// определение алиаса для БД
		if propTag != `` {
			propName = propTag
		}
		switch prop.Type().String() {
		case "bool":
			val := prop.Interface().(bool)
			if false == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, 1)
			}
			sql += "`" + propName + "` = ?, "
		case "int64":
			val := prop.Interface().(int64)
			if 0 == val {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "int32":
			val := prop.Interface().(int32)
			if 0 == val {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "int16":
			val := prop.Interface().(int16)
			if 0 == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "int8":
			val := prop.Interface().(int8)
			if 0 == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "uint64":
			val := prop.Interface().(uint64)
			if 0 == val {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "uint32":
			val := prop.Interface().(uint32)
			if 0 == val {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "uint16":
			val := prop.Interface().(uint16)
			if 0 == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "uint8":
			val := prop.Interface().(uint8)
			if 0 == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, prop.Interface())
			}
			sql += "`" + propName + "` = ?, "
		case "float32":
			val := prop.Interface().(float32)
			if 0 == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		case "float64":
			val := prop.Interface().(float64)
			if 0 == val {
				paramList = append(paramList, 0)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		case "string":
			val := prop.Interface().(string)
			if "" == val {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		case "[]string":
			val := strings.Join(prop.Interface().([]string), ",")
			if "" == val {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		case "[]uint8":
			val := prop.Interface().([]uint8)
			if 0 == len(val) {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		case "time.Time":
			val := prop.Interface().(time.Time)
			if val.IsZero() == true {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		case "time.Duration":
			val := prop.Interface().(time.Duration)
			if val <= 0 {
				paramList = append(paramList, nil)
			} else {
				paramList = append(paramList, val)
			}
			sql += "`" + propName + "` = ?, "
		}
	}
	sql = "UPDATE " + source + " SET " + sql[:len(sql)-2] + " WHERE `" + key + "` = ?"
	paramList = append(paramList, field.Interface())
	// запрос и параметры
	if stm, err = self.Connect.Prepare(sql); err != nil {
		return affectedRow, logs.Base.Error(801, sql, err).Err
	}
	stm.Bind(paramList...)
	// выполнение запроса
	if _, res, err = stm.Exec(); err != nil {
		return affectedRow, logs.Base.Error(802, sql, err).Err
	}
	if self.logs == true {
		logs.Base.Info(999, sql)
	}
	affectedRow = res.AffectedRows()
	return
}

// QueryByte Пользовательский запрос на обновление в БД (из дамп файлов).
func (self *Db) QueryByte(data []byte) (messages []string, err error) {
	var result mysql.Result
	_, result, err = self.Connect.Query(string(data))
	if err != nil {
		return nil, logs.Base.Error(809).Err
	} else {
		if result != nil {
			for result.MoreResults() == true {
				if message := result.Message(); message != "" {
					messages = append(messages, message)
				}
				result, err = result.NextResult()
				if err != nil {
					err = nil
					break
				}
			}
		}
	}
	return
}

// Query Пользовательский запрос на обновление в БД.
func (self *Db) Query(sql string, params ...interface{}) (err error) {
	if strings.LastIndex(sql, " ") == -1 {
		if sql, err = GetQuery(sql); err != nil {
			return
		}
	}
	// инициализация
	var stm mysql.Stmt
	defer func() {
		if stm != nil {
			stm.Delete()
		}
	}()
	// запрос
	//sql, params = sqlParse(sql, params)
	stm, err = self.Connect.Prepare(sql)
	if err != nil {
		return logs.Base.Error(801, sql, err).Err
	}
	stm.Bind(params...)
	if _, _, err = stm.Exec(); err != nil {
		return logs.Base.Error(802, sql, err).Err
	}
	if self.logs == true {
		logs.Base.Info(999, sql)
	}
	return
}

// sqlParse Разбор входных параметров
// корректировка параметров для запросов с параметрами IN
func sqlParse(sql string, param []interface{}) (string, []interface{}) {
	// инициализируем данные для встапвки и находим их местоположения в запросе
	var str []string
	var strPos []int
	var params []interface{}
	for i := range param {
		var typ = reflect.TypeOf(param[i])
		switch typ.String() {
		case "[]string":
			d := "'" + strings.Join(param[i].([]string), "', '") + "'"
			str = append(str, d)
			strPos = append(strPos, i)
		case "[]int64":
			var s []string
			for _, elm := range param[i].([]int64) {
				s = append(s, fmt.Sprintf("%d", elm))
			}
			d := strings.Join(s, ", ")
			str = append(str, d)
			strPos = append(strPos, i)
		case "[]int32":
			var s []string
			for _, elm := range param[i].([]int32) {
				s = append(s, fmt.Sprintf("%d", elm))
			}
			d := strings.Join(s, ", ")
			str = append(str, d)
			strPos = append(strPos, i)
		case "[]int16":
			var s []string
			for _, elm := range param[i].([]int16) {
				s = append(s, fmt.Sprintf("%d", elm))
			}
			d := strings.Join(s, ", ")
			str = append(str, d)
			strPos = append(strPos, i)
		case "[]int8":
			var s []string
			for _, elm := range param[i].([]int8) {
				s = append(s, fmt.Sprintf("%d", elm))
			}
			d := strings.Join(s, ", ")
			str = append(str, d)
			strPos = append(strPos, i)
		case "[]uint64":
			var s []string
			for _, elm := range param[i].([]uint64) {
				s = append(s, fmt.Sprintf("%d", elm))
			}
			d := strings.Join(s, ", ")
			str = append(str, d)
			strPos = append(strPos, i)
		case "[]uint32":
			var s []string
			for _, elm := range param[i].([]uint32) {
				s = append(s, fmt.Sprintf("%d", elm))
			}
			d := strings.Join(s, ", ")
			str = append(str, d)
			strPos = append(strPos, i)
		case "[]uint16":
			var s []string
			for _, elm := range param[i].([]uint16) {
				s = append(s, fmt.Sprintf("%d", elm))
			}
			d := strings.Join(s, ", ")
			str = append(str, d)
			strPos = append(strPos, i)
		case "[]uint8":
			return sql, param
			var s []string
			for _, elm := range param[i].([]uint8) {
				s = append(s, fmt.Sprintf("%d", elm))
			}
			d := strings.Join(s, ", ")
			str = append(str, d)
			strPos = append(strPos, i)
		case "[]float64":
			var s []string
			for _, elm := range param[i].([]float64) {
				s = append(s, fmt.Sprintf("%f", elm))
			}
			d := strings.Join(s, ", ")
			str = append(str, d)
			strPos = append(strPos, i)
		case "[]float32":
			var s []string
			for _, elm := range param[i].([]float32) {
				s = append(s, fmt.Sprintf("%f", elm))
			}
			d := strings.Join(s, ", ")
			str = append(str, d)
			strPos = append(strPos, i)
		default:
			params = append(params, param[i])
		}
	}
	// вставляем данные срезов
	var data = strings.Split(sql, "?")
	var dataNew []string
	var pos int
	for i := range strPos {
		dataNew = append(dataNew, strings.Join(data[pos:strPos[i]+1], "?"))
		dataNew = append(dataNew, str[i])
		pos = strPos[i] + 1
	}
	dataNew = append(dataNew, strings.Join(data[pos:], "?"))
	return strings.Join(dataNew, ""), params
}

func (self *Db) Call(typ interface{}, nameCall string, params ...interface{}) (err error) {
	return
}

func GetQuery(index string) (q string, err error) {
	l := strings.Split(index, `/`)
	//
	var section uint64
	if section, err = strconv.ParseUint(l[len(l)-1], 0, 64); err != nil {
		return q, logs.Base.Error(810, index).Err
	}
	index = strings.ToLower(strings.Join(l[:len(l)-1], `/`))
	if _, ok := Query[index]; ok == false {
		return q, logs.Base.Error(810, index).Err
	}
	if int(section) < len(Query[index]) {
		return Query[index][int(section)], err
	}
	return q, logs.Base.Error(810, index).Err
}

/*
     0  []*mysql.Field (len = 23) {
    36  .  ...
    37  .  3: *mysql.Field {
    38  .  .  Catalog: "def"
    39  .  .  Db: "Test"
    40  .  .  Table: "Test"
    41  .  .  OrgTable: "Test"
    42  .  .  Name: "LastName"
    43  .  .  OrgName: "LastName"
    44  .  .  DispLen: 64
    45  .  .  Flags: 16392
    46  .  .  Type: 8
    47  .  .  Scale: 0
    48  .  }
    49  .  ...
   277  }
*/
