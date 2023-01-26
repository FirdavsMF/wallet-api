package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/FirdavsMF/wallet-api/internal/entity"
)

type clientRepo struct {
	*sqlx.DB
}

func NewClientRepo(db *sqlx.DB) *clientRepo {
	return &clientRepo{db}
}

type ClientRepo interface {
	GetByWalletID(int, *sqlx.Tx) (entity.Client, error)
}

func (c *clientRepo) GetByWalletID(id int, tx *sqlx.Tx) (entity.Client, error) {
	client := entity.Client{}
	var err error
	if tx != nil {

		err = tx.Get(&client, `SELECT * FROM "Clients" WHERE wallet_id = $1`, id)
		fmt.Println(err)
		return client, err
	}
	err = c.Get(&client, "SELECT * FROM \"Clients\" WHERE wallet_id = '$1'", id)

	return client, err
}
