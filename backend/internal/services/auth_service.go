package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"mvp_multylink/backend/internal/models"
)

// AuthService предоставляет методы для аутентификации и авторизации
type AuthService struct {
	jwtSecret     string
	tokenDuration time.Duration
}

// NewAuthService создает новый экземпляр AuthService
func NewAuthService(jwtSecret string, tokenDuration time.Duration) *AuthService {
	return &AuthService{
		jwtSecret:     jwtSecret,
		tokenDuration: tokenDuration,
	}
}

// GenerateToken генерирует JWT токен для пользователя
func (s *AuthService) GenerateToken(user models.User) (string, int64, error) {
	expirationTime := time.Now().Add(s.tokenDuration)
	expiresAt := expirationTime.Unix()

	claims := models.TokenClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		Exp:      expiresAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  claims.UserID,
		"username": claims.Username,
		"email":    claims.Email,
		"is_admin": claims.IsAdmin,
		"exp":      claims.Exp,
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiresAt, nil
}

// ValidateToken проверяет JWT токен и возвращает данные пользователя
func (s *AuthService) ValidateToken(tokenString string) (models.TokenClaims, error) {
	var claims models.TokenClaims

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неверный метод подписи токена")
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return claims, err
	}

	if !token.Valid {
		return claims, errors.New("недействительный токен")
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return claims, errors.New("неверный формат данных токена")
	}

	// Извлечение данных из токена
	userID, ok := mapClaims["user_id"].(float64)
	if !ok {
		return claims, errors.New("неверный формат user_id")
	}
	claims.UserID = int64(userID)

	claims.Username, _ = mapClaims["username"].(string)
	claims.Email, _ = mapClaims["email"].(string)
	claims.IsAdmin, _ = mapClaims["is_admin"].(bool)

	exp, ok := mapClaims["exp"].(float64)
	if !ok {
		return claims, errors.New("неверный формат exp")
	}
	claims.Exp = int64(exp)

	// Проверка срока действия токена
	if time.Now().Unix() > claims.Exp {
		return claims, errors.New("срок действия токена истек")
	}

	return claims, nil
}
