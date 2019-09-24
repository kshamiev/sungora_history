package app

import (
	"time"
)

// шина чатов
type WSBus map[string]*WSClient

// NewWSBus создание шины чатов
func NewWSBus() WSBus {
	bus := make(WSBus)
	go bus.controlBus()
	return bus

}

// controlBus жизненный цикл шины чатов
func (bus WSBus) controlBus() {
	for range time.NewTicker(time.Minute).C {
		for i := range bus {
			if len(bus[i].clients) == 0 {
				delete(bus, i)
			}
		}
	}
}

// InitClient инициализация чата по условному идентификатору
func (bus WSBus) InitClient(wsbusID string) *WSClient {
	if b, ok := bus[wsbusID]; ok {
		return b
	}
	b := &WSClient{
		broadcast: make(chan interface{}),
		clients:   make(map[WSHandler]bool),
	}
	bus[wsbusID] = b
	go b.control()
	return b
}

// чат
type WSClient struct {
	broadcast chan interface{}   // канал рассылки сообщений клиентам
	clients   map[WSHandler]bool // массив всех клиентов чата
}

// start управление чатом
func (b *WSClient) control() {
	ticker := time.NewTicker(time.Second * 50)
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

// Start регистрация и старт работы нового клиента
func (b *WSClient) Start(handler WSHandler) {
	b.clients[handler] = true
	handler.HookStartClient(len(b.clients))
	for {
		msg, err := handler.HookGetMessage(len(b.clients))
		if err != nil {
			delete(b.clients, handler)
			return
		}
		b.broadcast <- msg // посылаем всем подключенным пользователям
	}
}
