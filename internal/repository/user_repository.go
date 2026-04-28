package repository

import (
	"github.com/rakabgs27/gin-self-project/internal/domain"
	"gorm.io/gorm"
)

// UserRepository mendefinisikan kontrak operasi database untuk User
type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindByID(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository membuat instance baru UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Order("created_at desc").Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}
