// Сессия
package core

import (
	typDb "types/db"
)

// Сессионные данные для контроллеров
type Session struct {
	Uri           *typDb.Uri           // Информация о текущем Uri
	User          *typDb.Users         // Пользователь
	UserGroupsMap map[uint64]uint64    // Группы пользователя
	Access        *typDb.GroupsUri     // Права в виде структуры. Это агрегировання запись
	AccessMap     map[string]bool      // Права. В ввиде хеша (метод = право).
	Controllers   []*typDb.Controllers // Контроллеры с учетом прав (before, main, after)
}
