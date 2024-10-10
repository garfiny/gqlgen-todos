package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/garfiny/gqlgen-todos/graph/model"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where("id =?", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil

}
