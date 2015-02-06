// TODO реализовать именованые конфиги
package database

import (
	"lib/database/face"
	"lib/database/mysql"
	"lib/logs"
)

type DbFace face.DbFace
type QubFace face.QubFace

// Доступные БД для использования
const (
	UseMysql int64 = 1 + iota // БД Mysql
)

func NewDb(useDb int64, idConfig int8) (db DbFace, err error) {
	switch useDb {
	case UseMysql:
		return mysql.NewDb(idConfig)
	}
	return nil, logs.Base.Fatal(811, useDb).Err
}

func NewQub(useDb int64) QubFace {
	switch useDb {
	case UseMysql:
		return mysql.NewQub()
	}
	return nil
}

// CheckConnect Проверка настроек, конфигарций и соединений с БД
func CheckConnect() (err error) {
	if err = mysql.CheckConnect(); err != nil {
		return
	}
	return
}
