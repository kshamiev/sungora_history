package controller

import (
	"lib/logs"
)

func debug(obj ...interface{}) {
	logs.Dumper(obj...)
}
