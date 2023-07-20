package models

import "time"

type License struct {
	ID         int       `json:"id" db:"id"`
	License    string    `json:"license" db:"license"`
	Redeemed   bool      `json:"redeemed" db:"redeemed"`
	Expiration time.Time `json:"expiration" db:"expiration"`
	UserID     int       `json:"user_id" db:"user_id"`
}
