package db

import (
	"company-keeper-go/models"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	portString := os.Getenv("POSTGRES_PORT")

	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatal("Error parsing POSTGRES_PORT")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}

	return database
}

func TruncateTables(db *gorm.DB, tables []string) {
	for _, v := range tables {
		_ = db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY;", v))
	}
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Company{},
		&models.User{},
	)
	if err != nil {
		panic(err)
	}
}
