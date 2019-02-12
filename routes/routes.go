package routes

import (
	"fmt"
	"net/http"

	"../controllers"
	"../models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Ruta accesible (pública)
func accessible(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>API - CLOUD</h1>")
}

// Rutas protegidas - valida token y devuelve el email
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.Claim)
	email := claims.Email
	return c.String(http.StatusOK, "Welcome "+email)
}

// InitRoutes - Init routes
func InitRoutes() {

	fmt.Printf("jjjj")

	e := echo.New()

	prefix := "api/v1"

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST(prefix+"/login", controllers.Login)

	// User route
	e.POST(prefix+"/user", controllers.UserCreate)

	// Unauthenticated route
	e.GET(prefix, accessible)
	// Restricted group
	r := e.Group(prefix + "/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &models.Claim{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":8080"))
}
