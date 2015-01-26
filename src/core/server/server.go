package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"lib/logs"
	typConfig "types/config"
)

// Канал передачи команд
var commandFactorControl = make(chan commandFactor)

// Команда
type commandFactor struct {
	action int
	name   string            // Уникальное имя сервера
	server *typConfig.Server // конфигурация запускаемого сервера
	result chan<- interface{}
}

// Допустимые команды (action)
const (
	factorStart int = iota
	factorStop
)

func init() {
	goServerFactory()
}

// goServerFactory Запуск сервера
func goServerFactory() {
	var store = make(map[string]net.Listener)
	var err error
	go func() {
		for command := range commandFactorControl {
			switch command.action {
			// запуск сервера
			case factorStart:
				go func(command commandFactor) {
					webserverAddress := fmt.Sprintf("%s:%d", command.server.Host, command.server.Port)
					Server := &http.Server{
						Addr:           webserverAddress,
						Handler:        newServer(command.server),
						ReadTimeout:    time.Second * time.Duration(command.server.ReadTimeout),
						WriteTimeout:   time.Second * time.Duration(command.server.WriteTimeout),
						MaxHeaderBytes: int(command.server.MaxHeaderBytes),
					}
					for i := 5; i > 0; i-- {
						store[command.name], err = net.Listen("tcp", webserverAddress)
						time.Sleep(time.Millisecond * 100)
						if err == nil {
							break
						}
						logs.Base.Error(127, webserverAddress)
					}
					if err == nil && store[command.name] != nil {
						logs.Base.Notice(130, webserverAddress)
						command.result <- true
						Server.Serve(store[command.name])
						logs.Base.Notice(128, webserverAddress)
						delete(store, command.name)
					} else {
						logs.Base.Error(129, webserverAddress)
						command.result <- false
						delete(store, command.name)
					}
				}(command)
			// остановка сервера
			case factorStop:
				if _, ok := store[command.name]; ok == false {
					command.result <- true
				} else if err := store[command.name].Close(); err != nil {
					command.result <- false
				} else {
					for ok {
						time.Sleep(time.Millisecond * 1)
						_, ok = store[command.name]
					}
					command.result <- true
				}
			}
		}
	}()
}

// ServerStartGo Запуск работы сервера
func GoStart(nameIndex string, serverConfig *typConfig.Server) bool {
	reply := make(chan interface{})
	commandFactorControl <- commandFactor{action: factorStart, name: nameIndex, server: serverConfig, result: reply}
	return (<-reply).(bool)
}

// ServerStopGo Завершение работы сервера
func GoStop(nameIndex string) bool {
	reply := make(chan interface{})
	commandFactorControl <- commandFactor{action: factorStop, name: nameIndex, result: reply}
	return (<-reply).(bool)
}
