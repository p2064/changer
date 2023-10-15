package logs

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	//LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	//APP_NAME := os.Getenv("APP_NAME")
	LOG_FILE_LOCATION := "/Users/pavelprodanov/Desktop/logs/app.log"
	if _, err := os.Stat(LOG_FILE_LOCATION); os.IsNotExist(err) {
		file, err1 := os.Create(LOG_FILE_LOCATION)
		if err1 != nil {
			panic(err1)
		}
		InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		file, err := os.OpenFile(LOG_FILE_LOCATION, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

}

/* func CreateLogger(app_name string) *log.Logger {
	//LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	//APP_NAME := os.Getenv("APP_NAME")
	LOG_FILE_LOCATION := "/Users/pavelprodanov/Desktop/logs/" + app_name + ".log"
	if _, err := os.Stat(LOG_FILE_LOCATION); os.IsNotExist(err) {
		file, err1 := os.Create(LOG_FILE_LOCATION)
		if err1 != nil {
			panic(err1)
		}
		return log.New(file, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
		//WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		//ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		file, err := os.OpenFile(LOG_FILE_LOCATION, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		return log.New(file, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

} */
