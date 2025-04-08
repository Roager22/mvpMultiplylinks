package models

import (
	"time"
)

// User представляет пользователя системы
type User struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:\"email\"`
	Password  string    `json:"-" db:"password_hash"`
	AvatarURL string    `json:"avatar_url,omitempty" db:"avatar_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	IsAdmin   bool      `json:"is_admin" db:"is_admin"`
}

// UpdateUserRequest представляет данные для обновления пользователя
type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// UpdateProfileRequest представляет данные для обновления профиля пользователя
type UpdateProfileRequest struct {
	DisplayName   string `json:"display_name,omitempty"`
	AvatarURL     string `json:"avatar_url,omitempty"`
	Bio           string `json:"bio,omitempty"`
	ThemeColor    string `json:"theme_color,omitempty"`
	BackgroundURL string `json:"background_url,omitempty"`
}

// UserProfile представляет публичную информацию о пользователе
type UserProfile struct {
	Username      string `json:"username" db:"username"`
	DisplayName   string `json:"display_name,omitempty" db:"display_name"`
	AvatarURL     string `json:"avatar_url,omitempty" db:"avatar_url"`
	Bio           string `json:"bio,omitempty" db:"bio"`
	ThemeColor    string `json:"theme_color,omitempty" db:"theme_color"`
	BackgroundURL string `json:"background_url,omitempty" db:"background_url"`
}
