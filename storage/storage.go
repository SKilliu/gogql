package storage

import (
	"github.com/SKilliu/gogql/graph/model"
)

type IStorage interface {
	GetAll() []*model.User
	GetByID(id string) (*model.User, error)
	Create(usr model.NewUser) (*model.User, error)
	AddFriend(userID, friendID string) (*model.User, error)
}

var users IStorage
var friends map[string][]*model.User

func InitStorage() {
	initUsers()

	friends = make(map[string][]*model.User)
}

func GetUsersStorage() IStorage {
	return users
}
