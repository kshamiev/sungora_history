// Базовый и общий функционал для работы с типами.
// Реализует хранение и получение сценариев для типов проецирующих структуру БД в программе.
// Реализует вспомогательные функции общего назначегия по работе с типами (к примеру копирование).
package types

import (
	"reflect"
	"strings"

	"lib/logs"
)

// Хранилище всех сценариев. Инициализируется при запуске программы. (init())
var scenarioMap = make(map[string]map[string]*Scenario)

// Сценарий типа проецирующего сущность в БД (таблица, коллекция, файл)
type Scenario struct {
	Name        string      // Имя сценария
	Description string      // Описание сценария
	Property    []Property  // Свойства типа
	Sample      interface{} // Пример
}

// Описание конкретного свойства в сценарии типа
type Property struct {
	Name        string            // Название свойства
	Title       string            // Заголовок свойства
	AliasDb     string            // Алиас свойства для запросов в БД (только для чтения, не менять)
	Required    string            // Свойство обязательно для заполнения
	Readonly    string            // Свойство для чтения
	Default     string            // Значение по умолчанию
	FormType    string            // Тип элемента формы
	FormMask    string            // Шаблон принимаемых данных в определенном формате
	Visible     string            // Видимость
	EnumSet     map[string]string // Коллекция вариантов для полей типа (enum, set в БД)
	Hint        string            // Подсказка, справка по свойству - полю (краткое)
	Help        string            // Подсказка, справка по свойству - полю (полная)
	Placeholder string            // Подсказка в пустом поле формы
	Tab         int8              // Вкладка
	Column      int8              // Колонка
	Uri         string            // uri от текущего, по работе с данным свойством (для данных по связям)
}

// Получение списка всех сценариев источника
func GetScenarioList(source string) (scenarioList []string) {
	if _, ok := scenarioMap[source]; ok == false {
		return
	}
	for i := range scenarioMap[source] {
		scenarioList = append(scenarioList, i)
	}
	return
}

// Сохранение сценария (как правило в момент инциализации, запуска программы)
func SetScenario(source string, scenario map[string]Scenario) {
	scRoot, ok := scenario[`root`]
	if ok == false {
		return
	}
	source = strings.ToLower(source)

	if _, ok := scenarioMap[source]; ok == false {
		scenarioMap[source] = make(map[string]*Scenario)
	}
	scenarioMap[source][`root`] = &scRoot
	delete(scenario, `root`)

	var objValue = reflect.ValueOf(scenarioMap[source][`root`].Sample)
	if objValue.Kind() != reflect.Ptr {
		panic(`Сценарий ` + source + `.root не имеет объекта примера Sample`)
	}
	objValue = objValue.Elem()

	for name := range scenario {
		var sc = Scenario{}
		// Общие настройки сценария
		if scenario[name].Description == `` {
			sc.Description = scRoot.Description
		} else {
			sc.Description = scenario[name].Description
		}
		if scenario[name].Name == `` {
			sc.Name = scRoot.Name
		} else {
			sc.Name = scenario[name].Name
		}
		if scenario[name].Sample == nil {
			sc.Sample = scRoot.Sample
		} else {
			sc.Sample = scenario[name].Sample
		}
		// Свойства
		for i := range scenario[name].Property {
			for j := range scRoot.Property {
				if scenario[name].Property[i].Name == scRoot.Property[j].Name {
					var p = scRoot.Property[j]
					p.Name = scenario[name].Property[i].Name
					// настройки свойств
					if scenario[name].Property[i].Title == `` {
						p.Title = scRoot.Property[j].Title
					} else {
						p.Title = scenario[name].Property[i].Title
					}
					if p.Title == `` {
						p.Title = p.Name
					}

					// алиас поля в БД
					f, fOk := objValue.Type().FieldByName(p.Name)
					if fOk == false {
						panic(`Свойство ` + source + `.` + p.Name + ` не алё в структуре`)
					}
					tag := f.Tag.Get(`db`)
					if tag == `` {
						p.AliasDb = "`" + p.Name + "`"
					} else {
						p.AliasDb = "`" + tag + "`"
					}

					if scenario[name].Property[i].Required == `` {
						p.Required = scRoot.Property[j].Required
					} else {
						p.Required = scenario[name].Property[i].Required
					}

					if scenario[name].Property[i].Readonly == `` {
						p.Readonly = scRoot.Property[j].Readonly
					} else {
						p.Readonly = scenario[name].Property[i].Readonly
					}

					if scenario[name].Property[i].Default == `` {
						p.Default = scRoot.Property[j].Default
					} else {
						p.Default = scenario[name].Property[i].Default
					}

					if scenario[name].Property[i].FormType == `` {
						p.FormType = scRoot.Property[j].FormType
					} else {
						p.FormType = scenario[name].Property[i].FormType
					}

					if scenario[name].Property[i].FormMask == `` {
						p.FormMask = scRoot.Property[j].FormMask
					} else {
						p.FormMask = scenario[name].Property[i].FormMask
					}

					if scenario[name].Property[i].Visible == `` {
						p.Visible = scRoot.Property[j].Visible
					} else {
						p.Visible = scenario[name].Property[i].Visible
					}

					if scenario[name].Property[i].EnumSet == nil {
						p.EnumSet = scRoot.Property[j].EnumSet
					} else {
						p.EnumSet = scenario[name].Property[i].EnumSet
					}

					if scenario[name].Property[i].Hint == `` {
						p.Hint = scRoot.Property[j].Hint
					} else {
						p.Hint = scenario[name].Property[i].Hint
					}

					if scenario[name].Property[i].Help == `` {
						p.Help = scRoot.Property[j].Help
					} else {
						p.Help = scenario[name].Property[i].Help
					}

					if scenario[name].Property[i].Placeholder == `` {
						p.Placeholder = scRoot.Property[j].Placeholder
					} else {
						p.Placeholder = scenario[name].Property[i].Placeholder
					}

					if scenario[name].Property[i].Tab == 0 {
						p.Tab = scRoot.Property[j].Tab
					} else {
						p.Tab = scenario[name].Property[i].Tab
					}

					if scenario[name].Property[i].Column == 0 {
						p.Column = scRoot.Property[j].Column
					} else {
						p.Column = scenario[name].Property[i].Column
					}

					if scenario[name].Property[i].Uri == `` {
						p.Uri = scRoot.Property[j].Uri
					} else {
						p.Uri = scenario[name].Property[i].Uri
					}

					sc.Property = append(sc.Property, p)
					break
				}
			}
		}
		scenarioMap[source][strings.ToLower(name)] = &sc
	}
}

// Получение выбранного сценария для выбранного источника
func GetScenario(source, scenarioName string) (scenario *Scenario, err error) {
	var ok bool
	source = strings.ToLower(source)
	scenarioName = strings.ToLower(scenarioName)
	if scenario, ok = scenarioMap[source][scenarioName]; ok == false {
		return nil, logs.Base.Error(104, source, scenarioName).Err
	}
	return
}
