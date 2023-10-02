package controllers

import (
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"golang-crud-rest-api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Echo Rules
// "product := new(entities.Product)" variable new(...) to get body from request

// Gorm Rules
// "var product entities.Product" variable to copy properties on the Product struct in database

func checkIfProductExists(productId string) bool {
	var product entities.Product
	database.Instance.First(&product, productId)
	if product.ID == 0 {
		return false
	}
	return true
}

func CreateProduct(c echo.Context) (err error) {
	product := new(entities.Product)

	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	database.Instance.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

func GetProductById(c echo.Context) error {
	productId := c.Param("id")

	if !checkIfProductExists(productId) {
		return echo.ErrNotFound
	}

	var product entities.Product
	if err := database.Instance.First(&product, productId).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.Report{
			Message: "Error retrieving product",
		})
	}

	database.Instance.First(&product, productId)
	return c.JSON(http.StatusOK, product)
}

func GetAllProducts(c echo.Context) error {
	var products []entities.Product

	database.Instance.Find(&products)
	return c.JSON(http.StatusOK, model.WithCount{
		Data:  products,
		Count: int64(len(products)),
	})
}

func UpdateProduct(c echo.Context) (err error) {
	productId := c.Param("id")
	product := new(entities.Product)

	// Validate
	if err = c.Bind(product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !checkIfProductExists(productId) {
		return echo.ErrNotFound
	}

	database.Instance.First(&product, productId)
	database.Instance.Save(&product)

	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	productId := c.Param("id")

	if !checkIfProductExists(productId) {
		return echo.ErrNotFound
	}

	var product entities.Product
	database.Instance.Delete(&product, productId)
	return c.JSON(http.StatusOK, model.Report{
		Message: "Product Deleted Successfully!",
	})
}
