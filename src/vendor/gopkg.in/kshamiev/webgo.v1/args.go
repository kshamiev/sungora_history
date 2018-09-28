// Инициализация параметров командной строки
package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Структура параметров командной строки
type arguments struct {
	Mode       string
	ConfigFile string
}

// getCmdArgs Инициализация параметров командной строки
func getCmdArgs() (args *arguments, err error) {
	args = new(arguments)
	if len(os.Args) > 1 {
		args.Mode = os.Args[1]
	}
	if len(os.Args) > 2 {
		args.ConfigFile = os.Args[2]
	}
	// - проверки
	if args.Mode == `-h` || args.Mode == `-help` || args.Mode == `--help` {
		var str string
		str += "Usage of %s: %s [mode] [configFile]\n"
		str += "\t mode: Режим запуска приложения\n"
		str += "\t\t install - Установка как сервиса в ОС\n"
		str += "\t\t uninstall - Удаление сервиса из ОС\n"
		str += "\t\t restart - Перезапуск ранее установленного сервиса\n"
		str += "\t\t start - Запуск ранее установленного сервиса\n"
		str += "\t\t stop - Остановка ранее установленного сервиса\n"
		str += "\t\t run - Прямой запуск (выход по 'Ctrl+C')\n"
		str += "\t\t если параметр опущен работает в режиме run\n"
		str += "\t configFile: Полный путь до конфигурационного файла\n"
		str += "\t\t если параметр опущен конфиг берется из директории запускного файла\n"
		fmt.Fprintf(os.Stderr, str, filepath.Base(os.Args[0]), filepath.Base(os.Args[0]))
		return nil, errors.New("Help startup request")
	}
	return
}
