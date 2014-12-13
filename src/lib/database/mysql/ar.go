package mysql

import (
	"lib/database/face"
	"strconv"
	"strings"
	"types"
)

type Ar struct {
	property []string
	from     string
	where    []string
	group    string
	having   string
	order    string
	limit    string
}

func NewAr() face.ArFace {
	var self = new(Ar)
	return self

}

func (self *Ar) Select(property string) face.ArFace {
	self.property = append(self.property, property)
	return self
}

func (self *Ar) SelectScenario(source, scenario string) face.ArFace {
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

func (self *Ar) From(from string) face.ArFace {
	self.from = from
	return self
}

func (self *Ar) Where(where string) face.ArFace {
	self.where = append(self.where, where)
	return self
}

func (self *Ar) Group(group string) face.ArFace {
	self.group = group
	return self
}

func (self *Ar) Having(having string) face.ArFace {
	self.having = having
	return self
}

func (self *Ar) Order(order string) face.ArFace {
	self.order = order
	return self
}

func (self *Ar) Limit(start, step int) face.ArFace {
	self.limit = strconv.Itoa(start) + `, ` + strconv.Itoa(step)
	return self
}

func (self *Ar) Get() (query string) {
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
