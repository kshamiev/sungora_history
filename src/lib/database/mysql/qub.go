package mysql

// Конструктор запрсов к БД

import (
	"strconv"
	"strings"

	"lib/database/face"
	"types"
)

// Структура конструктора
type Qub struct {
	property []string
	from     string
	where    []string
	group    string
	having   string
	order    string
	limit    string
}

// Создание объекта конструктора запросов
//    - *Qub конструктор запросов
func NewQub() face.QubFace {
	return new(Qub)
}

// Выборка
//    + property Строка полей таблицы через запятую получаемых из БД
//    - *Qub конструктор запросов
func (self *Qub) Select(property string) face.QubFace {
	self.property = append(self.property, property)
	return self
}

func (self *Qub) SelectScenario(source, scenario string) face.QubFace {
	sc, err := types.GetScenario(source, scenario)
	if err != nil {
		return self
	}
	for i := range sc.Property {
		if sc.Property[i].AliasDb != `-` {
			self.property = append(self.property, sc.Property[i].AliasDb)
		}
	}
	return self
}

func (self *Qub) From(from string) face.QubFace {
	self.from = from
	return self
}

func (self *Qub) Where(where string) face.QubFace {
	self.where = append(self.where, where)
	return self
}

func (self *Qub) Group(group string) face.QubFace {
	self.group = group
	return self
}

func (self *Qub) Having(having string) face.QubFace {
	self.having = having
	return self
}

func (self *Qub) Order(order string) face.QubFace {
	self.order = order
	return self
}

func (self *Qub) Limit(start, step int) face.QubFace {
	self.limit = strconv.Itoa(start) + `, ` + strconv.Itoa(step)
	return self
}

func (self *Qub) Get() (query string) {
	query += "SELECT\n\t" + strings.Join(self.property, `, `) + "\n"
	query += "FROM " + self.from + "\n"
	if len(self.where) > 0 {
		query += "WHERE 1\n\t" + strings.Join(self.where, "\n\t") + "\n"
	}
	if self.group != `` {
		query += `GROUP BY ` + self.group + "\n"
	}
	if self.having != `` {
		query += `HAVING ` + self.having + "\n"
	}
	if self.order != `` {
		query += `ORDER BY ` + self.order + "\n"
	}
	if self.limit != `` {
		query += `LIMIT ` + self.limit + "\n"
	}
	query = strings.TrimSpace(query)
	return
}
