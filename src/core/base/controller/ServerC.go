package controller

import (
	"core"
	"core/controller"
	"lib/uploader"
)

type Server struct {
	controller.Controller
}

// ApiPing Проверка доступности сервера
func (self *Server) ApiPing() (err error) {
	return self.RW.ResponseJson(nil, 200, 0)
}

////

// ApiUpload Загрузка бинарных данных
func (self *Server) ApiUpload() (err error) {
	if self.RW.Token == `` { // токен не передан
		return self.RW.ResponseJson(nil, 409, 171)
	} else if self.Session.User.Id == core.Config.Auth.GuestUID { // токен не верный
		return self.RW.ResponseJson(nil, 404, 139, self.RW.Token)
	}
	var key string
	key, err = uploader.Upload(self.RW.Request, `myFile`)
	if err != nil {
		return self.RW.ResponseJson(nil, 409, 590, err)
	}
	return self.RW.ResponseJson(key, 200, 0)
}
