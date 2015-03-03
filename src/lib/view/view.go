// Библиотека: Шаблонизатор.
//
// Компиляция или вставка данных в html шаблон.
package view

import (
	"bytes"
	"path/filepath"
	"text/template"
)

// Структура шаблонизатора
type View struct {
	Variables map[string]interface{} // Данные вставляемые в шаблон
	Functions map[string]interface{} // template.FuncMap (по умолчанию пустой)
}

// Создание объекта шаблонизатор
//	- *View шаблонизатор
func NewView() *View {
	var self = new(View)
	self.Variables = make(map[string]interface{})
	self.Functions = make(map[string]interface{})
	return self
}

// Выполнение шаблона (вставка данных в шаблон)
//	+ path string абсолютный путь до шаблона
//	- bytes.Buffer буфер собранного шаблонга с данными
//	- error ошибка операции
func (self *View) ExecuteFile(path string) (bytes.Buffer, error) {
	var err error
	var tpl *template.Template
	var ret bytes.Buffer
	tpl, err = template.New(filepath.Base(path)).Funcs(self.Functions).ParseFiles(path)
	if err != nil {
		return ret, err
	}
	err = tpl.Execute(&ret, self.Variables)
	if err != nil {
		return ret, err
	}
	return ret, err
}

// Выполнение шаблона (вставка данных в шаблон)
//	+ data []byte исходный шаблон
//	- bytes.Buffer буфер собранного шаблонга с данными
//	- error ошибка операции
func (self *View) ExecuteBytes(data []byte) (bytes.Buffer, error) {
	var err error
	var tpl *template.Template
	var ret bytes.Buffer

	tpl, err = template.New(`ExecuteStringTemplate`).Funcs(self.Functions).Parse(string(data))
	if err != nil {
		return ret, err
	}
	err = tpl.Execute(&ret, self.Variables)
	if err != nil {
		return ret, err
	}
	return ret, err
}

// Выполнение шаблона (вставка данных в шаблон)
//	+ str string исходный шаблон
//	- bytes.Buffer буфер собранного шаблонга с данными
//	- error ошибка операции
func (self *View) ExecuteString(str string) (bytes.Buffer, error) {
	var err error
	var tpl *template.Template
	var ret bytes.Buffer

	tpl, err = template.New(`ExecuteStringTemplate`).Funcs(self.Functions).Parse(str)
	if err != nil {
		return ret, err
	}
	err = tpl.Execute(&ret, self.Variables)
	if err != nil {
		return ret, err
	}
	return ret, err
}

////
