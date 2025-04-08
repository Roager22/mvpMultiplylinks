package services

import (
	"mvp_multylink/backend/internal/models"
	"mvp_multylink/backend/internal/repository"
)

// ButtonService предоставляет методы для работы с кнопками-ссылками
type ButtonService struct {
	buttonRepo  repository.ButtonRepository
	metricsRepo repository.MetricsRepository
}

// NewButtonService создает новый экземпляр ButtonService
func NewButtonService(buttonRepo repository.ButtonRepository, metricsRepo repository.MetricsRepository) *ButtonService {
	return &ButtonService{
		buttonRepo:  buttonRepo,
		metricsRepo: metricsRepo,
	}
}

// CreateButton создает новую кнопку-ссылку
func (s *ButtonService) CreateButton(button models.LinkButton) (int64, error) {
	return s.buttonRepo.CreateButton(button)
}

// GetButtonByID получает кнопку по ID
func (s *ButtonService) GetButtonByID(id int64) (models.LinkButton, error) {
	return s.buttonRepo.GetButtonByID(id)
}

// UpdateButton обновляет кнопку
func (s *ButtonService) UpdateButton(button models.LinkButton) error {
	return s.buttonRepo.UpdateButton(button)
}

// DeleteButton удаляет кнопку
func (s *ButtonService) DeleteButton(id int64) error {
	// Сначала удаляем метрики кнопки
	err := s.metricsRepo.DeleteMetricsByButtonID(id)
	if err != nil {
		return err
	}

	// Затем удаляем саму кнопку
	return s.buttonRepo.DeleteButton(id)
}

// GetButtonsCountByMultiLinkID получает количество кнопок для мультиссылки
func (s *ButtonService) GetButtonsCountByMultiLinkID(multiLinkID int64) (int, error) {
	return s.buttonRepo.GetButtonsCountByMultiLinkID(multiLinkID)
}

// UpdateButtonPosition обновляет позицию кнопки
func (s *ButtonService) UpdateButtonPosition(id int64, position int) error {
	return s.buttonRepo.UpdateButtonPosition(id, position)
}

// CreateButtonMetrics создает метрику для кнопки
func (s *ButtonService) CreateButtonMetrics(metrics models.LinkMetrics) (int64, error) {
	return s.metricsRepo.CreateLinkMetrics(metrics)
}

// GetButtonMetrics получает метрику для кнопки
func (s *ButtonService) GetButtonMetrics(buttonID int64) (models.LinkMetrics, error) {
	return s.metricsRepo.GetMetricsByButtonID(buttonID)
}
