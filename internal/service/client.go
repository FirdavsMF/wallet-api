package service

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/FirdavsMF/wallet-api/internal/entity"
)

type clientSer struct {
	*sqlx.DB
}

func NewClientSer(db *sqlx.DB) *clientSer {
	return &clientSer{db}
}

type ClientSer interface {
	GetByWalletID(int, *sqlx.Tx) (entity.Client, error)
}

func (c *clientSer) GetByWalletID(id int, tx *sqlx.Tx) (entity.Client, error) {
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
