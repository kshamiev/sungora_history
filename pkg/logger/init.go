package logger

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

// CreateLogger from config
func CreateLogger(config *Config) Logger {
	logger := initLogger(config)
	return logger
}

func initLogger(config *Config) Logger {
	logger := logrus.New()

	if config == nil {
		logger.SetOutput(os.Stdout)
		logger.SetLevel(logrus.Level(TraceLevel))
		logger.SetFormatter(&logrus.TextFormatter{})

		return logrusWrapper{logger.WithFields(map[string]interface{}{})}
	}

	switch config.Output {
	case Stdout:
		logger.SetOutput(os.Stdout)
	case Stderr:
		logger.SetOutput(os.Stderr)
	case Vacuum:
		logger.SetOutput(ioutil.Discard)
	default:
		f, err := os.OpenFile(config.Output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			logger.SetOutput(os.Stdout)
			logger.WithError(err).Debug("falling to stdout")
		} else {
			logger.SetOutput(f)
		}
	}

	switch config.Formatter {
	case JSONFormatter:
		logger.SetFormatter(&logrus.JSONFormatter{})
	default:
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	logger.SetLevel(logrus.Level(config.Level))
	addHooks(logger, config)

	if config.Title == "" {
		return logrusWrapper{logger.WithField(TitleField, config.Title)}
	}

	return logrusWrapper{logger.WithFields(map[string]interface{}{})}
}

func addHooks(logger *logrus.Logger, config *Config) {
	if config.Hooks.Sentry != nil {
		sh, err := sentryHook(config.Hooks.Sentry)
		if err == nil {
			logger.AddHook(sh)
		} else {
			logger.WithError(err).Debug("can't add hook sentry")
		}
	}

	if config.Hooks.Syslog != nil {
		sh, err := sysloggerHook(config.Hooks.Syslog)
		if err == nil {
			logger.AddHook(sh)
		} else {
			logger.WithError(err).Debug("can't add hook syslog")
		}
	}

	if config.Hooks.Logstash != nil {
		sh, err := logstashHook(config.Hooks.Logstash)
		if err == nil {
			logger.AddHook(sh)
		} else {
			logger.WithError(err).Debug("can't add hook logstash")
		}
	}

	if config.Hooks.FileName != nil {
		logger.AddHook(filenameHook(config.Hooks.FileName))
	}
}
