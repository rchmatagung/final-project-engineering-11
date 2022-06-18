package secure

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-11/backend/config"
)

func GetAuthentication(c *gin.Context) error {
	token, err := VerifyToken(c)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

func VerifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := TakeToken(c)
	if len(tokenString) == 0 {
		return nil, errors.New("Token not found")
	}
	tokendecode, err2 := Decrypt([]byte(tokenString), config.KEYPRIVATE)
	if err2 != nil {
		return nil, err2
	}
	token, err := jwt.Parse(tokendecode, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Configuration.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func TakeToken(c *gin.Context) string {
	keys, err := c.Cookie("token")
	if err == http.ErrNoCookie {
		return ""
	}

	return keys

}
