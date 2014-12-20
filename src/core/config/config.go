// Базовый контроллер является встроенным для всех контроллеров приложения.
// config.go Временная Конфигурация роутинга и контроллеров
// load.go Загрузка, инициализация и проверка перед запуском приложения
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/robfig/config"

	"core"
	"lib/database/mysql"
	typConfig "types/config"
)

// Конфигурация приложения
var conf *config.Config

// Init Инициализация приложения
func Init(args *typConfig.CmdArgs) (err error) {
	// Инициализация конфигурации
	if err = initConfig(args); err != nil {
		return
	}

	// Инициализация папок приложения
	if err = initFolders(0777); err != nil {
		return
	}

	// Инициализация библиотек
	if err = initLib(); err != nil {
		return
	}

	// Прекомпиляция запросов в БД
	if err = compileQuery(); err != nil {
		return
	}

	// Of the compiler while deploy transmitted version of the application assembly,
	// if the versions are not, then manual assembly
	var p, _ = filepath.Abs(os.Args[0])
	if fi, err := os.Stat(p); err == nil {
		core.Config.Main.VersionBuild += ` ` + fi.ModTime().String()
	}
	return nil
}

// initConfig Инициализация конфигурации
func initConfig(args *typConfig.CmdArgs) (err error) {
	// Поиск конфигурационного файла приложения
	var configFile string
	if configFile = searchConfigFile(args); configFile == `` {
		return errors.New(`Конфигурационный файл не найден`)
	}

	// Чтение конфигурации
	if conf, err = config.ReadDefault(configFile); err != nil {
		err = errors.New(`Read configuration file "` + configFile + `" error: ` + err.Error())
		return
	}

	// Инициализация конфигурации
	var self = new(typConfig.Configuration)

	var objValue = reflect.ValueOf(self)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	num := objValue.NumField()
	for i := 0; i < num; i++ {
		field := objValue.Field(i)
		if false == field.CanSet() {
			continue
		}
		name := objValue.Type().Field(i).Name
		section := objValue.Type().Field(i).Tag.Get(`config`)
		if section != `-` {
			parseConfig(self, name, section)
		}
	}

	var section string

	// инициализация конфигурации подключений к БД Mysql
	self.Mysql = make(map[int8]*mysql.CfgMysql)
	for i := int8(0); i < 100; i++ {
		section = fmt.Sprintf("MYSQL%d", i)
		if conf.HasSection(section) {
			obj := new(mysql.CfgMysql)
			parseConfig(obj, ``, section)
			self.Mysql[i] = obj
		}
	}

	// инициализация конфигурации серверов
	self.Server = make(map[int8]*typConfig.Server)
	for i := int8(0); i < 100; i++ {
		section = fmt.Sprintf("SERVER%d", i)
		if conf.HasSection(section) {
			obj := new(typConfig.Server)
			parseConfig(obj, ``, section)
			self.Server[i] = obj
		}
	}

	// Дефолтовые значения

	// Назначение параметров коммандной строки в конфигурацию
	self.Main.Mode = args.Mode
	self.Main.ConfigFile = configFile

	// Рабочая папка приложения
	if self.Main.WorkDir == "" {
		self.Main.WorkDir, _ = os.Getwd()
	}
	// Проверка существования рабочей папки приложения
	if _, err = os.Stat(self.Main.WorkDir); err != nil {
		return
	}
	// Установка временной зоны
	if loc, err := time.LoadLocation(self.Main.TimeZone); err == nil {
		self.Main.TimeLocation = loc
	} else {
		self.Main.TimeLocation = time.UTC
	}
	// Расположение загружаемых бинарных данных
	if self.Main.Upload == "" {
		self.Main.Upload = self.Main.WorkDir + "/upload"
	}

	// Сертификаты SSL
	if self.Main.Keys == "" {
		self.Main.Keys = self.Main.WorkDir + "/keys"
	}
	// Сертификаты SSL
	if self.Main.Pid == "" {
		self.Main.Pid = self.Main.WorkDir + `/run/` + filepath.Base(configFile) + `.pid`
	}
	// Шаблоны данных не привязанных к URI
	if self.View.Path == "" {
		self.View.Path = self.Main.WorkDir + "/templates"
	}
	if self.View.Tpl == "" {
		self.View.Tpl = self.View.Path + "/view"
	}

	// Файл системного журнала приложения
	if self.Logs.Mode == `` {
		self.Logs.Mode = `file`
	}
	if self.Logs.Level == 0 {
		self.Logs.Level = 6
	}
	if self.Logs.File == `` {
		self.Logs.Mode = `system`
	}

	// Гость и разработчик
	self.Auth.DevUID = 1
	self.Auth.GuestUID = 2
	if self.Auth.TokenCookie == `` {
		self.Auth.TokenCookie = `GKTf8VfoJBNFjREgU557`
	}
	if self.Auth.SessionTimeout == 0 { // Время жизни сессиии (в минутах)
		self.Auth.SessionTimeout = 3600
	}

	// Интернационализация
	if self.Main.Lang == `` {
		self.Main.Lang = `ru-ru`
	}

	core.Config = self
	return
}

// parseConfig Парсинг конфигурационного файла и инциализация структуры (рекурсивно)
// Взаимозависмо - простые поля (строки, числа, булевы) инициализируются от конфига, структуры, ссылки, срезы от струтуры
func parseConfig(obj interface{}, property, section string) {
	// Инициализация
	var objValue = reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	//
	var qwerty reflect.Value
	if property != `` {
		qwerty = objValue.FieldByName(property)
		if qwerty.Kind() == reflect.Ptr {
			qwerty = qwerty.Elem()
		}
	} else {
		qwerty = objValue
	}
	// Наполнение
	num := qwerty.NumField()
	for i := 0; i < num; i++ {
		field := qwerty.Field(i)
		if false == field.CanSet() {
			continue
		}
		option := qwerty.Type().Field(i).Name
		switch field.Type().Kind() {
		case reflect.Bool:
			val, _ := conf.Bool(section, option)
			field.SetBool(val)
		case reflect.Float64, reflect.Float32:
			val, _ := conf.Float(section, option)
			if 0 < val {
				field.SetFloat(val)
			}
		case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
			val, _ := conf.Int(section, option)
			if 0 < val {
				field.SetInt(int64(val))
			}
		case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
			val, _ := conf.Int(section, option)
			if 0 < val {
				field.SetUint(uint64(val))
			}
		case reflect.String:
			val, _ := conf.String(section, option)
			if "" != val {
				field.SetString(val)
			}
		}
	}
	reflect.ValueOf(obj).Elem().Set(objValue)
}

// searchConfigFile Писк конфигурационного файла
func searchConfigFile(args *typConfig.CmdArgs) (configFile string) {
	if args.ConfigFile != "" {
		if _, err := os.Stat(args.ConfigFile); err != nil {
			return
		} else {
			configFile = args.ConfigFile
		}
	}
	if configFile == "" {
		folders := searchConfigPaths()
		for i := range folders {
			if _, err := os.Stat(folders[i]); err == nil {
				configFile = folders[i]
				break
			}
		}
	}
	return
}

// GetCmdArgs Получение и разбор параметров коммандной строки
// Инициализация параметров командной строки
func GetCmdArgs() (args *typConfig.CmdArgs, err error) {
	args = new(typConfig.CmdArgs)
	if len(os.Args) > 1 {
		args.Mode = os.Args[1]
	}
	if len(os.Args) > 2 {
		args.ConfigFile = os.Args[2]
	}
	// - проверки
	if args.Mode == `-h` || args.Mode == `-help` || args.Mode == `--help` {
		var str string
		str += "Usage of %s: %s [-mode] [-cfgFile]\n"
		str += "  -mode=\"\": Режим запуска приложения:\n"
		str += "\tinstall - Установка сервиса в операционной системе, используется в windows, mac os x\n"
		str += "\tremove - Удаление сервиса из операционной системы\n"
		str += "\tstart - Запуск ранее установленного сервиса\n"
		str += "\tstop - Остановка ранее установленного сервиса\n"
		str += "\tupdate - Обновление приложенияv\n"
		str += "\trun - Запуск в режиме разработки выход по 'Ctrl+C'\n"
		str += "\ttest - Тестирование работоспособности и выход\n"
		str += "  -cfgFile=\"\": Полный путь до конфигурационного файла:\n"
		fmt.Fprintf(os.Stderr, str, filepath.Base(os.Args[0]), filepath.Base(os.Args[0]))
		return nil, errors.New("Help startup request")
	}
	return
}
