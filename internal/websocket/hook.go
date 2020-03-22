package websocket

import (
	"context"

	"github.com/gorilla/websocket"

	"github.com/kshamiev/sungora/internal/model"
	"github.com/kshamiev/sungora/pkg/logger"
)

// пример обработчика
type WSClient struct {
	Ws  *websocket.Conn
	Ctx context.Context
}

// HookStartClient метод при подключении и старте нового пользователя
func (h *WSClient) HookStartClient(cntClient int) error {
	logger.GetLogger(h.Ctx).Info("HookStartClient: ", cntClient)
	return nil
}

// HookGetMessage метод при получении данных из вебсокета пользователя
func (h *WSClient) HookGetMessage(cntClient int) (interface{}, error) {
	lg := logger.GetLogger(h.Ctx)
	msg := &model.Message{}
	if err := h.Ws.ReadJSON(msg); err != nil {
		return nil, err
	}
	lg.Info("HookGetMessage")
	return msg, nil
}

// HookSendMessage метод при отправке данных пользователю
func (h *WSClient) HookSendMessage(msg interface{}, cntClient int) error {
	lg := logger.GetLogger(h.Ctx)
	if err := h.Ws.WriteJSON(msg); err != nil {
		lg.Error("WS send message err: ", err.Error())
		return err
	}
	lg.Info("HookSendMessage")
	return nil
}

// Ping проверка соединения с пользователем
func (h *WSClient) Ping() error {
	logger.GetLogger(h.Ctx).Info("WS hook ping client")
	return h.Ws.WriteMessage(websocket.PingMessage, []byte{})
}
