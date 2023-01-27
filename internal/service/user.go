package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/FirdavsMF/wallet-api/internal/entity"
)

type userSer struct {
	*sqlx.DB
}

func NewUserSer(db *sqlx.DB) *userSer {
	return &userSer{db}
}

type UserSer interface {
	GetByID(int) (entity.User, error)
	GetByLogin(string) (entity.User, error)
}

func (u *userSer) GetByID(id int) (entity.User, error) {
	user := entity.User{}
	err := u.Get(&user, "SELECT * FROM \"Users\" WHERE id = $1", id)
	return user, err
}

func (u *userSer) GetByLogin(login string) (entity.User, error) {
	user := entity.User{}
	err := u.Get(&user, "SELECT * FROM \"Users\" WHERE login = $1", login)
	return user, err
}
