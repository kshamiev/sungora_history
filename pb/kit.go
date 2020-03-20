package pb

import (
	"encoding/json"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/volatiletech/null"
)

// NullStringIn перевод из примитива grpc
func NullStringIn(s string) null.String {
	if s == "" {
		return null.String{}
	}
	return null.StringFrom(s)
}

// TimeOut перевод в примитив grpc
func TimeOut(d time.Time) *timestamp.Timestamp {
	dp, _ := ptypes.TimestampProto(d)
	return dp
}

// TimeIn перевод из примитива grpc
func TimeIn(d *timestamp.Timestamp) time.Time {
	dp, _ := ptypes.Timestamp(d)
	return dp
}

// AnyOut перевод в примитив grpc
func AnyOut(d interface{}) *any.Any {
	v, _ := json.Marshal(d)
	return &any.Any{Value: v}
}

// AnyIn перевод из примитива grpc
func AnyIn(d *any.Any, obj interface{}) {
	_ = json.Unmarshal(d.Value, obj)
}
