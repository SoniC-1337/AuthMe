package models

type User struct {
	ID    int    `json:"id" db:"id"`
	UID   string `json:"uid" db:"uid"`
	Reset bool   `json:"reset" db:"reset"`
}
