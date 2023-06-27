package db

import (
	"log"

	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/config"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(c config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(c.DBDns()), &gorm.Config{})
	if err != nil {
		log.Fatalf("error opening database: %s\n", err.Error())
	}

	if err := db.AutoMigrate(&models.Customer{}, &models.WorkOrders{}); err != nil {
		log.Fatalf("error migrating db: %s\n", err.Error())
	}

	return db
}
