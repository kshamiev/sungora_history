package user

import (
	"encoding/json"
	"time"

	"github.com/volatiletech/null"

	"github.com/kshamiev/sungora/pb"
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/models"
	"github.com/kshamiev/sungora/pkg/typ"
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

func (us *User) ProtoSampleOut() *pb.TestReply {
	v, _ := json.Marshal(&us.SampleJS)
	return &pb.TestReply{
		Message:        us.Message.String,
		AdditionalTime: pb.TimeOut(us.CreatedAt),
		Data:           v,
	}
}

func (us *User) ProtoSampleIn(in *pb.TestReply) {
	us.CreatedAt = pb.TimeIn(in.AdditionalTime)
	us.Message = pb.NullStringIn(in.Message)
	_ = json.Unmarshal(in.Data, &us.SampleJS)
}

func (us *User) Dump() {
	app.Dumper(us.Message, us.CreatedAt.String(), us.SampleJS)
}
