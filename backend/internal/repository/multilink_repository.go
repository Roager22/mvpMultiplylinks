package repository

import (
	"mvp_multylink/backend/internal/models"
)

// MultiLinkRepository определяет интерфейс для работы с мультиссылками в базе данных
type MultiLinkRepository interface {
	// CreateMultiLink создает новую мультиссылку и возвращает ее ID
	CreateMultiLink(multiLink models.MultiLink) (int64, error)

	// GetMultiLinkByID получает мультиссылку по ID
	GetMultiLinkByID(id int64) (models.MultiLink, error)

	// GetMultiLinkBySlug получает мультиссылку по slug
	GetMultiLinkBySlug(slug string) (models.MultiLink, error)

	// GetMultiLinksByUserID получает все мультиссылки пользователя
	GetMultiLinksByUserID(userID int64) ([]models.MultiLink, error)

	// UpdateMultiLink обновляет мультиссылку
	UpdateMultiLink(multiLink models.MultiLink) error

	// DeleteMultiLink удаляет мультиссылку
	DeleteMultiLink(id int64) error

	// CheckSlugExists проверяет существование мультиссылки с указанным slug
	CheckSlugExists(slug string) (bool, error)
}
