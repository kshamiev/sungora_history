// Работа шаблонизатора
package view

import (
	"bytes"
	"path/filepath"
	"text/template"
)

// Шаблонизатор
type View struct {
	Variables map[string]interface{} // Variable (по умолчанию пустой)
	Functions map[string]interface{} // html/template.FuncMap (по умолчанию пустой)
}

// NewView Шаблонизатор (конструктор)
func NewView() *View {
	var self = new(View)
	self.Variables = make(map[string]interface{})
	self.Functions = make(map[string]interface{})
	return self
}

// ExecuteFile Выполнение шаблона (вставка данных в шаблон)
// path - путь до шаблона
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

// ExecuteBytes Выполнение шаблона (вставка данных в шаблон)
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

// ExecuteString Выполнение шаблона (вставка данных в шаблон)
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
