package setup

// Инициализация модулей.
// Сюда добавляются модули для инсталяции, инициализации и последующей работы в системе
import (
	baseService `core/base/service`
	_ `core/base/setup`
)

// Запуск работы служб модулей приложения
func GoServiceStart() {
	baseService.GoStart() // слжужбы базового модуля
}

// Завершение работы служб модулей приложения
func GoServiceStop() {
	baseService.GoStop() // слжужбы базового модуля
}
