package model

import (
	"strconv"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/kshamiev/sungora/pkg/typ"
)

type raw string

// IN реализация запросов с параметрами IN
func (q raw) IN(tag string, params interface{}) raw {
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
	return raw(strings.Replace(string(q), tag, "'"+strings.Join(s, "','")+"'", -1))
}

func (q raw) String() string {
	return strings.Replace(string(q), "\n", "", -1)
}

// Хранилище нативных запросов к БД
const (
	SQLAppVersion raw = `SELECT MAX(version_id) as version_id FROM goose_db_version WHERE is_applied = TRUE`
)
