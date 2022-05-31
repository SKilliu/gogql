package storage

import (
	"github.com/SKilliu/gogql/graph/model"
)

type IUsers interface {
	GetAll() []*model.User
	GetByID(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(usr model.NewUser) (*model.User, error)
	Follow(userID, followedID string) (*model.User, error)
}

type IPosts interface {
	GetAll(authorID *string, id *string) []*model.Post
	//GetByID(id string) (*model.Post, error)
	//GetByName(email string) (*model.Post, error)
	Create(usr model.NewPost, userID string) (*model.Post, error)
}

var users IUsers
var posts IPosts

func InitStorage() {
	initUsers()
	initPosts()
}

func GetUsersStorage() IUsers {
	return users
}

func GetPostsStorage() IPosts {
	return posts
}
