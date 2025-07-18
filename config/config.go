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

    // Auto migrate all models in one go
    err = db.AutoMigrate(
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
    )
    if err != nil {
        log.Fatal("Failed to auto-migrate tables: ", err)
    }

    DB = db
    fmt.Println("Connected to MySQL and Migrated")
}

