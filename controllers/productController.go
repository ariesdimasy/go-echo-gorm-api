package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	models "go-echo-gorm-api/models"
)

type ProductController struct {
	DB *gorm.DB
}

func (pc *ProductController) CreateProduct(c echo.Context) error {
	//fmt.Println(" hey look ==> ", c.Request().Body)
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	fmt.Println(product)

	result := pc.DB.Create(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	return c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) GetProducts(c echo.Context) error {
	var products []models.Product
	err := pc.DB.Find(&products).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductByID(c echo.Context) error {
	var product models.Product
	err := pc.DB.First(&product).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, product)
}
