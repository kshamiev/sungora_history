package lib

import (
	`lib/general/rw`
	`lib/general/slice`
	`lib/general/str`
	`lib/general/times`
)

var Time = times.NewTime(`Europe/Moscow`)

var String = str.NewString()

var Slice = slice.NewSlice()

var RW = rw.NewRW()
