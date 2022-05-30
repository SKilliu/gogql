package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/SKilliu/gogql/graph/generated"
	"github.com/SKilliu/gogql/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		ID:    uuid.New().String(),
		Name:  input.Name,
		Email: input.Email,
	}

	err := r.DB.UsersQ().Create(user)

	return &user, err
}

func (r *queryResolver) Users(ctx context.Context, id *string) ([]*model.User, error) {
	var (
		err   error
		users []*model.User
	)
	if id != nil {
		usr, err := r.DB.UsersQ().GetByID(*id)
		if err != nil {
			return nil, err
		}
		users = append(users, usr)
	} else {
		users, err = r.DB.UsersQ().GetAll()
	}

	return users, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
