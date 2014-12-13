/*
Базовый контроллер

*/
package controller

import (
	"core"
	typDb "types/db"
	//"core/i18n"
	//"lib/logs"
)

// Контроллеры
var Controllers = make(map[string]ControllerFace)

type ControllerFace interface {
	Init(rw *core.RW, session *core.Session, c *typDb.Controllers)
}

type Controller struct {
	Session     *core.Session      // Сессия
	RW          *core.RW           // Потоки ввода вывода
	Controllers *typDb.Controllers // Собственный контроллер
}

// Инициализация контроллера
func (self *Controller) Init(rw *core.RW, session *core.Session, c *typDb.Controllers) {
	self.RW = rw
	self.Session = session
	self.Controllers = c
}
