// Описание интерфейсов.
//
// Описание интерфейса для работы с абстрактной (любого типа) БД
// Описание интерфейса для формирования запросов к абстрактного БД
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
	// Загрузка одного объекта (записи) из БД
	Select(Object interface{}, sql string, params ...interface{}) (err error)
	// Загрузка объектов (записей) из БД в хеш
	SelectMap(ObjectMap interface{}, sql string, params ...interface{}) (err error)
	// Загрузка объектов (записей) из БД в срез
	SelectSlice(ObjectSlice interface{}, sql string, params ...interface{}) (err error)
	// Загрузка БД в память
	SelectData(object interface{}) (err error)
	// Сохранение объекта (записи) в БД
	Insert(Object interface{}, source string, properties ...map[string]string) (insertId uint64, err error)
	// Изменение объекта (записи) в БД
	Update(Object interface{}, source, key string, properties ...map[string]string) (affectedRow uint64, err error)
	// Пользовательский запрос на обновление в БД
	Query(sql string, params ...interface{}) (err error)
	// Пользовательский запрос на обновление в БД (как правило из файлов)
	QueryByte(data []byte) (messages []string, err error)
	// Выполнение функций
	CallFunc(Object interface{}, nameCall string, params ...interface{}) (err error)
	// Выполнение хранимых процедур
	CallExec(Object interface{}, nameCall string, params ...interface{}) (err error)
	// Освобождение коннекта
	Free()
}

////
