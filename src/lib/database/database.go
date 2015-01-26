// TODO реализовать именованые конфиги
package database

import (
	"lib/database/face"
	"lib/database/mysql"
	"lib/logs"
)

type DbFace face.DbFace
type ArFace face.ArFace

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

func NewAr(useDb int64) ArFace {
	switch useDb {
	case UseMysql:
		return mysql.NewAr()
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
