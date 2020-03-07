package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/app/response"
	"github.com/kshamiev/sungora/pkg/logger"
)

type Middleware struct {
	*config.Component
}

// NewMiddleware промежуточные обработчики запрсоов
func NewMiddleware(c *config.Component) *Middleware { return &Middleware{Component: c} }

// TimeoutContext инициализация таймаута контекста для контроля времени выполениня запроса
func (c *Middleware) TimeoutContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), c.Cfg.ServHTTP.RequestTimeout-time.Millisecond)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Logger формирование логера для запроса
func (c *Middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		lg := c.Lg.WithField(response.LogUUID, requestID).WithField(response.LogAPI, r.URL.Path)

		ctx := r.Context()
		ctx = context.WithValue(ctx, response.CtxUUID, requestID)
		ctx = logger.WithLogger(ctx, lg)
		ctx = boil.WithDebugWriter(ctx, lg.Writer())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Cors добавление заголовка ConfigCors
func (c *Middleware) Cors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   c.Cfg.Cors.AllowedOrigins,
		AllowedMethods:   c.Cfg.Cors.AllowedMethods,
		AllowedHeaders:   c.Cfg.Cors.AllowedHeaders,
		ExposedHeaders:   c.Cfg.Cors.ExposedHeaders,
		AllowCredentials: c.Cfg.Cors.AllowCredentials,
		MaxAge:           c.Cfg.Cors.MaxAge, // Maximum value not ignored by any of major browsers
	})
}

// Static статика или отдача существующего файла по запросу
func (c *Middleware) Static(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.Static(c.Cfg.App.DirStatic + r.URL.Path)
}
