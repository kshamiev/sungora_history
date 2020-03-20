package config

import (
	"database/sql"

	"github.com/kshamiev/sungora/pb"
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/logger"
)

type Component struct {
	Db            *sql.DB
	Lg            logger.Logger
	Cfg           *Config
	Wp            *app.Scheduler
	WsBus         app.WSBus
	SungoraClient pb.SungoraClient
}
