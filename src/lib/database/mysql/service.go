package mysql

// Службы и сервисы по работе с БД mysql

import (
	"time"

	"lib"
)

// Контроль коннектов.
// Закрытие неиспользуемых коннектов по таймауту.
func goControlConnect() {
	go func() {
		for {
			for i := range conn {
				for key := range conn[i] {
					t := conn[i][key].time.Add(time.Second * time.Duration(cfgMysql[i].TimeOut))
					if conn[i][key].free == true && 0 < lib.Time.Now().Sub(t) {
						conn[i][key].Connect.Close()
						delete(conn[i], key)
					}
				}
			}
			time.Sleep(time.Second * 1)
		}
	}()
}
