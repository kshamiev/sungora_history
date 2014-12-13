// запуск теста
// SET GOPATH=C:\Work\sungora
// go test -v core/server | go test -v -bench . core/server
package server_test

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"

	"core"
	_ `core/base/setup`
	coreConfig "core/config"
	_ `core/info/setup`
	"core/server"
	"lib"
	"lib/database"
	"lib/ensuring"
	"lib/logs"
)

func TestServerConfig(t *testing.T) {
	var err error

	// Setting to use the maximum number of sockets and cores
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Checking the available memory
	if err = ensuring.CheckMemory(1024 * 0.5); err != nil {
		t.Fatal(10, err.Error())
		return
	}
	runtime.GC()

	// Конфигурация приложения
	if err = coreConfig.Init(`C:/Work/zegota/bin/application.conf`); err != nil {
		t.Fatal(20, err.Error())
		return
	}

	// Запуск и остановка службы логирования
	logs.GoStart()
	defer logs.GoClose()

	// Create a PID file and lock on record, control run one copy of the application
	if err = ensuring.PidFileCreate(core.Config.Main.Pid); err != nil {
		t.Fatal(30, err.Error())
		return
	}
	defer ensuring.PidFileUnlock()

	// В режиме использования БД: проверяем, обновляем БД
	if core.Config.Main.TypeDb > 0 && core.Config.Main.DbCheck == true {
		// Checking availability of databases
		if err = database.CheckConnect(); err != nil {
			t.Fatal(40, err.Error())
			return
		}
	}

	// Инициализация системных данных
	if err = coreConfig.Load(); err != nil {
		t.Fatal(50, err.Error())
		return
	}

	// Running a web servers
	for i := range core.Config.Server {
		server.GoStart(fmt.Sprintf(`server%d`, i), core.Config.Server[i])
	}

	// Проверка работы серверов (запрашиваем главную страницу)
	for i, elm := range core.Config.Server {
		if data, err := lib.RW.RequestJson(`GET`, `http://localhost:`+strconv.Itoa(int(elm.Port))+`/`); err != nil {
			t.Log(err.Error())
		} else {
			t.Logf("done: %d, len(%d)\n", i, len(data))
		}
	}

	// The correctness of the application is closed by a signal
	for i := range core.Config.Server {
		server.GoStop(fmt.Sprintf(`server%d`, i))
	}

	return
}
