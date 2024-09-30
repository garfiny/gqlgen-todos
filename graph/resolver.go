package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"

	"github.com/garfiny/gqlgen-todos/graph/model"
)

type Resolver struct {
	TodoList []*model.Todo
	Wishes   []*model.Wish
}

func (r *wishResolver) WishList(ctx context.Context) ([]*model.Wish, error) {
	return r.Wishes, nil
}
