package services

import (
	"mvp_multylink/backend/internal/models"
	"mvp_multylink/backend/internal/repository"
)

// MultiLinkService предоставляет методы для работы с мультиссылками
type MultiLinkService struct {
	multiLinkRepo repository.MultiLinkRepository
	buttonRepo    repository.ButtonRepository
}

// NewMultiLinkService создает новый экземпляр MultiLinkService
func NewMultiLinkService(multiLinkRepo repository.MultiLinkRepository, buttonRepo repository.ButtonRepository) *MultiLinkService {
	return &MultiLinkService{
		multiLinkRepo: multiLinkRepo,
		buttonRepo:    buttonRepo,
	}
}

// CreateMultiLink создает новую мультиссылку
func (s *MultiLinkService) CreateMultiLink(multiLink models.MultiLink) (int64, error) {
	return s.multiLinkRepo.CreateMultiLink(multiLink)
}

// GetMultiLinkByID получает мультиссылку по ID
func (s *MultiLinkService) GetMultiLinkByID(id int64) (models.MultiLink, error) {
	return s.multiLinkRepo.GetMultiLinkByID(id)
}

// GetMultiLinkBySlug получает мультиссылку по slug
func (s *MultiLinkService) GetMultiLinkBySlug(slug string) (models.MultiLink, error) {
	return s.multiLinkRepo.GetMultiLinkBySlug(slug)
}

// GetMultiLinksByUserID получает все мультиссылки пользователя
func (s *MultiLinkService) GetMultiLinksByUserID(userID int64) ([]models.MultiLink, error) {
	return s.multiLinkRepo.GetMultiLinksByUserID(userID)
}

// UpdateMultiLink обновляет мультиссылку
func (s *MultiLinkService) UpdateMultiLink(multiLink models.MultiLink) error {
	return s.multiLinkRepo.UpdateMultiLink(multiLink)
}

// DeleteMultiLink удаляет мультиссылку
func (s *MultiLinkService) DeleteMultiLink(id int64) error {
	// Сначала удаляем все кнопки, связанные с мультиссылкой
	err := s.buttonRepo.DeleteButtonsByMultiLinkID(id)
	if err != nil {
		return err
	}

	// Затем удаляем саму мультиссылку
	return s.multiLinkRepo.DeleteMultiLink(id)
}

// CheckSlugExists проверяет существование мультиссылки с указанным slug
func (s *MultiLinkService) CheckSlugExists(slug string) (bool, error) {
	return s.multiLinkRepo.CheckSlugExists(slug)
}

// GetLinkButtonsByMultiLinkID получает все кнопки для мультиссылки
func (s *MultiLinkService) GetLinkButtonsByMultiLinkID(multiLinkID int64) ([]models.LinkButton, error) {
	return s.buttonRepo.GetButtonsByMultiLinkID(multiLinkID)
}

// GetActiveLinkButtonsByMultiLinkID получает все активные кнопки для мультиссылки
func (s *MultiLinkService) GetActiveLinkButtonsByMultiLinkID(multiLinkID int64) ([]models.LinkButton, error) {
	return s.buttonRepo.GetActiveButtonsByMultiLinkID(multiLinkID)
}
