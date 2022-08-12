package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/darkjoka/yodel/graph/generated"
	"github.com/darkjoka/yodel/graph/model"
	"github.com/google/uuid"
)

// ID is the resolver for the id field.
func (r *commentResolver) ID(ctx context.Context, obj *model.Comment) (string, error) {
	return obj.ID.String(), nil
}

// UserID is the resolver for the userId field.
func (r *commentResolver) UserID(ctx context.Context, obj *model.Comment) (string, error) {
	return obj.UserID.String(), nil
}

// PostID is the resolver for the postId field.
func (r *commentResolver) PostID(ctx context.Context, obj *model.Comment) (string, error) {
	return obj.PostID.String(), nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *commentResolver) CreatedAt(ctx context.Context, obj *model.Comment) (string, error) {
	return obj.CreatedAt.String(), nil
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (*model.User, error) {
	fmt.Println("got called")
	user := &model.User{Username: input.Username, Password: input.Password}
	err := r.UserScheme.Create(user, ctx)
	return user, err
}

// NewPost is the resolver for the newPost field.
func (r *mutationResolver) NewPost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	// TODO: run validations against input
	userId, _ := uuid.Parse(input.UserID)

	post := &model.Post{
		UserID:    userId,
		Content:   input.Content,
		Longitude: float32(input.Longitude),
		Latitude:  float32(input.Latitude),
		Location:  "Minneapolis",
	}

	_, err := r.PostScheme.DB.NewInsert().Model(post).Exec(ctx)
	return post, err
}

// NewComment is the resolver for the newComment field.
func (r *mutationResolver) NewComment(ctx context.Context, input *model.NewComment) (*model.Comment, error) {

	// TODO: post_id should be change to parent id and type possibly added.
	userId, _ := uuid.Parse(input.UserID)
	postId, _ := uuid.Parse(input.PostID)

	comment := &model.Comment{
		UserID:    userId,
		PostID:    postId,
		Content:   input.Content,
		Longitude: float32(input.Longitude),
		Latitude:  float32(input.Latitude),
		Location:  "Minneapolis",
	}

	_, err := r.CommentScheme.DB.NewInsert().Model(comment).Exec(ctx)
	// TODO: validate before output
	return comment, err
}

// ID is the resolver for the id field.
func (r *postResolver) ID(ctx context.Context, obj *model.Post) (string, error) {
	return obj.ID.String(), nil
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

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, postID string) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return obj.ID.String(), nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *postResolver) Comments(ctx context.Context, obj *model.Post) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) Karma(ctx context.Context, obj *model.User) (int, error) {
	return obj.Karma, nil
}
