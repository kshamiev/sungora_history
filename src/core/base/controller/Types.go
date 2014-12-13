package controller

import (
	"core/controller"
	"types"
)

type Types struct {
	controller.Controller
}

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
