package workers

import (
	"context"
	"fmt"
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

func (task *GrpcSample) Action(ctx context.Context) error {
	lg := logger.GetLogger(ctx)
	lg.Info("grpc client request")

	res, err := task.SungoraClient.HelloWorld(ctx, &proto.TestRequest{Name: "запрос от клиента"})
	if err != nil {
		return err
	}

	lg.Info("grpc client response: " + res.Message)
	fmt.Println()

	return nil
}

func (task *GrpcSample) WaitFor() time.Duration {
	return time.Second * 10
}

func (task *GrpcSample) Name() string {
	return GrpcSampleName
}
