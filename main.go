package main

import (
	"fmt"
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/core"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load Configurations from config.json using Viper
	core.LoadAppConfig()

	// Initialize Database with Gorm
	database.Connect(core.AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the echo
	e := echo.New()

	// Root level middleware of echo

	// Middleware to parse form data
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Configure middleware with the custom claims type
	// Experiment Mode TODO by Rama
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(core.JwtCustomClaims)
		},
		SigningKey: []byte(core.AppConfig.Secret),
		SuccessHandler: func(c echo.Context) {
			fmt.Println("matnap")
			user, ok := c.Get("user").(*jwt.Token)
			if ok {
				claims := user.Claims.(*core.JwtCustomClaims)
				//custom loader //https://echo.labstack.com/guide/context
				// appContext := c.(*common.AppContext)
				user := &entities.User{}
				user.Username = claims.Name
				// appContext.User = user
				fmt.Println("Claim Admin : ")
				fmt.Println(claims.Admin)
				fmt.Println("Standart Claim ID: ")
				fmt.Println(claims.ExpiresAt)
			}
		},
	}

	// Restricted Access Group
	r := e.Group("/")
	r.Use(echojwt.WithConfig(config))

	// Register Routes with echo
	AuthRoutes(e.Group("auth"), r.Group("auth"))
	ProductRoutes(r.Group("product"))

	// Start the server with echo
	log.Println(fmt.Sprintf("Starting Server on port %s:%s", core.AppConfig.Domain, core.AppConfig.Port))
	log.Fatal(e.Start(":" + core.AppConfig.Port))
}

func AuthRoutes(g *echo.Group, r *echo.Group) {
	// Un-Restricted
	g.POST("/login", controllers.AuthLogin)

	// Restricted
	r.POST("/me", controllers.Me)
}

func ProductRoutes(g *echo.Group) {
	// Restricted
	productsWithIdEndpoint := "/:id"
	g.GET("/with/count", controllers.GetAllProducts)
	g.GET(productsWithIdEndpoint, controllers.GetProductById)
	g.POST("", controllers.CreateProduct)
	g.PUT(productsWithIdEndpoint, controllers.UpdateProduct)
	g.DELETE(productsWithIdEndpoint, controllers.DeleteProduct)
}
