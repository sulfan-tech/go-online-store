package mysql

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

// LoadDatabaseConfig loads the database configuration from environment variables
func LoadDatabaseConfig() *DatabaseConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file on load db, falling back to environment variables")
	}

	dbURL := os.Getenv("CLEARDB_DATABASE_URL")
	parsedURL, err := url.Parse(dbURL)
	if err != nil {
		log.Fatal("Error parsing CLEARDB_DATABASE_URL: ", err)
	}

	dbUser := parsedURL.User.Username()
	dbPass, _ := parsedURL.User.Password()
	dbHost := parsedURL.Hostname()
	dbPort := parsedURL.Port()
	dbName := parsedURL.Path[1:]

	return &DatabaseConfig{
		Username: dbUser,
		Password: dbPass,
		Host:     dbHost,
		Port:     dbPort,
		DBName:   dbName,
	}
}

// ConnectDatabase initializes the database connection using GORM
func ConnectDatabase() (*gorm.DB, error) {
	cfg := LoadDatabaseConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
