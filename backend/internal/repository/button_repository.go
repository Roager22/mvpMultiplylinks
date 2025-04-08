package repository

import (
	"mvp_multylink/backend/internal/models"
)

// ButtonRepository определяет интерфейс для работы с кнопками-ссылками в базе данных
type ButtonRepository interface {
	// CreateButton создает новую кнопку-ссылку и возвращает ее ID
	CreateButton(button models.LinkButton) (int64, error)

	// GetButtonByID получает кнопку по ID
	GetButtonByID(id int64) (models.LinkButton, error)

	// GetButtonsByMultiLinkID получает все кнопки для мультиссылки
	GetButtonsByMultiLinkID(multiLinkID int64) ([]models.LinkButton, error)

	// GetActiveButtonsByMultiLinkID получает все активные кнопки для мультиссылки
	GetActiveButtonsByMultiLinkID(multiLinkID int64) ([]models.LinkButton, error)

	// UpdateButton обновляет кнопку
	UpdateButton(button models.LinkButton) error

	// UpdateButtonPosition обновляет позицию кнопки
	UpdateButtonPosition(id int64, position int) error

	// DeleteButton удаляет кнопку
	DeleteButton(id int64) error

	// DeleteButtonsByMultiLinkID удаляет все кнопки для мультиссылки
	DeleteButtonsByMultiLinkID(multiLinkID int64) error

	// GetButtonsCountByMultiLinkID получает количество кнопок для мультиссылки
	GetButtonsCountByMultiLinkID(multiLinkID int64) (int, error)
}
