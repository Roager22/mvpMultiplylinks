package repository

import (
	"time"

	"mvp_multylink/backend/internal/models"
)

// MetricsRepository определяет интерфейс для работы с метриками в базе данных
type MetricsRepository interface {
	// CreateLinkMetrics создает метрику для кнопки и возвращает ее ID
	CreateLinkMetrics(metrics models.LinkMetrics) (int64, error)

	// GetMetricsByButtonID получает метрику для кнопки
	GetMetricsByButtonID(buttonID int64) (models.LinkMetrics, error)

	// UpdateLinkMetrics обновляет метрику кнопки
	UpdateLinkMetrics(metrics models.LinkMetrics) error

	// DeleteMetricsByButtonID удаляет метрику для кнопки
	DeleteMetricsByButtonID(buttonID int64) error

	// CreateClickEvent создает событие клика и возвращает его ID
	CreateClickEvent(event models.ClickEvent) (int64, error)

	// GetClickEventsByButtonID получает все события кликов для кнопки
	GetClickEventsByButtonID(buttonID int64) ([]models.ClickEvent, error)

	// GetClickEventsByButtonIDsAndDateRange получает события кликов для кнопок за указанный период
	GetClickEventsByButtonIDsAndDateRange(buttonIDs []int64, startDate, endDate time.Time) ([]models.ClickEvent, error)

	// GetUTMSourceStatsByButtonIDs получает статистику по источникам трафика (utm_source) для кнопок
	GetUTMSourceStatsByButtonIDs(buttonIDs []int64) (map[string]int, error)

	// GetUTMMediumStatsByButtonIDs получает статистику по каналам трафика (utm_medium) для кнопок
	GetUTMMediumStatsByButtonIDs(buttonIDs []int64) (map[string]int, error)
}
