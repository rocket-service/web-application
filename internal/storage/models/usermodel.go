package models

type User struct {
	ID           int    `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"passwordhash"`
	IsAdmin      bool   `db:"isadmin"`
	IsBanned     bool   `db:"isbanned"`
	BanReason    string `db:"bannedreason"`
}
