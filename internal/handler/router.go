package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rakabgs27/gin-self-project/internal/repository"
	"github.com/rakabgs27/gin-self-project/internal/service"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong", "status": "ok"})
	})

	// Inisialisasi dependency (manual DI)
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userSvc)

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.POST("", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	return r
}
