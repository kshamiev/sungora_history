// Управление запуском и остановкой приложения
package app

import (
	"net"
	"os"
	"fmt"
	"os/signal"

	"gopkg.in/sungora/app.v1/config"
	"gopkg.in/sungora/app.v1/lg"
)

// Каналы управления запуском и остановкой приложения
var (
	chanelAppStop    = make(chan os.Signal, 1)
	chanelAppControl = make(chan os.Signal, 1)
)

// Start Launch an application
func Start() (code int) {

	// завершение работы приложения
	defer func() {
		chanelAppStop <- os.Interrupt
	}()

	var (
		err   error
		args  *arguments
		store net.Listener
		conf  *config.Config
	)

	// входные данные командной строки
	args, err = getCmdArgs()
	if err != nil {
		return 0
	}

	// конфигурация
	if conf, err = config.ReadToml(args.ConfigFile); err != nil {
		fmt.Println(err)
		return 1
	}

	// система логирования
	lg.Start(conf.Log)
	defer lg.Stop()

	// запуск веб сервера
	if store, err = newHTTP(conf.Server); err != nil {
		return 1

	}
	defer store.Close()

	// The correctness of the application is closed by a signal
	if args.Mode == "" || args.Mode == "run" {
		signal.Notify(chanelAppControl, os.Interrupt)
		<-chanelAppControl
	}
	return
}

// stop Stop an application
func stop() {
	chanelAppControl <- os.Interrupt
	<-chanelAppStop
}
