package grpcserver

import (
	"context"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/internal/model"
	"github.com/kshamiev/sungora/pb"
	"github.com/kshamiev/sungora/pkg/app"
)

type Server struct {
	*config.Component
}

func New(c *config.Component) pb.SungoraServer {
	return &Server{Component: c}
}

func (serv *Server) HelloWorld(ctx context.Context, req *pb.TestRequest) (*pb.TestReply, error) {
	_, lg := app.GRPCCtxIn(ctx, serv.Lg)
	lg.Info("grpc server ok (" + req.Name + ")")

	us := model.NewUser(serv.Component)

	return us.ProtoSampleOut(us.GetUser()), nil
}
