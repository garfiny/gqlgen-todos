package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUserByID(id string) (*model.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("id =?", id).Frist()
	if err != nil {
		return nil, err
	}
	return &user, nil

}
