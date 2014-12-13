// logs Таймер для замеров производительности
package logs

/*
import (
	"fmt"
	"time"

	"lib"
)

var timer = make(map[string]map[string]time.Time)
var timerLevel = make(map[string]int)
var timerLevelCounter int

// TimerStart старт таймера
func TimerStart(indexFormat string, params ...interface{}) {
	var message = fmt.Sprintf(indexFormat, params...)
	timer[message] = map[string]time.Time{
		`start`: lib.Time.Now(),
	}
	timerLevel[message] = timerLevelCounter
	timerLevelCounter++
}

// TimerStop стоп таймера
func TimerStop(indexFormat string, params ...interface{}) {
	var message = fmt.Sprintf(indexFormat, params...)
	timer[message][`stop`] = lib.Time.Now()
	timerLevelCounter--
}

type Timer struct {
	Message string
	Seconds float64
	Level   int
}

// TimerGet получение таймера
func TimerGet(indexFormat string, params ...interface{}) (t Timer) {
	t.Message = fmt.Sprintf(indexFormat, params...)
	if _, ok := timer[t.Message]; ok == true && len(timer[t.Message]) == 2 {
		t.Seconds = timer[t.Message][`stop`].Sub(timer[t.Message][`start`]).Seconds()
	} else {
		t.Seconds = time.Duration(-1).Seconds()
	}
	return
}

// TimerGetAll получение всех таймеров
func TimerGetAll() (timers map[string]Timer) {
	timers = make(map[string]Timer)
	for message := range timer {
		if len(timer[message]) == 2 {
			timers[message] = Timer{
				Message: message,
				Seconds: timer[message][`stop`].Sub(timer[message][`start`]).Seconds(),
				Level:   timerLevel[message],
			}
		}
	}
	return
}

func TimerGetAllString() (str string) {
	str = "\r\n"
	for _, t := range TimerGetAll() {
		for j := 0; j < t.Level; j++ {
			str += fmt.Sprintf("\t")
		}
		str += fmt.Sprintf("(%g) %s\r\n", t.Seconds, t.Message)
	}
	return
}
*/
