package logger

import (
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
)

type FileName struct {
	SourceField string
}

func filenameHook(config *FileName) logrus.Hook {
	filenameHook := filename.NewHook()
	filenameHook.Field = config.SourceField

	return filenameHook
}
