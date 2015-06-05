package config

import (
	"app"
	"core"
	"core/model"
	"lib"
	"lib/database"
	"lib/logs"
	typDb "types/db"
	typReq "types/request"
)

// Конфигурация контроллеров (Controllers)
var ConfigControllers = make(map[string]typDb.Controllers)

// Конфигурация разделов (Uri)
var ConfigUri = make(map[string]typDb.Uri)

// Инициализация данных и проверки
func App() (err error) {
	// Инициализация системных данных
	if core.Config.Main.UseDb > 0 {
		err = loadDataFromDb()
	} else {
		err = loadDataFromMemory()
	}
	if err != nil {
		return
	}

	// инициализация роутинга
	app.ReInitRoute()

	// проверка контроллеров
	app.CheckControllers()

	return
}

// Инициализация данных (с использованием БД)
func loadDataFromDb() (err error) {
	var controllersId = make(map[string]uint64)
	var uriId = make(map[string]uint64)
	var db database.DbFace
	if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
		return
	}
	defer db.Free()

	// загружаем системные данные из БД
	if err = db.SelectData(app.Data); err != nil {
		return
	}

	// Сортируем uri & controllers и получаем postion
	var res = new(typReq.DbResult)
	if err = db.Select(res, `base/uri/5`); err != nil {
		return
	}
	var postionUri int32 = int32(res.Max) + 1
	if err = db.Select(res, `base/controllers/4`); err != nil {
		return
	}
	var postionControllers int32 = int32(res.Max) + 1
	//
	app.Data.MaxPostion = make(map[string]int32)

	// (Контроллеры)
	for path := range ConfigControllers {
		var flag bool
		for _, c := range app.Data.Controllers {
			if c.Path == path {
				// обновляем информацию в БД
				c.Date = lib.Time.Now()
				if _, err = db.Update(c, `Controllers`, `Id`); err != nil {
					return
				}
				controllersId[path] = c.Id
				flag = true
				break
			}
		}
		// добавляем несуществующие из программы
		if flag == false {
			var ctrl = ConfigControllers[path]
			ctrl.Date = lib.Time.Now()
			ctrl.Path = path
			ctrl.Position = postionControllers
			if ctrl.Id, err = db.Insert(&ctrl, `Controllers`); err != nil {
				return
			}
			app.Data.Controllers = append(app.Data.Controllers, &ctrl)
			controllersId[path] = ctrl.Id
			postionControllers++
		}
	}
	app.Data.MaxPostion[`Controllers`] = postionControllers

	// (Роутинг Uri)
	for uri := range ConfigUri {
		var flag bool
		for _, u := range app.Data.Uri {
			if u.Uri == uri {
				// обновляем информацию в БД
				u.Uri = uri
				if _, err = db.Update(u, `Uri`, `Id`); err != nil {
					return
				}
				uriId[uri] = u.Id
				flag = true
				break
			}
		}
		// добавляем несуществующие из программы
		if flag == false {
			var u = ConfigUri[uri]
			u.Uri = uri
			if u.ContentType == `` {
				u.ContentType = `text/html`
			}
			if u.ContentEncode == `` {
				u.ContentEncode = `utf-8`
			}
			u.Position = postionUri
			if u.Id, err = db.Insert(&u, `Uri`); err != nil {
				return
			}
			app.Data.Uri = append(app.Data.Uri, &u)
			uriId[uri] = u.Id
			postionUri++
		}
	}
	app.Data.MaxPostion[`Uri`] = postionUri

	// Связи между роутингом и контроллерами
	for uri := range ConfigUri {
		for _, ctrlPath := range ConfigUri[uri].Controllers {
			var flag bool
			for _, acc := range app.Data.GroupsUri {
				if acc.Uri_Id == uriId[uri] && acc.Controllers_Id == controllersId[ctrlPath] {
					flag = true
					break
				}
			}
			// добавляем несуществующие из программы
			if flag == false {
				var acc = new(typDb.GroupsUri)
				acc.Uri_Id = uriId[uri]
				acc.Controllers_Id = controllersId[ctrlPath]
				if _, err = db.Insert(acc, `GroupsUri`); err != nil {
					return
				}
				app.Data.GroupsUri = append(app.Data.GroupsUri, acc)
			}
		}
	}

	// Пользователи
	var flagGuest, flagDev bool
	for i := range app.Data.Users {
		if app.Data.Users[i].Id == core.Config.Auth.GuestUID {
			flagGuest = true
		}
		if app.Data.Users[i].Id == core.Config.Auth.DevUID {
			flagDev = true
		}
	}
	if flagDev == false {
		var u = new(typDb.Users)
		u.Id = core.Config.Auth.DevUID
		u.Login = `developer`
		u.Email = `developer@developer.developer`
		if _, err = db.Insert(u, `Users`); err != nil {
			return
		}
		app.Data.Users = append(app.Data.Users, u)
	}
	if flagGuest == false {
		var u = new(typDb.Users)
		u.Id = core.Config.Auth.GuestUID
		u.Login = `guest`
		u.Email = `guest@guest.guest`
		if _, err = db.Insert(u, `Users`); err != nil {
			return
		}
		app.Data.Users = append(app.Data.Users, u)
	}

	// Группы
	flagGuest, flagDev = false, false
	for i := range app.Data.Groups {
		if app.Data.Groups[i].Id == core.Config.Auth.GuestUID {
			flagGuest = true
		}
		if app.Data.Groups[i].Id == core.Config.Auth.DevUID {
			flagDev = true
		}
	}
	if flagDev == false {
		var g = new(typDb.Groups)
		g.Id = core.Config.Auth.DevUID
		g.Name = `developer`
		if _, err = db.Insert(g, `Groups`); err != nil {
			return
		}
		app.Data.Groups = append(app.Data.Groups, g)
	}
	if flagGuest == false {
		var g = new(typDb.Groups)
		g.Id = core.Config.Auth.GuestUID
		g.Name = `guest`
		if _, err = db.Insert(g, `Groups`); err != nil {
			return
		}
		app.Data.Groups = append(app.Data.Groups, g)
	}

	// Связь пользователи и группы
	flagGuest, flagDev = false, false
	for _, acc := range app.Data.UsersGroups {
		if acc.Users_Id == core.Config.Auth.GuestUID && acc.Groups_Id == core.Config.Auth.GuestUID {
			flagGuest = true
		}
		if acc.Users_Id == core.Config.Auth.DevUID && acc.Groups_Id == core.Config.Auth.DevUID {
			flagDev = true
		}
	}
	if flagDev == false {
		var acc = new(typDb.UsersGroups)
		acc.Users_Id = core.Config.Auth.DevUID
		acc.Groups_Id = core.Config.Auth.DevUID
		app.Data.UsersGroups = append(app.Data.UsersGroups, acc)
		if _, err = db.Insert(acc, `UsersGroups`); err != nil {
			return
		}
	}
	if flagGuest == false {
		var acc = new(typDb.UsersGroups)
		acc.Users_Id = core.Config.Auth.GuestUID
		acc.Groups_Id = core.Config.Auth.GuestUID
		app.Data.UsersGroups = append(app.Data.UsersGroups, acc)
		if _, err = db.Insert(acc, `UsersGroups`); err != nil {
			return
		}
	}

	// Дефолтовые шаблоны
	app.Data.UriDefault = make(map[string]*typDb.Uri)

	// Очистка
	ConfigControllers = nil
	ConfigUri = nil

	return
}

// LoadDataFromMemory() (error)
// Инициализация данных (без использования БД)
func loadDataFromMemory() (err error) {
	var controllersId = make(map[string]uint64)
	var uriId = make(map[string]uint64)

	// Сортируем uri & controllers и получаем postion
	var postionUri int32 = 1
	var postionControllers int32 = 1
	app.Data.MaxPostion = make(map[string]int32)

	// (Контроллеры)
	for path := range ConfigControllers {
		var ctrl = ConfigControllers[path]
		ctrl.Id = model.GenerateId(`Controllers`)
		ctrl.Path = path
		ctrl.Position = postionControllers
		app.Data.Controllers = append(app.Data.Controllers, &ctrl)
		controllersId[path] = ctrl.Id
		postionControllers++
	}
	app.Data.MaxPostion[`Controllers`] = postionControllers

	// (Роутинг Uri)
	for uri := range ConfigUri {
		var u = ConfigUri[uri]
		u.Id = model.GenerateId(`Uri`)
		u.Uri = uri
		if u.ContentType == `` {
			u.ContentType = `text/html`
		}
		if u.ContentEncode == `` {
			u.ContentEncode = `utf-8`
		}
		u.Position = postionUri
		app.Data.Uri = append(app.Data.Uri, &u)
		uriId[uri] = u.Id
		postionUri++
	}
	app.Data.MaxPostion[`Uri`] = postionUri

	// Связи между роутингом и контроллерами
	for uri := range ConfigUri {
		for _, ctrlPath := range ConfigUri[uri].Controllers {
			var acc = new(typDb.GroupsUri)
			//acc.Id = model.GenerateId(`GroupsUri`)
			acc.Uri_Id = uriId[uri]
			acc.Controllers_Id = controllersId[ctrlPath]
			app.Data.GroupsUri = append(app.Data.GroupsUri, acc)
		}
	}

	// Пользователи
	var u *typDb.Users
	u = new(typDb.Users)
	u.Id = model.GenerateId(`Users`)
	u.Login = `developer`
	u.Email = `developer@developer.developer`
	if u.Id != core.Config.Auth.DevUID {
		return logs.Base.Fatal(1310).Err
	}
	app.Data.Users = append(app.Data.Users, u)
	//
	u = new(typDb.Users)
	u.Id = model.GenerateId(`Users`)
	u.Login = `guest`
	u.Email = `guest@guest.guest`
	if u.Id != core.Config.Auth.GuestUID {
		return logs.Base.Fatal(1320).Err
	}
	app.Data.Users = append(app.Data.Users, u)
	core.Config.Auth.GuestUID = u.Id

	// Группы
	var g *typDb.Groups
	g = new(typDb.Groups)
	g.Id = model.GenerateId(`Groups`)
	g.Name = `developer`
	if g.Id != core.Config.Auth.DevUID {
		return logs.Base.Fatal(1330).Err
	}
	app.Data.Groups = append(app.Data.Groups, g)
	core.Config.Auth.DevUID = g.Id
	//
	g = new(typDb.Groups)
	g.Id = model.GenerateId(`Groups`)
	g.Name = `guest`
	if g.Id != core.Config.Auth.GuestUID {
		return logs.Base.Fatal(1340).Err
	}
	app.Data.Groups = append(app.Data.Groups, g)
	core.Config.Auth.GuestUID = g.Id

	// Связь пользователи и группы
	acc := new(typDb.UsersGroups)
	acc.Users_Id = core.Config.Auth.DevUID
	acc.Groups_Id = core.Config.Auth.DevUID
	app.Data.UsersGroups = append(app.Data.UsersGroups, acc)
	//
	acc = new(typDb.UsersGroups)
	acc.Users_Id = core.Config.Auth.GuestUID
	acc.Groups_Id = core.Config.Auth.GuestUID
	app.Data.UsersGroups = append(app.Data.UsersGroups, acc)

	// Дефолтовые шаблоны
	app.Data.UriDefault = make(map[string]*typDb.Uri)

	// Очистка
	ConfigControllers = nil
	ConfigUri = nil
	return
}
