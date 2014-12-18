package controller

import (
	typDb "types/db"
)

// Конфигурация контроллеров (Controllers)
var ConfigControllers = make(map[string]typDb.Controllers)

// Конфигурация разделов (Uri)
var ConfigUri = make(map[string]typDb.Uri)
