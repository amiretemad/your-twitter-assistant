package repository

import "tweak_twitter/pkg/model"

type UserInterface interface {
	SaveList(user []model.User) (bool, error)
}

type UserRepository struct {
}

func (u *UserRepository) Save(user *model.User) bool {
	user.ID = 1
	return false
}
