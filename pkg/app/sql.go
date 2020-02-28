package app

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
	case []time.Time:
		s = make([]string, len(data))
		for i, v := range data {
			s[i] = v.Format(time.RFC3339)
		}
	case []int:
		s = make([]string, len(data))
		for i, v := range data {
			s[i] = strconv.Itoa(v)
		}
	case []decimal.Decimal:
		s = make([]string, len(data))
		for i, d := range data {
			s[i] = d.String()
		}
	case []typ.UUID:
		s = make([]string, len(data))
		for i, uid := range data {
			s[i] = uid.String()
		}
	case typ.UUIDS:
		s = make([]string, len(data))
		for i, uid := range data {
			s[i] = uid.String()
		}
	}

	return SQL(strings.ReplaceAll(string(q), tag, "'"+strings.Join(s, "','")+"'"))
}

// Replace замена на нужные параметры
func (q SQL) Replace(tag, value string) SQL {
	return SQL(strings.ReplaceAll(string(q), tag, value))
}

// String возвращает очищенный конечный запрос
// nolint: gocyclo
func (q SQL) String(p ...interface{}) string {
	query := strings.ReplaceAll(strings.ReplaceAll(string(q), "\n", " "), "\t", "")

	for i := range p {
		tag := "$" + strconv.Itoa(i+1)

		switch data := p[i].(type) {
		case time.Time:
			query = strings.Replace(query, tag, "'"+data.Format(time.RFC3339)+"'", -1)
		case string:
			query = strings.Replace(query, tag, "'"+data+"'", -1)
		case float32, float64:
			query = strings.Replace(query, tag, fmt.Sprintf("%f", data), -1)
		case int, int8, int16, int32, int64:
			query = strings.Replace(query, tag, fmt.Sprintf("%d", data), -1)
		case uint, uint8, uint16, uint32, uint64:
			query = strings.Replace(query, tag, fmt.Sprintf("%d", data), -1)
		case decimal.Decimal:
			query = strings.Replace(query, tag, data.String(), -1)
		case typ.UUID:
			query = strings.Replace(query, tag, "'"+data.String()+"'", -1)
		case []time.Time:
			s := make([]string, len(data))
			for i, v := range data {
				s[i] = v.Format(time.RFC3339)
			}

			query = strings.Replace(query, tag, "'"+strings.Join(s, "','")+"'", -1)
		case []string:
			query = strings.Replace(query, tag, "'"+strings.Join(data, "','")+"'", -1)
		case []int:
			s := make([]string, len(data))

			for i, v := range data {
				s[i] = strconv.Itoa(v)
			}

			query = strings.Replace(query, tag, "'"+strings.Join(s, "','")+"'", -1)
		case []decimal.Decimal:
			s := make([]string, len(data))
			for i, d := range data {
				s[i] = d.String()
			}

			query = strings.Replace(query, tag, "'"+strings.Join(s, "','")+"'", -1)
		case []typ.UUID:
			s := make([]string, len(data))
			for i, uid := range data {
				s[i] = uid.String()
			}

			query = strings.Replace(query, tag, "'"+strings.Join(s, "','")+"'", -1)
		case typ.UUIDS:
			s := make([]string, len(data))
			for i, uid := range data {
				s[i] = uid.String()
			}

			query = strings.Replace(query, tag, "'"+strings.Join(s, "','")+"'", -1)
		}
	}

	return query
}
