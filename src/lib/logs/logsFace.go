// logs programm interfaces
package logs

import (
	"errors"
	"fmt"
)

var cfglogs *Cfglogs

// Настройка лог службы
type Cfglogs struct {
	// Режим отладки приложения. Вывод всех логов в консоль
	Debug bool
	// Детализация дебага
	// 0 - детализация отключена (по умолчанию)
	// 1 - трейс вызвавшей функции, файла, номера строки в файле
	// 2 - трейс всего стека
	DebugDetail int64
	// Режим логирования сообщений системы (6 по умолчанию)
	// Info     = 6 - Сообщения о всех шагах работы системы, от Info до Fatal (по умолчанию)
	// Notice   = 5 - Сообщения о наиболее важных шагах работы системы, от Notice до Fatal
	// Warning  = 4 - Сообщения о не важных ошибках системы, от Warning до Fatal
	// Error    = 3 - Сообщения об ошибках системы последствия которых важны, но не приводят к деградации функционала системы, от Error до Fatal
	// Critical = 2 - Сообщения о критичных ошибках системы без которых возможно продолжения работы но с урезанным функционалом, от Critical до Fatal
	// Fatal    = 1 - Сообщения о фатальных ошибках системы, после фатальной ошибки работа приложения не возможна и немедленно завершается
	Level int64
	// Файл системного журнала приложения, если указан, то используется он,
	// если не указан, логи отправляются в журнал операционной системы
	File string
	// Режим записи логов
	// file - запись логов только в файл лога (по умолчанию), если файл лога не указан то режим переключается на system
	// system - запись логов только в системный журнал операционной системы
	// mixed - запись логов как в файл так и в системный журнал операционной системы
	Mode string
}

func Init(cfg *Cfglogs) {
	cfglogs = cfg
	if cfglogs.File == `` {
		cfglogs.Mode = `system`
	}
}

//var logsLevel = map[string]int8{
//	`[info]`:     6,
//	`[notice]`:   5,
//	`[warning]`:  4,
//	`[error]`:    3,
//	`[critical]`: 2,
//	`[fatal]`:    1,
//}

type Log struct {
	Code    int    // Log error code
	Message string // Log error message
	Error   error  // Log error as error
}

func NewLog(code int, message string) *Log {
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

// Direct Прямолинейный универсальный лог
//func Direct(code int, message, level string) *Log {
//	level = `[` + level + `]`
//	if cfglogs.Level >= logsLevel[level] {
//		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: level}
//	}
//	var self = new(Log)
//	self.Code = code
//	self.Message = message
//	self.Error = errors.New(message)
//	return self
//}

// Info Инофрмационное сообщение
func Info(code int, messages ...interface{}) *Log {
	var message = searchLog(code, messages...)
	if cfglogs.Level >= 6 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: `[info]`}
	}
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

// Notice Уведомление
func Notice(code int, messages ...interface{}) *Log {
	var message = searchLog(code, messages...)
	if cfglogs.Level >= 5 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: `[notice]`}
	}
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

// Warning Предупреждение
func Warning(code int, messages ...interface{}) *Log {
	var message = searchLog(code, messages...)
	if cfglogs.Level >= 4 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: `[warning]`}
	}
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

// Error Ошибка
func Error(code int, messages ...interface{}) *Log {
	var message = searchLog(code, messages...)
	if cfglogs.Level >= 3 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: `[error]`}
	}
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

// Critical Критическая ошибка
func Critical(code int, messages ...interface{}) *Log {
	var message = searchLog(code, messages...)
	if cfglogs.Level >= 2 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: `[critical]`}
	}
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

// Fatal Фатальная ошибка
func Fatal(code int, messages ...interface{}) *Log {
	var message = searchLog(code, messages...)
	if cfglogs.Level >= 1 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: `[fatal]`}
	}
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

func searchLog(code int, params ...interface{}) (message string) {
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

// logsStartGo Запуск работы службы логирования
func GoStart() {
	gologs()
}

// logsCloseGo Завершение работы службы логирования
func GoClose() bool {
	reply := make(chan interface{})
	commandlogsControl <- command{action: logsClose, result: reply}
	return (<-reply).(bool)
}
