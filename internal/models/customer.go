package models

import "time"

type Customer struct {
	ID        string    `json:"id" db:"id" example:"256c1214-3385-4235-9cfe-1dc85a5f2a46" format:"uuid"`
	FirstName string    `json:"firstName" db:"first_name" example:"Homer"`
	LastName  string    `json:"lastName" db:"last_name" example:"Simpson"`
	Address   string    `json:"address" db:"address" example:"742 Evergreen Terrace"`
	StartDate time.Time `json:"startDate" db:"start_date" example:"2023-06-26T17:45:00.408032Z" format:"iso"`
	EndDate   time.Time `json:"endDate" db:"end_date" example:"2023-06-27T17:45:00.408032Z" format:"iso"`
	IsActive  bool      `json:"isActive" db:"is_active" example:"false"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" example:"2023-06-26T17:45:00.408032Z" format:"iso"`
}
