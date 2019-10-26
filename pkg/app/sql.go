package app

import (
	"strconv"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/kshamiev/sungora/pkg/typ"
)

type SQL string

// IN реализация запросов с параметрами IN
func (q SQL) IN(tag string, params interface{}) SQL {
	var s []string
	switch data := params.(type) {
	case []string:
		s = data
	case []decimal.Decimal:
		s = make([]string, len(data))
		for i, d := range data {
			s[i] = d.String()
		}
	case []int:
		s = make([]string, len(data))
		for i, v := range data {
			s[i] = strconv.Itoa(v)
		}
	case []typ.UUID:
		s = make([]string, len(data))
		for i, uid := range data {
			s[i] = uid.String()
		}
	}
	return SQL(strings.Replace(string(q), tag, "'"+strings.Join(s, "','")+"'", -1))
}

// Replace замена на нужные параметры
func (q SQL) Replace(tag, value string) SQL {
	return SQL(strings.Replace(string(q), tag, value, 1))
}

// String возвращает очищенный конечный запрос
func (q SQL) String() string {
	s := strings.ReplaceAll(string(q), "\n", " ")
	return strings.ReplaceAll(s, "\t", "")
}
