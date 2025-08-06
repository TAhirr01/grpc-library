package database

import (
	"github.com/TAhirr01/grpc-library/internal/config"
	"github.com/TAhirr01/grpc-library/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	DB *gorm.DB
}

func NewConnection(cfg *config.Config) *Database {
	db, err := gorm.Open(mysql.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	return &Database{DB: db}
}

func (d *Database) AutoMigrate() error {
	return d.DB.AutoMigrate(
		&models.User{},
		&models.Book{},
	)
}

func (d *Database) GetDB() *gorm.DB {
	return d.DB
}