package app

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/kshamiev/sungora/pkg/logger"
)

// компонент
type GRPCServer struct {
	Addr      string        // адрес сервера grpc
	chControl chan struct{} // управление ожиданием завершения работы сервера
	lis       net.Listener
}

// NewGRPC создание компонента сервера GRPC
// Старт сервера HTTP(S)
func NewGRPC(cfg *ConfigGRPC, serve *grpc.Server, lg logger.Logger) (comp *GRPCServer, err error) {
	comp = &GRPCServer{
		Addr:      fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		chControl: make(chan struct{}),
	}

	if comp.lis, err = net.Listen("tcp", comp.Addr); err != nil {
		return
	}

	go func() {
		_ = serve.Serve(comp.lis)
		close(comp.chControl)
	}()

	lg.Info("start grpc server: ", comp.Addr)

	return comp, nil
}

// Wait завершение работы компонента
// Остановка сервера GRPC
func (comp *GRPCServer) Wait(lg logger.Logger) {
	if comp.lis == nil {
		return
	}

	if err := comp.lis.Close(); err != nil {
		return
	}

	<-comp.chControl
	lg.Info("stop grpc server: ", comp.Addr)
}
