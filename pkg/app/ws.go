package app

import (
	"net/http"
	"time"

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

// шина обработчиков вебсокетов по идентификаторам
type WSBus map[string]*wsHandler

// NewWSServer создание шины
func NewWSServer() WSBus {
	bus := make(WSBus)
	go bus.clearBus()

	return bus
}

// clearBus удаление пустых обработчиков
func (bus WSBus) clearBus() {
	for range time.NewTicker(time.Hour).C {
		for i := range bus {
			if len(bus[i].clients) == 0 {
				delete(bus, i)
			}
		}
	}
}

// RequestUpgrade переключение протокола на вебсокет
func (bus WSBus) RequestUpgrade(w http.ResponseWriter, r *http.Request, h http.Header) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return upgrader.Upgrade(w, r, h)
}

// RequestClose закрытие вебсокета
func (bus WSBus) RequestClose(ws *websocket.Conn, lg logger.Logger) {
	if err := ws.Close(); err != nil {
		lg.WithError(err).Error("WS close connect error")
	} else {
		lg.Info("WS close connect ok")
	}
}

// клиент обслуживающий вебсокет
type wsHandler struct {
	broadcast chan interface{}   // канал передачи данных всем обработчикам вебсокета
	clients   map[WSHandler]bool // массив всех обработчиков вебсокета
}

// StartClient инициализация обработчика вебсокета по условному идентификатору
// регистрация и старт работы нового пользователя
func (bus WSBus) StartClient(wsbusID string, handler WSHandler) {
	b, ok := bus[wsbusID]
	if !ok {
		b = &wsHandler{
			broadcast: make(chan interface{}),
			clients:   make(map[WSHandler]bool),
		}
		bus[wsbusID] = b
		go b.control()
	}
	// регистрация и старт работы нового пользователя
	b.clients[handler] = true
	defer delete(b.clients, handler)
	handler.HookStartClient(len(b.clients))

	for {
		msg, err := handler.HookGetMessage(len(b.clients))
		if err != nil {
			return
		}
		if msg != nil {
			b.broadcast <- msg // посылаем всем обработчикам (подключенным пользователям)
		}
	}
}

// control работа обработчика вебсокета
func (b *wsHandler) control() {
	ticker := time.NewTicker(time.Second * 55)

	for {
		select {
		// проверка соединений с клиентами
		case <-ticker.C:
			for handler := range b.clients {
				// если достучаться до клиента не удалось, то удаляем его
				if err := handler.Ping(); err != nil {
					delete(b.clients, handler)
					continue
				}
			}
		// каждому зарегистрированному клиенту шлем сообщение
		case message := <-b.broadcast:
			for handler := range b.clients {
				handler.HookSendMessage(message, len(b.clients))
			}
		}
	}
}
