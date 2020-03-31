package modelsun

import (
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null"

	"github.com/kshamiev/sungora/pb"
	"github.com/kshamiev/sungora/pb/typ"
)

func (o *User) Proto() *pb.User {
	return &pb.User{
		Id: o.ID.String(),
		Login: o.Login,
		Email: o.Email,
		Price: o.Price.String(),
		Message: o.Message.String,
		Metrika: o.Metrika.JSON,
%!(EXTRA string=Metrika)		CreatedAt: typ.PbFromTime(o.CreatedAt),
		UpdatedAt: typ.PbFromTime(o.UpdatedAt),
		DeletedAt: typ.PbFromTime(o.DeletedAt.Time),
	}
}

func NewUserProto(proto *pb.User) *User {
	return &User{
		ID: typ.UUIDMustParse(proto.Id),
		Login: proto.Login,
		Email: proto.Email,
		Price: decimal.RequireFromString(proto.Price),
		Message: typ.PbToNullString(proto.Message),
		Metrika: null.JSONFrom(proto.Metrika),
%!(EXTRA string=Metrika)		CreatedAt: typ.PbToTime(proto.CreatedAt),
		UpdatedAt: typ.PbToTime(proto.UpdatedAt),
		DeletedAt: typ.PbToNullTime(proto.DeletedAt),
	}
}

func (o *Order) Proto() *pb.Order {
	return &pb.Order{
		Id: o.ID.String(),
		UserId: o.UserID.String(),
		Status: typ.StatusValue[o.Status],
		CreatedAt: typ.PbFromTime(o.CreatedAt),
		UpdatedAt: typ.PbFromTime(o.UpdatedAt),
		DeletedAt: typ.PbFromTime(o.DeletedAt.Time),
	}
}

func NewOrderProto(proto *pb.Order) *Order {
	return &Order{
		ID: typ.UUIDMustParse(proto.Id),
		UserID: typ.UUIDMustParse(proto.UserId),
		Status: typ.StatusName[proto.Status],
		CreatedAt: typ.PbToTime(proto.CreatedAt),
		UpdatedAt: typ.PbToTime(proto.UpdatedAt),
		DeletedAt: typ.PbToNullTime(proto.DeletedAt),
	}
}

func (o *Role) Proto() *pb.Role {
	return &pb.Role{
		Id: o.ID.String(),
		Code: o.Code,
		Description: o.Description,
	}
}

func NewRoleProto(proto *pb.Role) *Role {
	return &Role{
		ID: typ.UUIDMustParse(proto.Id),
		Code: proto.Code,
		Description: proto.Description,
	}
}
