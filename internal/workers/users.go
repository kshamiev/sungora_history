package workers

import (
	"context"
	"database/sql"
	"time"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/logger"
)

const userOnlineOffName = "userOnlineOff"

// Обновление онлайн статуса ушедших пользователей
type userOnlineOff struct {
	db  *sql.DB
	cfg *config.Config
}

// NewUserOnlineOff
func newUserOnlineOff(db *sql.DB, cfg *config.Config) *userOnlineOff {
	w := &userOnlineOff{
		db:  db,
		cfg: cfg,
	}
	return w
}

func (task *userOnlineOff) Action(ctx context.Context) (err error) {
	lg := logger.GetLogger(ctx)
	lg.Infof("execute task: %s", userOnlineOffName)
	return
}

func (task *userOnlineOff) Info() (string, time.Duration) {
	return userOnlineOffName, time.Minute
}
