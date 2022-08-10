package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Post struct {
	bun.BaseModel `bun:"table:post,alias:p"`

	ID               uuid.UUID `bun:"id,type:uuid,pk,default:uuid_generate_v4()"`
	UserID           uuid.UUID `bun:"user_id,type:uuid"`
	Content          string    `bun:"content,notnull,unique"`
	IsFlagged        bool      `bun:"is_flagged,notnull,default: false"`
	NumberOfVotes    int       `bun:"number_of_votes,notnull, default:0"`
	NumberOfComments int       `bun:"number_of_comments,notnull, default:0"`
	Location         string    `bun:"location,notnull"`
	Longitude        float32   `bun:"longitude,notnull"`
	Latitude         float32   `bun:"latitude,notnull"`
	CreatedAt        time.Time `bun:"created_at,notnull, default:current_timestamp"`
}
