package servg

import (
	"context"
	"fmt"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/app/response"
	"github.com/kshamiev/sungora/pkg/logger"
	"github.com/kshamiev/sungora/protob"
)

type Server struct {
	*config.Component
}

func (s *Server) SayHello(ctx context.Context, req *protob.HelloRequest) (*protob.HelloReply, error) {
	lg := logger.GetLogger(ctx)

	fmt.Println(ctx.Value(response.CtxUUID))

	lg.Info("grpc server ok (" + req.Name + ")")

	return &protob.HelloReply{
		Message: "важное сообщение от сервера",
	}, nil
}

func New(c *config.Component) protob.SungoraServer {
	return &Server{Component: c}
}
