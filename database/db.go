package database

import (
	"log"

	"github.com/PedroVMB/API-GO-CHECKLIST-CONDOMINIO/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Failed to connect a database")
	}

	DB.AutoMigrate(&models.Administrador{})
	DB.AutoMigrate(&models.Sindico{})
	DB.AutoMigrate(&models.Condominio{})
}
