package handler

import (
	"company-keeper-go/company"
	"company-keeper-go/db"
	"company-keeper-go/models"
	"company-keeper-go/router"
	"company-keeper-go/store"
	"company-keeper-go/user"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	d  *gorm.DB
	us user.Store
	co company.Store
	h  *Handler
	e  *echo.Echo
)

func authHeader(token string) string {
	return "Bearer " + token
}

func responseMap(b []byte) map[string]interface{} {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil
	}
	return m
}

func setup() {
	var err = godotenv.Load("../.test.env")
	if err != nil {
		panic(err)
	}
	d = db.DatabaseInit()
	db.AutoMigrate(d)
	db.TruncateTables(d, []string{"users", "companies"})

	us = store.NewUserStore(d)
	co = store.NewCompanyStore(d)
	h = NewHandler(us, co)
	e = router.New()
	err = loadFixtures()
	if err != nil {
		panic(err)
	}
}

func loadFixtures() error {
	u1 := models.User{
		Username: "Jake",
		Email:    "jake@gmail.com",
	}
	u1.Password, _ = u1.HashPassword("secret")
	if err := us.Create(&u1); err != nil {
		return err
	}

	return nil
}
