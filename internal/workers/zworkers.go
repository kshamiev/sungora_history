package workers

import (
	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/app"
)

// New инициализация фоновых задач (воркеров)
func New(comp *config.Component) (*app.Scheduler, error) {
	wp := app.NewScheduler(comp.Lg)
	wp.AddStart(NewUserOnlineOff(comp))
	return wp, nil
}
