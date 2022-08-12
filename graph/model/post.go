package model

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Post struct {
	bun.BaseModel `bun:"table:posts,alias:p"`

	ID               uuid.UUID `bun:"id,type:uuid,pk,default:uuid_generate_v4()"`
	UserID           uuid.UUID `bun:"user_id,type:uuid"`
	Content          string    `bun:"content,notnull,unique"`
	IsFlagged        bool      `bun:"is_flagged,notnull,default:false"`
	NumberOfVotes    int       `bun:"number_of_votes,notnull,default:0"`
	NumberOfComments int       `bun:"number_of_comments,notnull,default:0"`
	Location         string    `bun:"location,notnull"`
	Longitude        float32   `bun:"longitude,notnull"`
	Latitude         float32   `bun:"latitude,notnull"`
	CreatedAt        time.Time `bun:"created_at,notnull,default:current_timestamp"`
}

type PostScheme struct {
	DB *bun.DB
}

func (p *PostScheme) Create(post *Post, ctx context.Context) error {
	_, err := p.DB.NewInsert().Model(post).Exec(ctx)
	return err
}

func (p *PostScheme) FindAll(ctx context.Context) (*[]Post, error) {
	var posts []Post
	_, err := p.DB.NewSelect().Model(&posts).Exec(ctx)
	return &posts, err
}
