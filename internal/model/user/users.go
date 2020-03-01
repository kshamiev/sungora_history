package user

import "github.com/kshamiev/sungora/pkg/models"

type Access struct {
	*models.User
	Roles models.RoleSlice `json:"roles"`
}
