package service

import (
	"context"

	"github.com/SKilliu/gogql/tools"

	"github.com/SKilliu/gogql/graph/model"
	"github.com/SKilliu/gogql/storage"
)

type UserService struct {
	storage storage.IUsers
}

func NewUserService(storage storage.IUsers) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (us *UserService) Login(login model.LoginData) (*model.User, error) {
	user, err := us.storage.GetByEmail(login.Email)
	if err != nil {
		return nil, err
	}

	err = tools.ComparePassword(user.Password, login.Password)
	if err != nil {
		return nil, err
	}

	token, err := JwtGenerate(context.Background(), user.ID)
	if err != nil {
		return nil, err
	}
	user.Token = token

	return user, err
}

func (us *UserService) GetAll() ([]*model.User, error) {
	return us.storage.GetAll(), nil
}

func (us *UserService) GetByID(uid string) (*model.User, error) {
	return us.storage.GetByID(uid)
}

func (us *UserService) Registration(nu model.NewUser) (*model.User, error) {
	user, err := us.storage.Create(nu)
	if err != nil {
		return nil, err
	}

	token, err := JwtGenerate(context.Background(), user.ID)
	if err != nil {
		return nil, err
	}
	user.Token = token

	return user, err
}

func (us *UserService) Follow(userID, followedID string) (*model.User, error) {
	return us.storage.Follow(userID, followedID)
}
