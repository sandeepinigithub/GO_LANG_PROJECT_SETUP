package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// LogLevel represents the logging level
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// LogEntry represents a structured log entry
type LogEntry struct {
	Timestamp string      `json:"timestamp"`
	Level     string      `json:"level"`
	Message   string      `json:"message"`
	Service   string      `json:"service"`
	TraceID   string      `json:"trace_id,omitempty"`
	UserID    string      `json:"user_id,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
}

// Logger interface for different logging implementations
type Logger interface {
	Debug(message string, data ...interface{})
	Info(message string, data ...interface{})
	Warn(message string, data ...interface{})
	Error(message string, err error, data ...interface{})
	Fatal(message string, err error, data ...interface{})
}

// JSONLogger implements structured JSON logging
type JSONLogger struct {
	service string
	level   LogLevel
}

// ConsoleLogger implements console logging
type ConsoleLogger struct {
	service string
	level   LogLevel
}

var (
	// Global logger instance
	GlobalLogger Logger
	// Default loggers for backward compatibility
	InfoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

// InitLogger initializes the global logger
func InitLogger(service, level, format string) {
	var logLevel LogLevel
	switch level {
	case "debug":
		logLevel = DEBUG
	case "info":
		logLevel = INFO
	case "warn":
		logLevel = WARN
	case "error":
		logLevel = ERROR
	case "fatal":
		logLevel = FATAL
	default:
		logLevel = INFO
	}

	if format == "json" {
		GlobalLogger = NewJSONLogger(service, logLevel)
	} else {
		GlobalLogger = NewConsoleLogger(service, logLevel)
	}
}

// NewJSONLogger creates a new JSON logger
func NewJSONLogger(service string, level LogLevel) *JSONLogger {
	return &JSONLogger{
		service: service,
		level:   level,
	}
}

// NewConsoleLogger creates a new console logger
func NewConsoleLogger(service string, level LogLevel) *ConsoleLogger {
	return &ConsoleLogger{
		service: service,
		level:   level,
	}
}

// JSONLogger methods
func (l *JSONLogger) Debug(message string, data ...interface{}) {
	if l.level <= DEBUG {
		l.log("DEBUG", message, nil, data...)
	}
}

func (l *JSONLogger) Info(message string, data ...interface{}) {
	if l.level <= INFO {
		l.log("INFO", message, nil, data...)
	}
}

func (l *JSONLogger) Warn(message string, data ...interface{}) {
	if l.level <= WARN {
		l.log("WARN", message, nil, data...)
	}
}

func (l *JSONLogger) Error(message string, err error, data ...interface{}) {
	if l.level <= ERROR {
		l.log("ERROR", message, err, data...)
	}
}

func (l *JSONLogger) Fatal(message string, err error, data ...interface{}) {
	if l.level <= FATAL {
		l.log("FATAL", message, err, data...)
		os.Exit(1)
	}
}

func (l *JSONLogger) log(level, message string, err error, data ...interface{}) {
	entry := LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     level,
		Message:   message,
		Service:   l.service,
	}

	if err != nil {
		entry.Error = err.Error()
	}

	if len(data) > 0 {
		entry.Data = data[0]
	}

	jsonData, _ := json.Marshal(entry)
	fmt.Println(string(jsonData))
}

// ConsoleLogger methods
func (l *ConsoleLogger) Debug(message string, data ...interface{}) {
	if l.level <= DEBUG {
		l.log("DEBUG", message, nil, data...)
	}
}

func (l *ConsoleLogger) Info(message string, data ...interface{}) {
	if l.level <= INFO {
		l.log("INFO", message, nil, data...)
	}
}

func (l *ConsoleLogger) Warn(message string, data ...interface{}) {
	if l.level <= WARN {
		l.log("WARN", message, nil, data...)
	}
}

func (l *ConsoleLogger) Error(message string, err error, data ...interface{}) {
	if l.level <= ERROR {
		l.log("ERROR", message, err, data...)
	}
}

func (l *ConsoleLogger) Fatal(message string, err error, data ...interface{}) {
	if l.level <= FATAL {
		l.log("FATAL", message, err, data...)
		os.Exit(1)
	}
}

func (l *ConsoleLogger) log(level, message string, err error, data ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] %s [%s] %s", timestamp, level, l.service, message)
	
	if err != nil {
		logMessage += fmt.Sprintf(" - Error: %v", err)
	}
	
	if len(data) > 0 {
		logMessage += fmt.Sprintf(" - Data: %+v", data[0])
	}
	
	fmt.Println(logMessage)
}

// Convenience functions for backward compatibility
func Info(message string) {
	if GlobalLogger != nil {
		GlobalLogger.Info(message)
	} else {
		InfoLogger.Println(message)
	}
}

func Error(message string, err error) {
	if GlobalLogger != nil {
		GlobalLogger.Error(message, err)
	} else {
		ErrorLogger.Printf("%s: %v", message, err)
	}
} 