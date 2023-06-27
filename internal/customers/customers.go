package customers

import "github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/models"

type service struct{}

type Service interface {
	List() ([]models.Customer, error)
}

func New() Service {
	return &service{}
}

func (s *service) List() ([]models.Customer, error) {
	return []models.Customer{
		{
			FirstName: "Luisa",
			Address:   "from customers package",
		},
	}, nil
}
