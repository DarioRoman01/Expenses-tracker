package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to read env.")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to connect to postgres: %s", err)
	}

	return db, nil
}
