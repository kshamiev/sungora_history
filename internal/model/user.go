package model

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/null"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pb/typ"
	"github.com/kshamiev/sungora/pb/typsun"
	"github.com/kshamiev/sungora/pkg/models"
)

// бизнес модель
type User struct {
	cm    *config.Component // служебный инструментарий
	Type  *typsun.User      // тип модели для приемки, отправки на фронт по grpc, валидации
	Model *models.User      // модель для работы с БД
	Order *Order            // зависимая бизнес модель
}

// NewUser создания безнес модели
func NewUser(cm *config.Component) *User { return &User{cm: cm} }

// Load
func (ml *User) Load(id typ.UUID) {
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
	ml.Type = &typsun.User{
		ID:        id,
		Login:     "pupkin",
		Price:     decimal.NewFromFloat(748.567),
		Alias:     []string{"one", "two", "tree"},
		Metrika:   null.JSONFrom(b),
		CreatedAt: time.Now(),
	}
}
