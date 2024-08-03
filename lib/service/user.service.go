package service

import (
	"github.com/ShoreLab/shorelab-backend/lib/dto"
)

func (s *Service) CreateUserService(user *dto.UserCreateRequest) error {
	return s.repository.CreateUser(user)
}
