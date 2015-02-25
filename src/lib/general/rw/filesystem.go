package rw

import (
	"io"
	"os"
)

// Тип для агрегирования методов работы с FS & IO
type RW struct {
}

// Конструктор
//	- *RW объект агрегатор функций-методов
func NewRW() *RW {
	var self = new(RW)
	return self
}

// Запись текстовых данных в конец указанного файла
//	+ filename string полнуй путь до файла
//	+ data string данные для записи
//	- error ошибка операции
func (self *RW) FileSaveAppend(filename string, data string) (err error) {
	dataByte := []byte(data)
	if err = os.MkdirAll(filename, 0777); err != nil {
		return
	}
	var f *os.File
	f, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	n, err := f.Write(dataByte)
	if err == nil && n < len(dataByte) {
		err = io.ErrShortWrite
	}
	return err
}
