package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/SKilliu/gogql/graph/generated"
	"github.com/SKilliu/gogql/graph/model"
	"github.com/SKilliu/gogql/middlewares"
	"github.com/SKilliu/gogql/tools"
	"github.com/google/uuid"
)

func (r *mutationResolver) Login(ctx context.Context, input model.NewUser) (*model.User, error) {
	uid := uuid.New().String()
	usr := &model.User{
		ID:              uid,
		Name:            input.Name,
		Password:        tools.HashPassword(input.Password),
		Status:          model.StatusActive,
		Email:           input.Email,
		FollowersAmount: 0,
		CreatedAt:       time.Now().String(),
	}

	usr, err := r.Service.Login(&input)
	if err != nil {
		return nil, err
	}

	return usr, err
}

func (r *mutationResolver) Registration(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.Service.Registration(input)
}

func (r *mutationResolver) Follow(ctx context.Context, input model.NewFollow) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, id *string, name *string, status *model.Status) ([]*model.User, error) {
	return r.Service.GetAll()
}

func (r *queryResolver) Profile(ctx context.Context) (*model.User, error) {
	jwt := middlewares.CtxValue(ctx)
	return r.Service.GetByID(jwt.ID)
}

func (r *queryResolver) Followers(ctx context.Context, id *string, status *model.Status) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
