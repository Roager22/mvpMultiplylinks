package api

import (
	"github.com/gin-gonic/gin"

	"mvp_multylink/backend/internal/handlers"
	"mvp_multylink/backend/internal/middleware"
)

// SetupRouter настраивает маршрутизацию API
func SetupRouter(authHandler *handlers.AuthHandler, multiLinkHandler *handlers.MultiLinkHandler, buttonHandler *handlers.ButtonHandler, metricsHandler *handlers.MetricsHandler, authMiddleware *middleware.AuthMiddleware) *gin.Engine {
	router := gin.Default()

	// Middleware для CORS
	router.Use(middleware.CORSMiddleware())

	// Группа для API
	api := router.Group("/api")
	{
		// Аутентификация
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Пользователи
		users := api.Group("/users")
		{
			// Публичные профили пользователей
			users.GET("/:username", authHandler.GetUserProfile)

			// Защищенные маршруты для пользователей
			protectedUsers := users.Group("/")
			protectedUsers.Use(authMiddleware.AuthRequired())
			{
				protectedUsers.GET("/me", authHandler.GetCurrentUser)
				protectedUsers.PUT("/me", authHandler.UpdateCurrentUser)
				protectedUsers.PUT("/me/profile", authHandler.UpdateUserProfile)
			}
		}

		// Мультиссылки
		multilinks := api.Group("/multilinks")
		{
			// Публичный доступ к мультиссылке по slug
			multilinks.GET("/s/:slug", multiLinkHandler.GetPublicMultiLink)

			// Защищенные маршруты для мультиссылок
			protectedMultilinks := multilinks.Group("/")
			protectedMultilinks.Use(authMiddleware.AuthRequired())
			{
				protectedMultilinks.POST("/", multiLinkHandler.CreateMultiLink)
				protectedMultilinks.GET("/", multiLinkHandler.GetUserMultiLinks)
				protectedMultilinks.GET("/:id", multiLinkHandler.GetMultiLink)
				protectedMultilinks.PUT("/:id", multiLinkHandler.UpdateMultiLink)
				protectedMultilinks.DELETE("/:id", multiLinkHandler.DeleteMultiLink)

				// Кнопки для мультиссылок
				protectedMultilinks.POST("/:multilink_id/buttons", buttonHandler.CreateButton)
				protectedMultilinks.PUT("/:multilink_id/buttons/reorder", buttonHandler.ReorderButtons)
			}
		}

		// Кнопки
		buttons := api.Group("/buttons")
		{
			// Запись клика по кнопке (публичный доступ)
			buttons.GET("/:button_id/click", metricsHandler.RecordClick)

			// Защищенные маршруты для кнопок
			protectedButtons := buttons.Group("/")
			protectedButtons.Use(authMiddleware.AuthRequired())
			{
				protectedButtons.PUT("/:id", buttonHandler.UpdateButton)
				protectedButtons.DELETE("/:id", buttonHandler.DeleteButton)
			}
		}

		// Метрики
		metrics := api.Group("/metrics")
		metrics.Use(authMiddleware.AuthRequired())
		{
			metrics.GET("/multilinks/:multilink_id", metricsHandler.GetMultiLinkMetrics)
		}

		// Админ-панель (только для администраторов)
		admin := api.Group("/admin")
		admin.Use(authMiddleware.AdminRequired())
		{
			// TODO: добавить маршруты для админ-панели
		}
	}

	// Обработка статических файлов для фронтенда
	router.Static("/static", "./frontend/build/static")
	router.StaticFile("/", "./frontend/build/index.html")

	return router
}
