/*
 * Copyright (c)  Nikita Cherkasov, 2023.
 * Spotter Project
 */

package storage

import (
	"SpotterBackend/src/internal/user/model"
	"SpotterBackend/src/pkg/client"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type UserStorage struct {
	client client.Client
	log    *logrus.Logger
}

func (s *UserStorage) CreateByEmail(ctx context.Context, user model.CreateByEmailDTO) (int64, error) {

	//TODO: bcrypt for hashing password
	qRow := s.client.QueryRow(
		ctx,
		"INSERT INTO public.\"user\" (email, password_hash) values ($1, $2) RETURNING id",
		user.Email, user.Password,
	)
	var id int64
	err := qRow.Scan(&id)
	if err != nil {
		s.log.Errorf("Unable to create new user by email: %s", err)
		return -1, err
	}
	return id, nil

}

func (s *UserStorage) CreateByPhoneNumber(ctx context.Context, user model.CreateByPhoneNumberDTO) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStorage) FindOne(ctx context.Context, id int) {
	//TODO implement me
	var name, email, hash string
	_ = s.client.QueryRow(ctx, "select name, email, password_hash from public.user where id=1").Scan(&name, &email, &hash)
	fmt.Printf("%s, %s, %s\n", name, email, hash)
}

func (s *UserStorage) Update(ctx context.Context, user model.User) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStorage) Delete(ctx context.Context, id int) {
	//TODO implement me
	panic("implement me")
}

func NewStorage(client client.Client, log *logrus.Logger) Storage {
	return &UserStorage{
		client: client,
		log:    log,
	}
}
