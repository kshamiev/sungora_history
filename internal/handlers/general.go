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

// WebSocketSample пример работы с вебсокетом
// @Summary пример работы с вебсокетом
// @Tags General
// @Router /api/v1/general/websocket [get]
// @Success 101 {string} string "Switching Protocols"
func (c *General) WebSocketSample(ws *websocket.Conn) {
	_, _ = io.Copy(ws, ws)
}
