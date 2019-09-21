package model

import (
	"context"
	"database/sql"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/models"
	"github.com/kshamiev/sungora/pkg/typ"
)

func TestUsers(t *testing.T) {
	var err error
	var db *sql.DB
	var cfg *config.Config

	// ConfigApp загрузка конфигурации
	if cfg, err = config.Get("../../config.yaml"); err != nil {
		t.Fatal(err)
	}

	// ConnectDB
	if db, err = app.NewConnectPostgres(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}
	boil.DebugMode = true

	var user = &models.User{
		Login: "qwerty",
	}
	js := typ.SampleJs{
		ID:   54687,
		Name: "Popcorn",
		Items: []typ.Item{
			{
				Price:    56.87,
				Quantity: 23,
			},
			{
				Price:    32.76,
				Quantity: 13,
			},
		},
	}
	user.SampleJS = js
	if err = user.Insert(context.Background(), db, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	user.Login = "test-test@test.ru"
	if _, err = user.Update(context.Background(), db, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if _, err = user.Delete(context.Background(), db); err != nil {
		t.Fatal(err)
	}
}
