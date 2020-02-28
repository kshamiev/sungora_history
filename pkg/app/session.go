package app

import (
	"crypto/rand"
	"io"
	"time"

	"github.com/kshamiev/sungora/pkg/app/response"
)

// Шина сессий
type SessionBus map[string]*Session

// NewSession создание шины сессий по таймауту
func NewSession(sessionTimeout time.Duration) SessionBus {
	var sessions = make(SessionBus)
	go sessions.controlSessionBus(sessionTimeout)

	return sessions
}

// controlSessionBus жизненный цикл сессии
func (bus SessionBus) controlSessionBus(sessionTimeout time.Duration) {
	for {
		time.Sleep(time.Minute)
		for key := range bus {
			if sessionTimeout < time.Since(bus[key].t) {
				delete(bus, key)
			}
		}
	}
}

// Set is a storage setter
func (bus SessionBus) Set(key string, val interface{}) {
	elm := new(Session)
	elm.t = time.Now()
	elm.data = map[string]interface{}{
		key: val,
	}
	bus[key] = elm
}

// Get is a storage getter
func (bus SessionBus) Get(key string) interface{} {
	if elm, ok := bus[key]; ok {
		if _, ok := elm.data[key]; ok {
			return elm.data[key]
		}
	}

	return nil
}

// Del removes storage
func (bus SessionBus) Del(key string) {
	delete(bus, key)
}

// GetSessionCookie Получение сессии по куке пришедшей из запроса
func (bus SessionBus) GetSessionCookie(rw *response.Response, cookieName string) *Session {
	token, _ := rw.CookieGet(cookieName)
	if token == "" {
		token = newRandomString(10)
		rw.CookieSet(cookieName, token)
	}

	return bus.GetSession(token)
}

// GetSession Получение сессии по токену
func (bus SessionBus) GetSession(token string) *Session {
	if elm, ok := bus[token]; ok {
		elm.t = time.Now()
		return elm
	}

	elm := new(Session)
	elm.t = time.Now()
	elm.data = make(map[string]interface{})
	bus[token] = elm

	return elm
}

// Сессия
type Session struct {
	t    time.Time
	data map[string]interface{}
}

// Get получение данных сессии
func (s *Session) Get(key string) interface{} {
	if _, ok := s.data[key]; ok {
		return s.data[key]
	}

	return nil
}

// Set сохранение данных в сессии
func (s *Session) Set(key string, value interface{}) {
	s.data[key] = value
}

// Del удаление данных из сессии
func (s *Session) Del(key string) {
	delete(s.data, key)
}

// ////

const (
	num     = "0123456789"
	strdown = "abcdefghijklmnopqrstuvwxyz"
	strup   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// newRandomString generates password key of a specified length
func newRandomString(length int) string {
	return randChar(length, []byte(strdown+strup+num))
}

func randChar(length int, chars []byte) string {
	pword := make([]byte, length)
	data := make([]byte, length+(length/4)) // storage for random bytes.
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0

	for {
		if _, err := io.ReadFull(rand.Reader, data); err != nil {
			panic(err)
		}

		for _, c := range data {
			if c >= maxrb {
				continue
			}

			pword[i] = chars[c%clen]
			i++

			if i == length {
				return string(pword)
			}
		}
	}
}
