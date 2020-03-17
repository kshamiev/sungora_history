package workers

import (
	"context"
	"time"

	"github.com/volatiletech/null"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/errs"
	"github.com/kshamiev/sungora/pkg/logger"
	"github.com/kshamiev/sungora/pkg/models"
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
	ctx = task.GRPCKit.CtxOut(ctx, nil)

	res, err := task.SungoraClient.HelloWorld(ctx, &proto.TestRequest{Name: "запрос от клиента"})
	if err != nil {
		return errs.NewBadRequest(err, "ошибка для пользователя")
	}

	or := models.Order{
		Status:    proto.TestReply_StatusOrders_name[int32(res.Status)],
		CreatedAt: task.GRPCKit.TimeFrom(res.AdditionalTime),
		Message:   null.StringFrom(res.Message),
	}
	task.GRPCKit.AnyFrom(res.Any, &or.SampleJS)

	app.Dumper(or.Status, or.CreatedAt.String(), or.Message, or.SampleJS)

	lg.Info("grpc client ok")
	return nil
}

func (task *GrpcSample) WaitFor() time.Duration {
	return time.Second * 10
}

func (task *GrpcSample) Name() string {
	return GrpcSampleName
}
