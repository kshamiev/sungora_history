package controller

import (
	"core"
	"types"
	typDb "types/db"
)

type Types struct {
	Session     *core.Session      // Сессия
	RW          *core.RW           // Управление вводом и выводом
	Controllers *typDb.Controllers // Соответсвующий контроллер из области данных по строковому ид.
}

// Создание контроллера
func NewTypes(rw *core.RW, s *core.Session, c *typDb.Controllers) interface{} {
	var self = new(Types)
	self.RW = rw
	self.Session = s
	self.Controllers = c
	return self
}

////

// ApiScenario Получение сценариев для всех типов
func (self *Types) ApiScenario() (err error) {
	var ok bool
	var typ, scenario string

	if typ, ok = self.RW.GetSegmentUriString(`typ`); ok == false {
		self.RW.ResponseJson(nil, 409, 120)
		return
	}
	//
	if scenario, ok = self.RW.GetSegmentUriString(`scenario`); ok == true {
		var options *types.Scenario
		if options, err = types.GetScenario(typ, scenario); err != nil {
			return self.RW.ResponseJson(nil, 404, 570, scenario, typ)
		}
		return self.RW.ResponseJson(options, 200, 0)
	}
	if scenarioList := types.GetScenarioList(typ); len(scenarioList) == 0 {
		self.RW.ResponseJson(nil, 404, 122, typ)
	} else {
		self.RW.ResponseJson(scenarioList, 200, 0)
	}
	return
}
