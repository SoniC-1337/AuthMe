package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	License string  `form:"license" binding:"required"`
	Hwid    string  `form:"hwid" binding:"required"`
	Product *string `form:"product"`
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
	router.POST("/login", login(db))
	router.GET("/download", downloadHandler)

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

func login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		// Fetch user information from the database
		query := "SELECT hwid FROM users WHERE license = ?"
		var hwid string
		err = db.QueryRow(query, user.License).Scan(&hwid)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(401, gin.H{
					"status": "failed",
					"error":  "User not found",
				})
				return
			}
			log.Println(err)
			c.JSON(500, gin.H{
				"status": "error",
				"error":  "Database error",
			})
			return
		}

		// Compare the hashed password
		err = bcrypt.CompareHashAndPassword([]byte(hwid), []byte(user.Hwid))
		if err != nil {
			c.JSON(401, gin.H{
				"status": "failed",
				"error":  "Invalid credentials",
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "success",
		})
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
