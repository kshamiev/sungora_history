/*
Реализовать:
Логирование в
файл
консоль
graylog
Возможность настройки куда логировать
Возможность переопределения реализации
 */
package lg

import (
	"errors"
	"runtime"
	"strings"
)

type callerInfo struct {
	FileName      string   // Имя исходного файла
	LineNumber    int      // Номер строки
	FuncName      string   // Название функции
	Version       string   // Текущая версия golang
	Gorutines     int      // Количество горутин работающих в настоящий момент
	GorutinesInfo []string // Информация по каждой горутине
}

// Получение информаци о вызвавшем лог
func getCallerInfo(level int) (*callerInfo, error) {
	var err error
	var ret *callerInfo = new(callerInfo)
	var pc uintptr
	var file string
	var line int
	var ok bool

	ret.Gorutines = runtime.NumGoroutine()
	ret.Version = runtime.Version()

	pc, file, line, ok = runtime.Caller(level)
	if ok == true {
		var fn *runtime.Func

		ret.LineNumber = line
		ret.FileName = file
		fn = runtime.FuncForPC(pc)
		if fn != nil {
			ret.FuncName = fn.Name()
		}

		// Информация о состоянии горутин
		var buf []byte = make([]byte, 1<<16)
		var i int = runtime.Stack(buf, true)
		var info string = string(buf[:i])
		var tmp []string

		tmp = strings.Split(info, "\n")
		i = -1
		for _, str := range tmp {
			if strings.Index(str, "goroutine") == 0 {
				i++
				ret.GorutinesInfo = append(ret.GorutinesInfo, "")
			}
			if i >= 0 && str != "" {
				if ret.GorutinesInfo[i] != "" {
					ret.GorutinesInfo[i] += "\n"
				}
				ret.GorutinesInfo[i] += str
			}
		}
	} else {
		err = errors.New("Не удалось получить информацию о вызвавшей лог функции")
	}
	return ret, err
}
