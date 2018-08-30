package lg

import (
	"os"
	"path/filepath"
)

var fp *os.File

func saveFile(com comand) {

	if fp == nil {
		os.MkdirAll(filepath.Dir(config.filePath), 0777)
		fp, _ = os.OpenFile(config.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	}
	if fp != nil {
		fp.WriteString(com.message + "\n")
	}

}
