package db

import (
	"fmt"
	"log"

	"github.com/p2064/pkg/config"
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
			Host:     "host.docker.internal",
			Port:     "5432",
			User:     "postgres",
			Password: "",
			DBName:   "postgres",
			SSLMode:  "disable",
		}
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
		var err error

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Print("Database not connected")
			panic(err)
		}

		log.Print("Database connected")
	} else {
		log.Print("Env file failure")
	}

}
