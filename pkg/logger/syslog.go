// +build !windows,!nacl,!plan9

package logger

import (
	"log/syslog"

	log "github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
)

type Syslog struct {
	NetworkType string `yaml:"network type" json:"network_type" toml:"network_type"`
	Host        string `yaml:"host" json:"host" toml:"host"`
	Severity    string `yaml:"severity" json:"severity" toml:"severity"`
	Facility    string `yaml:"facility" json:"facility" toml:"facility"`
	Port        string `yaml:"port" json:"port" toml:"port"`
}

func sysloggerHook(config *Syslog) (log.Hook, error) {
	var logSeverity = map[string]syslog.Priority{
		"LOG_EMERG":   syslog.LOG_EMERG,
		"LOG_ALERT":   syslog.LOG_ALERT,
		"LOG_CRIT":    syslog.LOG_CRIT,
		"LOG_ERR":     syslog.LOG_ERR,
		"LOG_WARNING": syslog.LOG_WARNING,
		"LOG_NOTICE":  syslog.LOG_NOTICE,
		"LOG_INFO":    syslog.LOG_INFO,
		"LOG_DEBUG":   syslog.LOG_DEBUG,
	}

	var logFacility = map[string]syslog.Priority{
		"LOG_KERN":     syslog.LOG_KERN,
		"LOG_USER":     syslog.LOG_USER,
		"LOG_MAIL":     syslog.LOG_MAIL,
		"LOG_DAEMON":   syslog.LOG_DAEMON,
		"LOG_AUTH":     syslog.LOG_AUTH,
		"LOG_SYSLOG":   syslog.LOG_SYSLOG,
		"LOG_LPR":      syslog.LOG_LPR,
		"LOG_NEWS":     syslog.LOG_NEWS,
		"LOG_UUCP":     syslog.LOG_UUCP,
		"LOG_CRON":     syslog.LOG_CRON,
		"LOG_AUTHPRIV": syslog.LOG_AUTHPRIV,
		"LOG_FTP":      syslog.LOG_FTP,

		"LOG_LOCAL0": syslog.LOG_LOCAL0,
		"LOG_LOCAL1": syslog.LOG_LOCAL1,
		"LOG_LOCAL2": syslog.LOG_LOCAL2,
		"LOG_LOCAL3": syslog.LOG_LOCAL3,
		"LOG_LOCAL4": syslog.LOG_LOCAL4,
		"LOG_LOCAL5": syslog.LOG_LOCAL5,
		"LOG_LOCAL6": syslog.LOG_LOCAL6,
		"LOG_LOCAL7": syslog.LOG_LOCAL7,
	}

	return logrus_syslog.NewSyslogHook(
		config.NetworkType,
		config.Host+":"+config.Port,
		logSeverity[config.Severity]|logFacility[config.Facility],
		"")
}
