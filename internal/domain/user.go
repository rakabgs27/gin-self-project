package domain

import "time"

// User adalah model database
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone     string    `json:"phone" gorm:"type:varchar(20)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest adalah DTO untuk membuat user baru
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=2"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone"`
}

// UpdateUserRequest adalah DTO untuk update user
type UpdateUserRequest struct {
	Name  string `json:"name" binding:"omitempty,min=2"`
	Email string `json:"email" binding:"omitempty,email"`
	Phone string `json:"phone"`
}
