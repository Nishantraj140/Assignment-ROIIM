package logger

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DBLogger    *log.Logger
	File        *os.File
)

func init() {
	var err error
	File, err = os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	DBFile, err := os.OpenFile("dbLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	InfoLogger = log.New(File, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(File, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DBLogger = log.New(DBFile, "DBLOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
