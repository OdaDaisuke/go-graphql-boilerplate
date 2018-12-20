package services

import (
	"github.com/OdaDaisuke/go-graphql-sample/repositories/user"
	"github.com/OdaDaisuke/go-graphql-sample/models"
)

func FindByUserId(userId string) (*models.User, error) {
	userRepo := repositories.NewUserRepository()
	return userRepo.FindById(userId)
}