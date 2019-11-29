package logger

import (
	"net"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

// Logstash hook configuration
type Logstash struct {
	NetworkType string `yaml:"network type" json:"network_type" toml:"network_type"`
	Host        string `yaml:"host" json:"host" toml:"host"`
	Port        string `yaml:"port" json:"port" toml:"port"`
}

func logstashHook(config *Logstash) (*logrustash.Hook, error) {
	conn, err := net.Dial(config.NetworkType, config.Host+":"+config.Port)
	if err != nil {
		return nil, err
	}
	return logrustash.NewHookWithFieldsAndConnAndPrefix(conn, "", logrus.Fields{"type": "logstash"}, "")
}
