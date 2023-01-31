package model

import "time"

type User struct {
	id               int
	name             string
	phoneNumber      int64
	email            string
	registrationDate time.Time
	rating           float64
}

type CreateByEmailDTO struct {
	Email    string
	Password string
}

type CreateByPhoneNumberDTO struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
