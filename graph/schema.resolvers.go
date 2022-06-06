package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/SKilliu/gogql/graph/generated"
	"github.com/SKilliu/gogql/graph/model"
	"github.com/SKilliu/gogql/middlewares"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginData) (*model.User, error) {
	return r.UserService.Login(input)
}

func (r *mutationResolver) Registration(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.UserService.Registration(input)
}

func (r *mutationResolver) Post(ctx context.Context, input model.NewPost) (*model.Post, error) {
	jwt := middlewares.CtxValue(ctx)
	return r.PostService.New(input, jwt.ID)
}

func (r *queryResolver) Users(ctx context.Context, id *string, name *string, status *model.Status) ([]*model.User, error) {
	return r.UserService.GetAll()
}

func (r *queryResolver) Profile(ctx context.Context) (*model.User, error) {
	jwt := middlewares.CtxValue(ctx)
	return r.UserService.GetByID(jwt.ID)
}

func (r *queryResolver) Posts(ctx context.Context, id *string, authorID *string) ([]*model.Post, error) {
	return r.PostService.GetAll(authorID, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
