package models

import (
	"time"
)

// MultiLink представляет основную страницу пользователя с мультиссылками
type MultiLink struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description,omitempty" db:"description"`
	Slug        string    `json:"slug" db:"slug"` // Уникальный идентификатор для URL
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// LinkButton представляет кнопку-ссылку на странице пользователя
type LinkButton struct {
	ID          int64     `json:"id" db:"id"`
	MultiLinkID int64     `json:"multilink_id" db:"multilink_id"`
	Title       string    `json:"title" db:"title"`
	URL         string    `json:"url" db:"url"`
	Icon        string    `json:"icon,omitempty" db:"icon"`
	Color       string    `json:"color,omitempty" db:"color"`
	Position    int       `json:"position" db:"position"` // Порядок отображения
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// LinkMetrics представляет метрики для кнопок-ссылок
type LinkMetrics struct {
	ID           int64     `json:"id" db:"id"`
	LinkButtonID int64     `json:"link_button_id" db:"link_button_id"`
	Clicks       int       `json:"clicks" db:"clicks"`
	LastClickAt  time.Time `json:"last_click_at,omitempty" db:"last_click_at"`
}

// ClickEvent представляет событие клика по ссылке с UTM-метками
type ClickEvent struct {
	ID           int64     `json:"id" db:"id"`
	LinkButtonID int64     `json:"link_button_id" db:"link_button_id"`
	IP           string    `json:"ip,omitempty" db:"ip"`
	UserAgent    string    `json:"user_agent,omitempty" db:"user_agent"`
	Referer      string    `json:"referer,omitempty" db:"referer"`
	UTMSource    string    `json:"utm_source,omitempty" db:"utm_source"`
	UTMMedium    string    `json:"utm_medium,omitempty" db:"utm_medium"`
	UTMCampaign  string    `json:"utm_campaign,omitempty" db:"utm_campaign"`
	UTMContent   string    `json:"utm_content,omitempty" db:"utm_content"`
	UTMTerm      string    `json:"utm_term,omitempty" db:"utm_term"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
