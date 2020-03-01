package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"

	"github.com/kshamiev/sungora/internal/model/wsocket"
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/app/response"
	"github.com/kshamiev/sungora/pkg/errs"
	"github.com/kshamiev/sungora/pkg/logger"
)

type Websocket struct {
	chatBus app.WSBus
	*Handler
}

// NewWebsocket общие запросы
func NewWebsocket(h *Handler) *Websocket { return &Websocket{Handler: h, chatBus: app.NewWSServer()} }

// WebSocketSample пример работы с вебсокетом (http://localhost:8080/gorilla/index.html)
// @Summary пример работы с вебсокетом (http://localhost:8080/gorilla/index.html)
// @Tags General
// @Router /api/v1/websocket/gorilla [get]
// @Success 101 {string} string "Switching Protocols to websocket"
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 403 {string} string
// @Failure 404 {string} string
// @Security ApiKeyAuth
func (c *Websocket) WebSocketSample(w http.ResponseWriter, r *http.Request) {
	var (
		lg         = logger.GetLogger(r.Context())
		ws         *websocket.Conn
		wsResponse http.Header
		rw         = response.New(r, w)
		err        error
	)

	// headers response from websocket
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	if ws, err = upgrader.Upgrade(w, r, wsResponse); err != nil {
		rw.JSONError(errs.NewBadRequest(err, "не удалось переключить протокол на ws"))
		return
	}

	defer func() {
		if err = ws.Close(); err != nil {
			lg.WithError(err).Error("WS close connect error")
		} else {
			lg.Info("WS close connect ok")
		}
	}()
	//
	client := &WSHandler{
		Ws:  ws,
		Log: lg,
	}
	bus := c.chatBus.InitClient(chi.URLParam(r, "id"))
	bus.Start(client)
}

// пример обработчика
type WSHandler struct {
	Ws  *websocket.Conn
	Log logger.Logger
}

// HookStartClient метод при подключении и старте нового пользователя
func (h *WSHandler) HookStartClient(cntClient int) {
	h.Log.Info("WS hook start client ", cntClient)
}

// HookGetMessage метод при получении данных из вебсокета пользователя
func (h *WSHandler) HookGetMessage(cntClient int) (interface{}, error) {
	msg := &wsocket.Message{}
	if err := h.Ws.ReadJSON(msg); err != nil {
		h.Log.Error("WS hook get message error: ", err.Error())
		return nil, err
	}
	h.Log.Info("WS hook get message: ", msg)
	// it`s work
	// ...
	return msg, nil
}

// HookSendMessage метод при отправке данных пользователю
func (h *WSHandler) HookSendMessage(msg interface{}, cntClient int) {
	if err := h.Ws.WriteJSON(msg); err != nil {
		h.Log.Error("WS send message err: ", err.Error())
		return
	}

	h.Log.Info("WS hook send message: ", msg)
}

// Ping проверка соединения с пользователем
func (h *WSHandler) Ping() error {
	h.Log.Info("WS hook ping client")
	return h.Ws.WriteMessage(websocket.PingMessage, []byte{})
}
