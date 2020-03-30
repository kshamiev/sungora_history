package modelcore

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/null"

	"github.com/kshamiev/sungora/pb"
	"github.com/kshamiev/sungora/pb/typ"
)

func NewUserGRPC(user *pb.User) *User {

	return &User{
		ID:        typ.UUID{},
		Login:     user.Login,
		Email:     user.Email,
		IsOnline:  false,
		SampleJS:  typ.SampleJs{},
		Price:     decimal.Decimal{},
		Summa:     0,
		CNT:       0,
		Message:   null.String{},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: null.Time{},
		R:         nil,
		L:         userL{},
	}
}
