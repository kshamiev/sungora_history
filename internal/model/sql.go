package model

import "github.com/kshamiev/sungora/pkg/app"

// Хранилище нативных запросов к БД
const (
	SQLAppVersion app.SQL = `SELECT MAX(version_id) as version_id FROM goose_db_version WHERE is_applied = TRUE`
)
