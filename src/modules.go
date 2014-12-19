package main

// Инициализация модулей.
// Сюда добавляются модули для регистрации, инициализации и последующей работы в системе
import (
	baseService `core/base/service`
	_ `core/base/setup`
)

// ServiceStartGo Запуск работы служб модулей приложения
func goServiceStart() {
	// слжужбы базового модуля
	baseService.GoStart()
}

// ServiceCloseGo Завершение работы служб модулей приложения
func goServiceStop() {
	// слжужбы базового модуля
	baseService.GoStop()
}
