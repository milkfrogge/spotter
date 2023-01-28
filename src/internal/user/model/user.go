package model

import "time"

type User struct {
	id               int
	name             string
	phoneNumber      int64
	email            string
	registrationDate time.Time
	passwordHash     string
}

type CreateByEmailDTO struct {
	Email    string
	Password string
}

type CreateByPhoneNumberDTO struct {
	PhoneNumber string
	Password    string
}
