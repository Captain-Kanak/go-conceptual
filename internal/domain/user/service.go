package user

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) RegisterUser() {
	var user User

	s.repo.Create(user)

}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repo.GetAll()

	return users, err
}
