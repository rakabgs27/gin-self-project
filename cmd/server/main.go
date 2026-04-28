package main

import (
	"log"

	"github.com/rakabgs27/gin-self-project/config"
	"github.com/rakabgs27/gin-self-project/internal/handler"
)

func main() {
	// Load config & koneksi database
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Gagal load config: %v", err)
	}

	// Setup router dengan dependency injection
	r := handler.NewRouter(cfg.DB)

	log.Printf("🚀 Server jalan di port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
