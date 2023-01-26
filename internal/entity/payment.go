package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/FirdavsMF/wallet-api/internal/entity"
)

type paymentRepo struct {
	*sqlx.DB
}

func NewPaymentRepo(db *sqlx.DB) *paymentRepo {
	return &paymentRepo{db}
}

type PaymentRepo interface {
	GetByExtID(*sqlx.Tx, string) (*entity.Payment, error)
	CreatePayment(*sqlx.Tx, string, string, float64, string) error
}

func (p *paymentRepo) GetByExtID(tx *sqlx.Tx, extID string) (*entity.Payment, error) {
	payment := &entity.Payment{}
	var err error
	if tx != nil {
		err = tx.Get(&payment, `SELECT * FROM "Payment" WHERE ext_id = $1`, extID)
		return payment, err
	}
	err = p.Get(&payment, `SELECT * FROM "Payment" WHERE ext_id = $1`, extID)

	return payment, err
}

func (p *paymentRepo) CreatePayment(tx *sqlx.Tx, src string, dest string, sum float64, extID string) error {
	now := time.Now()
	status := 200
	description := "Success"
	if tx != nil {
		_, err := tx.MustExec(`INSERT INTO "Payments" (src,dest,sum,created_at,updated_at, processed_at,status,description,ext_id) 
											VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
			src, dest, sum, now, now, now, status, description, extID).RowsAffected()
		return err

	}
	_, err := p.Exec(`INSERT INTO "Payments" (src,dest,sum,created_at,updated_at, processed_at,status,description,ext_id) 
						VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		src, dest, sum, now, now, now, status, description, extID)
	return err
}
