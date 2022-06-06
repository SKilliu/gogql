package service

import (
	"github.com/SKilliu/gogql/graph/model"
	"github.com/SKilliu/gogql/storage"
)

type PostService struct {
	storage storage.IPosts
}

func NewPostService(storage storage.IPosts) *PostService {
	return &PostService{
		storage: storage,
	}
}

func (ps *PostService) New(login model.NewPost, userID string) (*model.Post, error) {
	return ps.storage.Create(login, userID)
}

func (ps *PostService) GetAll(authorID, id *string) ([]*model.Post, error) {
	return ps.storage.GetAll(authorID, id), nil
}
