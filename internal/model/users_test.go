package model

import (
	"context"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/kshamiev/sungora/pkg/models"
	"github.com/kshamiev/sungora/pkg/typ"
	"github.com/kshamiev/sungora/test"
)

func TestUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	env := test.GetEnvironment(t)

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

	if err = user.Insert(ctx, env.DB, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	user.Login = "test-test@test.ru"

	if _, err = user.Update(ctx, env.DB, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if _, err = user.Delete(ctx, env.DB); err != nil {
		t.Fatal(err)
	}
}
