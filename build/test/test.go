package main

import (
	"lib/logs"
)

func main() {

	var t = new(Content)
	t.Description = `ghdfjghdfghdf`

	var b = t

	var t1 = b
	b = t1
	b.Description = `!!!!!!!!!!!!!!!!!!!!!!`

	logs.Dumper(t, b)
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
