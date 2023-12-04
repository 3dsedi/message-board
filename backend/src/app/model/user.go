package model

import (
	"time"
)

type CreateUserRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        string `json:"id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
}

func NewUser(userName, email, password string) *User {

	return &User{
		UserName:  userName,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
	}
}
