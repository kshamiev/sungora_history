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
	"lib/logs"
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
	FactorAccess func(session *core.Session, method string) (interruptHard bool)
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
	logs.Info(0, `--- [`+r.Method+`] `+r.URL.Path)

	// 1 Статика
	if core.Config.Main.ResponseStatic == true && responseStatic(w, r) == true {
		return
	}

	// 2 Uri. Поиск и инициализация URI (ОПРЕДЕЛЯЕТСЯ МОДУЛЕМ)
	if FactorNewUri == nil {
		core.NewRWSimple(w, r).ResponseError(500)
		return
	}
	var uri, uriSegment = FactorNewUri(r)
	// 404 и статика после ури
	if uri == nil {
		if responseStatic(w, r) == false {
			core.NewRWSimple(w, r).ResponseError(404)
		}
		return
	}
	if uri.IsDisable == true {
		core.NewRWSimple(w, r).ResponseError(404)
		return
	}
	// 301
	if uri.Redirect != `` {
		core.NewRWSimple(w, r).Redirect(uri.Redirect)
		return
	}
	var uriParams, _ = url.ParseQuery(r.URL.Query().Encode())
	logs.Info(131, uri.Id, uri.Uri)

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
	logs.Info(132, uriSegment[`token`])
	// пользователь
	var user *typDb.Users
	if FactorNewUsers == nil {
		core.NewRWSimple(w, r).ResponseError(500)
		return
	} else {
		user = FactorNewUsers(uriSegment[`token`])
	}
	logs.Info(134, user.Login, user.Id)
	// приоритетно переопределяем язык указанный в uri
	if user.Language != `` {
		uriSegment[`lang`] = user.Language
	}

	// 4 Инициализация управления потоком I/O
	var rw = core.NewRW(w, r, uri, uriSegment, uriParams)

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
	} else if FactorAccess(session, rw.Token) == false {
		rw.ResponseError(403)
		return
	}

	// 7 Работа контроллеров
	for _, c := range session.Controllers {
		if rw.Interrupt == true || rw.InterruptHard == true {
			break
		}
		logs.Info(126, c.Id, c.Path)
		controllersContentUpdate(c)
		if err := executeController(rw, session, c); err != nil {
			logs.Error(127, c.Id, c.Path)
		}
	}
	if rw.InterruptHard == true {
		return
	}

	// 8 Вывод (как правило html)
	if rw.Content.Type == `text/html` {
		// дефолтовый шаблон
		if uri.Layout != `` {
			if err = layoutUpdate(rw, uri); err != nil {
				rw.ResponseError(500)
				return
			}
			index := rw.DocumentFolder + `_` + uri.Layout
			if err := rw.Content.ExecuteBytes(app.Data.UriDefault[index].Content); err != nil {
				logs.Error(128, app.Data.UriDefault[index].Layout, err)
				rw.ResponseError(500)
				return
			}
		}
	}
	rw.Response(true)
	return
}

// executeController Выполнение контроллеров
func executeController(rw *core.RW, s *core.Session, c *typDb.Controllers) (err error) {
	defer core.RecoverErr(&err)

	// путь до контроллера и его метода в неправильном формате
	l := strings.Split(c.Path, `/`)
	if len(l) != 3 {
		return logs.Error(172, c.Path).Error
	}

	// нет такого контроллера
	ctrF, ok := app.Controller[l[0]+`/`+l[1]]
	if false == ok {
		return logs.Error(173, l[0], l[1]).Error
	}
	//rw.Logs.ModuleName = l[0]

	// нет такого метода
	var ctr = ctrF(rw, s, c)
	objValue := reflect.ValueOf(ctr)
	met := objValue.MethodByName(l[2])
	if met.IsValid() == false {
		return logs.Error(174, l[0], l[1], l[2]).Error
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
func layoutUpdate(rw *core.RW, uri *typDb.Uri) (err error) {
	var fi os.FileInfo
	// поиск шаблона в вверх по ФС
	var pathOrigin = rw.DocumentRoot + uri.Layout
	//if uri.Domain == `` {
	//	l := strings.Split(rw.Request.Host, `.`)
	//	path = core.Config.View.Path + `/` + l[len(l)-2] + `.` + l[len(l)-1] + uri.Layout
	//} else {
	//	path = rw.DocumentRoot + uri.Layout
	//}
	path := pathOrigin
	chunks := strings.Split(path, `/`)
	for fi, err = os.Stat(path); err != nil && 3 < len(chunks); fi, err = os.Stat(path) {
		fileName := chunks[len(chunks)-1]
		chunks = chunks[:len(chunks)-2]
		chunks = append(chunks, fileName)
		path = strings.Join(chunks, `/`)
	}
	if err != nil {
		return logs.Error(147, uri.Layout, pathOrigin).Error
	}
	var index = rw.DocumentFolder + `_` + uri.Layout
	if _, ok := app.Data.UriDefault[index]; ok == false {
		app.Data.UriDefault[index] = new(typDb.Uri)
	}
	var con []byte
	if fi.ModTime().Sub(app.Data.UriDefault[index].ContentTime) > 0 { // нашли, проверяем и обновляем
		if con, err = ioutil.ReadFile(path); nil != err {
			return logs.Error(147, uri.Layout, err).Error
		} else {
			app.Data.UriDefault[index].Layout = uri.Layout
			app.Data.UriDefault[index].Content = con
			app.Data.UriDefault[index].ContentTime = fi.ModTime()
		}
	}
	return
}

// controllersContentUpdate обновление контента контроллеров по отношению к файловой системе
func controllersContentUpdate(c *typDb.Controllers) (err error) {
	if c.Id == 0 {
		return
	}
	var fi os.FileInfo
	var path = core.Config.View.Tpl + `/` + c.Path + `.html`
	if fi, err = os.Stat(path); err != nil {
		// return logs.Error(145, path, err).Error
		return
	}
	var con []byte
	if fi.ModTime().Sub(c.ContentTime) > 0 {
		if con, err = ioutil.ReadFile(path); nil != err {
			return logs.Error(145, path, err).Error
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

// responseStatic(*http.Request) bool
// Отдаем статику (css, images, js, download ...)
func responseStatic(w http.ResponseWriter, r *http.Request) bool {
	var host = strings.Split(r.Host, `:`)[0]
	var pathAbs = core.Config.View.Path + `/` + host + r.URL.Path
	l := strings.Split(host, `.`)
	if len(l) > 2 {
		pathAbs = core.Config.View.Path + `/` + l[0] + `.` + l[1] + `.` + l[2] + r.URL.Path
	} else if host != `localhost` {
		pathAbs = core.Config.View.Path + `/www.` + host + r.URL.Path
	}
	pathAbs = strings.TrimRight(pathAbs, `/`)
	if fi, e := os.Stat(pathAbs); e == nil {
		if fi.IsDir() == false {
			core.NewRWSimple(w, r).ResponseFile(pathAbs, true)
			return true
		} else {
			pathAbs += `/index.html`
			if _, e := os.Stat(pathAbs); e == nil {
				core.NewRWSimple(w, r).ResponseFile(pathAbs, true)
				return true
			}
		}
	}
	return false
}
