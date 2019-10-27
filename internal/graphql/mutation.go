package graphql

import (
	"context"
	"time"

	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/gql"
	"github.com/kshamiev/sungora/pkg/typ"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input gql.NewTodo) (*gql.Todo, error) {
	app.Dumper(input)
	return &gql.Todo{
		ID:       typ.UUIDNew(),
		Text:     "popcorn 1",
		Done:     false,
		CreateAt: time.Now(),
		Item:     &gql.Item{},
	}, nil
}
