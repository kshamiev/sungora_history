package graphql

import (
	"context"

	"github.com/kshamiev/sungora/pkg/models"
)

type roleResolver struct{ *Resolver }

func (r *roleResolver) Users(ctx context.Context, obj *models.Role, limit, offset *int) ([]*models.User, error) {
	panic("not implemented")
}
