package main

import (
	"funtik"
	"lib/logs"
)

func NewTest() *Content {
	return new(Content)
}

func main() {

	funtik.FuntikOne()

	logs.Dumper()
}

type ContentFace interface {
	Set(r string)
}

// Контент
type Content struct {
	Uri_Id      uint64 // Uri
	Lang        string // Язык
	Title       string // Заголовок
	Keywords    string // Ключи
	Description string // Описание
	Content     []byte // Контент
	Block       string // Блок
}

func (self *Content) Set(t string) {
	self.Description = t
}
