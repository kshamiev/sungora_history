// запуск теста
// SET GOPATH=C:\Work\projectName
// go test -v lib/cache
// go test -v -bench . lib/cache
package cache_test

import (
	"testing"
	"time"

	"lib/cache"
)

// Пользователи
type Users struct {
	Id              uint64    // Id
	Users_Id        uint64    // Пользователь
	Login           string    // Логин пользователя
	Password        string    // Пароль пользователя (SHA256)
	Email           string    `json:"email"` // Email
	LastName        string    // Фамилия
	Name            string    // Имя
	MiddleName      string    // Отчество
	LastFailed      time.Time // Дата последней не удачной попытки входа
	IsAccess        bool      // Доступ разрешен
	IsCondition     bool      // Условия пользователя
	IsActivated     bool      // Активированный
	DateOnline      time.Time // Дата последнего посещения
	Date            time.Time // Дата регистрации
	Del             bool      // Запись удалена
	Hash            string    // Контрольная сумма для синхронизации (SHA256)
	CookieActivated string    // Кука активации и идентификации
}

func TestCache(t *testing.T) {

	// Несуществующий кеш
	userPointNotExists := cache.Get(`userPointNotExists`)
	if userPointNotExists != nil {
		t.Error(`Ошибка получение несуществующего кеша`)
	}

	// Кеш по ссылке
	var userPoint = new(Users)
	userPoint.Id = 2845
	userPoint.MiddleName = `Шариков`
	userPoint.Name = `Полиграф`
	userPoint.LastName = `Полиграфович`
	// сохранение
	cache.Set(`userPoint`, userPoint, cache.TS05)
	time.Sleep(time.Second * 2)
	if 1 != cache.Len() {
		t.Error(`Хранилище кеша не сохраняет данные`)
	}
	// получение
	userPointCheck := cache.Get(`userPoint`).(*Users)
	time.Sleep(time.Second * 4)
	if 0 != cache.Len() {
		t.Error(`Хранилище кеша не удаляет старые данные `)
	}
	// проверка ссылки
	userPointCheck.MiddleName = `Петров`
	userPointCheck.Name = `Василий`
	userPointCheck.LastName = `Николаевич`
	if userPoint.MiddleName != userPointCheck.MiddleName {
		t.Error(`Кеш не сохраняет данные по ссылке`)
	}

	// Кеш по значению
	var userValue = Users{}
	userValue.Id = 2845
	userValue.MiddleName = `Шариков`
	userValue.Name = `Полиграф`
	userValue.LastName = `Полиграфович`
	// сохранение
	cache.Set(`userValue`, userValue, cache.TS05)
	time.Sleep(time.Second * 2)
	if 1 != cache.Len() {
		t.Error(`Хранилище кеша не сохраняет данные`)
	}
	// получение
	userValueCheck := cache.Get(`userValue`).(Users)
	time.Sleep(time.Second * 4)
	if 0 != cache.Len() {
		t.Error(`Хранилище кеша не удаляет старые данные `)
	}
	// проверка ссылки
	userValueCheck.MiddleName = `Петров`
	userValueCheck.Name = `Василий`
	userValueCheck.LastName = `Николаевич`
	if userValue.MiddleName == userValueCheck.MiddleName {
		t.Error(`Кеш не сохраняет данные по значению`)
	}

	// Завершение
	if false == cache.GoClose() {
		t.Fatal(`Ошибка завершения службы кеширования`)
	}
}
