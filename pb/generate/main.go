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
	tplDoNotEdit  = "DO NOT EDIT"
	tplProtoExt   = ".proto"
	tplProtoBlank = "pb/generate/tpl.proto"
)

var source = map[string][]interface{}{
	"pb/modelcore": {
		&modelcore.User{},
		&modelcore.Order{},
		&modelcore.Role{},
	},
}

func main() {
	var err error
	var data []byte
	var proto, tplFull, tpl string

	for pkg := range source {
		// анализируем типы и формируем proto
		tplFull = "\n"
		for _, t := range source[pkg] {
			tpl, err = ParseType(t)
			tplFull += tpl
		}

		// формируем proto файлы
		if data, err = ioutil.ReadFile(pkg + tplProtoExt); err != nil {
			if data, err = ioutil.ReadFile(tplProtoBlank); err != nil {
				log.Fatal(err)
			}
		}
		proto = string(data)

		list := strings.Split(proto, tplDoNotEdit)
		list[1] = tplFull
		proto = strings.Join(list, tplDoNotEdit)

		if err = ioutil.WriteFile(pkg+tplProtoExt, []byte(proto), 0666); err != nil {
			log.Fatal(err)
		}

		//
	}
}

// ParseType Анализируем тип и формируем proto для него (Object = *TypeName)
func ParseType(Object interface{}) (tpl string, err error) {
	// разбираем тип
	var objValue = reflect.ValueOf(Object)
	if objValue.Kind() != reflect.Ptr {
		return tpl, errors.New("not ptr")
	}
	if objValue.IsNil() == true {
		return tpl, errors.New("is null")
	}
	objValue = objValue.Elem()

	list := strings.Split(objValue.Type().String(), ".")
	tpl = "\nmessage " + list[1] + " {\n"

	// разбираем свойства типа
	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		// пропускаем приватные свойства
		if false == field.IsValid() || false == field.CanSet() {
			continue
		}
		tpl += ParseField(objValue, i)
	}
	tpl += "}\n"

	return tpl, nil
}

// ParseField Анализируем свойство типа и формируем proto для него
func ParseField(objValue reflect.Value, i int) (tpl string) {
	field := objValue.Type().Field(i).Name
	fieldTagJSON := objValue.Type().Field(i).Tag.Get(`json`)

	// пропускаем исключенные и не обозначенные свойства
	if fieldTagJSON == `-` || fieldTagJSON == "" {
		return tpl
	}
	fieldTagJSON = strings.Split(fieldTagJSON, ",")[0]

	// формируем согласно типу
	prop := objValue.FieldByName(field)
	propType := prop.Type().String()
	propKind := prop.Type().Kind()
	subjErr := "implemented bytes: %s->%s [%s] %s"
	subjErr = fmt.Sprintf(subjErr, objValue.Type().String(), field, propKind.String(), propType)

	switch propKind {
	case reflect.String:
		tpl += "\tstring " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Bool:
		tpl += "\tbool " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Float32:
		tpl += "\tfloat " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Float64:
		tpl += "\tdouble " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Int, reflect.Int64:
		tpl += "\tint64 " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Int8, reflect.Int16, reflect.Int32:
		tpl += "\tint32 " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Uint, reflect.Uint64:
		tpl += "\tuint64 " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		tpl += "\tuint32 " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
	case reflect.Slice:
		if "[]string" == propType {
			tpl += "\trepeated string " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else if "[]uint8" == propType {
			tpl += "\tbytes " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else {
			// google.protobuf.Any
			tpl += "\tbytes " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
			fmt.Println(subjErr)
		}
	// custom type
	case reflect.Struct:
		if "typ.UUID" == propType || "decimal.Decimal" == propType {
			tpl += "\tstring " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else if "time.Time" == propType || "null.Time" == propType {
			tpl += "\tgoogle.protobuf.Timestamp " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else if "null.String" == propType {
			tpl += "\tstring " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else {
			// google.protobuf.Any
			tpl += "\tbytes " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
			fmt.Println(subjErr)
		}
	default:
		// google.protobuf.Any
		tpl += "\tbytes " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
		fmt.Println(subjErr)
	}
	return tpl
}
