package app

import (
	"time"
)

// интерфейс обработчика
type WSHandler interface {
	HookStartClient(cntClient int)
	HookGetMessage(cntClient int) (interface{}, error)
	HookSendMessage(msg interface{}, cntClient int)
	Ping() error
}

// шина обработчиков вебсокетов по идентификаторам
type WSBus map[string]*WSClient

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

// InitClient инициализация обработчика вебсокета по условному идентификатору
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

// клиент обслуживающий вебсокет
type WSClient struct {
	broadcast chan interface{}   // канал передачи данных всем обработчикам вебсокета
	clients   map[WSHandler]bool // массив всех обработчиков вебсокета
}

// control основной жизненый цикл клиента
func (b *WSClient) control() {
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
		case message := <-b.broadcast: // каждому зарегистрированному клиенту шлем сообщение
			for handler := range b.clients {
				handler.HookSendMessage(message, len(b.clients))
			}
		}
	}
}

// Start регистрация и старт работы нового обработчика
func (b *WSClient) Start(handler WSHandler) {
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
