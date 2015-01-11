// Контроллер, Сессия пользователя.
//
// Регистрация, авторизация, восстановление пароля...
package controller

import (
	"bytes"

	"github.com/dchest/captcha"

	"core"
	"core/base/model"
	"lib"
	"lib/cache"
	"lib/mailer"
	typDb "types/db"
	typReq "types/request"
	typResp "types/response"
)

type Session struct {
	Session     *core.Session      // Сессия
	RW          *core.RW           // Управление вводом и выводом
	Controllers *typDb.Controllers // Соответсвующий контроллер из области данных по строковому ид.
}

// Создание контроллера
func NewSession(rw *core.RW, s *core.Session, c *typDb.Controllers) interface{} {
	var self = new(Session)
	self.RW = rw
	self.Session = s
	self.Controllers = c
	return self
}

////

// ApiMain Авторизация, выход, проверка токена с его пролонгацией
func (self *Session) ApiMain() (err error) {
	// диспетчер методов
	switch self.RW.Request.Method {
	case "GET":
		return self.apiMainGet()
	case "PUT":
		return self.apiMainPut()
	case "DELETE":
		return self.apiMainDelete()
	}
	return self.RW.ResponseJson(nil, 404, 520, `base.Session.ApiMain`)
}

// apiMainDelete Выход
func (self *Session) apiMainDelete() (err error) {
	if self.RW.Token == `` {
		self.RW.ResponseJson(nil, 409, 171)
	} else if self.Session.User.Id != core.Config.Auth.GuestUID {
		user := model.NewUsersType(*self.Session.User)
		user.Type.Token = ``
		user.Save(`Online`, `Id`)
		self.RW.RemCookie(core.Config.Auth.TokenCookie)
	}
	self.RW.ResponseJson(nil, 200, 0)
	return
}

// apiMainGet Проверка токена с его пролонгацией
func (self *Session) apiMainGet() (err error) {
	if self.RW.Token == `` {
		self.RW.ResponseJson(nil, 409, 171)
	} else if self.Session.User.Id == core.Config.Auth.GuestUID {
		self.RW.ResponseJson(nil, 404, 166)
	} else {
		user := model.NewUsersType(*self.Session.User)
		user.Type.DateOnline = lib.Time.Now()
		user.Save(`Online`, `Id`)
		self.RW.ResponseJson(nil, 200, 0)
	}
	return
}

// apiMainPut Авторизация
func (self *Session) apiMainPut() (err error) {
	// входящие данные
	var reqUser = new(typReq.Users)
	if err = self.RW.RequestJsonParse(reqUser); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}

	// блокировка частых ошибочных авторизаций
	ch := cache.Get(`ApiLogin`+self.RW.Request.RemoteAddr, cache.TH1)
	if ch == nil {
		cache.Set(`ApiLogin`+self.RW.Request.RemoteAddr, 1, cache.TH1)
	} else {
		// капча проверка
		if self.checkCaptcha(reqUser.CaptchaHash, reqUser.CaptchaValue) == false {
			return nil
		}
	}

	// проверяем пользователя
	if reqUser.Login == `` {
		self.RW.ResponseJson(nil, 404, 165)
		return
	}
	var u = model.SearchUsersLogin(reqUser.Login)
	if u == nil {
		self.RW.ResponseJson(nil, 404, 166)
		return
	}
	if lib.String.CreatePasswordHash(reqUser.Password) != u.Password {
		self.RW.ResponseJson(nil, 404, 167)
		return
	}
	// временно отключено
	if u.IsAccess == false {
		self.RW.ResponseJson(nil, 404, 168)
		return
	}
	if u.Del == true {
		self.RW.ResponseJson(nil, 404, 169)
		return
	}

	// авторизация
	u.IsActivated = true
	u.DateOnline = lib.Time.Now()
	u.Language = reqUser.Language
	u.Token = lib.String.CreatePasswordHash(lib.String.CreatePassword())
	var user = model.NewUsersType(*u)
	if err = user.Save(`All`, `Id`); err != nil {
		return self.RW.ResponseJson(nil, 500, 170)
	}
	self.Session.User = u

	// сессионая кука
	self.RW.SetCookie(core.Config.Auth.TokenCookie, user.Type.Token)

	// сброс счетчика попыток авторизаций
	cache.Rem(`ApiLogin` + self.RW.Request.RemoteAddr)

	// Все хорошо
	return self.RW.ResponseJson(user.Type.Token, 200, 0)
}

////

// ApiRecovery Восстановление пароля пользователя
func (self *Session) ApiRecovery() (err error) {
	// входящие данные
	var reqUser = new(typReq.Users)
	if err = self.RW.RequestJsonParse(reqUser); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}

	// блокировка частых ошибочных авторизаций
	ch := cache.Get(`Recovery`+self.RW.Request.RemoteAddr, cache.TH1)
	if ch == nil {
		cache.Set(`Recovery`+self.RW.Request.RemoteAddr, 1, cache.TH1)
	} else {
		// капча проверка
		if self.checkCaptcha(reqUser.CaptchaHash, reqUser.CaptchaValue) == false {
			return nil
		}
	}

	// проверяем пользователя
	if reqUser.Email == `` {
		return self.RW.ResponseJson(nil, 404, 176)
	}
	var u = model.SearchUsersEmail(reqUser.Email)
	if u == nil {
		return self.RW.ResponseJson(nil, 404, 166)
	}

	// восстановление
	var user = model.NewUsersType(*u)
	var password = lib.String.CreatePassword()
	user.Type.Password = lib.String.CreatePasswordHash(password)
	if err = user.Save(`Recovery`, `Id`); err != nil {
		return self.RW.ResponseJson(nil, 409, 178)
	}

	// сброс счетчика попыток восстановлений
	cache.Rem(`Recovery` + self.RW.Request.RemoteAddr)

	// поточвый шаблон
	msg := mailer.NewMessageTpl(`Восстановление на сайте: `+self.RW.Request.Host, self.Controllers.Path)
	msg.To(user.Type.Email, user.Type.Name+` `+user.Type.LastName)
	msg.Variables["u"] = user.Type
	msg.Variables["Site"] = self.RW.Request.Host
	msg.Variables[`Password`] = password
	if _, err := msg.Send(); err != nil {
		return self.RW.ResponseJson(nil, 409, 180, user.Type.Email)
	}

	// восстановление
	return self.RW.ResponseJson(nil, 200, 0)
}

////

// ApiCaptcha Получение капчи
func (self *Session) ApiCaptcha() (err error) {
	// Блокировка частых запросов капчи
	ch := cache.Get(`ApiCaptcha`+self.RW.Request.RemoteAddr, cache.TS03)
	if ch == nil {
		cache.Set(`ApiCaptcha`+self.RW.Request.RemoteAddr, 1, cache.TS03)
	} else {
		self.RW.ResponseJson(nil, 409, 159)
		return
	}
	// Формирование капчи
	var ret bytes.Buffer
	var keystring string
	keystring = captcha.New()
	if err = captcha.WriteImage(&ret, keystring, captcha.StdWidth, captcha.StdHeight); err != nil {
		return self.RW.ResponseJson(nil, 409, 140, err)
	}
	var response = &typResp.Captcha{
		CaptchaImage: lib.String.Base64Encode(ret.Bytes()),
		CaptchaHash:  keystring,
	}
	return self.RW.ResponseJson(response, 200, 0)
}

////

// ApiRegistration Регистрация нового пользователя
func (self *Session) ApiRegistration() (err error) {
	// входящие данные
	var reqUser = new(typReq.Users)
	if err = self.RW.RequestJsonParse(reqUser); err != nil {
		return self.RW.ResponseJson(nil, 409, 510)
	}

	// блокировка частых ошибочных регистраций
	ch := cache.Get(`ApiRegistration`+self.RW.Request.RemoteAddr, cache.TH1)
	if ch == nil {
		cache.Set(`ApiRegistration`+self.RW.Request.RemoteAddr, 1, cache.TH1)
	} else {
		// капча проверка
		if self.checkCaptcha(reqUser.CaptchaHash, reqUser.CaptchaValue) == false {
			return
		}
	}

	// referrer
	if self.RW.Token != `` {
		if u := model.SearchUsersToken(self.RW.Token); u != nil {
			reqUser.Users_Id = u.Id
		} else {
			self.RW.ResponseJson(nil, 404, 139)
			return
		}
	}

	if model.SearchUsersEmail(reqUser.Email) != nil {
		self.RW.ResponseJson(nil, 409, 162, reqUser.Email)
		return
	}

	if model.SearchUsersLogin(reqUser.Email) != nil {
		self.RW.ResponseJson(nil, 409, 163, reqUser.Email)
		return
	}
	var password = lib.String.CreatePassword()

	// регистрация
	var user = model.NewUsers(0)
	user.Type.Name = reqUser.Email
	user.Type.Email = reqUser.Email
	user.Type.Login = reqUser.Email
	user.Type.Password = lib.String.CreatePasswordHash(password)
	//user.Type.Token = lib.String.CreatePasswordHash(lib.String.CreatePassword())
	user.Type.IsCondition = true
	user.Type.IsAccess = true
	user.Type.Date = lib.Time.Now()
	if err = user.Save(`Registration`, `Id`); err != nil {
		return self.RW.ResponseJson(nil, 500, 560, `User.Registration`)
	}
	self.Session.User = user.Type

	// сессионая кука
	// self.RW.SetCookie(core.Config.Auth.TokenCookie, user.Type.Token)

	// сброс счетчика попыток регистрации
	cache.Rem(`ApiRegistration` + self.RW.Request.RemoteAddr)

	// поточвый шаблон
	msg := mailer.NewMessageTpl(`Регистрация на сайте: `+self.RW.Request.Host, self.Controllers.Path)
	msg.To(user.Type.Email, user.Type.Name+` `+user.Type.LastName)
	msg.Variables["u"] = user.Type
	msg.Variables["Site"] = self.RW.Request.Host
	msg.Variables[`Password`] = password
	if _, err := msg.Send(); err != nil {
		return self.RW.ResponseJson(nil, 500, 142, user.Type.Email)
	}

	// Все хорошо
	return self.RW.ResponseJson(nil, 200, 0)
}

////

// ApiUserCurrent Получение текущего пользователя
func (self *Session) ApiUserCurrent() (err error) {
	return self.RW.ResponseJson(self.Session.User, 200, 0)
}

// checkCaptchaDdos капча проверка
func (self *Session) checkCaptcha(captchaHash, capchaValue string) bool {
	if captchaHash == "" || capchaValue == "" {
		self.RW.ResponseJson(nil, 400, 160)
		return false
	}
	if false == captcha.VerifyString(captchaHash, capchaValue) {
		self.RW.ResponseJson(nil, 404, 161, capchaValue)
		return false
	}
	return true
}
