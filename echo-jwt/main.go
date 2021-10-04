package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-study/echo-jwt/db"
	"go-study/echo-jwt/local"
)

var SIGNING_KEY = "secret"
var IS_DB = false

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var token string
	var err error
	if IS_DB {
		token, err = db.Login(username, password, SIGNING_KEY)
	} else {
		token, err = local.Login(username, password, SIGNING_KEY)
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	var name string
	if IS_DB {
		claims := db.GetUser(c)
		name = claims.Name

	} else {
		claims := local.GetUser(c)
		name = claims.Name
	}
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	if IS_DB {
		db.Migrate()
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	var config middleware.JWTConfig
	if IS_DB {
		config = middleware.JWTConfig{
			Claims:     &db.JwtCustomClaims{},
			SigningKey: []byte(SIGNING_KEY),
		}

	} else {
		config = middleware.JWTConfig{
			Claims:     &local.JwtCustomClaims{},
			SigningKey: []byte(SIGNING_KEY),
		}
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start("localhost" + ":1323"))
}
