/*
 * Copyright (c)  Nikita Cherkasov, 2023.
 * Spotter Project
 */

package storage

import (
	"SpotterBackend/src/internal/constants"
	"SpotterBackend/src/internal/domain/user/model"
	"SpotterBackend/src/pkg/client"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func (s *UserStorage) CreateByPhoneNumber(ctx context.Context, user model.CreateByPhoneNumberDTO) (int64, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), hashCost)
	if err != nil {
		return -1, errors.New(constants.InternalServerError)
	}
	qRow := s.client.QueryRow(
		ctx,
		"INSERT INTO public.\"user\" (phone_number, password_hash) values ($1, $2) RETURNING id",
		user.PhoneNumber, string(hashBytes),
	)
	var id int64
	err = qRow.Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			s.log.Errorf("SQL Error: %s, statuscode: %s", pgErr.Message, pgErr.Code)
			if pgErr.Code == "23505" {
				return -1, errors.New(constants.UserAlreadyExistsPhone)

			}
		}
		s.log.Errorf("Unable to create new user by phone number: %s", err)
		return -1, errors.New(constants.InternalServerError)
	}
	return id, nil
}

func (s *UserStorage) FindOne(ctx context.Context, id int) (model.User, error) {
	qRow := s.client.QueryRow(ctx,
		"select "+
			"id, "+
			"name, "+
			"phone_number,"+
			"email,"+
			"registration_date,"+
			"rating"+
			" from public.user where id=$1", id)
	var Id int
	var name string
	var phoneNumber string
	var email string
	var registrationDate time.Time
	var rating float64
	err := qRow.Scan(&Id, &name, &phoneNumber, &email, &registrationDate, &rating)
	if err != nil {
		fmt.Println(err)
		return model.User{}, nil
	}
	return model.User{Id: Id, Name: name, PhoneNumber: phoneNumber,
		Email: email, RegistrationDate: registrationDate, Rating: rating}, nil
}

func (s *UserStorage) Update(ctx context.Context, user model.UpdateUserDTO) {
	//TODO implement me
	panic("implement me")
}

func (s *UserStorage) Delete(ctx context.Context, id int) error {
	_, err := s.client.Exec(ctx, "delete from public.user where id=$1", id)
	if err != nil {
		return errors.New(constants.InternalServerError)
	}
	return nil
}

func NewStorage(client client.Client, log *logrus.Logger) Storage {
	return &UserStorage{
		client: client,
		log:    log,
	}
}
