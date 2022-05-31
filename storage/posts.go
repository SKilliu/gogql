package storage

import (
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/SKilliu/gogql/graph/model"
)

type Posts struct {
	mu      sync.Mutex
	storage map[string][]*model.Post
}

func initPosts() {
	posts = &Posts{
		storage: make(map[string][]*model.Post),
	}
}

func (p *Posts) Create(input model.NewPost, userID string) (*model.Post, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	user, err := users.GetByID(userID)
	if err != nil {
		return nil, err
	}

	post := &model.Post{
		ID:         uuid.New().String(),
		Name:       input.Name,
		Content:    input.Content,
		AuthorName: user.Name,
		AuthorID:   user.ID,
		CreatedAt:  time.Now().String(),
	}

	p.storage[user.ID] = append(p.storage[user.ID], post)

	return post, nil
}

func (p *Posts) GetAll(authorID *string, id *string) []*model.Post {
	var psts []*model.Post
	switch {
	case authorID != nil:
		psts = append(psts, p.storage[*authorID]...)
	case id != nil:
		for _, userPosts := range p.storage {
			for _, post := range userPosts {
				if post.ID == *id {
					psts = append(psts, post)
					break
				}
			}
		}
	default:
		for _, posts := range p.storage {
			psts = append(psts, posts...)
		}
	}

	return psts
}
