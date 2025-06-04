package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"stvCms/internal/models"
)

var (
	DB_USER     = os.Getenv("")
	DB_PASSWORD = os.Getenv("")
	DB_HOST     = os.Getenv("")
	DB_PORT     = os.Getenv("")
	DB          = os.Getenv("")
)

func Init() *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		panic("Error al cargar la bd")
	}

	db.AutoMigrate(
		&models.Post{},
	)

	return db
}
