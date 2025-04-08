package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"mvp_multylink/backend/internal/models"
	"mvp_multylink/backend/internal/services"
)

// ButtonHandler обрабатывает запросы, связанные с кнопками-ссылками
type ButtonHandler struct {
	multiLinkService *services.MultiLinkService
	buttonService    *services.ButtonService
}

// NewButtonHandler создает новый экземпляр ButtonHandler
func NewButtonHandler(multiLinkService *services.MultiLinkService, buttonService *services.ButtonService) *ButtonHandler {
	return &ButtonHandler{
		multiLinkService: multiLinkService,
		buttonService:    buttonService,
	}
}

// CreateButton обрабатывает запрос на создание новой кнопки-ссылки
func (h *ButtonHandler) CreateButton(c *gin.Context) {
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

	var req models.CreateLinkButtonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Если позиция не указана, устанавливаем последнюю
	if req.Position == 0 {
		count, err := h.buttonService.GetButtonsCountByMultiLinkID(multiLinkID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении количества кнопок"})
			return
		}
		req.Position = count + 1
	}

	button := models.LinkButton{
		MultiLinkID: multiLinkID,
		Title:       req.Title,
		URL:         req.URL,
		Icon:        req.Icon,
		Color:       req.Color,
		Position:    req.Position,
		IsActive:    req.IsActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	buttonID, err := h.buttonService.CreateButton(button)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании кнопки"})
		return
	}

	button.ID = buttonID

	// Инициализация метрики для новой кнопки
	metrics := models.LinkMetrics{
		LinkButtonID: buttonID,
		Clicks:       0,
	}

	_, err = h.buttonService.CreateButtonMetrics(metrics)
	if err != nil {
		// Логируем ошибку, но не прерываем выполнение
		// TODO: добавить логирование
	}

	c.JSON(http.StatusCreated, gin.H{"button": button})
}

// UpdateButton обрабатывает запрос на обновление кнопки-ссылки
func (h *ButtonHandler) UpdateButton(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	buttonID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID кнопки"})
		return
	}

	// Получение кнопки и проверка прав доступа
	button, err := h.buttonService.GetButtonByID(buttonID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Кнопка не найдена"})
		return
	}

	// Получение мультиссылки для проверки владельца
	multiLink, err := h.multiLinkService.GetMultiLinkByID(button.MultiLinkID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке прав доступа"})
		return
	}

	if multiLink.UserID != userID.(int64) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещен"})
		return
	}

	var req models.UpdateLinkButtonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновление полей кнопки
	if req.Title != "" {
		button.Title = req.Title
	}

	if req.URL != "" {
		button.URL = req.URL
	}

	button.Icon = req.Icon
	button.Color = req.Color

	if req.Position != 0 {
		button.Position = req.Position
	}

	button.IsActive = req.IsActive
	button.UpdatedAt = time.Now()

	err = h.buttonService.UpdateButton(button)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении кнопки"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"button": button})
}

// DeleteButton обрабатывает запрос на удаление кнопки-ссылки
func (h *ButtonHandler) DeleteButton(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	buttonID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID кнопки"})
		return
	}

	// Получение кнопки и проверка прав доступа
	button, err := h.buttonService.GetButtonByID(buttonID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Кнопка не найдена"})
		return
	}

	// Получение мультиссылки для проверки владельца
	multiLink, err := h.multiLinkService.GetMultiLinkByID(button.MultiLinkID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке прав доступа"})
		return
	}

	if multiLink.UserID != userID.(int64) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещен"})
		return
	}

	err = h.buttonService.DeleteButton(buttonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении кнопки"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Кнопка успешно удалена"})
}

// ReorderButtons обрабатывает запрос на изменение порядка кнопок
func (h *ButtonHandler) ReorderButtons(c *gin.Context) {
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

	var buttonOrder []struct {
		ID       int64 `json:"id"`
		Position int   `json:"position"`
	}

	if err := c.ShouldBindJSON(&buttonOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновление позиций кнопок
	for _, item := range buttonOrder {
		err := h.buttonService.UpdateButtonPosition(item.ID, item.Position)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении порядка кнопок"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Порядок кнопок успешно обновлен"})
}
