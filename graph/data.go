package graph

import "github.com/garfiny/gqlgen-todos/graph/model"

var todos = []*model.Todo{
	{
		ID:     "1",
		Text:   "test",
		Done:   false,
		UserID: "1",
	},
}

var users = []*model.User{
	{
		ID:   "1",
		Name: "garfiny",
	},
}
