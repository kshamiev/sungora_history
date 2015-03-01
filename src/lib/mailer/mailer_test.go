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
	var cfgMailer = new(mailer.CfgMailer)
	cfgMailer.Server = `mail.shamiev.ru`
	cfgMailer.Login = `konstantin@shamiev.ru`
	cfgMailer.Password = `LeRo_3riS`
	cfgMailer.FromAddress = `konstantin@shamiev.ru`
	cfgMailer.FromName = `Вася Пупкин`
	path, _ := os.Getwd()
	mailer.Init(cfgMailer)

	msg := mailer.NewMessageTpl(`Тема`, path+`/test`)
	msg.
		msg.To(`konstanta75@mail.ru`, `Шариков Полиграф Полиграфович`)
	msg.Variables[`content`] = `Тело шаблонного сообщения`
	if _, err := msg.Send(); err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("Успешно отправлено шаблонное сообщение на адрес [%s]", `konstanta75@mail.ru`)
	}

	msg = mailer.NewMessageBody(`Тема`, `Фирма веников не вяжет`)
	msg.To(`konstanta75@mail.ru`, `Шариков Полиграф Полиграфович`)
	if _, err := msg.Send(); err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("Успешно отправлено обычное сообщение на адрес [%s]", `konstanta75@mail.ru`)
	}
}
