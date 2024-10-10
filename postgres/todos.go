package postgres

type TodoRepo struct {
	DB * pg.DB
}

func (r *TodoRepo) todos(userId string) ([]*model.Todo, error) {
	return nil, nil
}
