package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"devsMailGo/models"
)

// Config holds all configuration
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
	Security SecurityConfig
	Logging  LoggingConfig
	Redis    RedisConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	DSN      string
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port        string
	Host        string
	Environment string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret      string
	ExpiryHours int
}

// SecurityConfig holds security configuration
type SecurityConfig struct {
	CORSAllowedOrigins string
	CORSAllowedMethods string
	CORSAllowedHeaders string
	RateLimitRequests  int
	RateLimitWindow    time.Duration
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string
	Format string
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

var (
	DB     *gorm.DB
	AppConfig *Config
)

// LoadConfig loads configuration from environment variables
func LoadConfig() error {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	AppConfig = &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "devsmailgo"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "devsmailgo"),
		},
		Server: ServerConfig{
			Port:        getEnv("SERVER_PORT", "8080"),
			Host:        getEnv("SERVER_HOST", "0.0.0.0"),
			Environment: getEnv("ENVIRONMENT", "development"),
		},
		JWT: JWTConfig{
			Secret:      getEnv("JWT_SECRET", "your-super-secret-jwt-key-change-in-production"),
			ExpiryHours: getEnvAsInt("JWT_EXPIRY_HOURS", 24),
		},
		Security: SecurityConfig{
			CORSAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "*"),
			CORSAllowedMethods: getEnv("CORS_ALLOWED_METHODS", "GET,POST,PUT,DELETE,OPTIONS"),
			CORSAllowedHeaders: getEnv("CORS_ALLOWED_HEADERS", "Content-Type,Authorization,X-Requested-With"),
			RateLimitRequests:  getEnvAsInt("RATE_LIMIT_REQUESTS", 100),
			RateLimitWindow:    time.Duration(getEnvAsInt("RATE_LIMIT_WINDOW_MINUTES", 1)) * time.Minute,
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
	}

	// Build DSN
	AppConfig.Database.DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.Host,
		AppConfig.Database.Port,
		AppConfig.Database.Name,
	)

	return validateConfig()
}

// InitDB initializes database connection
func InitDB() error {
	if AppConfig == nil {
		if err := LoadConfig(); err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
	}

	// Configure GORM logger based on environment
	var gormLogger logger.Interface
	if AppConfig.Server.Environment == "production" {
		gormLogger = logger.Default.LogMode(logger.Error)
	} else {
		gormLogger = logger.Default.LogMode(logger.Info)
	}

	// Open database connection
	db, err := gorm.Open(mysql.Open(AppConfig.Database.DSN), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Auto migrate models
	if err := autoMigrate(db); err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}

	DB = db
	log.Printf("Connected to MySQL database: %s@%s:%s/%s", 
		AppConfig.Database.User, 
		AppConfig.Database.Host, 
		AppConfig.Database.Port, 
		AppConfig.Database.Name)

	return nil
}

// autoMigrate performs database migrations
func autoMigrate(db *gorm.DB) error {
	models := []interface{}{
		&models.User{},
		&models.Domain{},
		&models.DomainAdmin{},
		&models.Alias{},
		&models.MailingList{},
		&models.SpamPolicy{},
		&models.Throttling{},
		&models.Wblist{},
		&models.Greylisting{},
		&models.GreylistingTracking{},
		&models.GreylistingWhitelistDomainSPF{},
		&models.GreylistingWhitelistDomain{},
		&models.GreylistingWhitelist{},
		&models.SenderscoreCache{},
		&models.SMTPSession{},
		&models.SRSExcludeDomain{},
		&models.Throttle{},
		&models.ThrottleTracking{},
		&models.WblistRDNS{},
		&models.Banned{},
		&models.Jail{},
		&models.LastLogin{},
		&models.Log{},
		&models.NewsletterSubunsubConfirm{},
		&models.Session{},
		&models.Setting{},
		&models.ShareFolder{},
		&models.Tracking{},
		&models.UpdateLog{},
		&models.UsedQuota{},
		&models.RoundcubeUser{},
	}

	return db.AutoMigrate(models...)
}

// validateConfig validates configuration
func validateConfig() error {
	if AppConfig.Database.Password == "" {
		return fmt.Errorf("database password is required")
	}
	if AppConfig.JWT.Secret == "your-super-secret-jwt-key-change-in-production" {
		log.Println("WARNING: Using default JWT secret. Change this in production!")
	}
	return nil
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// Legacy function for backward compatibility
func ConnectDB() {
	if err := InitDB(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

