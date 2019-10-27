package graphql

import (
	"github.com/kshamiev/sungora/pkg/gql"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Role() gql.RoleResolver {
	return &roleResolver{r}
}
func (r *Resolver) Todo() gql.TodoResolver {
	return &todoResolver{r}
}
func (r *Resolver) User() gql.UserResolver {
	return &userResolver{r}
}
