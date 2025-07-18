package logger

import (
	"log"
)

var (
	InfoLogger  = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
) 