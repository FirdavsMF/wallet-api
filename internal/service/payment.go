package service

import (
	"time"

	"github.com/FirdavsMF/wallet-api/internal/entity"
	"github.com/jmoiron/sqlx"
)

type paymentSer struct {
	*sqlx.DB
}

func NewPaymentSer(db *sqlx.DB) *paymentSer {
	return &paymentSer{db}
}

type PaymentSer interface {
	GetByExtID(*sqlx.Tx, string) (*entity.Payment, error)
	CreatePayment(*sqlx.Tx, string, string, float64, string) error
}

func (p *paymentSer) GetByExtID(tx *sqlx.Tx, extID string) (*entity.Payment, error) {
	payment := &entity.Payment{}
	var err error
	if tx != nil {
		err = tx.Get(&payment, `SELECT * FROM "Payment" WHERE ext_id = $1`, extID)
		return payment, err
	}
	err = p.Get(&payment, `SELECT * FROM "Payment" WHERE ext_id = $1`, extID)

	return payment, err
}

func (p *paymentSer) CreatePayment(tx *sqlx.Tx, src string, dest string, sum float64, extID string) error {
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
