package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Commentor struct {
	bun.BaseModel `bun:"table:commentors,alias:cr"`

	ID          uuid.UUID `bun:"id,type:uuid,pk,default:uuid_generate_v4()"`
	UserID      uuid.UUID `bun:"user_id,type:uuid,unique"`
	postID      uuid.UUID `bun:"post_id,type:uuid"`
	IncrementID int       `bun:"increment_id,autoincrement"`
}

type CommentorScheme struct {
	DB *bun.DB
}

func (c *CommentorScheme) Create(commentor *Commentor, ctx context.Context) error {
	_, err := c.DB.NewInsert().Model(commentor).Exec(ctx)
	return err
}
