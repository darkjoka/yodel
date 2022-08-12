package model

import (
	"context"

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

type UserScheme struct {
	DB *bun.DB
}

func (u *UserScheme) Create(user *User, ctx context.Context) error {
	_, err := u.DB.NewInsert().Model(user).Exec(ctx)
	return err
}

func (u *UserScheme) FindById(id string, ctx context.Context) (*User, error) {
	user := new(User)
	err := u.DB.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	return user, err
}
