package graph

import "github.com/garfiny/gqlgen-todos/graph/model"

var Wishes = []*model.Wish{
	{
		ID:   "1",
		Text: "This product will be successful",
		Todos: []model.Todo{
			{
				ID:     "3",
				Text:   "Work Hard",
				Done:   false,
				UserID: 1,
			},
			{
				ID:     "4",
				Text:   "Be Creative",
				Done:   false,
				UserID: 1,
			},
		},
	},
}

var Todos = []*model.Todo{
	{
		ID:     "1",
		Text:   "test",
		Done:   false,
		UserID: 2,
	},
	{
		ID:     "2",
		Text:   "test",
		Done:   false,
		UserID: 1,
	},
}

var Users = []*model.User{
	{
		ID:   1,
		Name: "garfiny",
	},
	{
		ID:   2,
		Name: "John",
	},
}
