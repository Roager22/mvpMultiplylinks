package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"mvp_multylink/backend/internal/models"
	"mvp_multylink/backend/internal/services"
)

// MultiLinkHandler обрабатывает запросы, связанные с мультиссылками
type MultiLinkHandler struct {
	multiLinkService *services.MultiLinkService
}

// NewMultiLinkHandler создает новый экземпляр MultiLinkHandler
func NewMultiLinkHandler(multiLinkService *services.MultiLinkService) *MultiLinkHandler {
	return &MultiLinkHandler{
		multiLinkService: multiLinkService,
	}
}

// CreateMultiLink обрабатывает запрос на создание новой мультиссылки
func (h *MultiLinkHandler) CreateMultiLink(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	var req models.CreateMultiLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Если slug не указан, генерируем его на основе имени пользователя или другой логики
	if req.Slug == "" {
		// Здесь должна быть логика генерации уникального slug
		// Например, на основе имени пользователя + случайная строка
	}

	// Проверка уникальности slug
	exists, err := h.multiLinkService.CheckSlugExists(req.Slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке slug"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Такой slug уже используется"})
		return
	}

	multiLink := models.MultiLink{
		UserID:      userID.(int64),
		Title:       req.Title,
		Description: req.Description,
		Slug:        req.Slug,
		IsActive:    req.IsActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	multiLinkID, err := h.multiLinkService.CreateMultiLink(multiLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании мультиссылки"})
		return
	}

	multiLink.ID = multiLinkID

	c.JSON(http.StatusCreated, gin.H{"multilink": multiLink})
}

// GetMultiLink обрабатывает запрос на получение мультиссылки по ID
func (h *MultiLinkHandler) GetMultiLink(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	multiLinkID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID мультиссылки"})
		return
	}

	multiLink, err := h.multiLinkService.GetMultiLinkByID(multiLinkID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Мультиссылка не найдена"})
		return
	}

	// Проверка, принадлежит ли мультиссылка текущему пользователю
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

	c.JSON(http.StatusOK, models.MultiLinkResponse{
		MultiLink: multiLink,
		Buttons:   buttons,
	})
}

// UpdateMultiLink обрабатывает запрос на обновление мультиссылки
func (h *MultiLinkHandler) UpdateMultiLink(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	multiLinkID, err := strconv.ParseInt(c.Param("id"), 10, 64)
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

	var req models.UpdateMultiLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновление полей мультиссылки
	if req.Title != "" {
		multiLink.Title = req.Title
	}

	multiLink.Description = req.Description

	if req.Slug != "" && req.Slug != multiLink.Slug {
		// Проверка уникальности нового slug
		exists, err := h.multiLinkService.CheckSlugExists(req.Slug)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке slug"})
			return
		}

		if exists {
			c.JSON(http.StatusConflict, gin.H{"error": "Такой slug уже используется"})
			return
		}

		multiLink.Slug = req.Slug
	}

	multiLink.IsActive = req.IsActive
	multiLink.UpdatedAt = time.Now()

	err = h.multiLinkService.UpdateMultiLink(multiLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении мультиссылки"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"multilink": multiLink})
}

// DeleteMultiLink обрабатывает запрос на удаление мультиссылки
func (h *MultiLinkHandler) DeleteMultiLink(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	multiLinkID, err := strconv.ParseInt(c.Param("id"), 10, 64)
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

	err = h.multiLinkService.DeleteMultiLink(multiLinkID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении мультиссылки"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Мультиссылка успешно удалена"})
}

// GetUserMultiLinks обрабатывает запрос на получение всех мультиссылок пользователя
func (h *MultiLinkHandler) GetUserMultiLinks(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	multiLinks, err := h.multiLinkService.GetMultiLinksByUserID(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении мультиссылок"})
		return
	}

	c.JSON(http.StatusOK, models.MultiLinkListResponse{
		MultiLinks: multiLinks,
		Total:      len(multiLinks),
	})
}

// GetPublicMultiLink обрабатывает запрос на получение публичной мультиссылки по slug
func (h *MultiLinkHandler) GetPublicMultiLink(c *gin.Context) {
	slug := c.Param("slug")

	multiLink, err := h.multiLinkService.GetMultiLinkBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Мультиссылка не найдена"})
		return
	}

	// Проверка активности мультиссылки
	if !multiLink.IsActive {
		c.JSON(http.StatusNotFound, gin.H{"error": "Мультиссылка не найдена или неактивна"})
		return
	}

	// Получение кнопок для мультиссылки
	buttons, err := h.multiLinkService.GetActiveLinkButtonsByMultiLinkID(multiLink.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении кнопок"})
		return
	}

	c.JSON(http.StatusOK, models.MultiLinkResponse{
		MultiLink: multiLink,
		Buttons:   buttons,
	})
}
