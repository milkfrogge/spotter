package model

import "time"

type User struct {
	Id               int
	Name             string
	PhoneNumber      string
	Email            string
	RegistrationDate time.Time
	Rating           float64
}

type CreateByEmailDTO struct {
	Email    string
	Password string
}

type CreateByPhoneNumberDTO struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UpdateUserDTO struct {
	Id          int    `json:"name"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
