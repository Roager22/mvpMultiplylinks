package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"mvp_multylink/backend/internal/models"
	"mvp_multylink/backend/internal/services"
)

// MetricsHandler обрабатывает запросы, связанные с метриками
type MetricsHandler struct {
	multiLinkService *services.MultiLinkService
	buttonService    *services.ButtonService
	metricsService   *services.MetricsService
}

// NewMetricsHandler создает новый экземпляр MetricsHandler
func NewMetricsHandler(multiLinkService *services.MultiLinkService, buttonService *services.ButtonService, metricsService *services.MetricsService) *MetricsHandler {
	return &MetricsHandler{
		multiLinkService: multiLinkService,
		buttonService:    buttonService,
		metricsService:   metricsService,
	}
}

// RecordClick обрабатывает запрос на запись клика по кнопке
func (h *MetricsHandler) RecordClick(c *gin.Context) {
	buttonID, err := strconv.ParseInt(c.Param("button_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID кнопки"})
		return
	}

	// Получение информации о кнопке
	button, err := h.buttonService.GetButtonByID(buttonID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Кнопка не найдена"})
		return
	}

	// Проверка активности кнопки
	if !button.IsActive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Кнопка неактивна"})
		return
	}

	// Получение UTM-меток из запроса
	utmSource := c.Query("utm_source")
	utmMedium := c.Query("utm_medium")
	utmCampaign := c.Query("utm_campaign")
	utmContent := c.Query("utm_content")
	utmTerm := c.Query("utm_term")

	// Создание события клика
	clickEvent := models.ClickEvent{
		LinkButtonID: buttonID,
		IP:           c.ClientIP(),
		UserAgent:    c.Request.UserAgent(),
		Referer:      c.Request.Referer(),
		UTMSource:    utmSource,
		UTMMedium:    utmMedium,
		UTMCampaign:  utmCampaign,
		UTMContent:   utmContent,
		UTMTerm:      utmTerm,
		CreatedAt:    time.Now(),
	}

	// Запись события клика
	_, err = h.metricsService.RecordClickEvent(clickEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при записи события клика"})
		return
	}

	// Обновление метрики кнопки
	err = h.metricsService.IncrementButtonClicks(buttonID)
	if err != nil {
		// Логируем ошибку, но не прерываем выполнение
		// TODO: добавить логирование
	}

	// Перенаправление на URL кнопки
	c.Redirect(http.StatusFound, button.URL)
}

// GetMultiLinkMetrics обрабатывает запрос на получение метрик мультиссылки
func (h *MetricsHandler) GetMultiLinkMetrics(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	multiLinkID, err := strconv.ParseInt(c.Param("multilink_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID мультиссылки"})
		return
	}

	// Проверка существования мультиссылки и прав доступа
	multiLink, err := h.multiLinkService.GetMultiLinkByID(multiLinkID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Мультиссылка не найдена"})
		return
	}

	if multiLink.UserID != userID.(int64) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещен"})
		return
	}

	// Получение кнопок для мультиссылки
	buttons, err := h.multiLinkService.GetLinkButtonsByMultiLinkID(multiLinkID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении кнопок"})
		return
	}

	// Получение метрик для каждой кнопки
	var totalClicks int
	buttonMetrics := make([]models.ButtonMetricsData, 0, len(buttons))

	for _, button := range buttons {
		metrics, err := h.metricsService.GetButtonMetrics(button.ID)
		if err != nil {
			continue // Пропускаем кнопку, если не удалось получить метрики
		}

		totalClicks += metrics.Clicks

		buttonMetrics = append(buttonMetrics, models.ButtonMetricsData{
			ButtonID:   button.ID,
			ButtonName: button.Title,
			Clicks:     metrics.Clicks,
			Percentage: 0, // Заполним после подсчета общего количества кликов
		})
	}

	// Расчет процентов для каждой кнопки
	if totalClicks > 0 {
		for i := range buttonMetrics {
			buttonMetrics[i].Percentage = float64(buttonMetrics[i].Clicks) / float64(totalClicks) * 100
		}
	}

	// Получение статистики по UTM-меткам
	utmSourceStats, err := h.metricsService.GetUTMSourceStats(multiLinkID)
	if err != nil {
		utmSourceStats = make(map[string]int)
	}

	utmMediumStats, err := h.metricsService.GetUTMMediumStats(multiLinkID)
	if err != nil {
		utmMediumStats = make(map[string]int)
	}

	c.JSON(http.StatusOK, models.MetricsResponse{
		TotalClicks:    totalClicks,
		ButtonMetrics:  buttonMetrics,
		UTMSourceStats: utmSourceStats,
		UTMMediumStats: utmMediumStats,
	})
}
