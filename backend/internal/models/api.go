package models

// CreateMultiLinkRequest представляет запрос на создание мультиссылки
type CreateMultiLinkRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Slug        string `json:"slug" binding:"omitempty,min=3,max=30"`
	IsActive    bool   `json:"is_active"`
}

// UpdateMultiLinkRequest представляет запрос на обновление мультиссылки
type UpdateMultiLinkRequest struct {
	Title       string `json:"title" binding:"omitempty"`
	Description string `json:"description"`
	Slug        string `json:"slug" binding:"omitempty,min=3,max=30"`
	IsActive    bool   `json:"is_active"`
}

// CreateLinkButtonRequest представляет запрос на создание кнопки-ссылки
type CreateLinkButtonRequest struct {
	Title    string `json:"title" binding:"required"`
	URL      string `json:"url" binding:"required,url"`
	Icon     string `json:"icon"`
	Color    string `json:"color"`
	Position int    `json:"position"`
	IsActive bool   `json:"is_active"`
}

// UpdateLinkButtonRequest представляет запрос на обновление кнопки-ссылки
type UpdateLinkButtonRequest struct {
	Title    string `json:"title" binding:"omitempty"`
	URL      string `json:"url" binding:"omitempty,url"`
	Icon     string `json:"icon"`
	Color    string `json:"color"`
	Position int    `json:"position"`
	IsActive bool   `json:"is_active"`
}

// MultiLinkResponse представляет ответ с данными мультиссылки и её кнопками
type MultiLinkResponse struct {
	MultiLink MultiLink    `json:"multilink"`
	Buttons   []LinkButton `json:"buttons,omitempty"`
}

// MultiLinkListResponse представляет ответ со списком мультиссылок пользователя
type MultiLinkListResponse struct {
	MultiLinks []MultiLink `json:"multilinks"`
	Total      int         `json:"total"`
}

// MetricsResponse представляет ответ с метриками для мультиссылки
type MetricsResponse struct {
	TotalClicks    int                 `json:"total_clicks"`
	ButtonMetrics  []ButtonMetricsData `json:"button_metrics"`
	UTMSourceStats map[string]int      `json:"utm_source_stats,omitempty"`
	UTMMediumStats map[string]int      `json:"utm_medium_stats,omitempty"`
}

// ButtonMetricsData представляет метрики для отдельной кнопки
type ButtonMetricsData struct {
	ButtonID   int64   `json:"button_id"`
	ButtonName string  `json:"button_name"`
	Clicks     int     `json:"clicks"`
	Percentage float64 `json:"percentage"`
}
