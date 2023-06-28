package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Customer struct {
	ID        string    `json:"id" example:"256c1214-3385-4235-9cfe-1dc85a5f2a46" format:"uuid" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	FirstName string    `json:"firstName" example:"Homer" gorm:"not null;type:varchar(100)"`
	LastName  string    `json:"lastName" example:"Simpson" gorm:"not null;type:varchar(100)"`
	Address   string    `json:"address" example:"742 Evergreen Terrace" gorm:"not null;type:varchar(100)"`
	StartDate time.Time `json:"startDate" example:"0001-01-01T00:00:00Z" format:"iso"`
	EndDate   time.Time `json:"endDate" example:"0001-01-01T00:00:00Z" format:"iso"`
	IsActive  bool      `json:"isActive" example:"false" gorm:"not null;default:false"`
	CreatedAt time.Time `json:"createdAt" example:"2023-06-26T17:45:00.408032Z" format:"iso" gorm:"default:(now() at time zone 'utc')"`
}

type CreateCustomer struct {
	FirstName string `json:"firstName" example:"Homer" binding:"required"`
	LastName  string `json:"lastName" example:"Simpson" binding:"required"`
	Address   string `json:"address" example:"742 Evergreen Terrace" binding:"required"`
}

func (c CreateCustomer) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.FirstName, validation.Required),
		validation.Field(&c.LastName, validation.Required),
		validation.Field(&c.Address, validation.Required),
	)
}
