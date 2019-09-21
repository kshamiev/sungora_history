package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// компонент
type Server struct {
	Server    *http.Server  // сервер HTTP
	chControl chan struct{} // управление ожиданием завершения работы сервера
}

// NewServer создание компонента вебсервера
func NewServer(cfg *ConfigServer, mux http.Handler) (com *Server, err error) {
	return &Server{
		Server: &http.Server{
			Addr:           fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler:        mux,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			IdleTimeout:    cfg.IdleTimeout,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
		},
	}, nil
}

// Run запуск компонента в работу
// Старт сервера HTTP(S)
func (comp *Server) Run() (err error) {
	comp.chControl = make(chan struct{})
	go func() {
		if err = comp.Server.ListenAndServe(); err != http.ErrServerClosed {
			return
		}
		close(comp.chControl)
	}()
	return
}

// Stop завершение работы компонента
// Остановка сервера HTTP(S)
func (comp *Server) Wait() {
	if comp.Server == nil {
		return
	}
	if err := comp.Server.Shutdown(context.Background()); err != nil {
		return
	}
	<-comp.chControl
}

// GetRoute получение обработчика запросов
func (comp *Server) GetRoute() *chi.Mux {
	return comp.Server.Handler.(*chi.Mux)
}
