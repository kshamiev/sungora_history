package app

import (
	"fmt"
	"net"

	"github.com/kshamiev/sungora/pkg/logger"
	"google.golang.org/grpc"
)

// клиент GRPC
type GRPCClient struct {
	Conn *grpc.ClientConn
}

// NewGRPCClient создание и старт клиента GRPC
func NewGRPCClient(cfg *ConfigGRPC) (comp *GRPCClient, err error) {
	comp = &GRPCClient{}
	comp.Conn, err = grpc.Dial(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), grpc.WithInsecure())
	return comp, err
}

// Wait завершение работы клиента GRPC
func (comp *GRPCClient) Wait() {
	_ = comp.Conn.Close()
}

// сервер GRPC
type GRPCServer struct {
	Addr      string        // адрес сервера grpc
	chControl chan struct{} // управление ожиданием завершения работы сервера
	lis       net.Listener
}

// NewGRPCServer создание и старт сервера GRPC
func NewGRPCServer(cfg *ConfigGRPC, serve *grpc.Server, lg logger.Logger) (comp *GRPCServer, err error) {
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

// Wait завершение работы сервера GRPC
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
