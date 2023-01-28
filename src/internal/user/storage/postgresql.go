/*
 * Copyright (c)  Nikita Cherkasov, 2023.
 * Spotter Project
 */

package storage

import (
	"SpotterBackend/src/internal/constants"
	"SpotterBackend/src/internal/user/model"
	"SpotterBackend/src/pkg/client"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const hashCost = 12

type UserStorage struct {
	client client.Client
	log    *logrus.Logger
}

func (s *UserStorage) CreateByEmail(ctx context.Context, user model.CreateByEmailDTO) (int64, error) {

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), hashCost)
	if err != nil {
		return -1, errors.New(constants.InternalServerError)
	}
	qRow := s.client.QueryRow(
		ctx,
		"INSERT INTO public.\"user\" (email, password_hash) values ($1, $2) RETURNING id",
		user.Email, string(hashBytes),
	)
	var id int64
	err = qRow.Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			s.log.Errorf("SQL Error: %s, statuscode: %s", pgErr.Message, pgErr.Code)
			if pgErr.Code == "23505" {
				return -1, errors.New(constants.UserAlreadyExistsEmail)

			}
		}
		s.log.Errorf("Unable to create new user by email: %s", err)
		return -1, errors.New(constants.InternalServerError)
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
