package storage

import "github.com/SKilliu/gogql/graph/model"

type UsersQ interface {
	GetAll() ([]*model.User, error)
	Create(user model.User) error
	GetByID(uid string) (*model.User, error)
}

type UsersWrapper struct {
	parent *DB
}

func (d *DB) UsersQ() UsersQ {
	return &UsersWrapper{
		parent: d,
	}
}

func (u *UsersWrapper) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := u.parent.DB().Find(&users).Error
	return users, err
}

func (u *UsersWrapper) Create(user model.User) error {
	return u.parent.DB().Create(&user).Error
}

func (u *UsersWrapper) GetByID(uid string) (*model.User, error) {
	var user *model.User
	err := u.parent.DB().First(&user, "id = ?", uid).Error
	return user, err
}
