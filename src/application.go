// TODO
// Проверит итерации выполнения контроллеров мультизапросы
// Посмотреть рефлексию на предмет разницы методов CanSet IsValid
package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"

	"bitbucket.org/kardianos/service"

	"core"
	coreConfig "core/config"
	"core/controller"
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

	// Запуск в выбранном режиме
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

var chanelAppStop = make(chan os.Signal, 1)
var chanelAppControl = make(chan os.Signal, 1)

// goAppStop Stop an application
func goAppStop() {
	chanelAppControl <- os.Interrupt
	<-chanelAppStop
}

// goAppStart Launch an application
func goAppStart(args *typConfig.CmdArgs) {

	defer func() {
		chanelAppStop <- os.Interrupt
	}()

	var err error

	// Запуск и остановка службы логирования
	logs.GoStart()
	defer logs.GoClose()

	//logs.Dumper(core.Config)

	// Setting to use the maximum number of sockets and cores
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Checking the available memory
	//if err = ensuring.CheckMemory(1024 * 0.5); err != nil {
	//	logs.Fatal(900, err.Error())
	//	return
	//}
	runtime.GC()

	// Create a PID file and lock on record, control run one copy of the application
	if err = ensuring.PidFileCreate(core.Config.Main.Pid); err != nil {
		logs.Fatal(910, err.Error())
		return
	}
	defer ensuring.PidFileUnlock()

	// В режиме использования БД: проверяем, обновляем БД
	if core.Config.Main.UseDb > 0 {
		// Checking availability of databases
		if err = database.CheckConnect(); err != nil {
			logs.Fatal(920, err.Error())
			return
		}
	}

	// Инициализация системных данных
	if err = controller.Init(); err != nil {
		//logs.Fatal(930, err.Error())
		return
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

	// Этот вариант кушает ресурсы больше
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

	// logs.Warning(143)
	return
}
