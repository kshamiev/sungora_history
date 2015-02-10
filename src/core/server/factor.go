// server Управляющий контроллер
package server

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"

	"app"
	"core"
	"lib/database"
	typConfig "types/config"
	typDb "types/db"
)

// Определяемые пользователем функции ядра
var (
	// Инициализация урла (правило роутинга)
	FactorNewUri func(r *http.Request) (self *typDb.Uri, uriSegment map[string]string)
	// Инициализация пользователя
	FactorNewUsers func(token string) *typDb.Users
	// Инициализация сессии
	FactorNewSession func(rw *core.RW, uri *typDb.Uri, user *typDb.Users) *core.Session
	// Проверка авторизации и прав доступа
	FactorAccess func(rw *core.RW, session *core.Session) (flag bool)
)

type Server struct {
	*typConfig.Server
}

// Создание УП. Создается пакетом роутера.
func newServer(cfg *typConfig.Server) *Server {
	var self = new(Server)
	self.Server = cfg
	return self
}

// ServeHTTP Точка входа запроса (в приложение).
func (self *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var err error

	// 0 Инициализация управления потоком I/O
	var rw, ok = core.NewRW(w, r, self.Server)
	if ok == false {
		rw.ResponseError(404)
		return
	}

	// 1 Статика
	if core.Config.Main.ResponseStatic == true && rw.ResponseStatic() == true {
		return
	}

	// 2 Uri. Поиск и инициализация URI (ОПРЕДЕЛЯЕТСЯ МОДУЛЕМ)
	if FactorNewUri == nil {
		rw.ResponseError(500)
		return
	}
	var uri, uriSegment = FactorNewUri(r)
	// 404 и статика после ури
	if uri == nil {
		if rw.ResponseStatic() == false {
			rw.ResponseError(404)
		}
		return
	} else if uri.IsDisable == true {
		rw.ResponseError(404)
		return
	}
	// 301
	if uri.Redirect != `` {
		rw.Redirect(uri.Redirect)
		return
	}
	var uriParams, _ = url.ParseQuery(r.URL.Query().Encode())
	rw.Log.Info(0, `--- [`+r.Host+`] [`+r.Method+`] `+r.URL.Path)
	rw.Log.Info(119, uri.Id, uri.Uri)

	// 3 Users. Поиск и инициализация пользователя (ОПРЕДЕЛЯЕТСЯ МОДУЛЕМ)
	// токен
	if _, ok := uriSegment[`token`]; ok == false {
		if l, ok := uriParams[`access-token`]; ok == true && l[0] != "" {
			uriSegment[`token`] = l[0]
		} else if r.Header.Get(`X-Access-Token`) != `` {
			uriSegment[`token`] = r.Header.Get(`X-Access-Token`)
		} else {
			if cookie, err := r.Cookie(core.Config.Auth.TokenCookie); err == nil {
				uriSegment[`token`] = cookie.Value
			}
		}
	}
	rw.Log.Info(120, uriSegment[`token`])
	// пользователь
	var user *typDb.Users
	if FactorNewUsers == nil {
		rw.ResponseError(500)
		return
	} else {
		user = FactorNewUsers(uriSegment[`token`])
	}
	rw.Log.Info(121, user.Login, user.Id)
	// приоритетно переопределяем язык указанный в uri
	if user.Language != `` {
		uriSegment[`lang`] = user.Language
	}

	// 4 Инициализация параметров I/O
	rw.InitParams(uri, uriSegment, uriParams)

	// 5 Инициализация сессии (ОПРЕДЕЛЯЕТСЯ МОДУЛЕМ)
	var session *core.Session
	if FactorNewSession == nil {
		rw.ResponseError(500)
		return
	} else {
		session = FactorNewSession(rw, uri, user)
	}

	// 6 Проверка прав (ОПРЕДЕЛЯЕТСЯ МОДУЛЕМ)
	if FactorAccess == nil {
		rw.ResponseError(500)
		return
	} else if FactorAccess(rw, session) == false {
		rw.ResponseError(403)
		return
	}

	// 7 Работа контроллеров
	for _, c := range session.Controllers {
		if rw.Interrupt == true || rw.InterruptHard == true {
			break
		}
		rw.Log.Info(122, c.Id, c.Path)
		self.controllersContentUpdate(rw, c)
		if err := self.executeController(rw, session, c); err != nil {
			rw.Log.Error(123, c.Id, c.Path)
		}
	}
	if rw.InterruptHard == true {
		return
	}

	// 8 Вывод (как правило html)
	if rw.Content.Type == `text/html` {
		// дефолтовый шаблон
		if uri.Layout != `` {
			if err = self.layoutUpdate(rw, uri); err != nil {
				rw.ResponseError(500)
				return
			}
			index := rw.DocumentFolder + `_` + uri.Layout
			if err := rw.Content.ExecuteBytes(app.Data.UriDefault[index].Content); err != nil {
				rw.Log.Error(124, app.Data.UriDefault[index].Layout, err)
				rw.ResponseError(500)
				return
			}
		}
	}
	rw.Response()
	return
}

// executeController Выполнение контроллеров
func (self *Server) executeController(rw *core.RW, s *core.Session, c *typDb.Controllers) (err error) {
	defer core.RecoverErr(&err)

	// путь до контроллера и его метода в неправильном формате
	l := strings.Split(c.Path, `/`)
	if len(l) != 3 {
		return rw.Log.Critical(105, c.Path).Err
	}

	// нет такого контроллера
	ctrF, ok := app.Controller[l[0]+`/`+l[1]]
	if false == ok {
		return rw.Log.Critical(106, l[0], l[1]).Err
	}
	rw.InitLog(l[0], s.User.Login)

	// нет такого метода
	var ctr = ctrF(rw, s, c)
	objValue := reflect.ValueOf(ctr)
	met := objValue.MethodByName(l[2])
	if met.IsValid() == false {
		return rw.Log.Error(107, l[0], l[1], l[2]).Err
	}

	// оставлено для примера передачи параметров в метод
	var params []interface{}
	var in = make([]reflect.Value, 0)
	for i := range params {
		in = append(in, reflect.ValueOf(params[i]))
	}

	// вызов
	out := met.Call(in)
	if nil == out[0].Interface() {
		return nil
	} else {
		return out[0].Interface().(error)
	}
}

// дефолтовый шаблон
func (self *Server) layoutUpdate(rw *core.RW, uri *typDb.Uri) (err error) {
	var fi os.FileInfo
	// поиск шаблона в вверх по ФС
	var pathOrigin = rw.DocumentRoot + uri.Layout
	path := pathOrigin
	chunks := strings.Split(path, `/`)
	for fi, err = os.Stat(path); err != nil && 3 < len(chunks); fi, err = os.Stat(path) {
		fileName := chunks[len(chunks)-1]
		chunks = chunks[:len(chunks)-2]
		chunks = append(chunks, fileName)
		path = strings.Join(chunks, `/`)
	}
	if err != nil {
		return rw.Log.Error(125, uri.Layout, pathOrigin).Err
	}
	var index = rw.DocumentFolder + `_` + uri.Layout
	if _, ok := app.Data.UriDefault[index]; ok == false {
		app.Data.UriDefault[index] = new(typDb.Uri)
	}
	var con []byte
	if fi.ModTime().Sub(app.Data.UriDefault[index].ContentTime) > 0 { // нашли, проверяем и обновляем
		if con, err = ioutil.ReadFile(path); nil != err {
			return rw.Log.Error(125, uri.Layout, err.Error()).Err
		} else {
			app.Data.UriDefault[index].Layout = uri.Layout
			app.Data.UriDefault[index].Content = con
			app.Data.UriDefault[index].ContentTime = fi.ModTime()
		}
	}
	return
}

// controllersContentUpdate обновление контента контроллеров по отношению к файловой системе
func (self *Server) controllersContentUpdate(rw *core.RW, c *typDb.Controllers) (err error) {
	if c.Id == 0 {
		return
	}
	var fi os.FileInfo
	var path = core.Config.View.Tpl + `/` + c.Path + `.html`
	if fi, err = os.Stat(path); err != nil {
		return
	}
	var con []byte
	if fi.ModTime().Sub(c.ContentTime) > 0 {
		if con, err = ioutil.ReadFile(path); nil != err {
			return rw.Log.Error(126, path, err).Err
		} else {
			c.Content = string(con)
			c.ContentTime = fi.ModTime()
			// сохраняем в БД
			if core.Config.Main.UseDb > 0 {
				var db database.DbFace
				if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
					return
				}
				// сохранение
				_, err = db.Update(c, `Controllers`, `Id`)
				db.Free()
			}
		}
	}
	return
}
