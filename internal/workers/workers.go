package workers

import (
	"database/sql"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/logger"
)

func Init(db *sql.DB, cfg *config.Config, lg logger.Logger) *app.Scheduler {
	wp := app.NewScheduler(lg)
	wp.Add(newUserOnlineOff(db, cfg))
	return wp
}
