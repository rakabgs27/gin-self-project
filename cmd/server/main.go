package main

import (
    "log"

    "github.com/rakabgs27/gin-self-project/config"
    "github.com/rakabgs27/gin-self-project/internal/handler"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }

    r := handler.NewRouter()

    log.Printf("Server jalan di port %s", cfg.AppPort)
    r.Run(":" + cfg.AppPort)
}