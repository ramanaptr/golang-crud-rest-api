package main

import (
	"fmt"
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/database"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database with Gorm
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the echo
	e := echo.New()

	// Root level middleware of echo
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	// TODO: JWT Auth https://echo.labstack.com/docs/middleware/jwt

	// Register Routes with echo
	RegisterProductRoutes(e.Group("/product"))

	// Start the server with echo
	log.Println(fmt.Sprintf("Starting Server on port %s:%s", AppConfig.Domain, AppConfig.Port))
	log.Fatal(e.Start(":" + AppConfig.Port))
}

func RegisterProductRoutes(g *echo.Group) {
	productsWithIdEndpoint := "/:id"
	g.GET("/with/count", controllers.GetAllProducts)
	g.GET(productsWithIdEndpoint, controllers.GetProductById)
	g.POST("", controllers.CreateProduct)
	g.PUT(productsWithIdEndpoint, controllers.UpdateProduct)
	g.DELETE(productsWithIdEndpoint, controllers.DeleteProduct)
}
