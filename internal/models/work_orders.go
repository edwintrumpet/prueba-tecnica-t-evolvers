package models

import (
	"time"
)

type WorkOrder struct {
	ID               string    `json:"id" example:"256c1214-3385-4235-9cfe-1dc85a5f2a46" format:"uuid" gorm:"column:id;type:uuid;primary_key;default:gen_random_uuid()"`
	CustomerID       string    `json:"customerId" example:"256c1214-3385-4235-9cfe-1dc85a5f2a46" format:"uuid" gorm:"column:customer_id;type:uuid;not null"`
	Title            string    `json:"title" example:"something" gorm:"column:title;type:varchar(255);not null"`
	PlannedDateBegin time.Time `json:"plannedTimeBegin" example:"2023-06-27T17:45:00.408032Z" gorm:"column:planned_title_begin;not null"`
	PlannedDateEnd   time.Time `json:"plannedTimeEnd" example:"2023-06-27T17:45:00.408032Z" gorm:"column:planned_title_end;not null"`
	Status           string    `json:"status" example:"ok" gorm:"column:status;type:varchar(10);not null;default:new"`
	CreatedAt        time.Time `json:"createdAt" example:"2023-06-26T17:45:00.408032Z" format:"iso" gorm:"column:created_at;default:(now() at time zone 'utc')"`
	Customer         Customer
}
