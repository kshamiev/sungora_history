package controller

import (
	"core"
	"lib/uploader"
	typDb "types/db"
)

type Server struct {
	Session     *core.Session      // Сессия
	RW          *core.RW           // Управление вводом и выводом
	Controllers *typDb.Controllers // Соответсвующий контроллер из области данных по строковому ид.
}

// Создание контроллера
func NewServer(rw *core.RW, s *core.Session, c *typDb.Controllers) interface{} {
	var self = new(Server)
	self.RW = rw
	self.Session = s
	self.Controllers = c
	return self
}

////

// ApiPing Проверка доступности сервера
func (self *Server) ApiPing() (err error) {
	return self.RW.ResponseJson(nil, 200, 0)
}

////

// ApiUpload Загрузка бинарных данных
func (self *Server) ApiUpload() (err error) {
	if self.RW.Token == `` { // токен не передан
		return self.RW.ResponseJson(nil, 409, 137)
	} else if self.Session.User.Id == core.Config.Auth.GuestUID { // токен не верный
		return self.RW.ResponseJson(nil, 404, 138, self.RW.Token)
	}
	var key string
	key, err = uploader.Upload(self.RW.Request, `myFile`)
	if err != nil {
		return self.RW.ResponseJson(nil, 409, 590, err)
	}
	return self.RW.ResponseJson(key, 200, 0)
}
