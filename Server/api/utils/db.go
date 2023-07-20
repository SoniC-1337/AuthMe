package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func ConnectToDB(config DBConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func GetDBConfigFromEnv() DBConfig {
	return DBConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}
}
