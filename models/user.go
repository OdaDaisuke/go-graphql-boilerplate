package models

import "errors"

type User struct {
	UserId      string `json:"userId"`
	Email       string `json:"email"`
}

func CreateUser(userId, email string) (*User, error){
	if userId == "" {
		return nil, errors.New("userName is empty")
	}
	if email == "" {
		return nil, errors.New("email is empty")
	}
	return &User{
		UserId: userId,
		Email: email,
	}, nil
}