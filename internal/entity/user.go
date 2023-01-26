package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/FirdavsMF/wallet-api/internal/entity"
)

type userRepo struct {
	*sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db}
}

type UserRepo interface {
	GetByID(int) (entity.User, error)
	GetByLogin(string) (entity.User, error)
}

func (u *userRepo) GetByID(id int) (entity.User, error) {
	user := entity.User{}
	err := u.Get(&user, "SELECT * FROM \"Users\" WHERE id = $1", id)
	return user, err
}

func (u *userRepo) GetByLogin(login string) (entity.User, error) {
	user := entity.User{}
	err := u.Get(&user, "SELECT * FROM \"Users\" WHERE login = $1", login)
	return user, err
}
