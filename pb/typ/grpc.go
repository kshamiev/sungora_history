package typ

import (
	"encoding/json"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/volatiletech/null"
)

// PbToNullString перевод из примитива grpc
func PbToNullString(s string) null.String {
	if s == "" {
		return null.String{}
	}
	return null.StringFrom(s)
}

// PbToNullTime перевод из примитива grpc
func PbToNullTime(d *timestamp.Timestamp) null.Time {
	dp, err := ptypes.Timestamp(d)
	if err != nil {
		return null.Time{}
	}
	return null.TimeFrom(dp)
}

// PbFromTime перевод в примитив grpc
func PbFromTime(d time.Time) *timestamp.Timestamp {
	dp, err := ptypes.TimestampProto(d)
	if err != nil {
		dp, _ = ptypes.TimestampProto(time.Time{})
	}
	return dp
}

// PbToTime перевод из примитива grpc
func PbToTime(d *timestamp.Timestamp) time.Time {
	dp, err := ptypes.Timestamp(d)
	if err != nil {
		dp = time.Time{}
	}
	return dp
}

// PbFromJSON перевод в примитив grpc
func PbFromJSON(d interface{}) []byte {
	v, _ := json.Marshal(d)
	return v



}

// PbToJSON перевод из примитива grpc
func PbToJSON(d *any.Any, obj interface{}) {
	_ = json.Unmarshal(d.Value, obj)
}


// PbFromAny перевод в примитив grpc
func PbFromAny(d interface{}) *any.Any {
	v, _ := json.Marshal(d)
	return &any.Any{Value: v}
}

// PbToAny перевод из примитива grpc
func PbToAny(d *any.Any, obj interface{}) {
	_ = json.Unmarshal(d.Value, obj)
}
