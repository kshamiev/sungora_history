package mysql

import (
	"time"

	"lib"
)

// goMysql отложенное выполнение запросов
func goMysql() {
	// контроль коннектов
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
