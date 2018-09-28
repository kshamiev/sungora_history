package lg

import (
	"fmt"
	"os"
	"path/filepath"
)

var fp *os.File

func saveFile(com comand) {

	var logLine = fmt.Sprintf("%s\t%s\t%s", com.datetime, com.level, com.message)

	if fp == nil {
		os.MkdirAll(filepath.Dir(conf.OutFilePath), 0777)
		fp, _ = os.OpenFile(conf.OutFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	}
	if fp != nil {
		fp.WriteString(logLine + "\n")

		for _, d := range com.traces {
			logLine = fmt.Sprintf("\t%s\t%d\t%s\n", d.FuncName, d.LineNumber, d.FileName)
			fp.WriteString(logLine)
		}

	}

}
