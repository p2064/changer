package config

import (
	"log"

	"github.com/joho/godotenv"
)

type DBStatus int64

const (
	GOOD DBStatus = iota
	ERROR
)

var Status DBStatus

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		Status = ERROR
		log.Print("No .env file found")
	}
	Status = GOOD
}
