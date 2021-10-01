package db

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"uniqueIndex"`
	Password string
}

type JwtCustomClaims struct {
	Name  string `json:"name"`
	jwt.StandardClaims
}

var migrates = []interface{}{
	&User{},
}

var db *gorm.DB

func Open() *gorm.DB{
	if db == nil {
		var err error
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
		  panic("failed to connect database")
		}
	}
	return db
}


// Migrate the schema
func Migrate() {
	log.Println("Migrate")
	// migrationファイルなどの作成はされず、追加や変更にのみ対応(migration処理は別ライブラリでやったほうが良い)
	db := Open()
	db.AutoMigrate(migrates...)

	// 初期データセット
	var user User
	res := db.Where(&User{Name: "jon"}).First(&user)
	if err := res.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("create user")
			user = User{Name: "jon", Password: "shhh!"}
			res := db.Create(&user)
			if err := res.Error; err != nil {
				log.Panic(err)
			}
		} else {
			log.Panic(err)
		}
	}
}

func Login(username string, password string, signingKey string) (string, error) {
	// Throws unauthorized error
	var user User
	res := Open().Where(&User{Name: username, Password: password}).First(&user)
	if err := res.Error; err != nil {
		return "", err
	}

	// Set custom claims(返却用の情報)
	claims := &JwtCustomClaims{
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return t, err
}

func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}