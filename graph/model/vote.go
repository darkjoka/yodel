package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Vote struct {
	bun.BaseModel `bun:"table:votes,alias:v"`

	ID       uuid.UUID `bun:"id,type:uuid,pk,default:uuid_generate_v4()"`
	UserID   uuid.UUID `bun:"user_id,type:uuid"`
	postID   uuid.UUID `bun:"post_id,type:uuid"`
	IsUpvote bool      `bun:"is_upvote,notnull"`
}
