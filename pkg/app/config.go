package app

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// основная общая конфигурация
type ConfigApp struct {
	SessionTimeout time.Duration `yaml:"sessionTimeout"` //
	Domain         string        `yaml:"domain"`         //
	Mode           string        `yaml:"mode"`           //
	DirWork        string        `yaml:"dirWork"`        //
	DirStatic      string        `yaml:"DirStatic"`      //
	ServiceName    string        `yaml:"serviceName"`    //
	ServiceID      string        `yaml:"serviceID"`      //
	Version        string
}

// конфигурация HTTP
type ConfigServer struct {
	Proto          string        `yaml:"proto"`          // Server Proto
	Host           string        `yaml:"host"`           // Server Host
	Port           int           `yaml:"port"`           // Server Port
	ReadTimeout    time.Duration `yaml:"readTimeout"`    // Время ожидания web запроса в секундах
	WriteTimeout   time.Duration `yaml:"writeTimeout"`   // Время ожидания окончания передачи ответа в секундах
	RequestTimeout time.Duration `yaml:"requestTimeout"` // Время ожидания окончания выполнения запроса
	IdleTimeout    time.Duration `yaml:"idleTimeout"`    // Время ожидания следующего запроса
	MaxHeaderBytes int           `yaml:"maxHeaderBytes"` // Максимальный размер заголовка получаемого от браузера клиента в байтах
	Header         string        `yaml:"header"`         // Пользовательский заголовок
}

// конфигурация GRPC
type ConfigGRPC struct {
	Host string `yaml:"host"` // Server Host
	Port int    `yaml:"port"` // Server Port
}

// HeaderCheck проверка валидности значения пользовательского заголовка
func (c *ConfigServer) HeaderCheck(r *http.Request, val string) bool {
	return val != "" && r.Header.Get(c.Header) == val
}

type ConfigCors struct {
	AllowedOrigins   []string `yaml:"allowedOrigins"`
	AllowedMethods   []string `yaml:"allowedMethods"`
	AllowedHeaders   []string `yaml:"allowedHeaders"`
	ExposedHeaders   []string `yaml:"exposedHeaders"`
	AllowCredentials bool     `yaml:"allowCredentials"`
	MaxAge           int      `yaml:"maxAge"`
}

type ConfigMysql struct {
	Host     string `yaml:"host"`     // протокол, хост и порт подключения
	Port     int    `yaml:"port"`     // Порт подключения по протоколу tcp/ip (3306 по умолчанию)
	DbName   string `yaml:"dbName"`   // Имя базы данных
	Login    string `yaml:"login"`    // Логин к базе данных
	Password string `yaml:"password"` // Пароль к базе данных
	Charset  string `yaml:"charset"`  // Кодировка данных (utf-8 - по умолчанию)
}

type ConfigPostgres struct {
	Host     string `yaml:"host"`     // Хост базы данных (localhost - по умолчанию)
	Port     int    `yaml:"port"`     // Порт подключения по протоколу tcp/ip (3306 по умолчанию)
	DbName   string `yaml:"dbName"`   // Имя базы данных
	Login    string `yaml:"login"`    // Логин к базе данных
	Password string `yaml:"password"` // Пароль к базе данных
	Charset  string `yaml:"charset"`  // Кодировка данных (utf-8 - по умолчанию)
	Ssl      string `yaml:"ssl"`      // Ssl
}

// ConfigLoad загрузка конфигурации
func ConfigLoad(path string, cfg interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, cfg)
}

// ConfigSetDefault инициализация дефолтовыми значениями
func ConfigSetDefault(cfg *ConfigApp) {
	if cfg == nil {
		cfg = &ConfigApp{}
	}
	// режим работы приложения
	if cfg.Mode == "" {
		cfg.Mode = "dev"
	}
	// пути
	sep := string(os.PathSeparator)

	if cfg.DirWork == "" {
		cfg.DirWork, _ = os.Getwd()
		sl := strings.Split(cfg.DirWork, sep)

		if sl[len(sl)-1] == "bin" {
			sl = sl[:len(sl)-1]
		}

		cfg.DirWork = strings.Join(sl, sep)
	}

	cfg.DirStatic = cfg.DirWork + cfg.DirStatic

	// сессия
	if cfg.SessionTimeout == 0 {
		cfg.SessionTimeout = time.Duration(14400) * time.Second
	}
	//
	if cfg.Domain == "" {
		cfg.Domain = "localhost"
	}
}
