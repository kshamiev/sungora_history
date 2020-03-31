package model

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/null"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pb/modelsun"
	"github.com/kshamiev/sungora/pb/typ"
)

type User struct {
	*config.Component
}

// NewUser создания безнес модели
func NewUser(comp *config.Component) *User {
	return &User{
		comp,
	}
}

// GetUser получение определенного пользователя
func (ml *User) GetUser() *modelsun.User {
	js := SampleJs{
		ID:   54687,
		Name: "Popcorn",
		Items: []Item{
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
	b, _ := json.Marshal(&js)
	return &modelsun.User{
		ID:        typ.UUIDNew(),
		CreatedAt: time.Now(),
		Login:     "pupkin",
		Price:     decimal.NewFromFloat(748.567),
		Metrika:   null.JSONFrom(b),
	}
}
