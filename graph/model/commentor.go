package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Commentor struct {
	bun.BaseModel `bun:"table:commentor,alias:cr"`

	ID          uuid.UUID `bun:"id,type:uuid,pk,default:uuid_generate_v4()"`
	UserID      uuid.UUID `bun:"user_id,type:uuid,unique"`
	postID      uuid.UUID `bun:"post_id,type:uuid"`
	IncrementID int       `bun:"increment_id,autoincrement"`
}
