package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/garfiny/gqlgen-todos/graph/model"
)

type TodoRepo struct {
	DB * pg.DB
}

func (r *TodoRepo) todos(userId string) ([]*model.Todo, error) {
	return nil, nil
}
