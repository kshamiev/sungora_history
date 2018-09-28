/*
Реализовать:
Логирование в
файл
консоль
graylog
Возможность настройки куда логировать
Возможность переопределения реализации
 */
package lg

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"gopkg.in/sungora/app.v1/config"
)

type comand struct {
	datetime string
	level    string
	message  string
	traces   []trace
}

var logCh = make(chan comand, 10000)
var logChClose = make(chan bool)
var conf config.Log

func Start(c config.Log) {
	conf = c
	if conf.OutFilePath == "" {
		if dir, err := os.Getwd(); err == nil {
			conf.OutFilePath = dir + "/logs/" + filepath.Base(os.Args[0]) + ".log"
		}
	}
	go func() {
		for com := range logCh {
			if conf.OutStd == true {
				saveStdout(com)
			}
			if conf.OutFile == true {
				saveFile(com)
			}
		}
		logChClose <- true
	}()

}

func Stop() {
	close(logCh)
	<-logChClose
}

type trace struct {
	FuncName   string // Название функции
	FileName   string // Имя исходного файла
	LineNumber int    // Номер строки внутри функции
}

// Получение информаци о вызвавшем лог
func getCallInfo(level int) (FuncName string, LineNumber int, FileName string) {
	pc, file, line, ok := runtime.Caller(level)
	if ok == true {
		LineNumber = line
		FileName = file
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			FuncName = fn.Name()
		}
	}
	return
}

// Получение информаци о вызвавшем лог
func getTrace() (traces []trace, err error) {
	buf := make([]byte, 1<<16)
	i := runtime.Stack(buf, true)
	info := string(buf[:i])
	infoList := strings.Split(info, "\n")
	infoList = infoList[5:]
	for i := 0; i < len(infoList)-1; i += 2 {
		if infoList[i] == "" {
			break
		}
		tmp := strings.Split(infoList[i+1], " ")
		tmp = strings.Split(tmp[0], ":")
		line, _ := strconv.Atoi(tmp[1])
		t := trace{
			FuncName:   infoList[i],
			FileName:   tmp[0],
			LineNumber: line,
		}
		traces = append(traces, t)
	}
	return
}

func datetimeLabal() string {
	var t time.Time
	loc, err := time.LoadLocation("Europe/Moscow")
	if err == nil {
		t = time.Now().In(loc)
	} else {
		t = time.Now().In(time.UTC)
	}
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())

}

func searchMsg(code int, params ...interface{}) (message string) {
	var ok bool
	if message, ok = messages[code]; ok == true {
		message = fmt.Sprintf(message, params...)
	} else if 0 < len(params) {
		if s, ok := params[0].(string); ok == true {
			message = fmt.Sprintf(s, params[1:]...)
		}
	}
	return
}

