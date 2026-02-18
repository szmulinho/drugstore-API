package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadFromEnv() StorageConfig {
	_ = godotenv.Load(".env")
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" {
		fmt.Println("Using DATABASE_URL")
		return StorageConfig{
			DatabaseURL: databaseURL,
		}
	}
	fmt.Println("Using individual DB_* variables")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "require"
	}
	fmt.Printf("Host: %s\n", host)
	fmt.Printf("User: %s\n", user)
	fmt.Printf("DBName: %s\n", dbname)
	fmt.Printf("Port: %s\n", port)
	return StorageConfig{
		Host:     host,
		User:     user,
		Password: password,
		Dbname:   dbname,
		Port:     port,
		SSLMode:  sslmode,
	}
}

type StorageConfig struct {
	DatabaseURL string `json:"database_url,omitempty"`
	Host        string `json:"host"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Dbname      string `json:"dbname"`
	Port        string `json:"port"`
	SSLMode     string `json:"sslmode"`
}

func (c StorageConfig) ConnectionString() string {
	if c.DatabaseURL != "" {
		return c.DatabaseURL
	}
	sslmode := c.SSLMode
	if sslmode == "" {
		sslmode = "require"
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Host, c.User, c.Password, c.Dbname, c.Port, sslmode)
}
