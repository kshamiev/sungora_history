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

	"github.com/kshamiev/sungora/pb/modelsun"
)

const (
	pathImport      = "github.com/kshamiev/sungora"
	tplProtoPkgName = "pb"
)

var source = map[string][]interface{}{
	"modelsun": {
		&modelsun.User{},
		&modelsun.Order{},
		&modelsun.Role{},
	},
}

func main() {
	var err error
	var tplPFull, tplMFull, tplP, tplM string

	for pkgName := range source {
		// анализируем типы и формируем сопряжение
		tplPFull = CreateProtoFile(pkgName)
		tplMFull = CreateTypeFile(pkgName)
		for _, t := range source[pkgName] {
			if tplP, tplM, err = ParseType(t); err != nil {
				log.Fatal(err)
			}
			tplPFull += tplP
			tplMFull += tplM
		}
		// описание прототипов
		if err = ioutil.WriteFile(pkgName+".proto", []byte(tplPFull), 0666); err != nil {
			log.Fatal(err)
		}
		// тметоды конвертации
		if err = ioutil.WriteFile(pkgName+"/grpc.go", []byte(tplMFull), 0666); err != nil {
			log.Fatal(err)
		}
	}
}

// ParseType Анализируем тип и формируем его сопряжение с grpc (Object = *TypeName)
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
	tplM = "\nfunc New" + list[1] + "GRPC(proto *" + tplProtoPkgName + "." + list[1] + ") *" + list[1] + " {\n"
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

// ParseField Анализируем свойство типа и формируем его конвертацию
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
		if "typ.UUID" == propType {
			tplP, tplM = GenerateFieldUUID(i, field, fieldJSON)
		} else if "decimal.Decimal" == propType {
			tplP, tplM = GenerateFieldDecimal(i, field, fieldJSON)
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

// GenerateFieldDecimal конвертация туда и обратно
func GenerateFieldDecimal(i int, field, fieldJSON string) (tplP, tplM string) {
	tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplM = fmt.Sprintf("\t\t%s: decimal.RequireFromString(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplM
}

// GenerateFieldUUID конвертация туда и обратно
func GenerateFieldUUID(i int, field, fieldJSON string) (tplP, tplM string) {
	tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplM = fmt.Sprintf("\t\t%s: typ.UUIDMustParse(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplM
}

// ValidNameField сопоставление свойств исходного типа с grpc прототипами через тег
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

// CreateTypeFile инициализация файла с методами конвертации типа
func CreateTypeFile(pkgName string) string {
	return `
package ` + pkgName + `

import (
	"github.com/shopspring/decimal"

	"` + pathImport + `/` + tplProtoPkgName + `"
	"` + pathImport + `/` + tplProtoPkgName + `/typ"
)
`
}

// CreateProtoFile инициализация файла с описанием прототипов
func CreateProtoFile(pkgName string) (proto string) {
	if data, err := ioutil.ReadFile(pkgName + ".proto"); err == nil {
		proto = string(data)
		list := strings.Split(proto, "DO NOT EDIT")
		proto = list[0] + "DO NOT EDIT" + "\n"
	} else {
		if data, err = ioutil.ReadFile("generate/template.proto"); err != nil {
			log.Fatal(err)
		}
		proto = string(data)
		proto = strings.ReplaceAll(proto, "TPLpackage", tplProtoPkgName)
		proto = strings.ReplaceAll(proto, "TPLservice", strings.Title(pkgName))
	}
	return proto
}

// func NewUserGRPC(proto *pb.User) *User {
// 	return &User{
// 		ID: typ.UUIDMustParse(proto.Id),
// 	}
// }
