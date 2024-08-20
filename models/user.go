package models

type User struct {
	Id       int    `db:"rowid" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
}
