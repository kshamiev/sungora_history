package modelsun

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null"

	"github.com/kshamiev/sungora/pb"
	"github.com/kshamiev/sungora/pb/typ"
)

func init() {
	_ = decimal.Decimal{}
	_ = typ.UUID{}
}

// pbFromNullString перевод из примитива grpc
func pbFromNullString(s string) null.String {
	if s == "" {
		return null.String{}
	}
	return null.StringFrom(s)
}

// pbFromNullTime перевод из примитива grpc
func pbFromNullTime(d *timestamp.Timestamp) null.Time {
	dp, err := ptypes.Timestamp(d)
	if err != nil {
		return null.Time{}
	}
	return null.TimeFrom(dp)
}

// pbToTime перевод в примитив grpc
func pbToTime(d time.Time) *timestamp.Timestamp {
	dp, err := ptypes.TimestampProto(d)
	if err != nil {
		dp, _ = ptypes.TimestampProto(time.Time{})
	}
	return dp
}

// pbFromTime перевод из примитива grpc
func pbFromTime(d *timestamp.Timestamp) time.Time {
	dp, err := ptypes.Timestamp(d)
	if err != nil {
		dp = time.Time{}
	}
	return dp
}

// Code generated. DO NOT EDIT

func (o *User) Proto() *pb.User {
	return &pb.User{
		Id: o.ID.String(),
		Login: o.Login,
		Email: o.Email,
		Price: o.Price.String(),
		SummaOne: o.SummaOne,
		SummaTwo: o.SummaTwo,
		Cnt2: int32(o.CNT2),
		Cnt4: int64(o.CNT4),
		Cnt8: o.CNT8,
		IsOnline: o.IsOnline,
		Alias: o.Alias,
		DataByte: o.DataByte.Bytes,
		Metrika: o.Metrika.JSON,
		CreatedAt: pbToTime(o.CreatedAt),
		UpdatedAt: pbToTime(o.UpdatedAt),
		DeletedAt: pbToTime(o.DeletedAt.Time),
	}
}

func (o UserSlice) ProtoS() []*pb.User {
	res := make([]*pb.User, len(o))
	for i := range o {
		res[i] = o[i].Proto()
	}
	return res
}

func NewUserProto(proto *pb.User) *User {
	return &User{
		ID: typ.UUIDMustParse(proto.Id),
		Login: proto.Login,
		Email: proto.Email,
		Price: decimal.RequireFromString(proto.Price),
		SummaOne: proto.SummaOne,
		SummaTwo: proto.SummaTwo,
		CNT2: int16(proto.Cnt2),
		CNT4: int(proto.Cnt4),
		CNT8: proto.Cnt8,
		IsOnline: proto.IsOnline,
		Alias: proto.Alias,
		DataByte: null.BytesFrom(proto.DataByte),
		Metrika: null.JSONFrom(proto.Metrika),
		CreatedAt: pbFromTime(proto.CreatedAt),
		UpdatedAt: pbFromTime(proto.UpdatedAt),
		DeletedAt: pbFromNullTime(proto.DeletedAt),
	}
}

func NewUserProtoS(protos []*pb.User) []*User {
	res := make([]*User, len(protos))
	for i := range protos {
		res[i] = NewUserProto(protos[i])
	}
	return res
}

func (o *Order) Proto() *pb.Order {
	return &pb.Order{
		Id: o.ID.String(),
		UserId: o.UserID.String(),
		Number: int64(o.Number),
		Status: typ.StatusValue[o.Status],
		CreatedAt: pbToTime(o.CreatedAt),
		UpdatedAt: pbToTime(o.UpdatedAt),
		DeletedAt: pbToTime(o.DeletedAt.Time),
	}
}

func (o OrderSlice) ProtoS() []*pb.Order {
	res := make([]*pb.Order, len(o))
	for i := range o {
		res[i] = o[i].Proto()
	}
	return res
}

func NewOrderProto(proto *pb.Order) *Order {
	return &Order{
		ID: typ.UUIDMustParse(proto.Id),
		UserID: typ.UUIDMustParse(proto.UserId),
		Number: int(proto.Number),
		Status: typ.StatusName[proto.Status],
		CreatedAt: pbFromTime(proto.CreatedAt),
		UpdatedAt: pbFromTime(proto.UpdatedAt),
		DeletedAt: pbFromNullTime(proto.DeletedAt),
	}
}

func NewOrderProtoS(protos []*pb.Order) []*Order {
	res := make([]*Order, len(protos))
	for i := range protos {
		res[i] = NewOrderProto(protos[i])
	}
	return res
}

func (o *Role) Proto() *pb.Role {
	return &pb.Role{
		Id: o.ID.String(),
		Code: o.Code,
		Description: o.Description,
	}
}

func (o RoleSlice) ProtoS() []*pb.Role {
	res := make([]*pb.Role, len(o))
	for i := range o {
		res[i] = o[i].Proto()
	}
	return res
}

func NewRoleProto(proto *pb.Role) *Role {
	return &Role{
		ID: typ.UUIDMustParse(proto.Id),
		Code: proto.Code,
		Description: proto.Description,
	}
}

func NewRoleProtoS(protos []*pb.Role) []*Role {
	res := make([]*Role, len(protos))
	for i := range protos {
		res[i] = NewRoleProto(protos[i])
	}
	return res
}
