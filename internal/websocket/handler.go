package websocket

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"

	"github.com/kshamiev/sungora/internal/config"
	"github.com/kshamiev/sungora/pkg/app/response"
	"github.com/kshamiev/sungora/pkg/errs"
	"github.com/kshamiev/sungora/pkg/logger"
)

type Websocket struct {
	*config.Component
}

// NewWebsocket общие запросы
func NewWebsocket(c *config.Component) *Websocket { return &Websocket{Component: c} }

// GetWebSocketSample пример работы с вебсокетом (http://localhost:8080/gorilla/index.html)
// @Summary пример работы с вебсокетом (http://localhost:8080/gorilla/index.html)
// @Tags General
// @Router /api/v1/websocket/gorilla/{id} [get]
// @Success 101 {string} string "Switching Protocols to websocket"
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 403 {string} string
// @Failure 404 {string} string
// @Security ApiKeyAuth
func (c *Websocket) GetWebSocketSample(w http.ResponseWriter, r *http.Request) {
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
