package workers

import (
	"database/sql"

	logger "github.com/kshamiev/sungora/pkg/logger"
	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/app"
)

func Init(db *sql.DB, cfg *config.Config, lg logger.Logger) *app.Scheduler {
	wp := app.NewScheduler(lg)
	wp.Add(newUserOnlineOff(db, lg, cfg))
	return wp
}
