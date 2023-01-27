package user

import (
	"SpotterBackend/src/internal/user/db"
	"context"
	"github.com/sirupsen/logrus"
)

type userService struct {
	storage db.UserStorage
	logger  *logrus.Logger
}

func (s *userService) CreateUserByEmail(dto CreateByEmailDTO) {
	s.storage.CreateByEmail(context.Background(), dto)
}
