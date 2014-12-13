// Контроллер. Информационный модуль.
package controller

import (
	"bytes"

	"app"
	"core"
	"core/controller"
	"lib/view"
)

type Page struct {
	controller.Controller
}

// Block Информационный блок. Контент блок
func (self *Page) Block() (err error) {
	var index = make(map[string][]byte)
	for _, c := range app.Data.Content {
		if c.Lang == self.RW.Lang {
			if c.Uri_Id == self.Session.Uri.Id {
				//v.Variables[`Block_`+c.Block] = c.Content
				index[`Block_`+c.Block] = c.Content
			} else if c.Uri_Id == 0 {
				if _, ok := index[`Block_`+c.Block]; ok == false {
					index[`Block_`+c.Block] = c.Content
					//v.Variables[`Block_`+c.Block] = c.Content
				}
			}
		}
	}
	//
	var con bytes.Buffer
	var v = view.NewView()
	for i := range index {
		v.Variables[`Contnet`] = index[i]
		if con, err = v.ExecuteString(self.Controllers.Content); err != nil {
			self.RW.ResponseError(500)
			return
		}
		self.RW.Content.Variables[i] = con.String()
	}
	return
}

// Content Информационная страница. Контент страницы
func (self *Page) Content() (err error) {
	var con bytes.Buffer
	var v = view.NewView()
	if self.RW.Lang == core.Config.Main.Lang {
		// язык по умолчанию в uri
		v.Variables[`Title`] = self.Session.Uri.Title
		v.Variables[`Content`] = string(self.Session.Uri.Content)
		self.RW.Content.Variables[`Title`] = self.Session.Uri.Title
		self.RW.Content.Variables[`Keywords`] = self.Session.Uri.KeyWords
		self.RW.Content.Variables[`Description`] = self.Session.Uri.Description
	} else {
		// дополнительные языки в контенте
		for _, c := range app.Data.Content {
			if c.Lang == self.RW.Lang && c.Block == `content` {
				if c.Uri_Id == self.Session.Uri.Id { // приоритет за контентом раздела
					self.RW.Content.Variables[`Title`] = c.Title
					self.RW.Content.Variables[`Keywords`] = c.Keywords
					self.RW.Content.Variables[`Description`] = c.Description
					v.Variables[`Title`] = c.Title
					v.Variables[`Content`] = c.Content
					break
				} else if c.Uri_Id == 0 { // если нет то контент шаблона
					self.RW.Content.Variables[`Title`] = c.Title
					self.RW.Content.Variables[`Keywords`] = c.Keywords
					self.RW.Content.Variables[`Description`] = c.Description
					v.Variables[`Title`] = c.Title
					v.Variables[`Content`] = c.Content
				}
			}
		}
	}
	if con, err = v.ExecuteString(self.Controllers.Content); err != nil {
		self.RW.ResponseError(500)
		return
	}
	//self.RW.Content.Content = con.Bytes()
	self.RW.Content.Content = con.Bytes()
	return
}
