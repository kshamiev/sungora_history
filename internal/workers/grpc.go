package workers

import (
	"context"
	"time"

	"github.com/kshamiev/sungora/internal/config"
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

func (task *GrpcSample) Action(ctx context.Context) {
	lg := logger.GetLogger(ctx)

	ctx = task.GRPCKit.CtxOut(ctx, nil)

	res, err := task.SungoraClient.HelloWorld(ctx, &proto.TestRequest{Name: "запрос от клиента"})
	if err != nil {
		lg.WithError(err).Error("ошибка обращения к серверу")
		return
	}
	lg.Info("grpc client ok (" + res.Message + ")")
}

func (task *GrpcSample) WaitFor() time.Duration {
	return time.Second * 10
}

func (task *GrpcSample) Name() string {
	return GrpcSampleName
}
