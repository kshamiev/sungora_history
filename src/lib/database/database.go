// Интерфейс по работе с БД.
//
// Интерфейс для работы с абстрактной (любого типа) БД
// Интерфейс для формирования запросов к абстрактного БД
package database

import (
	"lib/database/face"
	"lib/database/mysql"
	"lib/logs"
)

// Доступные БД для использования
const (
	UseMysql int64 = 1 + iota // БД Mysql
)

// Интерфейс для работы с абстрактной (любого типа) БД
type DbFace face.DbFace

// Создание соединения и объекта по работе с абстрактной БД
//    + useDb int64 идентификатор используемой БД
//    + idConfig int8 Номер конфигурации
//    - DbFace объект по работе с конкретной БД
//    - error Ошибка создания
func NewDb(useDb int64, idConfig int8) (db DbFace, err error) {
	switch useDb {
	case UseMysql:
		return mysql.NewDb(idConfig)
	}
	return nil, logs.Base.Fatal(811, useDb).Err
}

// Интерфейс для абстрактного формирования запросов к БД
type QubFace face.QubFace

// Создание конструктора запросов
//    + useDb int64 идентификатор используемой БД
//    - QubFace конструктор запросов конкретной БД
func NewQub(useDb int64) QubFace {
	switch useDb {
	case UseMysql:
		return mysql.NewQub()
	}
	return nil
}

// Проверка настроек, конфигарций и соединений с БД
//    - error Ошибка проверки
func CheckConnect() (err error) {
	if err = mysql.CheckConnect(); err != nil {
		return
	}
	return
}
