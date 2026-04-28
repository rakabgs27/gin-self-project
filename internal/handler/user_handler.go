package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rakabgs27/gin-self-project/internal/domain"
	"github.com/rakabgs27/gin-self-project/internal/service"
	"github.com/rakabgs27/gin-self-project/internal/repository"
	"github.com/rakabgs27/gin-self-project/pkg/response"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Tambahkan fungsi ini di file yang sama
func RegisterUserHandlers(r *gin.RouterGroup, db *gorm.DB) {
    // Inisialisasi internal modul user
    userRepo := repository.NewUserRepository(db)
    userSvc := service.NewUserService(userRepo)
    h := NewUserHandler(userSvc)

    // Definisi route untuk modul user
    users := r.Group("/users")
    {
        users.GET("", h.GetAllUsers)
        users.GET("/:id", h.GetUserByID)
        users.POST("", h.CreateUser)
        users.PUT("/:id", h.UpdateUser)
        users.DELETE("/:id", h.DeleteUser)
    }
}

// GetAllUsers godoc
// GET /api/v1/users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		response.InternalError(c, "Gagal mengambil data user")
		return
	}
	response.OK(c, "Berhasil mengambil data user", users)
}

// GetUserByID godoc
// GET /api/v1/users/:id
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "ID tidak valid")
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.OK(c, "Berhasil mengambil data user", user)
}

// CreateUser godoc
// POST /api/v1/users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req domain.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, err := h.userService.CreateUser(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Created(c, "User berhasil dibuat", user)
}

// UpdateUser godoc
// PUT /api/v1/users/:id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "ID tidak valid")
		return
	}

	var req domain.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, err := h.userService.UpdateUser(uint(id), &req)
	if err != nil {
		if err.Error() == "user tidak ditemukan" {
			response.NotFound(c, err.Error())
			return
		}
		response.BadRequest(c, err.Error())
		return
	}

	response.OK(c, "User berhasil diupdate", user)
}

// DeleteUser godoc
// DELETE /api/v1/users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "ID tidak valid")
		return
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		if err.Error() == "user tidak ditemukan" {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "Gagal menghapus user")
		return
	}

	response.OK(c, "User berhasil dihapus", nil)
}
