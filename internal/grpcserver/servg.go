package grpcserver

import (
	"context"
	"time"

	"github.com/volatiletech/null"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/models"
	"github.com/kshamiev/sungora/pkg/typ"
	"github.com/kshamiev/sungora/proto"
)

type Server struct {
	*config.Component
}

func New(c *config.Component) proto.SungoraServer {
	return &Server{Component: c}
}

func (serv *Server) HelloWorld(ctx context.Context, req *proto.TestRequest) (*proto.TestReply, error) {
	_, lg := serv.GRPCKit.CtxIn(ctx)
	lg.Info("grpc server ok (" + req.Name + ")")

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

	or := models.Order{
		Status:    "CANCEL",
		CreatedAt: time.Now(),
		Message:   null.StringFrom("важное сообщение от сервера"),
		SampleJS:  js,
	}

	return &proto.TestReply{
		Message:        or.Message.String,
		Status:         proto.TestReply_StatusOrders(proto.TestReply_StatusOrders_value[or.Status]),
		AdditionalTime: serv.GRPCKit.TimeTo(or.CreatedAt),
		Any:            serv.GRPCKit.AnyTo(&or.SampleJS),
	}, nil
}
