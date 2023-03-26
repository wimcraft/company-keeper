package main

import (
	"company-keeper-go/db"
	_ "company-keeper-go/docs"
	"company-keeper-go/handler"
	"company-keeper-go/router"
	"company-keeper-go/store"
	"fmt"
	"github.com/joho/godotenv"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
	"log"
	"os"
)

// @title Company Keeper API
// @version 1.0
// @description This is a sample server
// @termsOfService http://swagger.io/terms/
// @BasePath /api
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := router.New()

	r.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := r.Group("/api")

	d := db.DatabaseInit()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	co := store.NewCompanyStore(d)
	h := handler.NewHandler(us, co)

	h.Register(v1)
	PORT := os.Getenv("PORT")
	r.Logger.Fatal(r.Start(fmt.Sprintf(":%v", PORT)))
}
