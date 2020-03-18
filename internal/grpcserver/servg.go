package grpcserver

import (
	"context"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/internal/model/user"
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/proto"
)

type Server struct {
	*config.Component
}

func New(c *config.Component) proto.SungoraServer {
	return &Server{Component: c}
}

func (serv *Server) HelloWorld(ctx context.Context, req *proto.TestRequest) (*proto.TestReply, error) {
	_, lg := app.GRPCCtxIn(ctx, serv.Lg)
	lg.Info("grpc server ok (" + req.Name + ")")

	us := user.NewUserSet()
	return us.ProtoSampleOut(), nil
}
