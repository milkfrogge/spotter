package db

import (
	"SpotterBackend/src/internal/user"
	"SpotterBackend/src/pkg/client"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type storage struct {
	client client.Client
	log    *logrus.Logger
}

func (s *storage) CreateByEmail(ctx context.Context, user user.CreateByEmailDTO) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) CreateByPhoneNumber(ctx context.Context, user user.CreateByPhoneNumberDTO) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) FindOne(ctx context.Context, id int) {
	//TODO implement me
	var name, email, hash string
	_ = s.client.QueryRow(ctx, "select name, email, password_hash from public.user where id=1").Scan(&name, &email, &hash)
	fmt.Printf("%s, %s, %s\n", name, email, hash)
}

func (s *storage) Update(ctx context.Context, user user.User) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) Delete(ctx context.Context, id int) {
	//TODO implement me
	panic("implement me")
}

func NewStorage(client client.Client, log *logrus.Logger) user.Storage {
	return &storage{
		client: client,
		log:    log,
	}
}
