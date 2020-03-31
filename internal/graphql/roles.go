package graphql

import (
	"context"

	"github.com/kshamiev/sungora/pb/modelsun"
)

type roleResolver struct{ *Resolver }

func (r *roleResolver) Users(ctx context.Context, obj *modelsun.Role, limit, offset *int) ([]*modelsun.User, error) {
	panic("not implemented")
}
