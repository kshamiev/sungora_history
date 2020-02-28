package app

import (
	"github.com/gorilla/websocket"
	"github.com/kshamiev/sungora/pkg/logger"
)

// интерфейс обработчика
type WSHandler interface {
	HookStartClient(cntClient int)
	HookGetMessage(cntClient int) (interface{}, error)
	HookSendMessage(msg interface{}, cntClient int)
	Ping() error
}

// пример обработчика
type WSHandlerSample struct {
	Ws  *websocket.Conn
	Log logger.Logger
}

// HookStartClient метод при подключении и старте нового пользователя
func (h *WSHandlerSample) HookStartClient(cntClient int) {
	h.Log.Info("WS hook start client ", cntClient)
}

// HookGetMessage метод при получении данных из вебсокета пользователя
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

// HookSendMessage метод при отправке данных пользователю
func (h *WSHandlerSample) HookSendMessage(msg interface{}, cntClient int) {
	if err := h.Ws.WriteJSON(msg); err != nil {
		h.Log.Error("WS send message err: ", err.Error())
		return
	}

	h.Log.Info("WS hook send message: ", msg)
}

// Ping проверка соединения с пользователем
func (h *WSHandlerSample) Ping() error {
	h.Log.Info("WS hook ping client")
	return h.Ws.WriteMessage(websocket.PingMessage, []byte{})
}
