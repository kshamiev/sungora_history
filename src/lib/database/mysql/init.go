package mysql

import (
	"errors"
	"fmt"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
)

// Стек конфигураций БД
var cfgMysql map[int8]*CfgMysql

// Настройки mysql
type CfgMysql struct {
	Driver   string // Драйвер
	Host     string // Хост базы данных (localhost - по умолчанию)
	Port     int64  // Порт подключения по протоколу tcp/ip (3306 по умолчанию)
	Type     string // Тип подключения (socket | tcp (tcp по умолчанию)
	Socket   string // Путь к socket файлу
	Name     string // Имя базы данных
	Login    string // Логин к базе данных
	Password string // Пароль к базе данных
	Charset  string // Кодировка данных (utf-8 - по умолчанию)
	TimeOut  int64  // Таймаут использования соединения (в секундах) (5 - по умолчанию)
	Updates  string // Путь где лежат обновления БД
	CntConn  int64  // Максимальное допустимое количество коннектов (50 - по умолчанию)
	//Logs     bool   // Логирование выполняемых запросов
}

func InitMysql(cfg map[int8]*CfgMysql) {
	cfgMysql = cfg
	for i := range cfgMysql {
		if cfgMysql[i].Charset == `` {
			cfgMysql[i].Charset = `UTF-8`
		}
		if cfgMysql[i].Type == `` {
			cfgMysql[i].Type = `tcp`
		}
		if cfgMysql[i].Port == 0 {
			cfgMysql[i].Port = 3306
		}
		if cfgMysql[i].Host == `` {
			cfgMysql[i].Host = `localhost`
		}
		if cfgMysql[i].TimeOut == 0 {
			cfgMysql[i].TimeOut = 5
		}
		if cfgMysql[i].CntConn == 0 {
			cfgMysql[i].CntConn = 50
		}
	}
	goMysql()
}

// CheckConnect Проверка настроек, конфигарций и соединений с БД
func CheckConnect() (err error) {
	var db mysql.Conn
	for k, d := range cfgMysql {
		if d.Type == `tcp` {
			if "" == d.Host || 0 == d.Port || "" == d.Login || "" == d.Name {
				delete(cfgMysql, k)
				return errors.New("Настройки Mysql указаны не полностью")
			}
			db = mysql.New("tcp", "", fmt.Sprintf("%s:%d", d.Host, d.Port), d.Login, d.Password)
		} else {
			if "" == d.Socket || "" == d.Host || 0 == d.Port || "" == d.Login || "" == d.Name {
				delete(cfgMysql, k)
				return errors.New("Настройки Mysql указаны не полностью")
			}
			db = mysql.New("unix", "", d.Socket, d.Login, d.Password)
		}
		err = db.Connect()
		if err != nil {
			delete(cfgMysql, k)
			return
		}
		err = db.Use(d.Name)
		if err != nil {
			return
		}
		db.Close()
	}
	return
}
