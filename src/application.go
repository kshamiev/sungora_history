// Точка входа в программу, запуск на выполнение.
//
// Посмотреть использования стороннего решшеня сервиса для записей логов в системный журнал
// todo проверить загрузку бинарных файлов
// TODO Язык определяется по ссылке (убрать определение от пользователя)
// TODO подключение модулей убрать из корневого приложения (modules.go)
package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/kardianos/service"

	"app/setup"
	"core"
	coreConfig "core/config"
	"core/server"
	"lib/database"
	"lib/ensuring"
	"lib/logs"
	typConfig "types/config"
)

var logger service.Logger

// Service setup.
//   Define service config.
//   Create the service.
//   Setup the logger.
//   Handle service controls (optional).
//   Run the service.
func main() {

	var err error

	// Входные параметры командной строки
	p1 := flag.String("service", "", "Control the system service.")
	p2 := flag.String("config", "", "Path config file.")
	flag.Parse()
	var cmdArgs = new(typConfig.CmdArgs)
	cmdArgs.Service = (*p1)
	cmdArgs.ConfigFile = (*p2)

	// Глобальная конфигурация приложения
	if err = coreConfig.Init(cmdArgs); err != nil {
		fmt.Println(err)
		return
	}

	// Инициализация сервиса
	svcConfig := &service.Config{
		Name:        core.Config.Main.AppName,
		DisplayName: core.Config.Main.AppDisplay,
		Description: core.Config.Main.AppDescription,
	}

	prg := new(app)
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Запуск приложения в выбранном режиме
	if cmdArgs.Service != `` {
		err := service.Control(s, cmdArgs.Service)
		if err != nil {
			fmt.Printf("Valid actions: %q\n", service.ControlAction)
			fmt.Println(err)
		}
		return
	}
	err = s.Run()
	if err != nil || true {
		logger.Error(err)
	}
}

// Program structures.
//  Define Start and Stop methods.
type app struct {
	exit chan struct{}
}

// Запуск работы приложения
func (self *app) Start(s service.Service) error {
	self.exit = make(chan struct{})
	go self.run()
	return nil
}

// Приложение
func (self *app) run() (err error) {
	// Запуск и остановка службы логирования
	logs.GoStart(logger)
	logger.Info("Start application: " + core.Config.Main.AppName)
	defer logs.GoClose()
	defer logger.Info("Stop application: " + core.Config.Main.AppName)

	// Create a PID file and lock on record, control run one copy of the application
	if err = ensuring.PidFileCreate(core.Config.Main.Pid); err != nil {
		logs.Base.Fatal(9100, err)
		return
	}
	ensuring.PidFileUnlock()

	// Setting to use the maximum number of sockets and cores
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Checking the available memory (in mb)
	if err = ensuring.CheckMemory(core.Config.Main.Memory); err != nil {
		logs.Base.Fatal(9000, err)
		return
	}
	runtime.GC()

	// В режиме использования БД: проверяем, обновляем БД
	if core.Config.Main.UseDb > 0 {
		// Checking availability of databases
		if err = database.CheckConnect(); err != nil {
			return
		}
	}

	// Инициализация данных приложения
	if err = coreConfig.App(); err != nil {
		logs.Base.Fatal(9300, err)
		return
	}

	// Запуск и остановка служб модулей приложения
	// Destructor for the modules, daemons and etc.
	setup.GoServiceStart()
	defer setup.GoServiceStop()

	// Running a web servers
	for i := range core.Config.Server {
		server.GoStart(fmt.Sprintf(`server%d`, i), core.Config.Server[i])
	}

	// Контроль завершения работы приложения
	var flag bool
	for {
		select {
		case <-self.exit:
			flag = true
		default:
			time.Sleep(time.Second * 1)
			logger.Infof("Still running at %v...", time.Now().String())
		}
		if flag == true {
			break
		}
	}

	// Stopping a web servers
	for i := range core.Config.Server {
		server.GoStop(fmt.Sprintf(`server%d`, i))
	}

	return
}

// Завершение работы приложения
func (self *app) Stop(s service.Service) error {
	close(self.exit)
	return nil
}
