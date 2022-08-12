package model

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Comment struct {
	bun.BaseModel `bun:"table:comments,alias:c"`

	ID            uuid.UUID `bun:"id,type:uuid,pk,default:uuid_generate_v4()"`
	UserID        uuid.UUID `bun:"user_id,type:uuid"`
	postID        uuid.UUID `bun:"post_id,type:uuid"`
	Content       string    `bun:"content,notnull,unique"`
	IsFlagged     bool      `bun:"is_flagged,notnull,default:false"`
	NumberOfVotes int       `bun:"number_of_votes,notnull,default:0"`
	Location      string    `bun:"location,notnull"`
	Longitude     float32   `bun:"longitude,notnull"`
	Latitude      float32   `bun:"latitude,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
}

type CommentScheme struct {
	DB *bun.DB
}

func (c *CommentScheme) Create(comment *Comment, ctx context.Context) error {
	_, err := c.DB.NewInsert().Model(comment).Exec(ctx)
	return err
}
