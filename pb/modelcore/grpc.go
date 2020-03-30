package modelcore

import (
	"github.com/kshamiev/sungora/pb"
	"github.com/kshamiev/sungora/pb/typ"
)

func NewUserGRPC(proto *pb.User) *User {
	return &User{
		ID: typ.UUIDMustParse(proto.Id),
	}
}
