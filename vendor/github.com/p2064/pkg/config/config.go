package config

import (
	"github.com/joho/godotenv"
	"github.com/p2064/pkg/logs"
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
		logs.ErrorLogger.Print("No .env file found")
	}
	Status = GOOD
}
