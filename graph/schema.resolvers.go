package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/darkjoka/yodel/graph/generated"
	"github.com/darkjoka/yodel/graph/model"
)

// ID is the resolver for the id field.
func (r *commentResolver) ID(ctx context.Context, obj *model.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserID is the resolver for the userId field.
func (r *commentResolver) UserID(ctx context.Context, obj *model.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// PostID is the resolver for the postId field.
func (r *commentResolver) PostID(ctx context.Context, obj *model.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatedAt is the resolver for the createdAt field.
func (r *commentResolver) CreatedAt(ctx context.Context, obj *model.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// ID is the resolver for the id field.
func (r *postResolver) ID(ctx context.Context, obj *model.Post) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserID is the resolver for the userId field.
func (r *postResolver) UserID(ctx context.Context, obj *model.Post) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// IsFlagged is the resolver for the isFlagged field.
func (r *postResolver) IsFlagged(ctx context.Context, obj *model.Post) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Longitude is the resolver for the longitude field.
func (r *postResolver) Longitude(ctx context.Context, obj *model.Post) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Latitude is the resolver for the latitude field.
func (r *postResolver) Latitude(ctx context.Context, obj *model.Post) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreatedAt is the resolver for the createdAt field.
func (r *postResolver) CreatedAt(ctx context.Context, obj *model.Post) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *model.Post) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Karma is the resolver for the karma field.
func (r *userResolver) Karma(ctx context.Context, obj *model.User) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
