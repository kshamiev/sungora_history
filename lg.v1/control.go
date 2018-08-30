package lg

import (
	"os"
	"path/filepath"
)

type conf struct {
	file     bool
	filePath string
	graylog  bool
	stdout   bool
}

var config = new(conf)

type comand struct {
	action  string
	message string
}

var logCh = make(chan comand, 10000)
var logChClose = make(chan bool)

func Start() {

	if config.filePath == "" {
		if dir, err := os.Getwd(); err == nil {
			config.filePath = dir + "/logs/" + filepath.Base(os.Args[0]) + ".log"
		}
	}
	config.file = true
	config.graylog = true
	config.stdout = true

	go func() {
		for com := range logCh {
			if config.stdout == true {
				saveStdout(com)
			}
			if config.file == true {
				saveFile(com)
			}
			if config.graylog == true {
				saveGraylog(com)
			}

		}
		logChClose <- true
	}()

}

func Stop() {
	close(logCh)
	<-logChClose
}

func Info(msg string) {

	com := comand{
		action:  "Info",
		message: msg,
	}

	logCh <- com

}
