package app

import (
	"database/sql"
	"fmt"

	// драйвер работы с postgres
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql/driver"
)

// NewConnectPostgres создание соединения с postgres
func NewConnectPostgres(cfg *ConfigPostgres) (*sql.DB, error) {
	strCon := fmt.Sprintf("dbname=%s host=%s port=%d user=%s password=%s sslmode=%s",
		cfg.DbName,
		cfg.Host,
		cfg.Port,
		cfg.Login,
		cfg.Password,
		cfg.Ssl,
	)
	db, err := sql.Open("postgres", strCon)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	return db, err
}
