// core Ядро приложения. Система.

// Инициализация приложения

package core

import (
	"errors"
	"fmt"
	"lib"
	"runtime"
	"time"

	typConfig "types/config"
)

// Версия ядра
const VERSION string = `0.9.9`

// Время жизни сессиии (в минутах)
//var SessionTimeout time.Duration
func GetSessionTimeout(t time.Time) time.Duration {
	return lib.Time.Now().Sub(t.Add(time.Minute * time.Duration(Config.Auth.SessionTimeout)))
}

// Конфигурация приложения
var Config *typConfig.Configuration

// RecoverErr is the handler that turns panics into returns from the top
// level of Parse.
func RecoverErr(err *error) {
	if e := recover(); e != nil {
		switch errp := e.(type) {
		case runtime.Error:
			panic(e)
			//*err = errors.New(fmt.Sprintf("%v", e))
		case error:
			*err = errp
		default:
			//panic(e)
			*err = errors.New(fmt.Sprintf("%v", errp))
		}
	}
}
