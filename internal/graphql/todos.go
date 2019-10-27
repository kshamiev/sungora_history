package graphql

import (
	"context"

	"github.com/kshamiev/sungora/pkg/gql"
	"github.com/kshamiev/sungora/pkg/models"
)

type todoResolver struct{ *Resolver }

func (r *todoResolver) Users(ctx context.Context, obj *gql.Todo, limit, offset *int) ([]*models.User, error) {
	panic("not implemented")
}
func (r *todoResolver) Roles(ctx context.Context, obj *gql.Todo, limit, offset *int) ([]*models.Role, error) {
	panic("not implemented")
}
