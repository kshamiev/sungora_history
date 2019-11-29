package workers

import (
	"database/sql"
	"time"

 logger	"github.com/kshamiev/sungora/pkg/logger"

	"github.com/kshamiev/sungora/internal/config"
)

// Обновление онлайн статуса ушедших пользователей
type userOnlineOff struct {
	name string
	db   *sql.DB
	cfg  *config.Config
	lg   logger.Logger
}

// NewUserOnlineOff
func newUserOnlineOff(db *sql.DB, lg logger.Logger, cfg *config.Config) *userOnlineOff {
	w := &userOnlineOff{
		name: "userOnlineOff",
		db:   db,
		cfg:  cfg,
	}
	w.lg = lg.WithField("task", w.name)
	return w
}

func (task *userOnlineOff) Action() (err error) {
	task.lg.Infof("execute task: %s", task.name)
	return
}

func (task *userOnlineOff) Info() (string, time.Duration) {
	return task.name, time.Minute
}
