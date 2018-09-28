// Установка и запуск приложения как сервис
package app

import (
	"fmt"
	"os"

	"github.com/kardianos/service"
)

type program struct{}

// Start should not block. Do the actual work async.
func (p *program) Start(s service.Service) error {
	go Start()
	return nil
}

// Stop should not block. Return with a few seconds.
func (p *program) Stop(s service.Service) error {
	stop()
	return nil
}

// Service Установка и запуск приложения как сервис
func Service(svcConfig *service.Config) (code int) {
	var (
		err    error
		args   *arguments
		s      service.Service
		logger service.Logger
	)

	// входные данные командной строки
	args, err = getCmdArgs()
	if err != nil {
		return
	}

	prg := new(program)
	s, err = service.New(prg, svcConfig)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	if args.Mode != "" && args.Mode != "run" {
		err = service.Control(s, os.Args[1])
		if err != nil {
			fmt.Println(err)
			code = 1
		}
		return
	}
	logger, err = s.Logger(nil)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	err = s.Run()
	if err != nil {
		code = 1
		logger.Error(err)
	}
	return
}
