package service

import (
	"context"

	"github.com/SKilliu/gogql/graph/model"
	"github.com/SKilliu/gogql/storage"
)

type UserService struct {
	storage storage.IStorage
}

func NewUserService(storage storage.IStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (us *UserService) Login(user *model.NewUser) (*model.User, error) {
	//token, err := JwtGenerate(context.Background(), user.ID)
	//if err != nil {
	//	return nil, err
	//}
	//
	//user.Token = token

	//user, err = us.storage.Create(user)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
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
