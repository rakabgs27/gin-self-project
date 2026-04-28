package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rakabgs27/gin-self-project/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	AppPort string
	DB      *gorm.DB
}

func Load() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file tidak ditemukan, menggunakan environment variable sistem")
	}

	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		DB:      db,
	}, nil
}

func connectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASS", ""),
		getEnv("DB_HOST", "127.0.0.1"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "mydb"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("gagal konek database: %w", err)
	}

	// Auto migrate — otomatis buat/update tabel sesuai struct
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		return nil, fmt.Errorf("gagal auto migrate: %w", err)
	}

	log.Println("Database terhubung dan migration selesai")
	return db, nil
}

// getEnv mengambil env variable dengan fallback default value
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
