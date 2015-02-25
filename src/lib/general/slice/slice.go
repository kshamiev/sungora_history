// Библиотека для работы со срезами
package slice

import (
	"reflect"
	"sort"
	"time"
)

// Тип для агрегирования методов работы со срезами
type Slice struct {
}

// Конструктор
//	- *Slice объект агрегатор функций-методов
func NewSlice() *Slice {
	var self = new(Slice)
	return self
}

//// Сортировка срезов

// Сортировка срезов любых типов по возрастанию
//	+ slice interface{} срез значений структур или их ссылок
//	+ property string свойство по которому сортируем элементы среза (строки, числа, дата и время)
//	property должно быть публичным
func (self *Slice) SortingSliceAsc(slice interface{}, property string) {
	ps := &sorting{
		direction: true,
		data:      slice,
		property:  property,
	}
	sort.Sort(ps)
}

// Сортировка срезов любых типов по убыванию
//	+ slice interface{} срез значений структур или их ссылок
//	+ property string свойство по которому сортируем элементы среза (строки, числа, дата и время)
//	property должно быть публичным
func (self *Slice) SortingSliceDesc(slice interface{}, property string) {
	ps := &sorting{
		direction: false,
		data:      slice,
		property:  property,
	}
	sort.Sort(ps)
}

// Сортировка срезов любых типов по пользовательской функции
//	+ slice interface{} срез значений структур или их ссылок
//	+ f func(... пользовательская функция сортировки
func (self *Slice) SortingSliceFunc(slice interface{}, f func(left, right interface{}) bool) {
	ps := &sorting{
		data:     slice,
		function: f,
	}
	sort.Sort(ps)
}

type sorting struct {
	direction bool                               // направление сортировки
	data      interface{}                        // сортируемый срез данных
	property  string                             // свойство по которому сортируем
	function  func(left, right interface{}) bool // пользовательская функция сравнения двух значений
}

func (self sorting) Len() int {
	objValue := reflect.ValueOf(self.data)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	return objValue.Len()
}

func (self sorting) Less(i, j int) bool {
	var objValue = reflect.ValueOf(self.data)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	var objValue1 = objValue.Index(i)
	var objValue2 = objValue.Index(j)

	// сортировка по пользовательской функции
	if self.function != nil {
		return self.function(objValue1.Interface(), objValue2.Interface())
	}

	// прямая сортировка
	if objValue1.Kind() == reflect.Ptr {
		objValue1 = objValue1.Elem()
	}
	var vi = objValue1.FieldByName(self.property).Interface()
	if objValue2.Kind() == reflect.Ptr {
		objValue2 = objValue2.Elem()
	}
	var vj = objValue2.FieldByName(self.property).Interface()
	switch vi.(type) {
	case string:
		if self.direction {
			return vi.(string) < vj.(string)
		} else {
			return vi.(string) > vj.(string)
		}
	case float64:
		if self.direction {
			return vi.(float64) < vj.(float64)
		} else {
			return vi.(float64) > vj.(float64)
		}
	case int64:
		if self.direction {
			return vi.(int64) < vj.(int64)
		} else {
			return vi.(int64) > vj.(int64)
		}
	case uint64:
		if self.direction {
			return vi.(uint64) < vj.(uint64)
		} else {
			return vi.(uint64) > vj.(uint64)
		}
	case int32:
		if self.direction {
			return vi.(int32) < vj.(int32)
		} else {
			return vi.(int32) > vj.(int32)
		}
	case uint32:
		if self.direction {
			return vi.(uint32) < vj.(uint32)
		} else {
			return vi.(uint32) > vj.(uint32)
		}
	case int16:
		if self.direction {
			return vi.(int16) < vj.(int16)
		} else {
			return vi.(int16) > vj.(int16)
		}
	case uint16:
		if self.direction {
			return vi.(uint16) < vj.(uint16)
		} else {
			return vi.(uint16) > vj.(uint16)
		}
	case int8:
		if self.direction {
			return vi.(int8) < vj.(int8)
		} else {
			return vi.(int8) > vj.(int8)
		}
	case uint8:
		if self.direction {
			return vi.(uint8) < vj.(uint8)
		} else {
			return vi.(uint8) > vj.(uint8)
		}
	case time.Time:
		if self.direction {
			return vj.(time.Time).After(vi.(time.Time))
		} else {
			return vj.(time.Time).Before(vi.(time.Time))
		}
	case time.Duration:
		if self.direction {
			return vi.(time.Duration) < vj.(time.Duration)
		} else {
			return vi.(time.Duration) > vj.(time.Duration)
		}
	}
	return true
}

func (self sorting) Swap(i, j int) {
	var objValue = reflect.ValueOf(self.data)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	var objValue1 = objValue.Index(i)
	if objValue1.Kind() == reflect.Ptr {
		objValue1 = objValue1.Elem()
	}

	var objValue2 = objValue.Index(j)
	if objValue2.Kind() == reflect.Ptr {
		objValue2 = objValue2.Elem()
	}

	slc := reflect.New(objValue1.Type())
	if slc.Kind() == reflect.Ptr {
		slc = slc.Elem()
	}
	slc.Set(objValue1)

	objValue1.Set(objValue2)
	objValue2.Set(slc)
}

////
