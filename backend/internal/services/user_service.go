package services

import (
	"database/sql"
	"errors"

	"mvp_multylink/backend/internal/models"
	"mvp_multylink/backend/internal/repository"
)

// UserService предоставляет методы для работы с пользователями
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService создает новый экземпляр UserService
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// CreateUser создает нового пользователя
func (s *UserService) CreateUser(user models.User) (int64, error) {
	return s.userRepo.CreateUser(user)
}

// GetUserByID получает пользователя по ID
func (s *UserService) GetUserByID(id int64) (models.User, error) {
	return s.userRepo.GetUserByID(id)
}

// GetUserByEmail получает пользователя по email
func (s *UserService) GetUserByEmail(email string) (models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

// GetUserByUsername получает пользователя по имени пользователя
func (s *UserService) GetUserByUsername(username string) (models.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

// UpdateUser обновляет данные пользователя
func (s *UserService) UpdateUser(user models.User) error {
	return s.userRepo.UpdateUser(user)
}

// CheckUserExists проверяет существование пользователя с указанным email или username
func (s *UserService) CheckUserExists(email, username string) (bool, error) {
	// Проверка по email
	_, err := s.userRepo.GetUserByEmail(email)
	if err == nil {
		return true, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	// Проверка по username
	_, err = s.userRepo.GetUserByUsername(username)
	if err == nil {
		return true, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	return false, nil
}

// GetUserProfile получает профиль пользователя
func (s *UserService) GetUserProfile(username string) (models.UserProfile, error) {
	return s.userRepo.GetUserProfile(username)
}

// UpdateUserProfile обновляет профиль пользователя
func (s *UserService) UpdateProfile(user models.User) error {
	return s.userRepo.UpdateProfile(user)
}
