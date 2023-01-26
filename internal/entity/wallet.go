package entity

import "time"

type Wallet struct {
	ID        int
	Login     string
	Sum       float64
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
