package model

import (
	"app"
	"core"
	"core/base/config"
	"lib/database"
	"lib/logs"
	typDb "types/db"
	typReg "types/request"
)

func debug(obj ...interface{}) {
	logs.Dumper(obj...)
}

func SortingControllers(req *typReg.Position) (err error) {
	// проверки
	if req.Id == req.TargetId {
		return
	}
	var db database.DbFace
	if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
		return
	}
	defer db.Free()
	// инициализация
	var ctr = NewControllers(req.Id)
	if err = ctr.Model.Load(`Id`); err != nil {
		return
	}
	var ctrTarget = NewControllers(req.TargetId)
	if req.TargetId > 0 { // если не в начало
		if err = ctrTarget.Model.Load(`Id`); err != nil {
			return
		}
	}
	var step, stepAfter int32
	if ctr.Type.Position > ctrTarget.Type.Position {
		step = 1 // смещение сортируемого если сортируем в начало
	} else if ctr.Type.Position < ctrTarget.Type.Position {
		stepAfter = -1 // смещение сдвига если сортируем в конец
	} else {
		return
	}
	// сортировка в БД
	//debug(ctr, ctrTarget, step)
	if err = db.Query(`base/controllers/0`, ctr.Type.Position); err != nil {
		return
	}
	if err = db.Query(`base/controllers/1`, ctrTarget.Type.Position+stepAfter); err != nil {
		return
	}
	if err = db.Query(`base/controllers/2`, ctrTarget.Type.Position+step, req.Id); err != nil {
		return
	}

	//var tt = map[uint64]int32{}
	//for _, elm := range app.Data.Controllers {
	//	tt[elm.Id] = elm.Position
	//}
	//debug(tt)

	// сортировка в памяти
	for i := range app.Data.Controllers {
		//step = 0
		if app.Data.Controllers[i].Position > ctr.Type.Position {
			app.Data.Controllers[i].Position--

		}
		if app.Data.Controllers[i].Position > ctrTarget.Type.Position+stepAfter {
			app.Data.Controllers[i].Position++
		}
		if app.Data.Controllers[i].Id == ctr.Type.Id {
			app.Data.Controllers[i].Position = ctrTarget.Type.Position + step
		}
	}
	//tt = map[uint64]int32{}
	//for _, elm := range app.Data.Controllers {
	//	tt[elm.Id] = elm.Position
	//}
	//debug(tt)

	return
}

func SortingUri(req *typReg.Position) (err error) {
	// проверки
	if req.Id == req.TargetId {
		return
	}
	var db database.DbFace
	if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
		return
	}
	defer db.Free()
	// инициализация
	var uri = NewUri(req.Id)
	if err = uri.Model.Load(`Id`); err != nil {
		return
	}
	var uriTarget = NewUri(req.TargetId)
	if req.TargetId > 0 { // если не в начало
		if err = uriTarget.Model.Load(`Id`); err != nil {
			return
		}
	}
	var step, stepAfter int32
	if uri.Type.Position > uriTarget.Type.Position {
		step = 1 // смещение сортируемого если сортируем в начало
	} else if uri.Type.Position < uriTarget.Type.Position {
		stepAfter = -1 // смещение сдвига если сортируем в конец
	} else {
		return
	}
	// сортировка в БД
	//debug(uri, uriTarget, step)
	if err = db.Query(`base/uri/0`, uri.Type.Position); err != nil {
		return
	}
	if err = db.Query(`base/uri/1`, uriTarget.Type.Position+stepAfter); err != nil {
		return
	}
	if err = db.Query(`base/uri/2`, uriTarget.Type.Position+step, req.Id); err != nil {
		return
	}
	//
	for i := range app.Data.Uri {
		if app.Data.Uri[i].Position > uri.Type.Position {
			app.Data.Uri[i].Position--
		}
		if app.Data.Uri[i].Position > uriTarget.Type.Position+stepAfter {
			app.Data.Uri[i].Position++
		}
		if app.Data.Uri[i].Id == uri.Type.Id {
			app.Data.Uri[i].Position = uriTarget.Type.Position + step
		}
	}
	return
}

////

func GetUriGrid(page int) (data []*typDb.Uri, err error) {
	page = (page - 1) * config.PAGE_ITEM
	var ar = database.NewAr(core.Config.Main.UseDb).SelectScenario(`Uri`, `GridAdmin`)
	var query = ar.From(`Uri`).Order("`Name` ASC").Limit(page, config.PAGE_ITEM).Get()
	var db database.DbFace
	if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
		return
	}
	defer db.Free()
	//data = make([]*typDb.Uri, 0)
	err = db.SelectSlice(&data, query)
	return
}

func GetControllersGrid(page int) (data []*typDb.Controllers, err error) {
	page = (page - 1) * config.PAGE_ITEM
	var ar = database.NewAr(core.Config.Main.UseDb).SelectScenario(`Controllers`, `GridAdmin`)
	var query = ar.From(`Controllers`).Order("`Name` ASC").Limit(page, config.PAGE_ITEM).Get()
	var db database.DbFace
	if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
		return
	}
	defer db.Free()
	//data = make([]*typDb.Controllers, 0)
	err = db.SelectSlice(&data, query)
	return
}

func GetUsersGrid(page int) (data []*typDb.Users, err error) {
	page = (page - 1) * config.PAGE_ITEM
	var ar = database.NewAr(core.Config.Main.UseDb).SelectScenario(`Users`, `GridAdmin`)
	var query = ar.From(`Users`).Order("`Name` ASC").Limit(page, config.PAGE_ITEM).Get()
	var db database.DbFace
	if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
		return
	}
	defer db.Free()
	//data = make([]*typDb.Users, 0)
	err = db.SelectSlice(&data, query)
	return
}

func GetGroupsGrid(page int) (data []*typDb.Groups, err error) {
	page = (page - 1) * config.PAGE_ITEM
	var ar = database.NewAr(core.Config.Main.UseDb).SelectScenario(`Groups`, `GridAdmin`)
	var query = ar.From(`Groups`).Order("`Name` ASC").Limit(page, config.PAGE_ITEM).Get()
	debug(query)
	var db database.DbFace
	if db, err = database.NewDb(core.Config.Main.UseDb, 0); err != nil {
		return
	}
	defer db.Free()
	//data = make([]*typDb.Groups, 0)
	err = db.SelectSlice(&data, query)
	return
}

////

func GetGroupsMap(id uint64) (mapId map[uint64]uint64) {
	mapId = make(map[uint64]uint64)
	for i := range app.Data.UsersGroups {
		if app.Data.UsersGroups[i].Users_Id == id {
			mapId[app.Data.UsersGroups[i].Groups_Id] = app.Data.UsersGroups[i].Groups_Id
		}
	}
	return
}

// SearchUsersToken
func SearchUsersToken(token string) *typDb.Users {
	for i := range app.Data.Users {
		if app.Data.Users[i].Token == token {
			return app.Data.Users[i]
		}
	}
	return nil
}

// SearchUsersLogin
func SearchUsersLogin(login string) *typDb.Users {
	for i := range app.Data.Users {
		if app.Data.Users[i].Login == login {
			return app.Data.Users[i]
		}
	}
	return nil
}

// SearchUsersEmail
func SearchUsersEmail(email string) *typDb.Users {
	for i := range app.Data.Users {
		if app.Data.Users[i].Email == email {
			return app.Data.Users[i]
		}
	}
	return nil
}
