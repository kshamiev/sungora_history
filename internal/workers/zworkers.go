package workers

import (
	"github.com/kshamiev/sungora/internal/handlers"
	"github.com/kshamiev/sungora/pkg/app"
)

// New инициализация фоновых задач (воркеров)
func New(comp *handlers.Component) (*app.Scheduler, error) {
	wp := app.NewScheduler(comp.Lg)
	wp.AddStart(NewUserOnlineOff(comp.Db, comp.Cfg))
	return wp, nil
}
