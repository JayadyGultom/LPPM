package config

import (
    "fmt"
    "log"
    "os"

    "perpustakaan/internal/entity"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type Database struct {
    DB *gorm.DB
}

var dbInstance *Database

func InitDB() {
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        dsn = "root:@tcp(localhost:3306)/lppm?charset=utf8mb4&parseTime=True&loc=Local"
    }

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto Migrate - only for profile tables since others already exist
    err = db.AutoMigrate(
        &entity.VisiMisi{},
        &entity.StrukturOrganisasi{},
    )
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    dbInstance = &Database{DB: db}
    fmt.Println("Database connected successfully")
}

func GetDB() *Database {
    return dbInstance
}
