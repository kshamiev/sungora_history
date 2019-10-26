package graphql

import (
	"context"
	"time"

	"github.com/kshamiev/sungora/graphql/gen"
	"github.com/kshamiev/sungora/graphql/mod"
	"github.com/kshamiev/sungora/pkg/app"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input mod.NewTodo) (*mod.Todo, error) {
	app.Dumper(input)
	d := time.Now()
	return &mod.Todo{
		ID:       "uid 1",
		Text:     "popcorn 1",
		Done:     false,
		CreateAt: &d,
		Role:     &mod.Role{},
	}, nil
}

type queryResolver struct{ *Resolver }

// nolint[:dupl]
func (r *queryResolver) Todos(ctx context.Context) ([]*mod.Todo, error) {
	d := time.Now()
	data := []*mod.Todo{{
		ID:       "uid 1",
		Text:     "popcorn 1",
		Done:     false,
		CreateAt: &d,
		Role:     &mod.Role{},
	}, {
		ID:       "uid 2",
		Text:     "popcorn 2",
		Done:     false,
		CreateAt: &d,
		Role:     &mod.Role{},
	}, {
		ID:       "uid 3",
		Text:     "popcorn 3",
		Done:     false,
		CreateAt: &d,
		Role:     &mod.Role{},
	}}
	return data, nil
}

// nolint[:dupl]
func (r *queryResolver) Funtik(ctx context.Context) ([]*mod.Todo, error) {
	d := time.Now()
	data := []*mod.Todo{{
		ID:       "uid 1",
		Text:     "popcorn 1",
		Done:     false,
		CreateAt: &d,
		Role:     &mod.Role{},
	}, {
		ID:       "uid 2",
		Text:     "popcorn 2",
		Done:     false,
		CreateAt: &d,
		Role:     &mod.Role{},
	}, {
		ID:       "uid 3",
		Text:     "popcorn 3",
		Done:     false,
		CreateAt: &d,
		Role:     &mod.Role{},
	}}
	return data, nil
}
