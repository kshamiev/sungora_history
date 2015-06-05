package logs

// Сулжба логирования

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/kardianos/service"

	"lib"
)

var cntStart = uint8(0)

// Запуск работы службы логирования
//	+ logSys service.Logger интерфейс логирования в системный лог
func GoStart(logSys service.Logger) {
	cntStart++
	if cntStart == 1 {
		if cfg.Mode == `file` && fp == nil {
			os.MkdirAll(filepath.Dir(cfg.File), 0755)
			fp, _ = os.OpenFile(cfg.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		}
		sys = logSys
		gologs()
	}
}

// Завершение работы службы логирования
//	- bool флаг успешности операции
func GoClose() bool {
	if cntStart == 0 {
		reply := make(chan interface{})
		commandControl <- command{action: logsClose, result: reply}
		return (<-reply).(bool)
	}
	cntStart--
	return true
}

// Канал передачи данных
var commandControl = make(chan command, 1000)

// Структура пересылаемых данных
type command struct {
	action int
	log    Log
	level  uint8
	result chan<- interface{}
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
		for command := range commandControl {
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
				close(commandControl)
				command.result <- true
			}
		}
	}()

}

//// Сохранение лога

// Справочник уровня ошибок
var logsLevel = map[uint8]string{
	8: `database`,
	7: `journal `,
	6: `info    `,
	5: `notice  `,
	4: `warning `,
	3: `error   `,
	2: `critical`,
	1: `fatal   `,
}

var fp *os.File
var sys service.Logger

// Сохранение лога
//    + msg string сообщение
//    + level uint8 уровень сообщения
func logsSave(msg string, level uint8) {
	if cfg.Mode == `mixed` || cfg.Mode == `file` {
		fp.WriteString(msg)
		if inf, err := fp.Stat(); err == nil && inf.Size() > int64(cfg.Size)*1000000 {
			if err := fp.Close(); err == nil {
				t := lib.Time.Now()
				d := fmt.Sprintf("%04d_%02d_%02d_%02d_%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
				path := strings.Replace(fp.Name(), `.log`, `_`+d+`.log`, -1)
				os.Rename(fp.Name(), path)
				fp, _ = os.OpenFile(cfg.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
			}
		}
	}
	if sys != nil && (cfg.Mode == `mixed` || cfg.Mode == `system`) {
		switch level {
		case 1, 2, 3:
			sys.Error(msg)
		case 4, 5:
			sys.Warning(msg)
		case 6, 7, 8:
			sys.Info(msg)
		}
	}
	// В режиме дебаг пишем в stdout все
	if cfg.Debug == true {
		fmt.Print(msg)
	}

}

//// Подготовка сообщения к сохранению

// Формирование сообщения для логирования
func logsMessageCalculate(log Log, level uint8) string {
	// временная отметка
	var prefix = lib.Time.LabelFull()

	// формируем
	log.Message = strings.Replace(log.Message, "\r\n", " ", -1)
	log.Message = strings.Replace(log.Message, "\n", " ", -1)
	log.Message = strings.Replace(log.Message, "\t", " ", -1)
	var logLine = fmt.Sprintf("%s%s%s%s%d%s%s%s%s\n",
		prefix, cfg.Separator,
		log.label, cfg.Separator,
		log.Code, cfg.Separator,
		logsLevel[level], cfg.Separator,
		log.Message,
	)

	// информация режима debug
	if cfg.DebugDetail >= 1 {
		//var tr *trace
		//tr = newTrace(3)
		//logLine += fmt.Sprintf("%s %s (%d)\n", tr.function, tr.file, tr.line)
	}
	return logLine
}

type trace struct {
	function string
	file     string
	line     int
}

// Получение информаци о вызвавшем лог
func newTrace(level int) *trace {
	var ok bool
	var pc uintptr
	var self = new(trace)
	pc, self.file, self.line, ok = runtime.Caller(level)
	if ok == true {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			self.function = fn.Name()
		}
		//var buf = make([]byte, 1<<16)
		//i := runtime.Stack(buf, true)
		//var info string = string(buf[:i])
		//fmt.Println(info)
	}
	return self
}
