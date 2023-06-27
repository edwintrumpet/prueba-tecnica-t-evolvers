package models

import "time"

type Customer struct {
	ID        string    `json:"id" example:"256c1214-3385-4235-9cfe-1dc85a5f2a46" format:"uuid" gorm:"column:id;type:uuid;primary_key;default:gen_random_uuid()"`
	FirstName string    `json:"firstName" example:"Homer" gorm:"column:first_name;not null;type:varchar(100)"`
	LastName  string    `json:"lastName" example:"Simpson" gorm:"column:last_name;not null;type:varchar(100)"`
	Address   string    `json:"address" example:"742 Evergreen Terrace" gorm:"column:address;not null;type:varchar(100)"`
	StartDate time.Time `json:"startDate" example:"2023-06-26T17:45:00.408032Z" format:"iso" gorm:"column:start_date"`
	EndDate   time.Time `json:"endDate" example:"2023-06-27T17:45:00.408032Z" format:"iso" gorm:"column:end_date"`
	IsActive  bool      `json:"isActive" example:"false" gorm:"column:is_active;not null;default:false"`
	CreatedAt time.Time `json:"createdAt" example:"2023-06-26T17:45:00.408032Z" format:"iso" gorm:"column:created_at;default:(now() at time zone 'utc')"`
}
