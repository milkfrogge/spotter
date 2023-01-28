package service

import (
	"SpotterBackend/src/internal/user/model"
	"SpotterBackend/src/internal/user/storage"
	"context"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	storage storage.Storage
	logger  *logrus.Logger
}

func NewUserService(s storage.Storage, l *logrus.Logger) *UserService {
	return &UserService{
		storage: s,
		logger:  l,
	}
}

func (s *UserService) CreateUserByEmail(dto model.CreateByEmailDTO) (int64, error) {
	id, err := s.storage.CreateByEmail(context.Background(), dto)
	return id, err
}

func (s *UserService) CreateUserByPhone(dto model.CreateByPhoneNumberDTO) (int64, error) {
	id, err := s.storage.CreateByPhoneNumber(context.Background(), dto)
	return id, err
}
