// Инженеринг БД в типы.
//
// Создание програмного кода (структур объектов) на основе БД
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"lib/database/mysql"
	"lib/logs"
)

func main() {

	var cfgMysql = make(map[int8]*mysql.CfgMysql)
	cfgMysql[0] = new(mysql.CfgMysql)
	cfgMysql[0].Driver = "mysql"
	cfgMysql[0].Host = "localhost"
	cfgMysql[0].Port = 3306
	cfgMysql[0].Type = "tcp"
	cfgMysql[0].Socket = "/var/run/mysql/mysql.sock"
	cfgMysql[0].Name = "zegota"
	cfgMysql[0].Login = "root"
	cfgMysql[0].Password = "root"
	cfgMysql[0].Charset = "utf-8"
	cfgMysql[0].TimeOut = 5
	cfgMysql[0].CntConn = 1
	cfgMysql[0].Updates, _ = os.Getwd()

	mysql.InitMysql(cfgMysql)

	EngineTypes(`typesEngine`, 0)
	logs.Dumper()
}

// Таблица
type table struct {
	Name            string    // Название таблицы
	Engine          string    // Комментарий
	Version         int       // Комментарий
	Row_format      int       // Комментарий
	Rows            int       // Комментарий
	Avg_row_length  int       // Комментарий
	Data_length     int       // Комментарий
	Max_data_length int       // Комментарий
	Index_length    int       // Комментарий
	Data_free       int       // Комментарий
	Auto_increment  int       // Комментарий
	Create_time     time.Time // Комментарий
	Update_time     time.Time // Комментарий
	Check_time      time.Time // Комментарий
	Collation       string    // Комментарий
	Checksum        string    // Комментарий
	Create_options  string    // Комментарий
	Comment         string    // Комментарий
}

// Поля таблицы
type field struct {
	Field      string // имя поля
	Type       string // тип
	Collation  string // тип
	Null       string // не нулевое значение
	Key        string // индекс
	Extra      string // автоинкремент
	Default    string // дефолтовое значение
	Privileges string // дефолтовое значение
	Comment    string // комметарий
}

var modelBlank = `// Модель TABLECOMMENT
package model

import (
	"core"
	"core/model"
	typDb "FOLDER/types"
)

type TABLENAME struct {
	Model *model.Model
	Type  *typDb.TABLENAME
	Db    *model.Db
}

// NewTABLENAME Создание объекта модели
func NewTABLENAME(id uint64) *TABLENAME {
	var self = new(TABLENAME)
	self.Type = new(typDb.TABLENAME)
	self.Type.Id = id
	self.Model = model.NewModel(self.Type, self)
	self.Db = model.NewDb(self.Type, core.Config.Main.TypeDb, 0)
	return self
}

// NewTABLENAMEType Создание объекта модели
func NewTABLENAMEType(typ *typDb.TABLENAME) *TABLENAME {
	var self = new(TABLENAME)
	self.Type = typ
	self.Db = model.NewDb(self.Type, core.Config.Main.TypeDb, 0)
	return self
}

// VScenarioAll Пример общей проверки для сценария
func (self *TABLENAME) VScenarioAll(typ typDb.TABLENAME) (err error) {
	return
}

// VPropertySample Пример проверки свойства
func (self *TABLENAME) VPropertySample(scenario string, value uint64) (err error) {
	self.Type.Id = value
	return
}
`

// EngineTypes Создание типов по струтуре БД (id конфига, folder папка в которой буду созданы пакеты типов)
// Если тип существет в ФС он будет пропущен и не будет перезаписан
func EngineTypes(folder string, id int8) {
	var dirTypes, dirModel, path, path1, str, dataType, dataModel, pkgType string
	var db *mysql.Db
	var fp os.FileInfo
	var err error
	db, err = mysql.NewDb(id)
	if err != nil {
		fmt.Println("Ошибка коннекта к БД")
		return
	}
	defer db.Free()
	dir, _ := os.Getwd()
	dir += `/src/` + folder
	dirModel = dir + "/model"
	dirTypes = dir + "/types"

	var tables []*table
	if err := db.LoadArray(&tables, "SHOW TABLE STATUS"); err != nil {
		fmt.Println("Ошибка запроса получения таблиц из БД", err.Error())
		return
	}

	for i := range tables {
		// - вычисление пакета и типа

		/////
		//if tables[i].Name == "controllers" {
		//	tables[i].Name = "Controllers"
		//}
		//if tables[i].Name == "groups" {
		//	tables[i].Name = "Groups"
		//}
		//if tables[i].Name == "groupsuri" {
		//	tables[i].Name = "GroupsUri"
		//}
		//if tables[i].Name == "uri" {
		//	tables[i].Name = "Uri"
		//}
		//if tables[i].Name == "users" {
		//	tables[i].Name = "Users"
		//}
		/////

		pkgType = tables[i].Name

		//pkg := make([]rune, 0)
		//for i, b := range []rune(tables[i].Name) {
		//	if i == 0 {
		//		pkg = append(pkg, b)
		//		continue
		//	}
		//	if 90 < b {
		//		pkg = append(pkg, b)
		//		continue
		//	}
		//	break
		//}
		//pkgName = string(pkg)
		/////
		//pkgName = `access`

		var flagTemplatelibTime = false
		var dataInit = []string{}
		var dataInitAll = []string{}

		dataType = "// Тип " + tables[i].Comment + "\npackage db\n\nimport (\n\t\"types\"\n\tTEMPLATELIBTIME\n)\n\n"
		dataType += "// " + tables[i].Comment + "\n"
		dataType += "type " + pkgType + " struct {\n"

		var fields []*field
		db.LoadArray(&fields, "SHOW FULL COLUMNS FROM "+tables[i].Name)
		//logs.Dumper(fields)
		for i := range fields {
			fieldName := fields[i].Field
			fieldComment := fields[i].Comment
			fieldTypeFull := fields[i].Type
			fieldTypeDb := strings.Split(fieldTypeFull, "(")[0]
			fieldRequired := ""
			if fields[i].Null == "NO" {
				fieldRequired = "`yes`"
			}
			fieldDefault := fields[i].Default
			if fieldName != `Id` {
				dataInitAll = append(dataInitAll, fieldName)
			}
			// сопоставление типов из БД mysql в типы Го
			var fieldType, fieldForm string
			switch fieldTypeDb {
			case "bigint":
				if "Id" == fieldName || 0 < strings.LastIndex(fieldTypeFull, "unsigned") || 0 < strings.LastIndex(fieldName, "Id") {
					fieldType = "uint64"
				} else {
					fieldType = "int64"
				}
				fieldForm = "number"
				if 0 < strings.LastIndex(fieldName, "_Id") {
					fieldForm = "link"
				}
			case "int", "mediumint":
				if "Id" == fieldName || 0 < strings.LastIndex(fieldTypeFull, "unsigned") || 0 < strings.LastIndex(fieldName, "Id") {
					fieldType = "uint32"
				} else {
					fieldType = "int32"
				}
				fieldForm = "number"
				if 0 < strings.LastIndex(fieldName, "_Id") {
					fieldForm = "link"
				}
			case "smallint":
				if "Id" == fieldName || 0 < strings.LastIndex(fieldTypeFull, "unsigned") || 0 < strings.LastIndex(fieldName, "Id") {
					fieldType = "uint16"
				} else {
					fieldType = "int16"
				}
				fieldForm = "number"
			case "tinyint":
				if "tinyint(1)" == fieldTypeFull {
					fieldType = "bool"
					fieldForm = "bool"
					fieldDefault = ""
				} else if "Id" == fieldName || 0 < strings.LastIndex(fieldTypeFull, "unsigned") || 0 < strings.LastIndex(fieldName, "Id") {
					fieldType = "uint8"
					fieldForm = "number"
				} else {
					fieldType = "int8"
					fieldForm = "number"
				}
			case "float":
				fieldType = "float32"
				fieldForm = "float"
			case "decimal", "double":
				fieldType = "float64"
				fieldForm = "float"
			case "char", "varchar", "tinytext":
				fieldType = "string"
				fieldForm = "text"
				if "Name" == fieldName {
					fieldRequired = "`yes`"
				}
			case "text":
				fieldType = "string"
				fieldForm = "textarea"
			case "mediumtext", "longtext":
				fieldType = "string"
				fieldForm = "content"
			case "enum":
				fieldType = "string"
				if fieldRequired == "" {
					fieldForm = "select"
				} else {
					fieldForm = "radio"
				}
			case "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
				fieldType = "[]byte"
				fieldForm = "file"
			case "set":
				fieldType = "[]string"
				fieldForm = "checkbox"
			case "date":
				fieldType = "time.Time"
				fieldForm = "date"
				flagTemplatelibTime = true
			case "datetime":
				fieldType = "time.Time"
				fieldForm = "datetime"
				flagTemplatelibTime = true
			case "time":
				fieldType = "time.Duration"
				fieldForm = "time"
				flagTemplatelibTime = true
			default:
				fmt.Printf("Неизвестный тип поля: %s.%s\n", tables[i].Name, fieldName)
				continue
			}
			dataType += "\t" + fieldName + "\t" + fieldType + "\t// " + fieldComment + "\n"

			str = "\t\t\tName:\t\t\"" + fieldName + "\",\n"
			str += "\t\t\tTitle:\t\t\"" + fieldComment + "\",\n"
			if fieldName == "Id" {
				str += "\t\t\tReadonly:\t\t\t`yes`,\n"
			}
			if fieldRequired != `` {
				str += "\t\t\tRequired:\t\t" + fieldRequired + ",\n"
			}
			if fieldDefault != `` {
				str += "\t\t\tDefault:\t\t\"" + fieldDefault + "\",\n"
			}
			str += "\t\t\tFormType:\t\t\"" + fieldForm + "\",\n"
			if fieldTypeDb == "enum" || fieldTypeDb == "set" {
				l := strings.Split(fieldTypeFull, "(")
				l = strings.Split(l[1], ")")
				l[0] = strings.Replace(l[0], "'", "", -1)
				l = strings.Split(l[0], ",")
				str += "\t\t\tEnumSet:\t\tmap[string]string{"
				for i := range l {
					str += "\"" + l[i] + "\": \"" + l[i] + "\", "
				}

				str = str[0 : len(str)-2]
				str += "},\n"
			}
			dataInit = append(dataInit, str)
		}
		dataType += "}\n\n"
		if flagTemplatelibTime == true {
			dataType = strings.Replace(dataType, "TEMPLATELIBTIME", "\"time\"", 1)
		} else {
			dataType = strings.Replace(dataType, "TEMPLATELIBTIME", "", 1)
		}
		dataType += "func init() {\n"
		dataType += "\t// Набор сценариев для типа\n"
		dataType += "\ttypes.SetScenario(`" + tables[i].Name + "`, map[string]types.Scenario{\n"

		dataType += "\t\t`root`: types.Scenario{\n"
		dataType += "\t\t\tName:\t\t`" + tables[i].Comment + "`,\n"
		dataType += "\t\t\tDescription: `Базовая конфигурация всех свойств для всех сценарией указанного типа`,\n"
		dataType += "\t\t\tProperty: []types.Property{{\n"
		dataType += strings.Join(dataInit, "\t\t}, {\n")
		dataType += "\t\t\t}},\n"
		dataType += "\t\t},\n"
		dataType += "\t\t`All`: types.Scenario{\n"
		dataType += "\t\t\tDescription: `Все свойства`,\n"
		dataType += "\t\t\tProperty: []types.Property{{\n"
		dataType += "\t\t\t\tName: \t`" + strings.Join(dataInitAll, "`,\n\t\t\t}, {\n\t\t\t\tName:\t`") + "`,\n"
		dataType += "\t\t\t}},\n"
		dataType += "\t\t},\n"
		dataType += "\t})\n"
		dataType += "}\n"

		// Сохранение типа
		path = dirTypes
		if err := os.MkdirAll(path, 0777); err != nil {
			fmt.Printf("Ошибка создания папки пакета структуры: %s\n%v\n", path, err)
		}
		path += `/` + pkgType + "T.go"
		fp, _ = os.Stat(path1)
		if nil == fp {
			if err := ioutil.WriteFile(path, []byte(dataType), 0777); err != nil {
				fmt.Printf("Ошибка создания файла пакета структуры: %s\n%v\n", path, err)
			}
		}

		// Сохранение модели
		dataModel = strings.Replace(modelBlank, "FOLDER", folder, -1)
		dataModel = strings.Replace(dataModel, "TABLECOMMENT", tables[i].Comment, -1)
		dataModel = strings.Replace(dataModel, "TABLENAME", tables[i].Name, -1)
		path = dirModel
		if err := os.MkdirAll(path, 0777); err != nil {
			fmt.Printf("Ошибка создания папки пакета структуры: %s\n%v\n", path, err)
		}
		path += `/` + pkgType + "M.go"
		fp, _ = os.Stat(path1)
		if nil == fp {
			if err := ioutil.WriteFile(path, []byte(dataModel), 0777); err != nil {
				fmt.Printf("Ошибка создания файла пакета структуры: %s\n%v\n", path, err)
			}
		}
		//
	}
	fmt.Println("Модели созданы: " + dirModel)
	fmt.Println("Типы созданы: " + dirTypes)
}
