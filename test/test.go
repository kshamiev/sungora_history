package test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/app"
)

type ENV struct {
	Cfg *config.Config
	DB  *sql.DB
}

func GetEnvironment(t *testing.T) *ENV {
	var (
		err error
		db  *sql.DB
		cfg *config.Config
	)
	// ConfigApp загрузка конфигурации
	if cfg, err = config.Get(os.Getenv("CONF")); err != nil {
		t.Fatal(err)
	}
	// ConnectDB SqlBoiler
	if db, err = app.NewPostgresBoiler(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}

	boil.DebugMode = true

	return &ENV{cfg, db}
}
