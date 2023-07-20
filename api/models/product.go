package models

type Product struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	LicenseID int    `json:"license_id" db:"license_id"`
}
