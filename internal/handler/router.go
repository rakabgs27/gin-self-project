package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong", "status": "ok"})
	})


	// API v1 routes
    v1 := r.Group("/api/v1")
    {
        // Cukup panggil fungsinya saja
        RegisterUserHandlers(v1, db)
        
        // Nanti kalau ada modul Order, tinggal:
        // RegisterOrderHandlers(v1, db)
    }

	return r
}
