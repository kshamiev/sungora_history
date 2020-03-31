package graphql

import (
	"context"

	"github.com/kshamiev/sungora/pb/modelsun"
	"github.com/kshamiev/sungora/pkg/gql"
)

type todoResolver struct{ *Resolver }

func (r *todoResolver) Users(ctx context.Context, obj *gql.Todo, limit, offset *int) ([]*modelsun.User, error) {
	panic("not implemented")
}
func (r *todoResolver) Roles(ctx context.Context, obj *gql.Todo, limit, offset *int) ([]*modelsun.Role, error) {
	panic("not implemented")
}
