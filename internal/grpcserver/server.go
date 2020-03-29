package grpcserver

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/internal/model"
	"github.com/kshamiev/sungora/pb"
	"github.com/kshamiev/sungora/pkg/app/response"
	"github.com/kshamiev/sungora/pkg/errs"
	"github.com/kshamiev/sungora/pkg/logger"
)

type Server struct {
	*config.Component
}

func New(c *config.Component) pb.SungoraServer {
	return &Server{Component: c}
}

func (ser *Server) HelloWorld(ctx context.Context, req *pb.TestRequest) (*pb.TestReply, error) {
	_, lg := ser.getContextData(ctx)
	lg.Info("grpc server ok (" + req.Name + ")")

	us := model.NewUser(ser.Component)

	// sample OK
	// return us.ProtoSampleOut(us.GetUser()), nil

	// sample error
	err := errs.GRPC(codes.InvalidArgument, errors.New("library error"), "ошибка для пользователя")
	return us.ProtoSampleOut(us.GetUser()), err
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
