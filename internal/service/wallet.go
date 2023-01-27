package service

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/FirdavsMF/wallet-api/internal/entity"
)

type walletSer struct {
	*sqlx.DB
}

func NewWalletSer(db *sqlx.DB) *walletSer {
	return &walletSer{db}
}

type WalletHistory struct {
	Count int
	Sum   float64
}

type WalletSer interface {
	GetByID(int) (entity.Wallet, error)
	GetByLogin(string, *sqlx.Tx) (entity.Wallet, error)
	ChargeWallet(*sqlx.Tx, string, float64) error
	GetWalletHistory(string) (WalletHistory, error)
}

func (w *walletSer) GetByID(id int) (entity.Wallet, error) {
	wallet := entity.Wallet{}
	err := w.Get(&wallet, "SELECT * FROM \"Wallets\" WHERE id = '$1'", id)
	return wallet, err
}

func (w *walletSer) GetByLogin(login string, tx *sqlx.Tx) (entity.Wallet, error) {
	wallet := entity.Wallet{}
	var err error
	if tx != nil {
		err = tx.Get(&wallet, "SELECT * FROM \"Wallets\" WHERE login = $1", login)

		return wallet, err
	}

	err = w.Get(&wallet, "SELECT * FROM \"Wallets\" WHERE login = $1", login)
	return wallet, err
}

func (w *walletSer) ChargeWallet(tx *sqlx.Tx, login string, sum float64) error {
	var err error
	if tx != nil {
		_, err = tx.MustExec(`UPDATE "Wallets" SET sum=sum+$1 WHERE login = $2`, sum, login).RowsAffected()
		return err
	}

	_, err = w.MustExec(`UPDATE "Wallets" SET sum=sum+$1 WHERE login = $2`, sum, login).RowsAffected()
	return err
}

func (w *walletSer) GetWalletHistory(login string) (WalletHistory, error) {
	walletHistory := WalletHistory{}
	err := w.Get(&walletHistory, `SELECT sum(sum),count(*) FROM "Payments" WHERE src=$1 or dest=$1   
	  and created_at BETWEEN  date_trunc('month', current_date)::date and  (date_trunc('month', now()) + interval '1 month - 1 day')::date
	`, login)
	return walletHistory, err
}
