package config

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "GO_LANG_PROJECT_SETUP/models"
)

var DB *gorm.DB

func ConnectDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to DB: ", err)
    }

    // 🔥 Auto migrate your models (creates tables if not exists)
    db.AutoMigrate(&models.User{})

    DB = db
    fmt.Println("Connected to MySQL and Migrated")
}

