package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type Config struct {
    AppPort string
    DB      *gorm.DB
}

func Load() (*Config, error) {
    godotenv.Load()

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("gagal konek database: %w", err)
    }

    return &Config{
        AppPort: os.Getenv("APP_PORT"),
        DB:      db,
    }, nil
}