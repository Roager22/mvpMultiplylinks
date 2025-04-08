package repository

import (
	"mvp_multylink/backend/internal/models"
)

// UserRepository определяет интерфейс для работы с пользователями в базе данных
type UserRepository interface {
	// CreateUser создает нового пользователя и возвращает его ID
	CreateUser(user models.User) (int64, error)

	// GetUserByID получает пользователя по ID
	GetUserByID(id int64) (models.User, error)

	// GetUserByEmail получает пользователя по email
	GetUserByEmail(email string) (models.User, error)

	// GetUserByUsername получает пользователя по имени пользователя
	GetUserByUsername(username string) (models.User, error)

	// UpdateUser обновляет данные пользователя
	UpdateUser(user models.User) error

	// GetUserProfile получает профиль пользователя
	GetUserProfile(username string) (models.UserProfile, error)

	// UpdateUserProfile обновляет профиль пользователя
	UpdateUserProfile(username string, profile models.UserProfile) error
}
