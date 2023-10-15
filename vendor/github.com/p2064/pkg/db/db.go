package db

import (
	"fmt"
	"os"

	"github.com/p2064/pkg/config"
	"github.com/p2064/pkg/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var (
	DB *gorm.DB
)

func init() {
	if config.Status == config.GOOD {
		cfg := Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		}
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
		var err error

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logs.ErrorLogger.Print("Database not connected")
			panic(err)
		}

		logs.InfoLogger.Print("Database connected")
	} else {
		logs.ErrorLogger.Print("Env file failure")
	}

}
