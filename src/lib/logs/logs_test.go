// запуск теста
// SET GOPATH=C:\Work\projectName
// go test -v lib/logs | go test -v -bench . lib/logs
package logs_test

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"testing"
	"time"

	"lib"
	"lib/logs"
)

func Testlogs(t *testing.T) {
	var cfglogs = new(logs.Cfglogs)
	cfglogs.Debug = true
	cfglogs.DebugDetail = 0
	cfglogs.Level = 6
	cfglogs.Mode = `file`
	cfglogs.File, _ = os.Getwd()
	cfglogs.File = filepath.Dir(cfglogs.File)
	cfglogs.File = filepath.Dir(cfglogs.File)
	cfglogs.File = filepath.Dir(cfglogs.File) + `/log`
	os.MkdirAll(cfglogs.File, 0777)
	cfglogs.File += `/logs.log`
	logs.Init(cfglogs)
	logs.GoStart()
	t.Logf("путь до файл лога %s", cfglogs.File)

	logs.Info(0, `logsInfo`)
	logs.Notice(0, `logsNotice`)
	logs.Warning(0, `logsWarning`)
	logs.Error(0, `logsError`)
	logs.Critical(0, `logsCritical`)
	logs.Fatal(0, `logsFatal`)

	var flag = logs.GoClose()
	if flag == false {
		t.Fatal(`Ошибка остановки службы логирования`)
	}
	fmt.Println()
}

//func TestTimer(t *testing.T) {
//	logs.TimerStart(`Millisecond %d`, 543)
//	time.Sleep(time.Millisecond * 543)
//	logs.TimerStop(`Millisecond %d`, 543)

//	logs.TimerStart(`Microsecond %d`, 345)
//	time.Sleep(time.Microsecond * 345)
//	logs.TimerStop(`Microsecond %d`, 345)

//	logs.TimerStart(`Nanosecond %d`, 785)
//	time.Sleep(time.Nanosecond * 785)
//	logs.TimerStop(`Nanosecond %d`, 785)

//	var result = logs.TimerGetAllString()
//	t.Log(result)
//}

func BenchmarkBlank(b *testing.B) {
	b.StopTimer()
	var dataSlice []*Users
	for i := 0; i < 1000000; i++ {
		var u = &Users{
			Id:    i,
			Login: `login` + strconv.Itoa(int(i)),
		}
		dataSlice = append(dataSlice, u)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			if dataSlice[i].Login == `tttt` && dataSlice[i].Id > 0 {
				i++
				fmt.Println()
			}
		}
		//time.Sleep(time.Microsecond * 1000)
	}
}

func BenchmarkBlank2(b *testing.B) {
	b.StopTimer()
	var dataSlice []*Users
	for i := 0; i < 100; i++ {
		var u = &Users{
			Id:    i,
			Login: `login` + strconv.Itoa(int(i)),
		}
		dataSlice = append(dataSlice, u)
	}
	//fmt.Println(dataSlice[457854].Login)
	fn := func(n int) bool {
		return dataSlice[n].Login >= `Login80`
	}
	lib.Slice.SortingSliceAsc(dataSlice, `Login`)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		j := sort.Search(len(dataSlice), fn)
		if j < len(dataSlice) && dataSlice[j].Login == `Login80` {
			//b.Log("FOUND", dataSlice[i].Login)
			fmt.Println(`FOUND Login457854`)
		}
		//time.Sleep(time.Microsecond * 1000)
	}
}

type Users struct {
	Id           int       // Id
	Users_Id     uint64    // Пользователь
	Login        string    `json:"login"`    // Логин пользователя
	Password     string    `json:"password"` // Пароль пользователя (SHA256)
	PasswordR    string    `db:"-"`          // Пароль пользователя (SHA256) (повторно)
	Email        string    `json:"email"`    // Email
	LastName     string    // Фамилия
	Name         string    // Имя
	MiddleName   string    // Отчество
	IsAccess     bool      // Доступ разрешен
	IsCondition  bool      // Условия пользователя
	IsActivated  bool      // Активированный Email
	DateOnline   time.Time // Дата последнего посещения
	Date         time.Time // Дата регистрации
	Del          bool      // Запись удалена
	Hash         string    // Контрольная сумма для синхронизации (SHA256)
	Token        string    // Кука активации и идентификации
	CaptchaValue string    `db:"-" json:"captchaValue"`
	CaptchaHash  string    `db:"-" json:"captchaHash"`
	Language     string    `db:"-" json:"language"`
}
