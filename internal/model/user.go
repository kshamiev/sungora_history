package model

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/null"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pb"
	"github.com/kshamiev/sungora/pb/typ"
	"github.com/kshamiev/sungora/pkg/models"
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
func (ml *User) GetUser() *models.User {
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
	return &models.User{
		CreatedAt: time.Now(),
		Message:   null.StringFrom("важное сообщение от сервера"),
		SampleJS:  js,
	}
}

// ProtoSampleOut конвертируем для передачи по GRPC
func (ml *User) ProtoSampleOut(us *models.User) *pb.TestReply {
	v, _ := json.Marshal(&us.SampleJS)
	price := decimal.NewFromFloat(468.435).String()
	return &pb.TestReply{
		CreatedAt: typ.PbFromTime(us.CreatedAt),      // дата и время
		Message:   us.Message.String,                 // строка
		Price:     price,                             // дробные числа
		Data:      v,                                 // бинарные данные, JSON
		Status:    typ.StatusValue[typ.Status_CLOSE], // ENUM
	}
}

// ProtoSampleIn конвертируем обратно из GRPC
func (ml *User) ProtoSampleIn(in *pb.TestReply) (*models.User, *models.Order) {
	us := &models.User{
		CreatedAt: typ.PbToTime(in.CreatedAt),
		Message:   typ.PbToNullString(in.Message),
		Price:     decimal.RequireFromString(in.Price),
	}
	_ = json.Unmarshal(in.Data, &us.SampleJS)

	or := &models.Order{
		Status: typ.StatusName[in.Status],
	}
	return us, or
}
