package graphql

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	data := []*Todo{{
		ID:   "uid 1",
		Text: "popcorn 1",
		Done: false,
		User: nil,
	}, {
		ID:   "uid 2",
		Text: "popcorn 2",
		Done: false,
		User: nil,
	}, {
		ID:   "uid 3",
		Text: "popcorn 3",
		Done: false,
		User: nil,
	},
	}
	return data, nil
}
