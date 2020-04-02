package grpcserver

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/metadata"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/internal/model"
	"github.com/kshamiev/sungora/pb"
	"github.com/kshamiev/sungora/pkg/app/response"
	"github.com/kshamiev/sungora/pkg/logger"
)

type Server struct {
	*config.Component
}

func New(c *config.Component) pb.SunServer {
	return &Server{Component: c}
}

func (ser *Server) Ping(context.Context, *empty.Empty) (*pb.TestSun, error) {
	return &pb.TestSun{
		Text:      "Pong",
		CreatedAt: ptypes.TimestampNow(),
	}, nil
}

func (ser *Server) GetUser(ctx context.Context, _ *empty.Empty) (*pb.User, error) {
	_, lg := ser.getContextData(ctx)
	lg.Info("grpc server ok")

	us := model.NewUser(ser.Component).GetUser()

	// sample OK
	return us.Proto(), nil

	// sample error
	// err := errs.GRPC(codes.InvalidArgument, errors.New("library error"), "ошибка для пользователя")
	// return us.Proto(), err
}

// getContextData получение данных (метаданных) и логера из контекста grpc
func (ser *Server) getContextData(ctx context.Context) (metadata.MD, logger.Logger) {
	lg := ser.Lg
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && md.Get(response.LogUUID) != nil && md.Get(response.LogAPI) != nil {
		lg = lg.WithField(response.LogUUID, md.Get(response.LogUUID)[0])
		lg = lg.WithField(response.LogAPI, md.Get(response.LogAPI)[0])
	}
	return md, lg
}
