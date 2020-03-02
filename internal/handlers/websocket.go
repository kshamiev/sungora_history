package handlers

import (
	"context"
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
	wsBus app.WSBus
	*Handler
}

// NewWebsocket общие запросы
func NewWebsocket(h *Handler) *Websocket { return &Websocket{Handler: h, wsBus: app.NewWSServer()} }

// WebSocketSample пример работы с вебсокетом (http://localhost:8080/gorilla/index.html)
// @Summary пример работы с вебсокетом (http://localhost:8080/gorilla/index.html)
// @Tags General
// @Router /api/v1/websocket/gorilla/{id} [get]
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

	if ws, err = c.wsBus.RequestUpgrade(w, r, wsResponse); err != nil {
		rw.JSONError(errs.NewBadRequest(err, "не удалось переключить протокол на ws"))
		return
	}
	defer c.wsBus.RequestClose(ws, lg)
	//
	client := &WSHandler{
		Ws:  ws,
		Ctx: r.Context(),
	}
	c.wsBus.StartClient(chi.URLParam(r, "id"), client)
}

// пример обработчика
type WSHandler struct {
	Ws  *websocket.Conn
	Ctx context.Context
}

// HookStartClient метод при подключении и старте нового пользователя
func (h *WSHandler) HookStartClient(cntClient int) {
	lg := logger.GetLogger(h.Ctx)
	lg.Info("WS hook start client ", cntClient)
}

// HookGetMessage метод при получении данных из вебсокета пользователя
func (h *WSHandler) HookGetMessage(cntClient int) (interface{}, error) {
	lg := logger.GetLogger(h.Ctx)
	msg := &wsocket.Message{}
	if err := h.Ws.ReadJSON(msg); err != nil {
		lg.Error("WS hook get message error: ", err.Error())
		return nil, err
	}
	lg.Info("WS hook get message: ", msg)
	// it`s work
	// ...
	return msg, nil
}

// HookSendMessage метод при отправке данных пользователю
func (h *WSHandler) HookSendMessage(msg interface{}, cntClient int) {
	lg := logger.GetLogger(h.Ctx)
	if err := h.Ws.WriteJSON(msg); err != nil {
		lg.Error("WS send message err: ", err.Error())
		return
	}

	lg.Info("WS hook send message: ", msg)
}

// Ping проверка соединения с пользователем
func (h *WSHandler) Ping() error {
	lg := logger.GetLogger(h.Ctx)
	lg.Info("WS hook ping client")
	return h.Ws.WriteMessage(websocket.PingMessage, []byte{})
}
