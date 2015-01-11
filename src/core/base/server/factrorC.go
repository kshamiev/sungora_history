// Переопределение функционала ядра
package server

import (
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"

	"app"
	"core"
	"core/base/model"
	"lib"
	"lib/cache"
	"lib/database"
	"lib/logs"
	typDb "types/db"
)

// FactorUriDefault Поиск и инциализация запрошенного урла
// Возвращает найденный урл
func FactorNewUri(r *http.Request) (self *typDb.Uri, uriSegment map[string]string) {
	// главная страница
	path := strings.Replace(r.URL.Path, `index.html`, ``, 1)
	if path == `/` {

		for i := range app.Routes {
			if (app.Routes[i].Uri == `/` || app.Routes[i].Uri == `/index.html`) &&
				(app.Routes[i].Domain == `` || -1 != strings.LastIndex(r.Host, app.Routes[i].Domain)) {

				index := sort.Search(len(app.Data.Uri), func(j int) bool { return app.Data.Uri[j].Id >= app.Routes[i].Id })
				if index >= len(app.Data.Uri) || app.Data.Uri[index].Id != app.Routes[i].Id {
					return
				}
				uriContentUpdate(r, app.Data.Uri[index])
				return app.Data.Uri[index], uriSegment
			}
		}
		return
	}
	// остальные страницы
	lreq := strings.Split(strings.Trim(path, `/`), `/`)
	lreqCnt := len(lreq)
	lreqLastIndex := len(lreq) - 1
	for i := range app.Routes {

		if app.Routes[i].Domain != `` && -1 == strings.LastIndex(r.Host, app.Routes[i].Domain) {
			continue
		}
		uriSegment = make(map[string]string)
		var flag = true

		luri := strings.Split(strings.Trim(strings.Replace(app.Routes[i].Uri, `index.html`, ``, 1), `/`), `/`)
		luriCnt := len(luri)

		for index := range luri {
			// получаем сегент
			var segment string
			if 2 < len(luri[index]) {
				segment = luri[index][1 : len(luri[index])-1]
			} else {
				segment = luri[index]

			}
			// запрос закончился
			if lreqLastIndex < index {
				// [] - необязательный параметр, * - все и вся // допустимо
				if `[`+segment+`]` == luri[index] || `*` == luri[index] {
					continue
				}
				flag = false
				break
			}
			// защита от дурака и кривых рук (склейка слешов)
			if lreq[index] == `` || luri[index] == `` {
				flag = false
				break
			}
			// прямое соответствие
			if luri[index] == lreq[index] {
				continue
			}
			// шаблоны {} обязательный, [] необязательный
			if `[`+segment+`]` == luri[index] || `{`+segment+`}` == luri[index] { // {} обязательный
				uriSegment[segment] = lreq[index]
				continue
			} else if `*` == luri[index] { // * - все и вся
				lreqCnt = 0
				luriCnt = 0
				break
			}
			// запрос не соответствует шаблону
			flag = false
			break
		}

		// соответствие найдено
		// lreqCnt <= luriCnt проверка что запрос не длинее чем роутинг и все совпало
		// 0 == lreqCnt && 0 == luriCnt - это * (все и вся)
		// поиск найденного Uri
		var index int
		if (true == flag && lreqCnt <= luriCnt) || (0 == lreqCnt && 0 == luriCnt) {
			index = sort.Search(len(app.Data.Uri), func(j int) bool { return app.Data.Uri[j].Id >= app.Routes[i].Id })
			if index >= len(app.Data.Uri) || app.Data.Uri[index].Id != app.Routes[i].Id {
				// TODO дописать логирование данной ошибки
				return
			}
			uriContentUpdate(r, app.Data.Uri[index])
			return app.Data.Uri[index], uriSegment
		}
	}
	return
}

// uriContentUpdate обновление контента uri по отношению к файловой системе
func uriContentUpdate(r *http.Request, u *typDb.Uri) (err error) {
	// определение пути
	var path = core.Config.View.Path
	l := strings.Split(r.Host, `.`)
	if len(l) > 2 {
		path += `/` + l[0] + `.` + l[1] + `.` + l[2]
	} else {
		path += `/` + r.Host
	}
	l = strings.Split(u.Uri, `[`)
	l = strings.Split(l[0], `/`)
	pos := len(l) - 1
	if l[pos] == `` {
		l[pos] = `index.html`
	} else if strings.LastIndex(l[pos], `.`) == -1 {
		l = append(l, `index.html`)
	}
	path += strings.Join(l, `/`)
	//
	var fi os.FileInfo
	if fi, err = os.Stat(path); err != nil {
		return
	}
	var con []byte
	if fi.ModTime().Sub(u.ContentTime) > 0 {
		if con, err = ioutil.ReadFile(path); nil != err {
			return logs.Error(146, path, err).Error
		} else {
			u.Content = con
			u.ContentTime = fi.ModTime()
			// сохраняем в БД
			if core.Config.Main.UseDb > 0 {
				var db database.DbFace
				if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
					return
				}
				// сохранение
				_, err = db.Update(u, `Uri`, `Id`)
				db.Free()
			}
		}
	}
	return
}

// FactorUsersDefault Поиск и инициализация пользователя
func FactorNewUsers(token string) *typDb.Users {
	var self *typDb.Users
	// поиск пользователя
	if token != `` {
		for i := range app.Data.Users {
			// пользователь найден
			if app.Data.Users[i].Token == token {
				self = app.Data.Users[i]
				if self.DateOnline.IsZero() == false && 0 < core.GetSessionTimeout(self.DateOnline) {
					// забываем пользователя по таймауту
					self.Token = ``
				} else {
					// фиксируем (обновляем) время активности
					self.DateOnline = lib.Time.Now()
				}
				if core.Config.Main.UseDb > 0 {
					db, err := database.NewDb(core.Config.Main.UseDb, 0)
					if err != nil {
						break
					}
					defer db.Free()
					db.Query(`base/users/2`, self.DateOnline, self.Token, self.Id)
					//self.Db.Save(`Online`, `Id`)
				}
				break
			}
		}
	}
	// гость (если пользователя не нашли)
	if self == nil || self.Token == `` {
		for i := range app.Data.Users {
			if app.Data.Users[i].Id == core.Config.Auth.GuestUID {
				self = app.Data.Users[i]
				break
			}
		}
	}
	return self
}

func FactorNewSession(rw *core.RW, uri *typDb.Uri, user *typDb.Users) *core.Session {

	// Получение сессии из кеша
	var self *core.Session
	var indexCache = `session-` + uri.Uri + `-` + user.Login
	var cacheSession = cache.Get(indexCache, cache.TH1)
	if cacheSession != nil {
		self = cacheSession.(*core.Session)
	} else {
		self = new(core.Session)

		// Группы пользовыателя
		self.UserGroupsMap = model.GetGroupsMap(user.Id)

		// Права
		self.Access = new(typDb.GroupsUri)
		self.AccessMap = make(map[string]bool)
		var controllerMapId = make(map[uint64]bool)
		var controllerMapIdDenied = make(map[uint64]bool)
		for _, val := range app.Data.GroupsUri {
			_, ok := self.UserGroupsMap[val.Groups_Id]
			if uri.Id == val.Uri_Id {
				// права доступа на этот урл и его действия
				if true == ok {
					if val.Get {
						self.Access.Get = true
						self.AccessMap[`GET`] = true
					}
					if val.Post {
						self.Access.Post = true
						self.AccessMap[`POST`] = true
					}
					if val.Put {
						self.Access.Put = true
						self.AccessMap[`PUT`] = true
					}
					if val.Delete {
						self.Access.Delete = true
						self.AccessMap[`DELETE`] = true
					}
					if val.Options {
						self.Access.Options = true
						self.AccessMap[`OPTIONS`] = true
					}
				}
				// контроллеры данного урла
				if 0 < val.Controllers_Id {
					controllerMapId[val.Controllers_Id] = true
				}
				continue
			}
			// запрещенные контроллеры для пользователя
			if true == ok && 0 < val.Controllers_Id && true == val.Disable {
				controllerMapIdDenied[val.Controllers_Id] = true
			}
		}
		// для группы разработчиков можно все
		if _, ok := self.UserGroupsMap[core.Config.Auth.DevUID]; ok == true {
			self.Access.Get = true
			self.Access.Post = true
			self.Access.Put = true
			self.Access.Delete = true
			self.Access.Options = true
			self.AccessMap[`GET`] = true
			self.AccessMap[`POST`] = true
			self.AccessMap[`PUT`] = true
			self.AccessMap[`DELETE`] = true
			self.AccessMap[`OPTIONS`] = true
		}

		// Контроллеры
		self.Controllers = make([]*typDb.Controllers, 0)
		var controllersBefore = make([]*typDb.Controllers, 0)
		var controllersMain = make([]*typDb.Controllers, 0)
		var controllersAfter = make([]*typDb.Controllers, 0)
		for _, c := range app.Data.Controllers {
			// пропускаем запрещенные контроллеры
			if _, ok := controllerMapIdDenied[c.Id]; true == ok {
				continue
			}
			// проверяем контроллер по домену
			if c.Domain != `` && -1 == strings.LastIndex(rw.Request.Host, c.Domain) {
				continue
			}
			// по умолчанию
			if true == c.IsDefault {
				if true == c.IsBefore {
					controllersBefore = append(controllersBefore, c)
				} else {
					controllersAfter = append(controllersAfter, c)
				}
				continue
			}
			// разрешенные контроллеры данного урла
			if _, ok := controllerMapId[c.Id]; true == ok {
				controllersMain = append(controllersMain, c)
			}
		}
		lib.Slice.SortingSliceAsc(controllersBefore, `Position`)
		lib.Slice.SortingSliceAsc(controllersMain, `Position`)
		lib.Slice.SortingSliceAsc(controllersAfter, `Position`)
		self.Controllers = append(self.Controllers, controllersBefore...)
		self.Controllers = append(self.Controllers, controllersMain...)
		self.Controllers = append(self.Controllers, controllersAfter...)

		cache.Set(indexCache, self, cache.TH1)
	}

	// Пользователь
	self.User = user

	// URI
	self.Uri = uri

	return self
}

// FactorAccess Проверка авторизации и прав доступа
func FactorAccess(rw *core.RW, session *core.Session) bool {
	// если не разработчик и uri авторизованный
	_, ok := session.UserGroupsMap[core.Config.Auth.DevUID]
	if ok == false && session.Uri.IsAuthorized == true {
		// если пользователь гость или нет права на метод запроса
		if _, ok := session.AccessMap[rw.Request.Method]; ok == false || session.User.Id == core.Config.Auth.GuestUID {
			return false
		}
		// Здесь можно реализовывать проверку прав бизнес логики
	}
	return true
}
