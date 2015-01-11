package logs

// Сулжба логирования

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"bitbucket.org/kardianos/service"

	"lib"
)

// Запуск работы службы логирования
func GoStart(logSys service.Logger) {
	if fp == nil {
		os.MkdirAll(filepath.Dir(cfg.File), 0755)
		fp, _ = os.OpenFile(cfg.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	}
	sys = logSys
	gologs()
}

// Завершение работы службы логирования
func GoClose() bool {
	reply := make(chan interface{})
	commandlogsControl <- command{action: logsClose, result: reply}
	return (<-reply).(bool)
}

// Cammand channel
var commandlogsControl = make(chan command, 1000)

// Command structure
type command struct {
	action int
	log    *Log
	level  uint8
	result chan<- interface{}
}

// Справочник уровня ошибок
var logsLevel = map[uint8]string{
	6: `[info]`,
	5: `[notice]`,
	4: `[warning]`,
	3: `[error]`,
	2: `[critical]`,
	1: `[fatal]`,
}

// Допустимые команды (action)
const (
	logsMessage int = iota // Сообщение (лог сообщение)
	logsClose              // Закрытие службы логирования
)

// Служба логирования
func gologs() {

	go func() {
		//msg := fmt.Sprintf("%s\t[start]\r", lib.Time.Label())
		//logsSave(msg)
		for command := range commandlogsControl {
			switch command.action {
			case logsMessage:
				msg := logsMessageCalculate(command.log, command.level)
				logsSave(msg, command.level)
			case logsClose:
				//msg := fmt.Sprintf("%s\t[stop]\r", lib.Time.Label())
				//logsSave(msg)
				if fp != nil {
					fp.Close()
				}
				close(commandlogsControl)
				command.result <- true
			}
		}
	}()

}

//// Сохранение лога

var fp *os.File
var sys service.Logger

// Сохранение лога
//    + msg string сообщение
//    + level uint8 уровень сообщения
func logsSave(msg string, level uint8) {
	if cfg.Mode == `mixed` || cfg.Mode == `file` {
		fp.WriteString(msg)
	}
	if sys != nil && (cfg.Mode == `mixed` || cfg.Mode == `system`) {
		switch level {
		case 1, 2, 3:
			sys.Error(msg)
		case 4, 5:
			sys.Warning(msg)
		case 6:
			sys.Info(msg)
		}
	}
	// В режиме дебаг пишем в stdout все
	if cfg.Debug == true {
		fmt.Print(msg)
	}

}

//// Подготовка сообщения к сохранению

// формирование сообщения для логирования
func logsMessageCalculate(log *Log, level uint8) string {
	// временная отметка
	var prefix = lib.Time.Label()

	// формируем
	var logLine = fmt.Sprintf("%s [%d] %s\t%s\n", prefix, log.Code, logsLevel[level], log.Message)

	// информация режима debug
	if cfg.DebugDetail >= 1 {
		// информация о вызвавшей программе
		var info, _ = getCallerInfo(3)
		var debugLine = fmt.Sprintf("version: %s func: %s line: %d file: %s run: %d\n",
			info.Version,
			info.FuncName,
			info.LineNumber,
			info.FileName,
			info.Gorutines)
		if cfg.DebugDetail >= 2 {
			debugLine += "gorutines:\n"
			for i := range info.GorutinesInfo {
				debugLine += info.GorutinesInfo[i] + "\n"
			}
		}
		logLine += debugLine
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
