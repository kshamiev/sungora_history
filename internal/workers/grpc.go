package workers

import (
	"context"
	"time"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/internal/model/user"
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/errs"
	"github.com/kshamiev/sungora/pkg/logger"
	"github.com/kshamiev/sungora/proto"
)

const GrpcSampleName = "GrpcSample"

// Пример работы с GRPC
type GrpcSample struct {
	*config.Component
}

// NewGrpcSample
func NewGrpcSample(c *config.Component) *GrpcSample { return &GrpcSample{Component: c} }

func (task *GrpcSample) Action(ctx context.Context) error {
	lg := logger.GetLogger(ctx)
	ctx = app.GRPCCtxOut(ctx, nil)

	res, err := task.SungoraClient.HelloWorld(ctx, &proto.TestRequest{Name: "запрос от клиента"})
	if err != nil {
		return errs.NewBadRequest(err, "ошибка для пользователя")
	}

	us := user.NewUser()
	us.ProtoSampleIn(res)
	us.Dump()

	lg.Info("grpc client ok")
	return nil
}

func (task *GrpcSample) WaitFor() time.Duration {
	return time.Second * 10
}

func (task *GrpcSample) Name() string {
	return GrpcSampleName
}
