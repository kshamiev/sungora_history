package main

// Инициализация модулей.
// Сюда добавляются модули для инсталяции, инициализации и последующей работы в системе
import (
	baseService `core/base/service`
	_ `core/base/setup`
)

// Запуск работы служб модулей приложения
func goServiceStart() {
	baseService.GoStart() // слжужбы базового модуля
}

// Завершение работы служб модулей приложения
func goServiceStop() {
	baseService.GoStop() // слжужбы базового модуля
}
