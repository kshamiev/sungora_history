// Генерация описаний прототипов и методов конвертации типа в обе стороны
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/kshamiev/sungora/pb/modelcore"
)

const (
	tplPkg       = "pb"
	tplProtoExt  = ".proto"
	tplDoNotEdit = "DO NOT EDIT"
)

var source = map[string][]interface{}{
	"modelcore": {
		&modelcore.User{},
		&modelcore.Order{},
		&modelcore.Role{},
	},
}

func main() {
	var err error
	var data []byte
	var proto, tplPFull, tplMFull, tplP, tplM string

	for pkg := range source {
		// анализируем типы и формируем proto
		tplPFull = "\n"
		tplMFull = "\n"
		for _, t := range source[pkg] {
			if tplP, tplM, err = ParseType(t); err != nil {
				log.Fatal(err)
			}
			tplPFull += tplP
			tplMFull += tplM
		}

		// формируем proto файлы
		if data, err = ioutil.ReadFile(tplPkg + "/" + pkg + tplProtoExt); err == nil {
			proto = string(data)
		} else {
			if data, err = ioutil.ReadFile(tplPkg + "/generate/tpl.proto"); err != nil {
				log.Fatal(err)
			}
			proto = string(data)
			proto = strings.ReplaceAll(proto, "TPLpackage", tplPkg)
			proto = strings.ReplaceAll(proto, "TPLservice", strings.Title(pkg))
		}

		list := strings.Split(proto, tplDoNotEdit)
		list[1] = tplPFull
		proto = strings.Join(list, tplDoNotEdit)

		if err = ioutil.WriteFile(tplPkg+"/"+pkg+tplProtoExt, []byte(proto), 0666); err != nil {
			log.Fatal(err)
		}

		//
	}
}

// ParseType Анализируем тип и формируем его описание (Object = *TypeName)
func ParseType(Object interface{}) (tplP, tplM string, err error) {
	// разбираем тип
	var objValue = reflect.ValueOf(Object)
	if objValue.Kind() != reflect.Ptr {
		return tplP, tplM, errors.New("not ptr")
	}
	if objValue.IsNil() == true {
		return tplP, tplM, errors.New("is null")
	}
	objValue = objValue.Elem()

	list := strings.Split(objValue.Type().String(), ".")
	tplP = "\nmessage " + list[1] + " {\n"
	tplM = "\nfunc New" + list[1] + "GRPC(proto *" + tplPkg + "." + list[1] + ") *" + list[1] + " {\n"
	tplM += "\n\treturn &" + list[1] + "{\n"

	// разбираем свойства типа
	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		// пропускаем приватные свойства
		if false == field.IsValid() || false == field.CanSet() {
			continue
		}
		tplP_, tplM_ := ParseField(objValue, i)
		tplP += tplP_
		tplM += tplM_
	}
	tplP += "}\n"
	tplM += "\t}\n}\n"

	return tplP, tplM, nil
}

// ParseField Анализируем свойство типа и формируем его описание
func ParseField(objValue reflect.Value, i int) (tplP, tplM string) {
	field := objValue.Type().Field(i).Name
	fieldJSON := objValue.Type().Field(i).Tag.Get(`json`)

	// пропускаем исключенные и не обозначенные свойства
	if fieldJSON == `-` || fieldJSON == "" {
		return tplP, tplM
	}
	fieldJSON = strings.Split(fieldJSON, ",")[0]

	// формируем согласно типу
	prop := objValue.FieldByName(field)
	propType := prop.Type().String()
	propKind := prop.Type().Kind()
	subjErr := "implemented bytes: %s->%s [%s] %s"
	subjErr = fmt.Sprintf(subjErr, objValue.Type().String(), field, propKind.String(), propType)

	switch propKind {
	case reflect.String:
		tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Bool:
		tplP += "\tbool " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Float32:
		tplP += "\tfloat " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Float64:
		tplP += "\tdouble " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Int, reflect.Int64:
		tplP += "\tint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Int8, reflect.Int16, reflect.Int32:
		tplP += "\tint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Uint, reflect.Uint64:
		tplP += "\tuint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		tplP += "\tuint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Slice:
		if "[]string" == propType {
			tplP += "\trepeated string " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else if "[]uint8" == propType {
			tplP += "\tbytes " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else {
			// google.protobuf.Any
			tplP += "\tbytes " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
			fmt.Println(subjErr)
		}
	// custom type
	case reflect.Struct:
		if "typ.UUID" == propType || "decimal.Decimal" == propType {
			tplP, tplM = GenerateFieldUUID(i, field, fieldJSON)

		} else if "time.Time" == propType || "null.Time" == propType {
			tplP += "\tgoogle.protobuf.Timestamp " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else if "null.String" == propType {
			tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else {
			// google.protobuf.Any
			tplP += "\tbytes " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
			fmt.Println(subjErr)
		}
	default:
		// google.protobuf.Any
		tplP += "\tbytes " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		fmt.Println(subjErr)
	}
	return tplP, tplM
}

// GenerateFieldUUID
func GenerateFieldUUID(i int, field, fieldJSON string) (tplP, tplM string) {
	tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplM = fmt.Sprintf("\t\t%s: typ.UUIDMustParse(proto.%s),", field, ValidNameField(fieldJSON))
	return tplP, tplM
}

// ValidNameField сопоставление названий свойств с grpc типами
func ValidNameField(field string) string {
	if field == "id" {
		return "Id"
	}
	list := strings.Split(field, "_")
	for i := range list {
		list[i] = strings.Title(list[i])
	}
	return strings.Join(list, "")
}
