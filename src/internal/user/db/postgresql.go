package db

import (
	"SpotterBackend/src/internal/user"
	"SpotterBackend/src/pkg/client"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type UserStorage struct {
	client client.Client
	log    *logrus.Logger
}

func (s *UserStorage) CreateByEmail(ctx context.Context, user user.CreateByEmailDTO) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStorage) CreateByPhoneNumber(ctx context.Context, user user.CreateByPhoneNumberDTO) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStorage) FindOne(ctx context.Context, id int) {
	//TODO implement me
	var name, email, hash string
	_ = s.client.QueryRow(ctx, "select name, email, password_hash from public.user where id=1").Scan(&name, &email, &hash)
	fmt.Printf("%s, %s, %s\n", name, email, hash)
}

func (s *UserStorage) Update(ctx context.Context, user user.User) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStorage) Delete(ctx context.Context, id int) {
	//TODO implement me
	panic("implement me")
}

func NewStorage(client client.Client, log *logrus.Logger) user.Storage {
	return &UserStorage{
		client: client,
		log:    log,
	}
}
