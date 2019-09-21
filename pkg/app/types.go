package app

import "time"

// основная общая конфигурация
type ConfigApp struct {
	SessionTimeout time.Duration `yaml:"sessionTimeout"` //
	Domain         string        `yaml:"domain"`         //
	Mode           string        `yaml:"mode"`           //
	DirWork        string        `yaml:"dirWork"`        //
	ServiceName    string        `yaml:"serviceName"`    //
	Version        string
}

// конфигурация
type ConfigServer struct {
	Proto          string        `yaml:"proto"`          // Server Proto
	Host           string        `yaml:"host"`           // Server Host
	Port           int           `yaml:"port"`           // Server Port
	ReadTimeout    time.Duration `yaml:"readTimeout"`    // Время ожидания web запроса в секундах
	WriteTimeout   time.Duration `yaml:"writeTimeout"`   // Время ожидания окончания передачи ответа в секундах
	RequestTimeout time.Duration `yaml:"requestTimeout"` // Время ожидания окончания выполнения запроса
	IdleTimeout    time.Duration `yaml:"idleTimeout"`    // Время ожидания следующего запроса
	MaxHeaderBytes int           `yaml:"maxHeaderBytes"` // Максимальный размер заголовка получаемого от браузера клиента в байтах
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
