// Интерфейс всех котроллеров и базовый контроллер его реализующий.
//
// Базовый контроллер является встроенным для всех контроллеров приложения.
// config.go Временная Конфигурация роутинга и контроллеров
// load.go Загрузка, инициализация и проверка перед запуском приложения
package controller

import (
	"core"
	typDb "types/db"
)

// Контроллеры
var Controllers = make(map[string]ControllerFace)

// Интерфейс для всех котроллеров
type ControllerFace interface {
	Init(rw *core.RW, session *core.Session, c *typDb.Controllers)
}

// Базовый контроллер
type Controller struct {
	Session     *core.Session      // Сессия
	RW          *core.RW           // Управление вводом и выводом
	Controllers *typDb.Controllers // Соответсвующий контроллер из области данных по строковому ид.
}

// Init(*core.RW, *core.Session, *typDb.Controllers)
// Инициализация контроллера
func (self *Controller) Init(rw *core.RW, session *core.Session, c *typDb.Controllers) {
	self.RW = rw
	self.Session = session
	self.Controllers = c
}
