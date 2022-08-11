package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       uuid.UUID `bun:"id,type:uuid,pk,default:uuid_generate_v4()"`
	Username string    `bun:"username,notnull,unique"`
	Password string    `bun:"password,notnull"`
	Karma    int       `bun:"karma,notnull,default:0"`
}

// type UserScheme struct {
// 	db *bun.DB
// }

// func (u *UserScheme) {
// 	u.dbUser()
// }
