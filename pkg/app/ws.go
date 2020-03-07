package app

import (
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/kshamiev/sungora/pkg/logger"
)

// интерфейс клиента для взаимодействия с вебсокетом
type WSClient interface {
	// Вызывается единожды в момент подключения нового клиента (заходит на страницу с вебсокетом)
	HookStartClient(cntClient int) error
	// Читает сообщения вебсокета
	HookGetMessage(cntClient int) (interface{}, error)
	// Отправляет сообщения в вебсокет (централизованно вызывается для всех подключенных)
	HookSendMessage(msg interface{}, cntClient int) error
	// Проверка и пролонгация подключения
	Ping() error
}

// шина обработчиков вебсокетов по идентификаторам
type WSBus map[string]*WSHandler

// NewWSServer создание шины
func NewWSServer() WSBus { return make(WSBus) }

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

// GetWSHandler получение обработчика
func (bus WSBus) GetWSHandler(wsbusID string) (*WSHandler, error) {
	if _, ok := bus[wsbusID]; ok {
		return bus[wsbusID], nil
	}
	return nil, errors.New("not found ws handler")
}

// StartClient инициализация клиента по условному идентификатору
// инициализация и регистрация обработчика в шине
// регистрация и старт работы нового клиента в обратчике
// управление работой клиента
func (bus WSBus) StartClient(wsbusID string, client WSClient) {
	// инициализация обработчика в шине
	room, ok := bus[wsbusID]
	if !ok {
		room = &WSHandler{
			broadcast: make(chan interface{}),
			clients:   make(map[WSClient]bool),
		}
		bus[wsbusID] = room
		go room.control()
	}

	// регистрация клиента и завершение работы
	room.clients[client] = true
	defer func() {
		delete(room.clients, client)
		if len(room.clients) == 0 {
			room.isClose = true
			delete(bus, wsbusID)
		}
	}()

	// старт работы клиента
	if err := client.HookStartClient(len(room.clients)); err != nil {
		_ = client.HookSendMessage(err, len(room.clients))
		return
	}

	// здесь мы лочимся и обрабатываем входящие сообщения до выхода
	for {
		msg, err := client.HookGetMessage(len(room.clients))
		if err != nil {
			if _, ok := err.(*websocket.CloseError); !ok {
				_ = client.HookSendMessage(err, len(room.clients))
			}
			return
		}
		if msg != nil {
			room.broadcast <- msg // посылаем всем обработчикам (подключенным пользователям)
		}
	}
}

// обработчик клиентов
type WSHandler struct {
	broadcast chan interface{}  // канал передачи данных всем клиентам обработчика
	clients   map[WSClient]bool // массив всех клиентов обработчика
	isClose   bool              // признак завершения работы обработчика
}

// control здесь мы отправляем сообщения всем подключенным клиентам и пингуем
func (room *WSHandler) control() {
	ticker := time.NewTicker(time.Minute)
	for {
		if room.isClose == true {
			return
		}
		select {
		// проверка соединений с клиентами
		case <-ticker.C:
			for handler := range room.clients {
				// если достучаться до клиента не удалось, то удаляем его
				if err := handler.Ping(); err != nil {
					delete(room.clients, handler)
					continue
				}
			}
		// каждому зарегистрированному клиенту шлем сообщение
		case message := <-room.broadcast:
			for handler := range room.clients {
				_ = handler.HookSendMessage(message, len(room.clients))
			}
		}
	}
}

// SendMessage отправка сообщений всем покдлюченным клиентам
func (room *WSHandler) SendMessage(msg interface{}) {
	if msg != nil {
		room.broadcast <- msg
	}
}

// GetClientCnt получение количества подключенных клиентов
func (room *WSHandler) GetClientCnt() int {
	return len(room.clients)
}
