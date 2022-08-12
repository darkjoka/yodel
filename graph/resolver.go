package graph

import "github.com/darkjoka/yodel/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserScheme    model.UserScheme
	PostScheme    model.PostScheme
	CommentScheme model.CommentScheme
	VoteScheme    model.VoteScheme
}
