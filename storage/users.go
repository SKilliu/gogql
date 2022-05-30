package storage

import (
	"errors"
	"sync"
	"time"

	"github.com/SKilliu/gogql/tools"
	"github.com/google/uuid"

	"github.com/SKilliu/gogql/graph/model"
)

type Users struct {
	mu      sync.Mutex
	storage map[string]*model.User
}

func initUsers() {
	users = &Users{
		storage: make(map[string]*model.User),
	}
}

func (u *Users) GetAll() []*model.User {
	u.mu.Lock()
	defer u.mu.Unlock()
	var usrs []*model.User
	for _, v := range u.storage {
		usrs = append(usrs, v)
	}

	return usrs
}

func (u *Users) GetByID(id string) (*model.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()
	usr, ok := u.storage[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return usr, nil
}

func (u *Users) Create(usr model.NewUser) (*model.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	_, err := u.findByEmail(usr.Email)
	if err != nil {
		uid := uuid.New().String()
		usr := &model.User{
			ID:              uid,
			Name:            usr.Name,
			Password:        tools.HashPassword(usr.Password),
			Status:          model.StatusActive,
			Email:           usr.Email,
			FollowersAmount: 0,
			CreatedAt:       time.Now().String(),
		}

		u.storage[usr.ID] = usr
		return usr, nil
	}

	return nil, errors.New("user already exists")
}

func (u *Users) AddFriend(userID, friendID string) (*model.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	friend, ok := u.storage[friendID]
	if !ok {
		return nil, errors.New("user not found")
	}

	friends[userID] = append(friends[userID], friend)
	return nil, nil
}

func (u *Users) findByEmail(uemail string) (*model.User, error) {
	for _, user := range u.storage {
		if user.Email == uemail {
			return user, nil
		}
	}

	return nil, errors.New("user doesn't exist")
}

func countFriends(uid string) int {
	return len(friends[uid])
}
