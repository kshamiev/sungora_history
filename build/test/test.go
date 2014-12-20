package main

import (
	"lib/logs"
)

func NewTest() *Content {
	return new(Content)
}

func main() {

	var ttt = make(map[string]func() *Content)
	ttt[`qqq`] = NewTest
	//t = new(Content)
	//t.Set(`ghdfjghdfghdf`)

	var t = ttt[`qqq`]()
	t.Description = `wwwwwwwwwwwwwwwwwwwwwwwww`
	//b = t

	//var t1 = *b
	//b = &t1
	//b.Set(`!!!!!!!!!!!!!!!!!!!!!!`)

	//logs.Dumper(t, b)
	logs.Dumper(ttt, t)
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
