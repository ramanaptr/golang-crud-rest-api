package controllers

import (
	"golang-crud-rest-api/core"
	"golang-crud-rest-api/entities"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthLogin(c echo.Context) error {
	user := new(entities.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	// Throws unauthorized error
	if user.Username != "ramanaptr" || user.Password != "lupapassword" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &core.JwtCustomClaims{
		Name:  "Ramana Putra",
		Admin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(core.AppConfig.Secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func Me(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*core.JwtCustomClaims)
	name := claims.Name
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Welcome Back Your Name is " + name,
		"name":    name,
	})
}
