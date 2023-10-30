package services

import "github.com/KemalBekir/Go-Tutorials/go-mongodb/models"

type AuthService interface {
	SignUpUser(*models.SignInInput) (*models.DBResponse, error)
	SignInUser(*models.SignInInput) (*models.DBResponse, error)
}
