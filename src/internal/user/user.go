package user

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
	email    string
	password string
}

type CreateByPhoneNumberDTO struct {
	phoneNumber string
	password    string
}
