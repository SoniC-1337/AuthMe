package handlers

import (
	"database/sql"
	"github.com/Xoro-1337/AuthMe/api/models"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	LicenseID int    `json:"license_id" db:"license_id"`
}

func Authenticate(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		var product Product

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"status": "error",
				"error":  "Failed to bind JSON data",
			})
			return
		}

		stmt := "SELECT p.* FROM products p JOIN licenses l ON l.id = p.license_id JOIN users u ON l.user_id = u.id WHERE u.uid = ? AND l.expiration >= NOW();"
		if err := db.QueryRow(stmt, user.UID).Scan(&product.ID, &product.Name, &product.LicenseID); err != nil {
			if err == sql.ErrNoRows {
				c.JSON(404, gin.H{
					"status": "error",
					"error":  "UID not found",
				})
				return
			} else {
				c.JSON(403, gin.H{
					"status": "error",
					"error":  "License expired",
				})
				return
			}
		}

		c.JSON(200, gin.H{
			"status": "success",
			"product": gin.H{
				"id":        product.ID,
				"name":      product.Name,
				"licenseID": product.LicenseID,
			},
		})
	}
}
