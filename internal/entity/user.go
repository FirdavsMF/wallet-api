package entity

type User struct {
	ID        int
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Login     string
	Password  string
	Salt      string
	PartnerID string `db:"partner_id"`
	WalletID  string `db:"wallet_id"`
}
