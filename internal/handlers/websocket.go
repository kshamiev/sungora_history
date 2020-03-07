package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/internal/model/wsocket"
	"github.com/kshamiev/sungora/pkg/app/response"
	"github.com/kshamiev/sungora/pkg/errs"
	"github.com/kshamiev/sungora/pkg/logger"
)

type Websocket struct {
	*config.Component
}

// NewWebsocket общие запросы
func NewWebsocket(c *config.Component) *Websocket { return &Websocket{Component: c} }

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

	if ws, err = c.WsBus.RequestUpgrade(w, r, wsResponse); err != nil {
		rw.JSONError(errs.NewBadRequest(err, "не удалось переключить протокол на ws"))
		return
	}
	defer c.WsBus.RequestClose(ws, lg)
	//
	client := &WSClient{
		Ws:  ws,
		Ctx: r.Context(),
	}
	c.WsBus.StartClient(chi.URLParam(r, "id"), client)
}

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
	msg := &wsocket.Message{}
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
