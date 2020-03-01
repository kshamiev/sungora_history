package handlers

import (
	"io"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/kshamiev/sungora/pkg/app/response"
)

type General struct {
	*Handler
}

// NewGeneral общие запросы
func NewGeneral(h *Handler) *General { return &General{Handler: h} }

// Ping ping
// @Summary ping
// @Tags General
// @Router /api/v1/general/ping [get]
// @Success 200 {string} string "OK"
// @Failure 500 {string} string
func (c *General) Ping(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON("OK")
}

// GetVersion получение версии приложения
// @Summary получение версии приложения
// @Tags General
// @Router /api/v1/general/version [get]
// @Success 200 {string} string "version"
// @Failure 500 {string} string
func (c *General) GetVersion(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON(c.cfg.App.Version)
}

// WebSocketSample пример работы с вебсокетом (http://localhost:8080/websocket/index.html)
// @Summary пример работы с вебсокетом (http://localhost:8080/websocket/index.html)
// @Tags General
// @Router /api/v1/general/websocket [get]
// @Success 101 {string} string "Switching Protocols to websocket"
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 403 {string} string
// @Failure 404 {string} string
// @Security ApiKeyAuth
func (c *General) WebSocketSample(ws *websocket.Conn) {
	_, _ = io.Copy(ws, ws)
}
