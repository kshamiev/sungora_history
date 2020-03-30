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

var source = map[string][]interface{}{
	"pb/modelcore": {
		&modelcore.User{},
		&modelcore.Order{},
		&modelcore.Role{},
	},
}

const (
	doNotEdit = "DO NOT EDIT"
	protoExt  = ".proto"
)

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
		if data, err = ioutil.ReadFile(pkg + protoExt); err != nil {
			log.Fatal(err)
		}
		proto = string(data)

		list := strings.Split(proto, doNotEdit)
		list[1] = tplFull
		proto = strings.Join(list, doNotEdit)

		if err = ioutil.WriteFile(pkg+protoExt, []byte(proto), 0666); err != nil {
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
	fieldName := objValue.Type().Field(i).Name
	fieldTagJSON := objValue.Type().Field(i).Tag.Get(`json`)
	// пропускаем исключенные и не обозначенные свойства
	if fieldTagJSON == `-` || fieldTagJSON == "" {
		return tpl
	}
	fieldTagJSON = strings.Split(fieldTagJSON, ",")[0]
	// формируем согласно типу
	prop := objValue.FieldByName(fieldName)
	subjErr := "not implemented: " + fieldName + " [" + prop.Type().Kind().String() + "] " + prop.Type().String()
	switch prop.Type().Kind() {
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
		if "[]string" == prop.Type().String() {
			tpl += "\trepeated string " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else {
			tpl += "\tgoogle.protobuf.Any " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
			fmt.Println(subjErr)
		}
	case reflect.Struct:
		if "typ.UUID" == prop.Type().String() || "decimal.Decimal" == prop.Type().String() {
			tpl += "\tstring " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else if "time.Time" == prop.Type().String() || "null.Time" == prop.Type().String() {
			tpl += "\tgoogle.protobuf.Timestamp " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else if "null.String" == prop.Type().String() {
			tpl += "\tstring " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
		} else {
			tpl += "\tgoogle.protobuf.Any " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
			fmt.Println(subjErr)
		}
	default:
		tpl += "\tgoogle.protobuf.Any " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
		fmt.Println(subjErr)
	}
	return tpl
}
