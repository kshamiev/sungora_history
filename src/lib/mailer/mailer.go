// Бибилотека для раюоты с электронной почтой (RFC-822).
//
// Источник: http://github.com/gogits/gogs/modules/mailer
package mailer

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"

	"lib/logs"
	"lib/view"
)

// Конфигурация
var cfgMailer *CfgMailer

// Структура конфигурации
type CfgMailer struct {
	Mailer      string // Строка, которой представляется почтовый клиент. Например 'Mac OS-X Mail'
	FromAddress string // Почтовый адрес от которого приложение отправляет все письма
	FromName    string // Имя почтового адреса от которого приложение отправляет все письма
	Mode        string // Режим отправки писем, возможные режимы: sendmail, tcp
	Sendmail    string // Путь и параметры запуска программы sendmail для отправки писем.
	Server      string // Адрес почтового сервера. Можно указывать доменное имя либо ip адрес.
	Port        int64  // Порт почтового сервера для отправки писем. Возможные порты:
	SSL         bool   // Булевое значение указывающее использовать ли ключи для шифрованной связи с почтовым сервером по SSL протоколу
	CertPublic  string // Путь и имя файла с публичном ключом, для отправки почты по SSL протоколу (порт 465 или 587)
	CertPrivate string // Путь и имя файла с приватным ключом, для отправки почты по SSL протоколу (порт 465 или 587)
	Login       string // Логин подключения к почтовому серверу
	Password    string // Пароль подключения к почтовому серверу
}

// Инициализация библиотеки
//	+ cfg *CfgMailer конфигурация
func Init(cfg *CfgMailer) {
	if cfg.Port == 0 {
		cfg.Port = 25
	}
	if cfg.Mode == `` {
		cfg.Mode = `tcp`
	}
	cfgMailer = cfg
}

// Сообщение
type Message struct {
	to        []string
	fromMail  string
	fromName  string
	subject   string
	body      string
	typ       string
	tpl       string
	Massive   bool
	Variables map[string]interface{}
	Functions map[string]interface{}
}

// Create new mail message from html
//	+ subject string Тема сообщения
//	+ tpl string абсолютный путь до почтового шаблона
//	- *Message объект сообщения
func NewMessageTpl(subject, tpl string) *Message {
	var self = new(Message)
	self.fromMail = cfgMailer.FromAddress
	self.fromName = cfgMailer.FromName
	self.subject = subject
	self.tpl = tpl + `.tpl`
	self.typ = `html`
	self.Variables = make(map[string]interface{})
	self.Functions = make(map[string]interface{})
	return self
}

// Create new mail message from body
//	+ subject string Тема сообщения
//	+ body string Тело (шаблон) сообщения
//	- *Message объект сообщения
func NewMessageBody(subject, body string) *Message {
	var self = new(Message)
	self.fromMail = cfgMailer.FromAddress
	self.fromName = cfgMailer.FromName
	self.subject = subject
	self.body = body
	self.typ = `html`
	return self
}

// Отправитель сообщения
//	+ fromMail string Адрес отправителя
//	+ fromName string Имя отправителя
func (self *Message) From(fromMail, fromName string) {
	self.fromMail = fromMail
	self.fromName = fromName
}

// Получатель сообщения
//	+ toMail string Адрес получателя
//	+ toName string Имя получателя
func (self *Message) To(toMail, toName string) {
	self.to = append(self.to, toName+" <"+toMail+">")
}

// Отправка сообщения
//	- num int количество отправленых сообщений
//	- error ошибка операции
func (self *Message) Send() (num int, err error) {
	defer func() {
		if err != nil {
			logs.Base.Error(141, err)
			return
		}
	}()

	var server = cfgMailer.Server + `:` + strconv.Itoa(int(cfgMailer.Port))

	// поточвый шаблон
	var data bytes.Buffer
	var tpl = view.NewView()
	tpl.Functions = self.Functions
	tpl.Variables = self.Variables
	if self.tpl != `` && self.body == `` {
		if data, err = tpl.ExecuteFile(self.tpl); err != nil {
			return
		}
	} else {
		if data, err = tpl.ExecuteBytes([]byte(self.body)); err != nil {
			return
		}
	}
	self.body = data.String()

	// get message body
	content := self.content()

	if len(self.to) == 0 {
		return 0, fmt.Errorf("empty receive emails")
	} else if len(self.body) == 0 {
		return 0, fmt.Errorf("empty email body")
	}

	auth := smtp.PlainAuth("", cfgMailer.Login, cfgMailer.Password, cfgMailer.Server)

	if self.Massive {
		// send mail to multiple emails one by one
		num := 0
		for i := range self.to {
			body := []byte("To: " + self.to[i] + "\r\n" + content)
			err := smtp.SendMail(server, auth, self.fromMail, []string{self.to[i]}, body)
			if err != nil {
				return num, err
			}
			num++
		}
		return num, nil
	} else {
		body := []byte("To: " + strings.Join(self.to, ";") + "\r\n" + content)
		// send to multiple emails in one message
		err := smtp.SendMail(server, auth, self.fromMail, self.to, body)
		if err != nil {
			return 0, err
		} else {
			return len(self.to), nil
		}
	}
}

// Формирование заголовков и тела сообщения
//	- content string сформированное тело сообщения
func (self *Message) content() (content string) {
	contentType := "text/plain; charset=utf-8"
	if self.typ == "html" {
		contentType = "text/html; charset=utf-8"
	}
	content = "From: " + self.fromName + " <" + self.fromMail + ">\r\n"
	content += "Subject: " + self.subject + "\r\nContent-Type: " + contentType + "\r\n\r\n"
	content += self.body
	return
}

////
