package workers

import (
	"context"
	"time"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/internal/model"
	"github.com/kshamiev/sungora/pb"
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/errs"
	"github.com/kshamiev/sungora/pkg/logger"
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

	res, err := task.SungoraClient.HelloWorld(ctx, &pb.TestRequest{Name: "запрос от клиента"})
	if err != nil {
		return errs.NewGRPC(err)
	}

	us := model.NewUser(task.Component)
	user, order := us.ProtoSampleIn(res)

	app.Dumper(user.Price.String(), user.CreatedAt.String(), user.Message, user.SampleJS, order.Status)

	lg.Info("grpc client ok")
	return nil
}

func (task *GrpcSample) WaitFor() time.Duration {
	return time.Second * 10
}

func (task *GrpcSample) Name() string {
	return GrpcSampleName
}
