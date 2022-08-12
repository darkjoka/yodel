package model

import (
	"context"

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

type VoteScheme struct {
	DB *bun.DB
}

func (v *VoteScheme) Create(vote *Vote, ctx context.Context) error {
	_, err := v.DB.NewInsert().Model(vote).Exec(ctx)
	return err
}

func (v *VoteScheme) Delete(vote *Vote, ctx context.Context) error {
	_, err := v.DB.NewDelete().Model(vote).WherePK().Exec(ctx)
	return err
}

func (v *VoteScheme) FindOne(user, post uuid.UUID, ctx context.Context) error {
	vote := new(Vote)
	err := v.DB.NewSelect().Model(vote).Where("user_id = ? AND post_id = ?", user, post).Scan(ctx)
	return err
}
