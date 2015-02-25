// Библиотека кеширования
package cache

import (
	"reflect"
	"time"
)

// Таймауты хранения кеша. Для использования в программе
const (
	TS03 = time.Second * 3
	TS05 = time.Second * 5
	TS10 = time.Second * 10
	TS20 = time.Second * 20
	TS30 = time.Second * 30
	TM1  = time.Minute * 1
	TM3  = time.Minute * 3
	TM5  = time.Minute * 5
	TM10 = time.Minute * 10
	TM20 = time.Minute * 20
	TM30 = time.Minute * 30
	TH1  = time.Hour * 1
	TH3  = time.Hour * 3
	TH6  = time.Hour * 6
	TH12 = time.Hour * 12
	TD1  = time.Hour * 24
	TD3  = time.Hour * 24 * 3
	TD5  = time.Hour * 24 * 5
	TD7  = time.Hour * 24 * 7
	TD10 = time.Hour * 24 * 10
	TD20 = time.Hour * 24 * 20
	TD30 = time.Hour * 24 * 30
)

// Канал передачи команд
var commandControl = make(chan commandCache)

// Команда
type commandCache struct {
	action  int
	key     string
	value   interface{}
	timeout time.Duration
	result  chan<- interface{}
}

// Допустимые команды (action)
const (
	set int = iota
	get
	rem
	length
	cacheClose
)

var (
	store        = make(map[string]reflect.Value)
	storeTimeOut = make(map[string]time.Time)
	storeLink    = make(map[string][]string)
)

func init() {
	goCache()
}

// Служба кеширования
func goCache() {
	// Запуск механизма кеширования
	go func() {
		for command := range commandControl {
			switch command.action {
			case set:
				store[command.key] = reflect.ValueOf(command.value)
				if command.timeout == 0 {
					storeTimeOut[command.key] = time.Time{}
				} else {
					storeTimeOut[command.key] = time.Now().Add(command.timeout)
				}
			case get:
				if value, found := store[command.key]; found == true {
					if command.timeout > 0 {
						//storeTimeOut[command.key] = storeTimeOut[command.key].Add(command.timeout)
						storeTimeOut[command.key] = time.Now().Add(command.timeout)
					}
					command.result <- value.Interface()
				} else {
					command.result <- nil
				}
			case rem:
				for i := range storeLink[command.key] {
					delete(store, storeLink[command.key][i])
				}
				delete(storeLink, command.key)
				delete(store, command.key)
			case length:
				command.result <- len(store)
			case cacheClose:
				close(commandControl)
				store = nil
				storeTimeOut = nil
				storeLink = nil
				command.result <- true
			}
		}
	}()
	// Запуск чистки кеша
	go func() {
		for storeTimeOut != nil {
			for i := range storeTimeOut {
				if storeTimeOut[i].IsZero() != true && 0 < time.Now().Sub(storeTimeOut[i]).Nanoseconds() {
					delete(store, i)
					delete(storeTimeOut, i)
				}
			}
			time.Sleep(time.Second * 1)
		}
	}()
}

// Связывание кешируемых данных между собой по сбросу кеша
// При удалении кеша по целевому индексу будет удален кеш исходного индекса
//    + key string Индекс, кеш которого будет удален при удалении кеша целевого индекса
//    + keyTarget string Индекс целевого кеша
func SetLink(key, keyTarget string) {
	storeLink[keyTarget] = append(storeLink[keyTarget], key)
}

// Добавление в кеш (Если timeout не указан - вечный кеш)
//    + key string индекс кеша
//    + value interface{} данные сохраняемые в кеш
//    + timeout ...time.Duration время жизни кеша в секундах
func Set(key string, value interface{}, timeout ...time.Duration) {
	if len(timeout) == 0 {
		commandControl <- commandCache{action: set, key: key, value: value}
	} else {
		commandControl <- commandCache{action: set, key: key, value: value, timeout: timeout[0]}
	}
}

// Получение из кеша с проверкой на существование
// Если timeout указан то время жизни кеша переустанавливается в указанное, при усолвии что кеш существует.
//    + key string индекс кеша
//    + timeout ...time.Duration время жизни кеша в секундах
//    - value interface{} данные получаемые из кеша по индексу
func Get(key string, timeout ...time.Duration) (value interface{}) {
	reply := make(chan interface{})
	if len(timeout) == 0 {
		commandControl <- commandCache{action: get, key: key, result: reply}
	} else {
		commandControl <- commandCache{action: get, key: key, result: reply, timeout: timeout[0]}
	}
	return <-reply
}

// Удаление из кеша
//    + key string индекс кеша
func Rem(key string) {
	commandControl <- commandCache{action: rem, key: key}
}

// Количество данных в кеше
//    - int количество элементов хранимых в кеше
func Len() int {
	reply := make(chan interface{})
	commandControl <- commandCache{action: length, result: reply}
	return (<-reply).(int)
}

// Завершение работы службы кеширования
//    - bool признак успешности заверешения работы
func GoClose() bool {
	reply := make(chan interface{})
	commandControl <- commandCache{action: cacheClose, result: reply}
	return (<-reply).(bool)
}
