package user

import (
	"go-conceptual/internal/auth"
	"go-conceptual/internal/domain/user/dto"
)

type service struct {
	repo Repository
	jwt  auth.JWTService
}

func NewService(repo Repository, jwt auth.JWTService) *service {
	return &service{repo, jwt}
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

func (s *service) LoginUser(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)

	if err != nil {
		return nil, err
	}

	if err := user.checkPassword(req.Password); err != nil {
		return nil, err
	}

	token, err := s.jwt.GenerateToken(user.ID, user.Name, user.Email)

	if err != nil {
		return nil, err
	}

	res := &dto.LoginResponse{
		Token: token,
		User:  *user.ToResponse(),
	}

	return res, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repo.GetAll()

	return users, err
}
