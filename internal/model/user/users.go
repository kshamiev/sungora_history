package user

import (
	"time"

	"github.com/volatiletech/null"

	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/models"
	"github.com/kshamiev/sungora/pkg/typ"
	"github.com/kshamiev/sungora/proto"
)

type User struct {
	User *models.User
}

// NewUser пример создания безнес модели
func NewUser() *User {
	return &User{
		User: &models.User{},
	}
}

// NewUserSet пример создания безнес модели
func NewUserSet() *User {
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
	return &User{
		User: &models.User{
			CreatedAt: time.Now(),
			Message:   null.StringFrom("важное сообщение от сервера"),
			SampleJS:  js,
		},
	}
}

func (us *User) ProtoSampleOut() *proto.TestReply {
	return &proto.TestReply{
		Message:        us.User.Message.String,
		AdditionalTime: proto.TimeOut(us.User.CreatedAt),
		Any:            proto.AnyOut(&us.User.SampleJS),
	}
}

func (us *User) ProtoSampleIn(in *proto.TestReply) {
	us.User.CreatedAt = proto.TimeIn(in.AdditionalTime)
	us.User.Message = proto.NullStringIn(in.Message)
	proto.AnyIn(in.Any, &us.User.SampleJS)
}

func (us *User) Dump() {
	app.Dumper(us.User.Message, us.User.CreatedAt.String(), us.User.SampleJS)
}
