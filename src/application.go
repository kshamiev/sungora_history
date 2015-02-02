// Точка входа в программу, запуск на выполнение.
//
// Посмотреть использования стороннего решшеня сервиса для записей логов в системный журнал
// todo проверить загрузку бинарных файлов
package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"

	"bitbucket.org/kardianos/service"

	"core"
	coreConfig "core/config"
	"core/server"
	"lib/database"
	"lib/ensuring"
	"lib/logs"
	typConfig "types/config"
)

const (
	APP_NAME         string = "Sungora"
	APP_DISPLAY_NAME string = "Sungora CMF"
	APP_DESCRIPTION  string = "Description Sungora CMF"
)

var log service.Logger

func main() {

	// Входные данные командной строки
	var args, err = coreConfig.GetCmdArgs()
	if err != nil {
		return
	}

	// Глобальная конфигурация приложения
	if err = coreConfig.Init(args); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Инициализация сервиса
	var s service.Service
	s, err = service.NewService(APP_NAME, APP_DISPLAY_NAME, APP_DESCRIPTION)
	if err != nil {
		fmt.Printf("%s unable to start: %s", APP_DISPLAY_NAME, err)
		return
	}
	log = s

	// Запуск приложения в выбранном режиме
	switch args.Mode {
	case "install":
		err = s.Install()
		if err != nil {
			fmt.Printf("Failed to install: %s\n", err)
			return
		}
		fmt.Printf("Service \"%s\" installed.\n", APP_DISPLAY_NAME)
	case "remove":
		err = s.Remove()
		if err != nil {
			fmt.Printf("Failed to remove: %s\n", err)
			return
		}
		fmt.Printf("Service \"%s\" removed.\n", APP_DISPLAY_NAME)
	case "start":
		err = s.Start()
		if err != nil {
			fmt.Printf("Failed to start: %s\n", err)
			return
		}
		fmt.Printf("Service \"%s\" started.\n", APP_DISPLAY_NAME)
	case "stop":
		err = s.Stop()
		if err != nil {
			fmt.Printf("Failed to stop: %s\n", err)
			return
		}
		fmt.Printf("Service \"%s\" stopped.\n", APP_DISPLAY_NAME)
	case "run", "test":
		goAppStart(args)
	default:
		err = s.Run(func() error {
			// start
			go goAppStart(args)
			return nil
		}, func() error {
			// stop
			goAppStop()
			return nil
		})
		if err != nil {
			s.Error(err.Error())
		}
	}
}

// Каналы управления запуском и остановом приложения
var (
	chanelAppStop    = make(chan os.Signal, 1)
	chanelAppControl = make(chan os.Signal, 1)
)

// goAppStop Stop an application
func goAppStop() {
	chanelAppControl <- os.Interrupt
	<-chanelAppStop
}

// goAppStart Launch an application
func goAppStart(args *typConfig.CmdArgs) (err error, code int) {
	defer func() {
		if code != 910 {
			ensuring.PidFileUnlock()
		}
		chanelAppStop <- os.Interrupt
	}()

	// Запуск и остановка службы логирования
	logs.GoStart(log)
	defer logs.GoClose()

	// Create a PID file and lock on record, control run one copy of the application
	if err = ensuring.PidFileCreate(core.Config.Main.Pid); err != nil {
		logs.Base.Fatal(910, err.Error())
		return err, 910
	}

	// Setting to use the maximum number of sockets and cores
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Checking the available memory (in mb)
	if err = ensuring.CheckMemory(core.Config.Main.Memory); err != nil {
		logs.Base.Fatal(900, err.Error())
		return err, 900
	}
	runtime.GC()

	// В режиме использования БД: проверяем, обновляем БД
	if core.Config.Main.UseDb > 0 {
		// Checking availability of databases
		if err = database.CheckConnect(); err != nil {
			logs.Base.Fatal(920, err.Error())
			return err, 920
		}
	}

	// Инициализация системных данных
	if err = coreConfig.App(); err != nil {
		logs.Base.Fatal(930, err.Error())
		return err, 930
	}

	// Запуск и остановка служб модулей приложения
	// Destructor for the modules, daemons and etc.
	goServiceStart()
	defer goServiceStop()

	// Running a web servers
	for i := range core.Config.Server {
		server.GoStart(fmt.Sprintf(`server%d`, i), core.Config.Server[i])
	}

	// The correctness of the application is closed by a signal
	if args.Mode == `` || args.Mode == `run` {
		signal.Notify(chanelAppControl, os.Interrupt)
		<-chanelAppControl
	}

	for i := range core.Config.Server {
		server.GoStop(fmt.Sprintf(`server%d`, i))
	}

	// The correctness of the application is closed by a signal
	//var appExit bool
	//signal.Notify(chanelServerExit, os.Interrupt)
	//for appExit == false {
	//	select {
	//	case <-chanelServerExit:
	//		for i := range core.Config.Server {
	//			server.GoStop(fmt.Sprintf(`server%d`, i))
	//		}
	//		appExit = true
	//	default:
	//		time.Sleep(time.Second * 1)
	//	}
	//}
	return
}
