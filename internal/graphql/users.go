package graphql

import (
	"context"

	"github.com/kshamiev/sungora/pkg/models"
)

type userResolver struct{ *Resolver }

func (r *userResolver) Roles(ctx context.Context, obj *models.User, limit, offset *int) ([]*models.Role, error) {
	panic("not implemented")
}
