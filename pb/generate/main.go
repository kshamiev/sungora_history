package main

import (
	"errors"
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
			tpl, err = TypParse(t)
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

// TypParse Анализируем тип и формируем proto для него (Object = *TypeName)
func TypParse(Object interface{}) (tpl string, err error) {
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
		// пропускаем исключенные свойства
		fieldTagJSON := objValue.Type().Field(i).Tag.Get(`json`)
		if fieldTagJSON == `-` {
			continue
		}
		// формируем согласно типу
		fieldName := objValue.Type().Field(i).Name
		prop := objValue.FieldByName(fieldName)
		switch prop.Type().Kind() {
		case reflect.String:
			tpl += "\tstring " + fieldTagJSON + " = " + strconv.Itoa(i+1) + ";\n"
		}
	}
	tpl += "}\n"

	return tpl, nil
}
