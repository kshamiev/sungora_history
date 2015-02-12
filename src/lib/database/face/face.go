// Описание интерфейсов.
//
// Интерфейс для работы с абстрактной (любого типа) БД
// Интерфейс для формирования запросов к абстрактного БД
package face

// Интерфейс для формирования запросов к абстрактного БД
type QubFace interface {
	Select(property string) QubFace
	SelectScenario(source, scenario string) QubFace
	From(from string) QubFace
	Where(where string) QubFace
	Group(group string) QubFace
	Having(having string) QubFace
	Order(order string) QubFace
	Limit(start, step int) QubFace
	Get() (query string)
}

// Интерфейс для работы с абстрактной (любого типа) БД
type DbFace interface {
	Select(Object interface{}, sql string, params ...interface{}) (err error)
	SelectMap(ObjectMap interface{}, sql string, params ...interface{}) (err error)
	SelectSlice(ObjectSlice interface{}, sql string, params ...interface{}) (err error)
	SelectData(object interface{}) (err error)
	Insert(Object interface{}, source string, properties ...map[string]string) (insertId uint64, err error)
	Update(Object interface{}, source, key string, properties ...map[string]string) (affectedRow uint64, err error)
	Query(sql string, params ...interface{}) (err error)
	QueryByte(data []byte) (messages []string, err error)
	CallFunc(Object interface{}, nameCall string, params ...interface{}) (err error)
	CallExec(Object interface{}, nameCall string, params ...interface{}) (err error)
	Free()
}
