package grpcserver

import (
	"context"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/proto"
)

type Server struct {
	*config.Component
}

func New(c *config.Component) proto.SungoraServer {
	return &Server{Component: c}
}

func (s *Server) HelloWorld(ctx context.Context, req *proto.TestRequest) (*proto.TestReply, error) {
	md := s.GRPCKit.CtxIn(ctx)
	lg := s.GRPCKit.GetLog(md)

	lg.Info("grpc server ok (" + req.Name + ")")

	return &proto.TestReply{
		Message: "важное сообщение от сервера",
	}, nil
}
