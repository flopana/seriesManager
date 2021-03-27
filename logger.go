package main

import (
	"log"
	"os"
)

func GetLogger() *log.Logger {
	f, err := os.OpenFile("logfile.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	logger := log.New(f, "", log.Ldate|log.Ltime|log.Lshortfile)

	return logger
}
