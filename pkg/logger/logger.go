// Package logger provide safe logger wrapper
package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"

	"github.com/sirupsen/logrus"
)

// Fields is a log fields type
type Fields map[string]interface{}

const ErrorField = "error"
const TitleField = "title"
const Stdout = "stdout"
const Stderr = "stderr"
const Vacuum = "vacuum"
const JSONFormatter = "json"

type Logger interface {
	// Level calls
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})

	// Levelf calls
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	// StdLogger calls
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})

	// Levelln calls
	Traceln(...interface{})
	Debugln(...interface{})
	Infoln(...interface{})
	Warningln(...interface{})
	Errorln(...interface{})

	// Native log with level
	Log(level Level, args ...interface{})
	Logf(level Level, format string, args ...interface{})
	Logln(level Level, args ...interface{})

	// Writer
	Writer() *io.PipeWriter

	// Context
	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
	WithError(err error) Logger
}

// Config is a configuration for logger
type Config struct {
	Title     string `yaml:"title" json:"title" toml:"title"`             // no field "title" if value is empty
	Output    string `yaml:"output" json:"output" toml:"output"`          // enum (stdout|stderr|vacuum|path/to/file)
	Formatter string `yaml:"formatter" json:"formatter" toml:"formatter"` // enum (json|text)
	Level     Level  `yaml:"level" json:"level" toml:"level"`             // enum (panic|fatal|error|warning|info|debug|trace)
	Hooks     Hooks  `yaml:"hooks" json:"hooks" toml:"hooks"`
}

// Hooks is a set of hooks for logger
type Hooks struct {
	Sentry   *Sentry   `yaml:"sentry" json:"sentry" toml:"sentry"`
	Syslog   *Syslog   `yaml:"syslog" json:"syslog" toml:"syslog"`
	Logstash *Logstash `yaml:"logstash" json:"logstash" toml:"logstash"`
	FileName *FileName `yaml:"filename" json:"filename" toml:"filename"`
}

type ctxlog struct{}

// WithLogger put logger to context
func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, ctxlog{}, logger)
}

// GetLogger get logger from context
func GetLogger(ctx context.Context) Logger {
	l, ok := ctx.Value(ctxlog{}).(Logger)
	if !ok {
		l = initLogger(&Config{Output: "stdout", Level: InfoLevel})
	}
	return l
}

type Level logrus.Level

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	// config: panic
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	// config: fatal
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	// config: error
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	// config: warning
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	// config: info
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	// config: debug
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	// config: trace
	TraceLevel
)

func (level Level) MarshalText() ([]byte, error) {
	var levelArray = []string{
		"panic",
		"fatal",
		"error",
		"warning",
		"info",
		"debug",
		"trace",
	}
	if int(level) < len(levelArray) {
		return []byte(levelArray[level]), nil
	}
	return nil, fmt.Errorf("not a valid level %d", level)
}
func (level *Level) UnmarshalText(text []byte) error {
	var levelMap = map[string]Level{
		"panic":   PanicLevel,
		"fatal":   FatalLevel,
		"error":   ErrorLevel,
		"warning": WarnLevel,
		"info":    InfoLevel,
		"debug":   DebugLevel,
		"trace":   TraceLevel,
	}
	l, ok := levelMap[string(text)]
	if !ok {
		return fmt.Errorf("not a valid level %d", level)
	}

	*level = l

	return nil
}

func (level *Level) UnmarshalYAML(value *yaml.Node) error {
	return level.UnmarshalText([]byte(value.Value))
}

func (level Level) MarshalYAML() (interface{}, error) {
	data, err := level.MarshalText()
	return string(data), err
}
func (level *Level) UnmarshalJSON(value []byte) error {
	var temp string
	err := json.Unmarshal(value, &temp)
	if err != nil {
		return err
	}
	return level.UnmarshalText([]byte(temp))
}

func (level Level) MarshalJSON() ([]byte, error) {
	data, err := logrus.Level(level).MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(data))
}
