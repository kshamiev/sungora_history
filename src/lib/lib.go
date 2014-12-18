// Библиотека агрегатор общих и широко использумых повсеместно функций
package lib

import (
	`lib/general/rw`
	`lib/general/slice`
	`lib/general/str`
	`lib/general/times`
)

var (
	// Работа с датой и временем
	Time = times.NewTime(`Europe/Moscow`)
	// Работа со строками
	String = str.NewString()
	// Работа со срезами
	Slice = slice.NewSlice()
	// Работа с вводом и выводом (IO)
	RW = rw.NewRW()
)
