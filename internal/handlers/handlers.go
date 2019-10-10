package handlers

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"gitlab.services.mts.ru/libs/logger"

	"github.com/kshamiev/sungora/internal/config"
)

type Main struct {
	db         *sql.DB
	cfg        *config.Config
	middleware *Middleware
	general    *General
}

// NewMain
func NewMain(db *sql.DB, lg logger.Logger, cfg *config.Config) *chi.Mux {
	c := &Main{
		db:  db,
		cfg: cfg,
	}
	c.middleware = NewMiddleware(c)

	router := chi.NewRouter()
	router.Use(c.middleware.Cors().Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(cfg.ServHTTP.RequestTimeout))
	router.Use(c.middleware.Logger(lg))
	router.NotFound(c.middleware.Static())
	router.Get("/api/docs/*", httpSwagger.Handler()) // swagger

	c.general = NewGeneral(c)
	router.Get("/api/v1/general/ping", c.general.Ping)
	router.Get("/api/v1/general/version", c.general.GetVersion)

	return router
}
