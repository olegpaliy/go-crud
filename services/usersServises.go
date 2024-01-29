package services

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UsersService interface {
	CreateUser(user User)
	GetAllUsers() []User
	GetUserByID(id string) (User, error)
	UpdateUser(user User) error
	DeleteUser(id string) error
}

type InMemoryUserService struct {
	users map[string]User
}

func NewInMemoryUserService() *InMemoryUserService {
	return &InMemoryUserService{
		users: make(map[string]User),
	}
}

func (s *InMemoryUserService) CreateUser(user User) {
	s.users[user.ID] = user
}

func (s *InMemoryUserService) GetAllUsers() []User {
	var result []User
	for _, user := range s.users {
		result = append(result, user)
	}
	return result
}

func (s *InMemoryUserService) GetUserByID(id string) (User, error) {
	user, exists := s.users[id]
	if !exists {
		return User{}, NotFoundError("User not found with id: " + id)
	}
	return user, nil
}

func (s *InMemoryUserService) UpdateUser(user User) error {
	_, exists := s.users[user.ID]
	if !exists {
		return NotFoundError("User not found with id: " + user.ID)
	}

	s.users[user.ID] = user
	return nil
}

func (s *InMemoryUserService) DeleteUser(id string) error {
	_, exists := s.users[id]
	if !exists {
		return NotFoundError("User not found with id: " + id)
	}

	delete(s.users, id)
	return nil
}

// NotFoundError represents an error when an user is not found.
type NotFoundError string

func (e NotFoundError) Error() string {
	return string(e)
}
