package local

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func Login(username string, password string, signingKey string) (string, error) {

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return "", echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 必要であればここでkidを付与できる
	token.Header["kid"] = "abc"

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return t, err
}

func GetUser(c echo.Context) *JwtCustomClaims {
	i := c.Get("user")
	log.Println(i)

	user := i.(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	header := user.Header
	log.Println("claims", claims)
	log.Println("header", header)
	return claims
}