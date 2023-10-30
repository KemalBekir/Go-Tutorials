package services

import "github.com/KemalBekir/Go-Tutorials/go-mongodb/models"

type UserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
}
