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
	*models.User
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
		Message:        us.Message.String,
		AdditionalTime: proto.TimeOut(us.CreatedAt),
		Any:            proto.AnyOut(&us.SampleJS),
	}
}

func (us *User) ProtoSampleIn(in *proto.TestReply) {
	us.CreatedAt = proto.TimeIn(in.AdditionalTime)
	us.Message = proto.NullStringIn(in.Message)
	proto.AnyIn(in.Any, &us.SampleJS)
}

func (us *User) Dump() {
	app.Dumper(us.Message, us.CreatedAt.String(), us.SampleJS)
}
