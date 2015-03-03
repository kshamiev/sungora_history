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
	t.Logf(`Несуществующий кеш`)
	userPointNotExists := cache.Get(`userPointNotExists`)
	if userPointNotExists != nil {
		t.Error(`Ошибка получение несуществующего кеша`)
	} else {
		t.Logf(`Ok`)
	}

	// Сохранение кеша по ссылке
	t.Logf(`Сохранение кеша по ссылке`)
	var userPoint = new(Users)
	userPoint.Id = 2845
	userPoint.MiddleName = `Шариков`
	userPoint.Name = `Полиграф`
	userPoint.LastName = `Полиграфович`
	cache.Set(`userPoint`, userPoint, cache.TS05)
	time.Sleep(time.Second * 2)
	if 1 != cache.Len() {
		t.Error(`Хранилище кеша не сохраняет данные`)
	} else {
		t.Logf(`Ok`)
	}

	// Получение кеша
	t.Logf(`Получение кеша`)
	userPointCheck, ok := cache.Get(`userPoint`).(*Users)
	if ok == false {
		t.Error(`Ошибка получение ранее сохраненного кеша`)
	} else {
		t.Logf(`Ok`)
	}

	// Проверка полученного кеша по ссылке
	t.Logf(`Проверка полученного кеша по ссылке`)
	userPointCheck.MiddleName = `Петров`
	userPointCheck.Name = `Василий`
	userPointCheck.LastName = `Николаевич`
	if userPoint.MiddleName != userPointCheck.MiddleName {
		t.Error(`Кеш не сохраняет данные по ссылке`)
	} else {
		t.Logf(`Ok`)
	}

	// Проверка авто-удаления устаревшего кеша
	t.Logf(`Проверка авто-удаления устаревшего кеша`)
	time.Sleep(time.Second * 4)
	if 0 != cache.Len() {
		t.Error(`Хранилище кеша не удаляет старые данные `)
	} else {
		t.Logf(`Ok`)
	}

	////

	// Сохранение кеша по значению
	t.Logf(`Сохранение кеша по значению`)
	var userValue = Users{}
	userValue.Id = 2845
	userValue.MiddleName = `Шариков`
	userValue.Name = `Полиграф`
	userValue.LastName = `Полиграфович`
	cache.Set(`userValue`, userValue, cache.TS05)
	time.Sleep(time.Second * 2)
	if 1 != cache.Len() {
		t.Error(`Хранилище кеша не сохраняет данные`)
	} else {
		t.Logf(`Ok`)
	}

	// Получение кеша
	t.Logf(`Получение кеша`)
	userValueCheck, ok := cache.Get(`userValue`).(Users)
	if ok == false {
		t.Error(`Ошибка получение ранее сохраненного кеша`)
	} else {
		t.Logf(`Ok`)
	}

	// Проверка полученного кеша по значению
	t.Logf(`Проверка полученного кеша по значению`)
	userValueCheck.MiddleName = `Петров`
	userValueCheck.Name = `Василий`
	userValueCheck.LastName = `Николаевич`
	if userValue.MiddleName == userValueCheck.MiddleName {
		t.Error(`Кеш не сохраняет данные по значению`)
	} else {
		t.Logf(`Ok`)
	}

	// Проверка авто-удаления устаревшего кеша
	t.Logf(`Проверка авто-удаления устаревшего кеша`)
	time.Sleep(time.Second * 4)
	if 0 != cache.Len() {
		t.Error(`Хранилище кеша не удаляет старые данные `)
	} else {
		t.Logf(`Ok`)
	}

	// Завершение
	if false == cache.GoClose() {
		t.Fatal(`Ошибка завершения службы кеширования`)
	}
}

////
