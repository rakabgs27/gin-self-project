package service

import (
	"errors"

	"github.com/rakabgs27/gin-self-project/internal/domain"
	"github.com/rakabgs27/gin-self-project/internal/repository"
	"gorm.io/gorm"
)

// UserService mendefinisikan kontrak business logic untuk User
type UserService interface {
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	CreateUser(req *domain.CreateUserRequest) (*domain.User, error)
	UpdateUser(id uint, req *domain.UpdateUserRequest) (*domain.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

// NewUserService membuat instance baru UserService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}
	return user, nil
}

func (s *userService) CreateUser(req *domain.CreateUserRequest) (*domain.User, error) {
	// Cek apakah email sudah dipakai
	existing, _ := s.repo.FindByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("email sudah digunakan")
	}

	user := &domain.User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(id uint, req *domain.UpdateUserRequest) (*domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}

	// Cek email duplikat kalau email diubah
	if req.Email != "" && req.Email != user.Email {
		existing, _ := s.repo.FindByEmail(req.Email)
		if existing != nil {
			return nil, errors.New("email sudah digunakan")
		}
		user.Email = req.Email
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user tidak ditemukan")
		}
		return err
	}
	return s.repo.Delete(id)
}
