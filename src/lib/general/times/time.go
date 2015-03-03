// Библиотека для работы с датой и временем
package times

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	FORMAT string = "01.02.2006 15:04:05.000000000"
)

// Тип для агрегирования методов работы с датой и временем
type Time struct {
	Location *time.Location
}

// Конструктор
//	+ timeZone string имя временной зоны
//	- *Time объект агрегатор функций-методов
func NewTime(timeZone string) *Time {
	var self = new(Time)
	if loc, err := time.LoadLocation(timeZone); err == nil {
		self.Location = loc
	} else {
		self.Location = time.UTC
	}
	return self
}

// Получение текущей даты и времени
//	- time.Time Дата и время в определенной конфигом временной зоне
func (self *Time) Now() time.Time {
	return time.Now().In(self.Location)
}

// Формирование даты и времени с наносекундами для записи в лог файлы
//	- string ГГГГ.ММ.ДД ЧЧ:ММ:СС:МЛМКНС
func (self *Time) LabelFull(t ...time.Time) string {
	if len(t) == 0 {
		t = append(t, self.Now())
	}
	return fmt.Sprintf("%04d.%02d.%02d %02d:%02d:%02d:%09d",
		t[0].Year(), t[0].Month(), t[0].Day(), t[0].Hour(), t[0].Minute(), t[0].Second(), t[0].Nanosecond())
}

// Формирование даты и времени
//	- string ГГГГ.ММ.ДД ЧЧ:ММ:СС
func (self *Time) LabelTime(t ...time.Time) string {
	if len(t) == 0 {
		t = append(t, self.Now())
	}
	return fmt.Sprintf("%04d.%02d.%02d %02d:%02d:%02d",
		t[0].Year(), t[0].Month(), t[0].Day(), t[0].Hour(), t[0].Minute(), t[0].Second())
}

// Формирование даты
//	- string ГГГГ.ММ.ДД
func (self *Time) LabelDate(t ...time.Time) string {
	if len(t) == 0 {
		t = append(t, self.Now())
	}
	return fmt.Sprintf("%04d.%02d.%02d", t[0].Year(), t[0].Month(), t[0].Day())
}

// Перевод времени из человеческого формата в time.Duration (07:23:12 -> time.Duration)
//	+ duration string 07:23:12
//	- time.Duration
//	- bool флаг успешности операции
func (self *Time) ParseDuration(duration string) (t time.Duration, flag bool) {
	list := strings.Split(duration, ":")
	if len(list) != 3 {
		return t, false
	}
	h, _ := strconv.ParseInt(list[0], 10, 8)
	m, _ := strconv.ParseInt(list[1], 10, 8)
	s, _ := strconv.ParseInt(list[2], 10, 8)
	t = time.Hour * time.Duration(h)
	t = t + time.Minute*time.Duration(m)
	t = t + time.Second*time.Duration(s)
	return t, true
}

// Перевод даты и времени из человеческого формата в time.Time
//	+ datetime string y[-./]m[-./]d h:m:s || d[-./]m[-./]y h:m:s
//	- time.Time
//	- bool флаг успешности операции
func (self *Time) Parse(datetime string) (t time.Time, flag bool) {
	datetime = strings.Trim(datetime, " ")
	listRoot := strings.Split(datetime, " ")
	if len(listRoot) != 2 {
		datetime += " 00:00:00"
		listRoot = strings.Split(datetime, " ")
	}
	var list []string
	for _, sep := range []string{".", "-", "/"} {
		list = strings.Split(listRoot[0], sep)
		if len(list) == 3 {
			break
		}
	}
	if len(list) != 3 {
		return t, false
	}
	if len(list[0]) == 4 {
		list[0], list[1], list[2] = list[1], list[2], list[0]
	} else {
		list[0], list[1], list[2] = list[1], list[0], list[2]
	}
	datetime = strings.Join(list, ".") + " " + listRoot[1]
	t, _ = time.ParseInLocation(FORMAT[:19], datetime, self.Location)
	if t.Year() == 1 {
		return t, false
	}
	return t, true
}

////
