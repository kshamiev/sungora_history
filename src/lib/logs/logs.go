// Логирование и журналирвоание работы приложения.
package logs

import (
	"errors"
	"lib/i18n"
)

// Конфигурация лога
var cfg *Cfglogs

// Настройки конфигурации лога
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
	// максимальный размер файла лога для его дефрагметации.
	Size int8
}

func Init(cfgLogs *Cfglogs) {
	cfg = cfgLogs
	Base = NewLog(`en-en`, `base`)
}

var Base *Log

// Лог
type Log struct {
	Code       int    // Error code
	Message    string // Error message
	Err        error  // Error as error
	Lang       string // Префикс языка
	ModuleName string // Имя модуля
}

/*
// Инофрмационное сообщение
func Info(code int, messages ...interface{}) *Log {
	var message = searchMsg(code, messages...)
	if cfg.Level >= 6 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: 6}
	}
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

// Уведомление
func Notice(code int, messages ...interface{}) *Log {
	var message = searchMsg(code, messages...)
	if cfg.Level >= 5 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: 5}
	}
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

// Предупреждение
func Warning(code int, messages ...interface{}) *Log {
	var message = searchMsg(code, messages...)
	if cfg.Level >= 4 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: 4}
	}
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

// Ошибка
func Error(code int, messages ...interface{}) *Log {
	var message = searchMsg(code, messages...)
	if cfg.Level >= 3 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: 3}
	}
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

// Критическая ошибка
func Critical(code int, messages ...interface{}) *Log {
	var message = searchMsg(code, messages...)
	if cfg.Level >= 2 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: 2}
	}
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

// Фатальная ошибка
func Fatal(code int, messages ...interface{}) *Log {
	var message = searchMsg(code, messages...)
	if cfg.Level >= 1 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: 1}
	}
	var self = new(Log)
	self.Code = code
	self.Message = message
	self.Error = errors.New(message)
	return self
}

// Формирование сообщения
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
*/

////

func NewLog(lang, moduleName string) *Log {
	var self = new(Log)
	self.Lang = lang
	self.ModuleName = moduleName
	return self
}

// Инофрмационное сообщение
func (self *Log) Info(codeLocal int, params ...interface{}) *Log {
	var code, message = i18n.Message(self.ModuleName, self.Lang, codeLocal, params...)
	if cfg.Level >= 6 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: 6}
	}
	var s = NewLog(self.Lang, self.ModuleName)
	s.Code = code
	s.Message = message
	s.Err = errors.New(message)
	return s
}

// Уведомление
func (self *Log) Notice(codeLocal int, params ...interface{}) *Log {
	var code, message = i18n.Message(self.ModuleName, self.Lang, codeLocal, params...)
	if cfg.Level >= 5 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: 5}
	}
	var s = NewLog(self.Lang, self.ModuleName)
	s.Code = code
	s.Message = message
	s.Err = errors.New(message)
	return s
}

// Предупреждение
func (self *Log) Warning(codeLocal int, params ...interface{}) *Log {
	var code, message = i18n.Message(self.ModuleName, self.Lang, codeLocal, params...)
	if cfg.Level >= 4 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: 4}
	}
	var s = NewLog(self.Lang, self.ModuleName)
	s.Code = code
	s.Message = message
	s.Err = errors.New(message)
	return s
}

// Ошибка
func (self *Log) Error(codeLocal int, params ...interface{}) *Log {
	var code, message = i18n.Message(self.ModuleName, self.Lang, codeLocal, params...)
	if cfg.Level >= 3 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: 3}
	}
	var s = NewLog(self.Lang, self.ModuleName)
	s.Code = code
	s.Message = message
	s.Err = errors.New(message)
	return s
}

// Критическая ошибка
func (self *Log) Critical(codeLocal int, params ...interface{}) *Log {
	var code, message = i18n.Message(self.ModuleName, self.Lang, codeLocal, params...)
	if cfg.Level >= 2 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: 2}
	}
	var s = NewLog(self.Lang, self.ModuleName)
	s.Code = code
	s.Message = message
	s.Err = errors.New(message)
	return s
}

// Фатальная ошибка
func (self *Log) Fatal(codeLocal int, params ...interface{}) *Log {
	var code, message = i18n.Message(self.ModuleName, self.Lang, codeLocal, params...)
	if cfg.Level >= 1 {
		commandlogsControl <- command{action: logsMessage, code: code, message: message, level: 1}
	}
	var s = NewLog(self.Lang, self.ModuleName)
	s.Code = code
	s.Message = message
	s.Err = errors.New(message)
	return s
}
