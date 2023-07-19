package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int    `json:"id" db:"id"`
	UID   string `json:"uid" db:"uid"`
	Reset bool   `json:"reset" db:"reset"`
}

type License struct {
	ID         int       `json:"id" db:"id"`
	License    string    `json:"license" db:"license"`
	Redeemed   bool      `json:"redeemed" db:"redeemed"`
	Expiration time.Time `json:"expiration" db:"expiration"`
	UserID     int       `json:"user_id" db:"user_id"`
}

type Product struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	LicenseID int    `json:"license_id" db:"license_id"`
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func main() {
	dbConfig := DBConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}

	db, err := connectToDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.POST("/login", authenticate(db))
	router.GET("/download", downloadHandler)
	router.POST("/register", registerHandler(db))

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func connectToDB(config DBConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func authenticate(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		var product Product
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  "Failed to bind form data",
			})
			return
		}

		err = db.QueryRow("SELECT p.* FROM products p JOIN licenses l ON l.id = p.license_id JOIN users u ON l.user_id = u.id WHERE u.uid = ? AND l.expiration >= NOW();", user.UID).Scan(&product.ID, &product.Name, &product.LicenseID)
		if err != nil {
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

func registerHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  "Failed to bind form data",
			})
			return
		}

		var existingUser User
		err = db.QueryRow("SELECT * FROM users WHERE uid = ?;", user.UID).Scan(&existingUser.ID, &existingUser.UID, &existingUser.Reset)
		if err != nil {
			if err != sql.ErrNoRows {
				c.JSON(500, gin.H{
					"status": "error",
					"error":  "Failed to query database",
				})
				return
			}
		}

		if existingUser.ID == 0 {
			res, err := db.Exec("INSERT INTO users (uid) VALUES (?);", user.UID)
			if err != nil {
				c.JSON(500, gin.H{
					"status": "error",
					"error":  "Failed to create new user",
				})
				return
			}
			userID, err := res.LastInsertId()
			if err != nil {
				c.JSON(500, gin.H{
					"status": "error",
					"error":  "Failed to create new user",
				})
				return
			}
			c.JSON(200, gin.H{
				"status": "success",
				"user": gin.H{
					"id":  userID,
					"uid": user.UID,
				},
			})
		} else {
			c.JSON(200, gin.H{
				"error": "The UID already exists",
			})
		}
	}
}

func downloadHandler(c *gin.Context) {
	filePath := "file.bin"
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"status": "error",
			"error":  "Failed to open file",
		})
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=file.bin")

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		log.Println(err)
	}
}
