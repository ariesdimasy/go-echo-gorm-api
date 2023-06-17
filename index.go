package main

import (
	"go-echo-gorm-api/config"
	"go-echo-gorm-api/controllers"
	"go-echo-gorm-api/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	db, err := config.Connection()
	if err != nil {
		panic("error connection " + err.Error())
	}

	err2 := db.AutoMigrate(&models.Product{})
	if err2 != nil {
		panic("error connection " + err.Error())
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	pc := controllers.ProductController{
		DB: db,
	}

	e.GET("products", pc.GetProducts)
	e.POST("products", pc.CreateProduct)

	// Start server
	e.Start(":4570")

}
