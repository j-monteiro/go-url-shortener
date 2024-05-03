package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func GetDBConnection() *sql.DB {
	c := getDBConfig()

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.DbName,
		c.Password,
		fetchSslMode(),
	)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("Error pinging database")
	}

	return conn
}

func getDBConfig() DbConfig {
	return DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}
}

func fetchSslMode() string {
	env := os.Getenv("ENV")

	switch env {
	case "development":
		return "disable"
	default:
		return "verify-full"
	}
}
