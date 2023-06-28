package customers

import (
	"encoding/json"
	"net/http"

	"github.com/ansel1/merry"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/models"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

type Service interface {
	Create(models.CreateCustomer) (*models.Customer, error)
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s *service) Create(c models.CreateCustomer) (*models.Customer, error) {
	/* -------------------------- Data validation -------------------------- */
	if err := c.Validate(); err != nil {
		return nil, merry.Wrap(err).
			WithHTTPCode(http.StatusBadRequest).
			WithUserMessage(err.Error())
	}

	/* -------------------------- Type conversion -------------------------- */
	// Convert a CreateCustomer into a Customer to save it
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merry.Wrap(err)
	}

	data := models.Customer{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, merry.Wrap(err)
	}

	/* -------------------------- Save on database -------------------------- */
	res := s.db.Create(&data)
	if err := res.Error; err != nil {
		return nil, merry.Wrap(err)
	}

	if res.RowsAffected != 1 {
		errMsg := "the customer was not created"
		return nil, merry.New(errMsg).WithUserMessage(errMsg)
	}

	return &data, nil
}
