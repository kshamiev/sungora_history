package graphql

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/null"

	"github.com/kshamiev/sungora/pkg/gql"
	"github.com/kshamiev/sungora/pkg/typ"
)

type queryResolver struct{ *Resolver }

// nolint[:dupl]
func (r *queryResolver) Todos(ctx context.Context) ([]*gql.Todo, error) {
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
		Role:     &gql.Role{Code: "111"},
		Roles: []*gql.Role{
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
		Role:     &gql.Role{Code: "222"},
		LinkID:   typ.UUID{},
	}, {
		ID:       typ.UUIDNew(),
		Text:     "popcorn 3",
		Number:   768,
		Price:    86.78,
		Done:     false,
		CreateAt: time.Now(),
		Access:   gql.AccessAdmin,
		Role:     &gql.Role{Code: "333"},
		LinkID:   typ.UUID{},
	}}
	return data, nil
}

// nolint[:dupl]
func (r *queryResolver) Funtik(ctx context.Context) ([]*gql.Todo, error) {
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
		Role:     &gql.Role{Code: "111"},
		Roles: []*gql.Role{
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
		Role:     &gql.Role{Code: "222"},
		LinkID:   typ.UUID{},
	}, {
		ID:       typ.UUIDNew(),
		Text:     "popcorn 3",
		Number:   768,
		Price:    86.78,
		Done:     false,
		CreateAt: time.Now(),
		Access:   gql.AccessAdmin,
		Role:     &gql.Role{Code: "333"},
		LinkID:   typ.UUID{},
	}}
	return data, nil
}
