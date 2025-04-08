package repository

type UserRepository struct {
	// Здесь будет подключение к базе данных
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindAll() ([]User, error) {
	// TODO: реализовать запрос к базе данных
	return nil, nil
}

func (r *UserRepository) Create(user User) error {
	// TODO: реализовать создание записи в базе данных
	return nil
}

func (r *UserRepository) FindByID(id string) (*User, error) {
	// TODO: реализовать поиск по ID
	return nil, nil
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
