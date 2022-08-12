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

// OrderIndex is the resolver for the orderIndex field.
func (r *commentResolver) OrderIndex(ctx context.Context, obj *model.Comment) (int, error) {
	commentor := new(model.Commentor)
	query := r.CommentorScheme.DB.NewSelect().Model(commentor).Where("user_id = ? AND post_id = ?", obj.UserID, obj.PostID)
	if err := query.Scan(ctx); err == nil {
		return commentor.IncrementID, nil
	}

	// post author is commentor
	return 0, nil
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (*model.User, error) {
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
	// TODO: after jwt integration remove user id from input and refine author check
	userId, _ := uuid.Parse(input.UserID)
	postId, _ := uuid.Parse(input.PostID)

	// get post author and check if the same as commentor
	post := new(model.Post)
	r.PostScheme.DB.NewSelect().Model(post).Where("id = ?", postId).Scan(ctx)
	// not author
	if post.UserID != userId {
		commentor := new(model.Commentor)
		query := r.CommentorScheme.DB.NewSelect().Model(commentor).Where("user_id = ? AND post_id = ?", userId, postId)
		if err := query.Scan(ctx); err != nil {
			newCommentor := &model.Commentor{
				UserID: userId,
				PostID: postId,
			}
			r.CommentScheme.DB.NewInsert().Model(newCommentor).Exec(ctx)
		}
	}

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
	return obj.UserID.String(), nil
}

// IsFlagged is the resolver for the isFlagged field.
func (r *postResolver) IsFlagged(ctx context.Context, obj *model.Post) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Longitude is the resolver for the longitude field.
func (r *postResolver) Longitude(ctx context.Context, obj *model.Post) (float64, error) {
	return float64(obj.Latitude), nil
}

// Latitude is the resolver for the latitude field.
func (r *postResolver) Latitude(ctx context.Context, obj *model.Post) (float64, error) {
	return float64(obj.Latitude), nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var posts []*model.Post
	err := r.PostScheme.DB.NewSelect().Model(&posts).Scan(ctx)
	return posts, err
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	post := new(model.Post)
	err := r.PostScheme.DB.NewSelect().Model(post).Where("id = ?", id).Scan(ctx)
	return post, err
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, postID string) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := r.CommentScheme.DB.NewSelect().Model(&comments).Where("post_id = ?", postID).Scan(ctx)
	return comments, err
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
