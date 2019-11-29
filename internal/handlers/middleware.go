package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/cors"
	"github.com/google/uuid"

	"github.com/kshamiev/sungora/pkg/app/response"
	"github.com/kshamiev/sungora/pkg/logger"
)

type Middleware struct {
	*Main
}

// NewMiddleware промежуточные обработчики запрсоов
func NewMiddleware(main *Main) *Middleware { return &Middleware{main} }

// Logger формирование логера для запроса
func (c *Middleware) Logger(lg logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(logger.WithLogger(r.Context(), lg.WithField("request", uuid.New().String()))))
		})
	}
}

// TimeoutContext
// инициализация таймаута контекста для контроля времени выполениня запроса
func (c *Middleware) TimeoutContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), c.cfg.ServHTTP.RequestTimeout-time.Millisecond)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ConfigCors добавление заголовка ConfigCors
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
func (c *Middleware) Static() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := response.New(r, w)
		rw.Static(c.cfg.App.DirWork + r.URL.Path)
	})
}

// SampleOne пример middleware
func (c *Middleware) SampleOne(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lg := logger.GetLogger(r.Context())
		lg.Info("Middleware SampleOne")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// SampleTwo пример middleware
func (c *Middleware) SampleTwo(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lg := logger.GetLogger(r.Context())
		lg.Info("Middleware SampleTwo")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
