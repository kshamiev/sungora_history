package workers

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/internal/model"
	"github.com/kshamiev/sungora/pb/typsun"
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/app/request"
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
	ctx = request.ContextGRPC(ctx, nil)

	res, err := task.SungoraClient.GetUser(ctx, &empty.Empty{})
	if err != nil {
		return errs.NewGRPC(err)
	}

	user := model.User{
		Type: typsun.NewUserProto(res),
	}

	app.Dumper(
		user.Type.Price.String(),
		user.Type.CreatedAt.String(),
		user.Type.Login,
		user.Type.Metrika,
		user.Type.ID.String(),
		user.Type.Alias,
	)

	lg.Info("grpc client ok")
	return nil
}

func (task *GrpcSample) WaitFor() time.Duration {
	return time.Second * 10
}

func (task *GrpcSample) Name() string {
	return GrpcSampleName
}
