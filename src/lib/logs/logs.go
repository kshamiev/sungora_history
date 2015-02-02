// Логирование и журналирвоание работы приложения.
package logs

import (
	"errors"
	"lib"
	"lib/i18n"
)

// Конфигурация лога
var cfg *Cfglogs

// Структура конфигурации лога
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
	// Префикс языка для системных логов
	Lang string
}

// Инициализация библиотеки, системного или базового лога.
//    + cfgLogs *Cfglogs конфигурация лога
func Init(cfgLogs *Cfglogs) {
	cfg = cfgLogs
	Base = NewLog(`base`, cfg.Lang)
	Base.label = `base`
}

// Системный или базовый лог
var Base *Log

////

// Структура лога
type Log struct {
	Code       int    // Error code
	Message    string // Error message
	Err        error  // Error as error
	lang       string // Префикс языка
	moduleName string // Имя модуля
	label      string // метка лога в рамках обработки запроса (сессии)
}

// Создание лога
//    + moduleName string имя модуля
//    - *Log объект лога
func NewLog(moduleName, lang string) *Log {
	var self = new(Log)
	self.lang = lang
	self.moduleName = moduleName
	self.label = lib.String.CreatePassword()
	return self
}

// Инициализация лога
//    + moduleName string имя модуля
//    + lang string язык
func (self *Log) Init(moduleName, lang string) {
	self.lang = lang
	self.moduleName = moduleName
}

// Инофрмационное сообщение
//    + codeLocal int локальный (в рамках модуля) код сообщения
//    + params ...interface{} параметры вставляемые в сообщение
//    - *Log объект лога
func (self *Log) Info(codeLocal int, params ...interface{}) *Log {
	self.Code, self.Message = i18n.Message(self.moduleName, self.lang, codeLocal, params...)
	if cfg.Level >= 6 {
		commandlogsControl <- command{action: logsMessage, log: self, level: 6}
	}
	self.Err = errors.New(self.Message)
	return self
}

// Уведомление
//    + codeLocal int локальный (в рамках модуля) код сообщения
//    + params ...interface{} параметры вставляемые в сообщение
//    - *Log объект лога
func (self *Log) Notice(codeLocal int, params ...interface{}) *Log {
	self.Code, self.Message = i18n.Message(self.moduleName, self.lang, codeLocal, params...)
	if cfg.Level >= 5 {
		commandlogsControl <- command{action: logsMessage, log: self, level: 5}
	}
	self.Err = errors.New(self.Message)
	return self
}

// Предупреждение
//    + codeLocal int локальный (в рамках модуля) код сообщения
//    + params ...interface{} параметры вставляемые в сообщение
//    - *Log объект лога
func (self *Log) Warning(codeLocal int, params ...interface{}) *Log {
	self.Code, self.Message = i18n.Message(self.moduleName, self.lang, codeLocal, params...)
	if cfg.Level >= 4 {
		commandlogsControl <- command{action: logsMessage, log: self, level: 4}
	}
	self.Err = errors.New(self.Message)
	return self
}

// Ошибка
//    + codeLocal int локальный (в рамках модуля) код сообщения
//    + params ...interface{} параметры вставляемые в сообщение
//    - *Log объект лога
func (self *Log) Error(codeLocal int, params ...interface{}) *Log {
	self.Code, self.Message = i18n.Message(self.moduleName, self.lang, codeLocal, params...)
	if cfg.Level >= 3 {
		commandlogsControl <- command{action: logsMessage, log: self, level: 3}
	}
	self.Err = errors.New(self.Message)
	return self
}

// Критическая ошибка
//    + codeLocal int локальный (в рамках модуля) код сообщения
//    + params ...interface{} параметры вставляемые в сообщение
//    - *Log объект лога
func (self *Log) Critical(codeLocal int, params ...interface{}) *Log {
	self.Code, self.Message = i18n.Message(self.moduleName, self.lang, codeLocal, params...)
	if cfg.Level >= 2 {
		commandlogsControl <- command{action: logsMessage, log: self, level: 2}
	}
	self.Err = errors.New(self.Message)
	return self
}

// Фатальная ошибка
//    + codeLocal int локальный (в рамках модуля) код сообщения
//    + params ...interface{} параметры вставляемые в сообщение
//    - *Log объект лога
func (self *Log) Fatal(codeLocal int, params ...interface{}) *Log {
	self.Code, self.Message = i18n.Message(self.moduleName, self.lang, codeLocal, params...)
	if cfg.Level >= 1 {
		commandlogsControl <- command{action: logsMessage, log: self, level: 1}
	}
	self.Err = errors.New(self.Message)
	return self
}
