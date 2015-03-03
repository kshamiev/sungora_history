// запуск теста
// SET GOPATH=C:\Work\projectName
// go test -v lib/mailer
// go test -v -bench . lib/mailer
package mailer_test

import (
	"os"
	"testing"

	"lib/mailer"
)

func TestMailer(t *testing.T) {

	var to = `test@sungora.ru`
	var toName = `Шариков Полиграф Полиграфович`

	var cfgMailer = new(mailer.CfgMailer)
	cfgMailer.Server = `mail.sungora.ru`
	cfgMailer.Login = `test@sungora.ru`
	cfgMailer.Password = `test`
	cfgMailer.FromAddress = `test@sungora.ru`
	cfgMailer.FromName = `Вася Пупкин`
	path, _ := os.Getwd()
	mailer.Init(cfgMailer)

	msg := mailer.NewMessageTpl(`Тема`, path+`/test`)
	msg.To(to, toName)
	msg.Variables[`content`] = `Вставленное в шаблон тело сообщения`
	if cnt, err := msg.Send(); err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("Успешно отправлено: [%d] сообщение на адрес %s", cnt, to)
	}

	msg = mailer.NewMessageBody(`Тема`, `Фирма веников не вяжет`)
	msg.To(to, `Шариков Полиграф Полиграфович`)
	if cnt, err := msg.Send(); err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("Успешно отправлено: [%d] сообщение на адрес %s", cnt, to)
	}
}
