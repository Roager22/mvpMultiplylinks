package services

import (
	"time"

	"mvp_multylink/backend/internal/models"
	"mvp_multylink/backend/internal/repository"
)

// MetricsService предоставляет методы для работы с метриками
type MetricsService struct {
	metricsRepo repository.MetricsRepository
	buttonRepo  repository.ButtonRepository
}

// NewMetricsService создает новый экземпляр MetricsService
func NewMetricsService(metricsRepo repository.MetricsRepository, buttonRepo repository.ButtonRepository) *MetricsService {
	return &MetricsService{
		metricsRepo: metricsRepo,
		buttonRepo:  buttonRepo,
	}
}

// RecordClickEvent записывает событие клика по кнопке
func (s *MetricsService) RecordClickEvent(event models.ClickEvent) (int64, error) {
	return s.metricsRepo.CreateClickEvent(event)
}

// IncrementButtonClicks увеличивает счетчик кликов для кнопки
func (s *MetricsService) IncrementButtonClicks(buttonID int64) error {
	// Получаем текущие метрики кнопки
	metrics, err := s.metricsRepo.GetMetricsByButtonID(buttonID)
	if err != nil {
		// Если метрики не существуют, создаем новые
		metrics = models.LinkMetrics{
			LinkButtonID: buttonID,
			Clicks:       0,
		}
		_, err = s.metricsRepo.CreateLinkMetrics(metrics)
		if err != nil {
			return err
		}
	}

	// Увеличиваем счетчик кликов
	metrics.Clicks++
	metrics.LastClickAt = time.Now()

	// Обновляем метрики
	return s.metricsRepo.UpdateLinkMetrics(metrics)
}

// GetButtonMetrics получает метрики для кнопки
func (s *MetricsService) GetButtonMetrics(buttonID int64) (models.LinkMetrics, error) {
	return s.metricsRepo.GetMetricsByButtonID(buttonID)
}

// GetUTMSourceStats получает статистику по источникам трафика (utm_source)
func (s *MetricsService) GetUTMSourceStats(multiLinkID int64) (map[string]int, error) {
	// Получаем все кнопки для мультиссылки
	buttons, err := s.buttonRepo.GetButtonsByMultiLinkID(multiLinkID)
	if err != nil {
		return nil, err
	}

	// Собираем ID всех кнопок
	buttonIDs := make([]int64, len(buttons))
	for i, button := range buttons {
		buttonIDs[i] = button.ID
	}

	// Получаем статистику по UTM-меткам
	return s.metricsRepo.GetUTMSourceStatsByButtonIDs(buttonIDs)
}

// GetUTMMediumStats получает статистику по каналам трафика (utm_medium)
func (s *MetricsService) GetUTMMediumStats(multiLinkID int64) (map[string]int, error) {
	// Получаем все кнопки для мультиссылки
	buttons, err := s.buttonRepo.GetButtonsByMultiLinkID(multiLinkID)
	if err != nil {
		return nil, err
	}

	// Собираем ID всех кнопок
	buttonIDs := make([]int64, len(buttons))
	for i, button := range buttons {
		buttonIDs[i] = button.ID
	}

	// Получаем статистику по UTM-меткам
	return s.metricsRepo.GetUTMMediumStatsByButtonIDs(buttonIDs)
}

// GetClickEventsByDateRange получает события кликов за указанный период
func (s *MetricsService) GetClickEventsByDateRange(multiLinkID int64, startDate, endDate time.Time) ([]models.ClickEvent, error) {
	// Получаем все кнопки для мультиссылки
	buttons, err := s.buttonRepo.GetButtonsByMultiLinkID(multiLinkID)
	if err != nil {
		return nil, err
	}

	// Собираем ID всех кнопок
	buttonIDs := make([]int64, len(buttons))
	for i, button := range buttons {
		buttonIDs[i] = button.ID
	}

	// Получаем события кликов
	return s.metricsRepo.GetClickEventsByButtonIDsAndDateRange(buttonIDs, startDate, endDate)
}
