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
	fmt.Println("OK")
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

	tplMFrom := "\nfunc New" + list[1] + "Proto(proto *" + tplProtoPkgName + "." + list[1] + ") *" + list[1] + " {\n"
	tplMFrom += "\treturn &" + list[1] + "{\n"

	tplMTo := "\nfunc (o *" + list[1] + ") Proto() *" + tplProtoPkgName + "." + list[1] + " {\n"
	tplMTo += "\treturn &" + tplProtoPkgName + "." + list[1] + "{\n"

	// разбираем свойства типа
	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		// пропускаем приватные свойства
		if false == field.IsValid() || false == field.CanSet() {
			continue
		}
		tplP_, tplMFrom_, tplMTo_ := ParseField(objValue, i)
		tplP += tplP_
		tplMFrom += tplMFrom_
		tplMTo += tplMTo_
	}
	tplP += "}\n"

	tplMFrom += "\t}\n}\n"
	tplMTo += "\t}\n}\n"

	return tplP, tplMTo + tplMFrom, nil
}

// ParseField Анализируем свойство типа и формируем его конвертацию
func ParseField(objValue reflect.Value, i int) (tplP, tplMFrom, tplMTo string) {
	field := objValue.Type().Field(i).Name
	fieldJSON := objValue.Type().Field(i).Tag.Get(`json`)

	// пропускаем исключенные и не обозначенные свойства
	if fieldJSON == `-` || fieldJSON == "" {
		return tplP, tplMFrom, tplMTo
	}
	fieldJSON = strings.Split(fieldJSON, ",")[0]

	// формируем согласно типу
	prop := objValue.FieldByName(field)
	propType := prop.Type().String()
	propKind := prop.Type().Kind()
	subjErr := "not implemented undefined property: %s->%s [%s] %s"
	subjErr = fmt.Sprintf(subjErr, objValue.Type().String(), field, propKind.String(), propType)

	switch propKind {
	case reflect.String:
		if "string" != propType {
			tplP, tplMFrom, tplMTo = GenerateFieldEnum(i, field, fieldJSON, propType)
		} else {
			tplP, tplMFrom, tplMTo = GenerateFieldString(i, field, fieldJSON)
		}

	case reflect.Bool:
		tplP, tplMFrom, tplMTo = GenerateFieldBool(i, field, fieldJSON)

	case reflect.Float32:
		tplP += "\tfloat " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = GenerateFieldScalar(i, field, fieldJSON)
	case reflect.Float64:
		tplP += "\tdouble " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = GenerateFieldScalar(i, field, fieldJSON)

	case reflect.Int:
		tplP, tplMFrom, tplMTo = GenerateFieldInt(i, field, fieldJSON)
	case reflect.Int8:
		tplP, tplMFrom, tplMTo = GenerateFieldInt8(i, field, fieldJSON)
	case reflect.Int16:
		tplP, tplMFrom, tplMTo = GenerateFieldInt16(i, field, fieldJSON)
	case reflect.Int32:
		tplP += "\tint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = GenerateFieldScalar(i, field, fieldJSON)
	case reflect.Int64:
		tplP += "\tint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = GenerateFieldScalar(i, field, fieldJSON)

	case reflect.Uint:
		tplP, tplMFrom, tplMTo = GenerateFieldUint(i, field, fieldJSON)
	case reflect.Uint8:
		tplP, tplMFrom, tplMTo = GenerateFieldUint8(i, field, fieldJSON)
	case reflect.Uint16:
		tplP, tplMFrom, tplMTo = GenerateFieldUint16(i, field, fieldJSON)
	case reflect.Uint32:
		tplP += "\tuint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = GenerateFieldScalar(i, field, fieldJSON)
	case reflect.Uint64:
		tplP += "\tuint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = GenerateFieldScalar(i, field, fieldJSON)

	case reflect.Slice:
		if "[]string" == propType {
			tplP, tplMFrom, tplMTo = GenerateFieldStringArray(i, field, fieldJSON)
		} else if "[]uint8" == propType {
			tplP, tplMFrom, tplMTo = GenerateFieldBytes(i, field, fieldJSON)
		} else if "types.StringArray" == propType {
			tplP, tplMFrom, tplMTo = GenerateFieldStringArray(i, field, fieldJSON)
		} else {
			fmt.Println(subjErr)
		}

	case reflect.Struct:
		if "typ.UUID" == propType {
			tplP, tplMFrom, tplMTo = GenerateFieldUUID(i, field, fieldJSON)
		} else if "decimal.Decimal" == propType {
			tplP, tplMFrom, tplMTo = GenerateFieldDecimal(i, field, fieldJSON)
		} else if "time.Time" == propType {
			tplP, tplMFrom, tplMTo = GenerateFieldTime(i, field, fieldJSON)
		} else if "null.Time" == propType {
			tplP, tplMFrom, tplMTo = GenerateFieldNullTime(i, field, fieldJSON)
		} else if "null.String" == propType {
			tplP, tplMFrom, tplMTo = GenerateFieldNullString(i, field, fieldJSON)
		} else if "null.Bytes" == propType {
			tplP, tplMFrom, tplMTo = GenerateFieldNullBytes(i, field, fieldJSON)
		} else if "null.JSON" == propType {
			tplP, tplMFrom, tplMTo = GenerateFieldNullJSON(i, field, fieldJSON)
		} else {
			fmt.Println(subjErr)
		}

	default:
		fmt.Println(subjErr)
	}
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldScalar конвертация туда и обратно
func GenerateFieldScalar(i int, field, fieldJSON string) (tplMFrom, tplMTo string) {
	tplMTo = fmt.Sprintf("\t\t%s: o.%s,\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: proto.%s,\n", field, ValidNameField(fieldJSON))
	return tplMFrom, tplMTo
}

// GenerateFieldUint8 конвертация туда и обратно
func GenerateFieldUint8(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tuint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: uint32(o.%s),\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: uint8(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldUint16 конвертация туда и обратно
func GenerateFieldUint16(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tuint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: uint32(o.%s),\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: uint16(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldUint конвертация туда и обратно
func GenerateFieldUint(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tuint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: uint64(o.%s),\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: uint(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldInt8 конвертация туда и обратно
func GenerateFieldInt8(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: int32(o.%s),\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: int8(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldInt16 конвертация туда и обратно
func GenerateFieldInt16(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: int32(o.%s),\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: int16(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldInt конвертация туда и обратно
func GenerateFieldInt(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: int64(o.%s),\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: int(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldNullBytes конвертация туда и обратно
func GenerateFieldNullBytes(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tbytes " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: o.%s.Bytes,\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: null.BytesFrom(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldBytes конвертация туда и обратно
func GenerateFieldBytes(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tbytes " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: o.%s,\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: proto.%s,\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldStringArray конвертация туда и обратно
func GenerateFieldStringArray(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\trepeated string " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: o.%s,\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: proto.%s,\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldBool конвертация туда и обратно
func GenerateFieldBool(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tbool " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: o.%s,\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: proto.%s,\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldNullJSON конвертация туда и обратно
func GenerateFieldNullJSON(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tbytes " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: o.%s.JSON,\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: null.JSONFrom(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldEnum конвертация туда и обратно
func GenerateFieldEnum(i int, field, fieldJSON, propType string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: %sValue[o.%s],\n", ValidNameField(fieldJSON), propType, field)
	tplMFrom = fmt.Sprintf("\t\t%s: %sName[proto.%s],\n", field, propType, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldNullString конвертация туда и обратно
func GenerateFieldNullString(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: o.%s.String,\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: typ.PbToNullString(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldString конвертация туда и обратно
func GenerateFieldString(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: o.%s,\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: proto.%s,\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldTime конвертация туда и обратно
func GenerateFieldTime(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tgoogle.protobuf.Timestamp " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: typ.PbFromTime(o.%s),\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: typ.PbToTime(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldNullTime конвертация туда и обратно
func GenerateFieldNullTime(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tgoogle.protobuf.Timestamp " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: typ.PbFromTime(o.%s.Time),\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: typ.PbToNullTime(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldDecimal конвертация туда и обратно
func GenerateFieldDecimal(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: o.%s.String(),\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: decimal.RequireFromString(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldUUID конвертация туда и обратно
func GenerateFieldUUID(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: o.%s.String(),\n", ValidNameField(fieldJSON), field)
	tplMFrom = fmt.Sprintf("\t\t%s: typ.UUIDMustParse(proto.%s),\n", field, ValidNameField(fieldJSON))
	return tplP, tplMFrom, tplMTo
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
	return `package ` + pkgName + `

import (
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null"

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
