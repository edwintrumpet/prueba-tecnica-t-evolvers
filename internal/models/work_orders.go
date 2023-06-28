package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type workOrderStatus string

const (
	New       workOrderStatus = "new"
	Done                      = "done"
	Cancelled                 = "cancelled"
)

type WorkOrder struct {
	ID               string          `json:"id" example:"cd9bde09-5374-4749-86a6-34866c100e6e" format:"uuid" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CustomerID       string          `json:"customerId" example:"256c1214-3385-4235-9cfe-1dc85a5f2a46" format:"uuid" gorm:"type:uuid;not null"`
	Title            string          `json:"title" example:"something" gorm:"type:varchar(255);not null"`
	PlannedDateBegin time.Time       `json:"plannedTimeBegin" example:"2023-06-27T17:45:00.408032Z" gorm:"not null"`
	PlannedDateEnd   time.Time       `json:"plannedTimeEnd" example:"2023-06-27T17:45:00.408032Z" gorm:"not null"`
	Status           workOrderStatus `json:"status" example:"ok" gorm:"type:varchar(10);not null;default:new"`
	CreatedAt        time.Time       `json:"createdAt" example:"2023-06-26T17:45:00.408032Z" format:"iso" gorm:"default:(now() at time zone 'utc')"`
	Customer         Customer        `json:"customer"`
}

type CreateWorkOrder struct {
	CustomerID       string    `json:"customerId" example:"256c1214-3385-4235-9cfe-1dc85a5f2a46" format:"uuid" binding:"required"`
	Title            string    `json:"title" example:"something" binding:"required"`
	PlannedDateBegin time.Time `json:"plannedTimeBegin" example:"2023-06-27T17:45:00.408032Z" binding:"required"`
	PlannedDateEnd   time.Time `json:"plannedTimeEnd" example:"2023-06-27T17:45:00.408032Z" binding:"required"`
}

type FinishWorkOrder struct {
	CustomerID  string `json:"customerId" example:"256c1214-3385-4235-9cfe-1dc85a5f2a46" format:"uuid" binding:"required"`
	WorkOrderID string `json:"workOrderId" example:"cd9bde09-5374-4749-86a6-34866c100e6e" format:"uuid" binding:"required"`
}

func (c CreateWorkOrder) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.CustomerID, validation.Required, is.UUID),
		validation.Field(&c.Title, validation.Required),
		validation.Field(&c.PlannedDateBegin, validation.Required),
		validation.Field(&c.PlannedDateEnd, validation.Required),
	)
}

func (c FinishWorkOrder) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.CustomerID, validation.Required, is.UUID),
		validation.Field(&c.WorkOrderID, validation.Required, is.UUID),
	)
}
