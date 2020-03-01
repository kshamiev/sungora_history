package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/kshamiev/sungora/pkg/app/response"
	"github.com/kshamiev/sungora/pkg/logger"
)

type Middleware struct {
	*Handler
}

// NewMiddleware промежуточные обработчики запрсоов
func NewMiddleware(h *Handler) *Middleware { return &Middleware{Handler: h} }

// TimeoutContext инициализация таймаута контекста для контроля времени выполениня запроса
func (c *Middleware) TimeoutContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), c.cfg.ServHTTP.RequestTimeout-time.Millisecond)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Logger формирование логера для запроса
func (c *Middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		lg := c.lg.WithField(response.LogUUID, requestID).WithField(response.LogAPI, r.URL.Path)

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
		AllowedOrigins:   c.cfg.Cors.AllowedOrigins,
		AllowedMethods:   c.cfg.Cors.AllowedMethods,
		AllowedHeaders:   c.cfg.Cors.AllowedHeaders,
		ExposedHeaders:   c.cfg.Cors.ExposedHeaders,
		AllowCredentials: c.cfg.Cors.AllowCredentials,
		MaxAge:           c.cfg.Cors.MaxAge, // Maximum value not ignored by any of major browsers
	})
}

// Static статика или отдача существующего файла по запросу
func (c *Middleware) Static(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.Static(c.cfg.App.DirStatic + r.URL.Path)
}
