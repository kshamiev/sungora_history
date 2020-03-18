package user

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/kshamiev/sungora/pkg/models"
	"github.com/kshamiev/sungora/pkg/typ"
	"github.com/kshamiev/sungora/test"
)

func TestUser(t *testing.T) {
	var err error

	ctx := context.Background()
	env := test.GetEnvironment(t)

	var us = &models.User{
		Login: "qwerty",
		Email: "test-test@test.ru",
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
	us.SampleJS = js

	if err = us.Insert(ctx, env.DB, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	us.Price = decimal.NewFromFloat(345.876)

	if _, err = us.Update(ctx, env.DB, boil.Whitelist(models.UserColumns.Price)); err != nil {
		t.Fatal(err)
	}

	if _, err = us.Delete(ctx, env.DB); err != nil {
		t.Fatal(err)
	}
}
