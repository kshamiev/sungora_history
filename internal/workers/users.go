package workers

import (
	"context"
	"time"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/logger"
)

const UserOnlineOffName = "UserOnlineOff"

// Обновление онлайн статуса ушедших пользователей
type UserOnlineOff struct {
	*config.Component
}

// NewUserOnlineOff
func NewUserOnlineOff(c *config.Component) *UserOnlineOff { return &UserOnlineOff{Component: c} }

func (task *UserOnlineOff) Action(ctx context.Context) {
	lg := logger.GetLogger(ctx)
	lg.Infof("execute task: %s", UserOnlineOffName)

	return
}

func (task *UserOnlineOff) WaitFor() time.Duration {
	return time.Minute
}

func (task *UserOnlineOff) Name() string {
	return UserOnlineOffName
}
