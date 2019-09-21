package app

import (
	"database/sql"
	"fmt"

	// драйвер работы с postgres
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql/driver"
)

// NewConnectPostgres создание соединения с postgres
func NewConnectPostgres(cfg *ConfigPostgres) (db *sql.DB, err error) {
	strCon := fmt.Sprintf("dbname=%s host=%s port=%d user=%s password=%s sslmode=%s",
		cfg.DbName,
		cfg.Host,
		cfg.Port,
		cfg.Login,
		cfg.Password,
		cfg.Ssl,
	)
	if db, err = sql.Open("postgres", strCon); err != nil {
		return
	}
	err = db.Ping()
	return
}

// NewConnectMysql создание соединения с mysql
func NewConnectMysql(cfg *ConfigMysql) (db *sql.DB, err error) {
	strCon := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Login,
		cfg.Password,
		cfg.DbName,
	)
	if db, err = sql.Open("mysql", strCon); err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	return
}
