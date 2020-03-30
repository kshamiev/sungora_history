package graphql

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/null"

	"github.com/kshamiev/sungora/pb/typ"
	"github.com/kshamiev/sungora/pkg/gql"
	"github.com/kshamiev/sungora/pkg/models"
)

type queryResolver struct{ *Resolver }

// nolint[:dupl]
func (r *queryResolver) Todos(ctx context.Context, limit, offset *int) ([]*gql.Todo, error) {
	data := []*gql.Todo{{
		ID:       typ.UUIDNew(),
		Number:   34,
		Price:    45.78,
		Decimal:  decimal.NewFromFloat(56.87),
		Done:     false,
		Access:   gql.AccessAdmin,
		Text:     "popcorn 1",
		TextNull: null.String{},
		CreateAt: time.Now(),
		DeleteAt: null.Time{},
		Item:     &gql.Item{Code: "111"},
		Items: []*gql.Item{
			{Code: "ONE"},
			{Code: "TWO"},
			{Code: "TREE"},
		},
		LinkID: typ.UUIDNew(),
	}, {
		ID:       typ.UUIDNew(),
		Text:     "popcorn 2",
		Number:   876,
		Price:    45.80,
		Done:     false,
		CreateAt: time.Now(),
		Access:   gql.AccessAdmin,
		Item:     &gql.Item{Code: "222"},
		LinkID:   typ.UUID{},
	}, {
		ID:       typ.UUIDNew(),
		Text:     "popcorn 3",
		Number:   768,
		Price:    86.78,
		Done:     false,
		CreateAt: time.Now(),
		Access:   gql.AccessAdmin,
		Item:     &gql.Item{Code: "333"},
		LinkID:   typ.UUID{},
	}}
	return data, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Roles(ctx context.Context) ([]*models.Role, error) {
	panic("not implemented")
}
func (r *queryResolver) Interfaces(ctx context.Context) ([]gql.Characters, error) {
	panic("not implemented")
}
func (r *queryResolver) Union(ctx context.Context) ([]gql.SearchResult, error) {
	panic("not implemented")
}
