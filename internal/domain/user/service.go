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
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	res := &dto.Response{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		Phone:       user.Phone,
		Address:     user.Address,
		CreatedAt:   user.CreatedAt,
	}

	return res, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repo.GetAll()

	return users, err
}
