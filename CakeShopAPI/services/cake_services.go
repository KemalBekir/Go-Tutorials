package services

import (
	"context"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/models"
)

type CakeService struct {
}

func (s *CakeService) CreateCake(ctx context.Context, cake *models.Cake) (*models.Cake, error) {

}
