package customers

import (
	"github.com/ansel1/merry/v2"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/models"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

type Service interface {
	List() ([]models.Customer, error)
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s *service) List() ([]models.Customer, error) {
	customers := []models.Customer{}
	res := s.db.Find(&customers)

	if res.Error != nil {
		return nil, merry.Wrap(res.Error)
	}

	return customers, nil
}
