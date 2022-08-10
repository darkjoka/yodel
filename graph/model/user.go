package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`

	ID       uuid.UUID `bun:"id,type:uuid,pk,default:uuid_generate_v4()"`
	Username string    `bun:"username,notnull,unique"`
	password string    `bun:"password,notnull"`
	karma    int       `bun:"karma,notnull,default:0"`
}
