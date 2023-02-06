package service

import (
	"SpotterBackend/src/internal/domain/user/model"
	"SpotterBackend/src/internal/domain/user/storage"
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

func (s *UserService) AboutUser(id int) (model.User, error) {
	user, err := s.storage.FindOne(context.Background(), id)
	return user, err
}

func (s *UserService) DeleteUser(id int) error {
	err := s.storage.Delete(context.Background(), id)
	return err
}
