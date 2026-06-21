package user

import "go-conceptual/internal/domain/user/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) RegisterUser(req dto.CreateRequest) (*dto.Response, error) {
	user := &User{
		Name:  req.Name,
		Email: req.Email,
	}

	err := user.hashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repo.GetAll()

	return users, err
}
