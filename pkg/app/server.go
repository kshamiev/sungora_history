package app

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/kshamiev/sungora/pkg/logger"
)

// сервер http(s)
type Server struct {
	Server    *http.Server  // сервер HTTP
	chControl chan struct{} // управление ожиданием завершения работы сервера
	lis       net.Listener
}

// NewServer создание и старт вебсервера (HTTP(S))
func NewServer(cfg *ConfigServer, mux http.Handler, lg logger.Logger) (comp *Server, err error) {
	comp = &Server{
		Server: &http.Server{
			Addr:           fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler:        mux,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			IdleTimeout:    cfg.IdleTimeout,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
		},
		chControl: make(chan struct{}),
	}

	if comp.lis, err = net.Listen("tcp", comp.Server.Addr); err != nil {
		return
	}

	go func() {
		_ = comp.Server.Serve(comp.lis)
		close(comp.chControl)
	}()

	lg.Info("start web server: ", comp.Server.Addr)

	return comp, nil
}

// Wait завершение работы сервера (HTTP(S))
func (comp *Server) Wait(lg logger.Logger) {
	if comp.lis == nil {
		return
	}

	if err := comp.lis.Close(); err != nil {
		return
	}

	<-comp.chControl
	lg.Info("stop web server: ", comp.Server.Addr)
}

// GetRoute получение обработчика запросов
func (comp *Server) GetRoute() *chi.Mux {
	return comp.Server.Handler.(*chi.Mux)
}
