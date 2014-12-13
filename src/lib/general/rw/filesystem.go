package rw

import (
	"io"
	"os"
)

type RW struct {
}

func NewRW() *RW {
	var self = new(RW)
	return self
}

func (self *RW) FileSaveAppend(filename string, data string) error {
	dataByte := []byte(data)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
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
