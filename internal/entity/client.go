package entity

type Client struct {
	ID             int
	FirstName      string `db:"first_name"`
	LastName       string `db:"last_name"`
	Identified     bool
	WalletID       string `db:"wallet_id"`
	PassportNumber string `db:"passport_number"`
}
