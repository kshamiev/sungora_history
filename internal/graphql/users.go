package graphql

import (
	"context"

	"github.com/kshamiev/sungora/pb/modelsun"
)

type userResolver struct{ *Resolver }

func (r *userResolver) Roles(ctx context.Context, obj *modelsun.User, limit, offset *int) ([]*modelsun.Role, error) {
	panic("not implemented")
}
