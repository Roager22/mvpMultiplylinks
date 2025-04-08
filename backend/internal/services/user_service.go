package services

type UserService struct {
	// Здесь будут зависимости репозитория
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAllUsers() ([]User, error) {
	// TODO: реализовать получение всех пользователей
	return nil, nil
}

func (s *UserService) CreateUser(user User) error {
	// TODO: реализовать создание пользователя
	return nil
}

func (s *UserService) GetUserByID(id string) (*User, error) {
	// TODO: реализовать получение пользователя по ID
	return nil, nil
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
