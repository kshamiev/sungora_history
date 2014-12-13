// logs daemon

package logs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"lib"
)

// cammand channel
var commandlogsControl = make(chan command, 1000)

// Command structure
type command struct {
	action  int
	code    int
	message string
	level   string
	result  chan<- interface{}
}

// Допустимые команды (action)
const (
	logsMessage int = iota // Сообщение (лог сообщение)
	logsClose              // Закрытие службы логирования
)

// gologs Служба логирования
func gologs() {

	go func() {
		msg := fmt.Sprintf("%s\t[start]\r", lib.Time.Label())
		logsSave(msg)
		for command := range commandlogsControl {
			switch command.action {
			case logsMessage:
				msg := logsMessageCalculate(command.code, command.message, command.level)
				logsSave(msg)
			case logsClose:
				msg := fmt.Sprintf("%s\t[stop]\r", lib.Time.Label())
				logsSave(msg)
				if fp != nil {
					fp.Close()
				}
				close(commandlogsControl)
				command.result <- true
			}
		}
	}()

}

// logsMessageCalculate формирование сообщения для логирования
func logsMessageCalculate(code int, message string, level string) string {
	// временная отметка
	var prefix = lib.Time.Label()

	// формируем
	var logLine = fmt.Sprintf("%s\t%s[%d]", prefix, level, code)
	if cfglogs.Debug == true {
		logLine += `[debug]`
	}
	logLine += "\t" + message

	// информация режима debug
	if cfglogs.Debug == true && cfglogs.DebugDetail >= 1 {
		// информация о вызвавшей программе
		var info, _ = getCallerInfo(3)
		var debugLine = fmt.Sprintf("%s\t[%s]\t%s func:%s line:%d file:%s run:%d",
			lib.Time.Label(),
			prefix,
			info.Version,
			info.FuncName,
			info.LineNumber,
			info.FileName,
			info.Gorutines)
		if cfglogs.DebugDetail >= 2 {
			debugLine += " gorutines:\r\n"
			for i := range info.GorutinesInfo {
				debugLine += info.GorutinesInfo[i] + "\r\n"
			}
		}
		logLine += "\t" + debugLine
	}

	return logLine
}

type callerInfo struct {
	FileName      string   // Имя исходного файла
	LineNumber    int      // Номер строки
	FuncName      string   // Название функции
	Version       string   // Текущая версия golang
	Gorutines     int      // Количество горутин работающих в настоящий момент
	GorutinesInfo []string // Информация по каждой горутине
}

// Получение информаци о вызвавшем лог
func getCallerInfo(level int) (*callerInfo, error) {
	var err error
	var ret *callerInfo = new(callerInfo)
	var pc uintptr
	var file string
	var line int
	var ok bool

	ret.Gorutines = runtime.NumGoroutine()
	ret.Version = runtime.Version()

	pc, file, line, ok = runtime.Caller(level)
	if ok == true {
		var fn *runtime.Func

		ret.LineNumber = line
		ret.FileName = file
		fn = runtime.FuncForPC(pc)
		if fn != nil {
			ret.FuncName = fn.Name()
		}

		// Информация о состоянии горутин
		var buf []byte = make([]byte, 1<<16)
		var i int = runtime.Stack(buf, true)
		var info string = string(buf[:i])
		var tmp []string

		tmp = strings.Split(info, "\n")
		i = -1
		for _, str := range tmp {
			if strings.Index(str, "goroutine") == 0 {
				i++
				ret.GorutinesInfo = append(ret.GorutinesInfo, "")
			}
			if i >= 0 && str != "" {
				if ret.GorutinesInfo[i] != "" {
					ret.GorutinesInfo[i] += "\n"
				}
				ret.GorutinesInfo[i] += str
			}
		}
	} else {
		err = errors.New("Не удалось получить информацию о вызвавшей лог функции")
	}
	return ret, err
}

var fp *os.File

// logsSave непосредственное сохранение лога
func logsSave(msg string) {
	if cfglogs.Mode == `mixed` || cfglogs.Mode == `file` {
		if fp == nil {
			os.MkdirAll(filepath.Dir(cfglogs.File), 0777)
			fp, _ = os.OpenFile(cfglogs.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
		}
		if fp != nil {
			fp.WriteString(msg + "\n")
		}
	}
	if cfglogs.Mode == `mixed` || cfglogs.Mode == `system` {
		// TODO Реализовать
	}
	// В режиме дебаг пишем в stdout все
	if cfglogs.Debug == true {
		fmt.Println(msg)
	}

}
