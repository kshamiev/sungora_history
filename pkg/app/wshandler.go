package app

import (
	"github.com/gorilla/websocket"

	"github.com/kshamiev/sungora/pkg/logger"
)

// интерфейс обработчика сообщений
type WSHandler interface {
	HookStartClient(cntClient int)
	HookGetMessage(cntClient int) (interface{}, error)
	HookSendMessage(msg interface{}, cntClient int)
	Ping() error
}

// пример обработчика сообщений
type WSHandlerSample struct {
	Ws  *websocket.Conn
	Log logger.Logger
}

// HookStartClient метод при старте чата
func (h *WSHandlerSample) HookStartClient(cntClient int) {
	h.Log.Info("WS hook start client ", cntClient)
}

// HookGetMessage метод при получении сообщения чата для конкретного клиента
func (h *WSHandlerSample) HookGetMessage(cntClient int) (interface{}, error) {
	var msg interface{}
	if err := h.Ws.ReadJSON(&msg); err != nil {
		h.Log.Error("WS hook get message error: ", err.Error())
		return nil, err
	}

	h.Log.Info("WS hook get message: ", msg)
	// it`s work
	// ...
	return &msg, nil
}

// HookSendMessage метод при отправке сообщения клиенту (выполняется для всех пользователей конкретного чата)
func (h *WSHandlerSample) HookSendMessage(msg interface{}, cntClient int) {
	if err := h.Ws.WriteJSON(msg); err != nil {
		h.Log.Error("WS send message err: ", err.Error())
		return
	}

	h.Log.Info("WS hook send message: ", msg)
}

// Ping проверка соединения с клиентом
func (h *WSHandlerSample) Ping() error {
	h.Log.Info("WS hook ping client")
	return h.Ws.WriteMessage(websocket.PingMessage, []byte{})
}
