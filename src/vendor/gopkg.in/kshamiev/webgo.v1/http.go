// Стандартный вебсервер работающий по протоколу http
package app

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"gopkg.in/sungora/app.v1/config"
	"gopkg.in/sungora/app.v1/lg"
)

// newHTTP создание и запуск сервера
func newHTTP(conf config.Server) (store net.Listener, err error) {
	Server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Handler:        newHTTPHandler(),
		ReadTimeout:    time.Second * time.Duration(300),
		WriteTimeout:   time.Second * time.Duration(300),
		MaxHeaderBytes: 1048576,
	}
	for i := 5; i > 0; i-- {
		store, err = net.Listen("tcp", Server.Addr)
		time.Sleep(time.Millisecond * 100)
		if err == nil {
			break
		}
		fmt.Println("connected retry ", Server.Addr)
	}
	if err == nil && store != nil {
		fmt.Println("connected success")
		go Server.Serve(store)
	} else {
		fmt.Println("connected error")
	}
	return
}

type httpHandler struct {
}

// newHTTPHandler Создание УП. Создается пакетом роутера.
func newHTTPHandler() *httpHandler {
	var self = new(httpHandler)
	return self
}

// ServeHTTP Точка входа запроса (в приложение).
func (self *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response(w)
}

func response(Writer http.ResponseWriter) {
	// Тип и Кодировка документа
	Writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	content := []byte("popcorn QWERTY POIUY")
	var err error
	var loc *time.Location

	if loc, err = time.LoadLocation(`Europe/Moscow`); err != nil {
		loc = time.UTC
	}
	t := time.Now().In(loc)
	d := t.Format(time.RFC1123)

	// запрет кеширования
	Writer.Header().Set("Cache-Control", "no-cache, must-revalidate")
	Writer.Header().Set("Pragma", "no-cache")
	Writer.Header().Set("Date", d)
	Writer.Header().Set("Last-Modified", d)
	// размер контента
	Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
	// Статус ответа
	Writer.WriteHeader(200)
	// Тело документа
	Writer.Write(content)
	//

	lg.Info(126,10, "ssssss")
	lg.Warning(126,15, "sfdsflkmhnghsssss")
	lg.Error(126,30, "ssssshfgmh;lgfmhlmfghs")
	lg.Dumper(loc)

	return
}
