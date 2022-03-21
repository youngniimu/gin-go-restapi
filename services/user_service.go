package services

import "go-rest-api/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUser(string) (*models.User, error)
	GetAll() ([]models.User, error)
	UpdateUser(*models.User, string) error
	DeleteUser(string) error
}
