package repositories

import "github.com/OdaDaisuke/go-graphql-sample/models"

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (_ *UserRepository) FindById(userId string) (*models.User, error) {
	return &models.User{
		UserId: "aaaa",
		Email: "sample@gmail.com",
	}, nil
}