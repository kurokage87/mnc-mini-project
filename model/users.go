package model

import (
	"github.com/google/uuid"
)

type RequestRegister struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Pin         string `json:"pin"`
}

type DataResponseRegister struct {
	UserId      uuid.UUID `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Pin         string    `json:"pin"`
	CreatedDate string    `json:"created_date"`
}

type RequestLogin struct {
	PhoneNumber string `json:"phone_number"`
	Pin         string `json:"pin"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RequestUpdateProfile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

type DataResponseUpdateProfile struct {
	UserId     string `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Address    string `json:"address"`
	UpdateDate string `json:"update_date"`
}
