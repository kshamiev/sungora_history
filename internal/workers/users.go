package workers

import (
	"context"
	"database/sql"
	"time"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/logger"
)

const UserOnlineOffName = "UserOnlineOff"

// Обновление онлайн статуса ушедших пользователей
type UserOnlineOff struct {
	db  *sql.DB
	cfg *config.Config
}

// NewUserOnlineOff
func NewUserOnlineOff(db *sql.DB, cfg *config.Config) *UserOnlineOff {
	w := &UserOnlineOff{
		db:  db,
		cfg: cfg,
	}

	return w
}

func (task *UserOnlineOff) Action(ctx context.Context) (err error) {
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
