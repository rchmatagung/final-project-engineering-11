package secure

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-11/backend/model"
)

func ExtractAuthToken(c *gin.Context) (*model.Authorize, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		id, ok := claims["id"].(float64)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		username, ok := claims["username"].(string)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		role, ok := claims["role"].(string)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		cookierole, _ := c.Cookie("RLPP")
		_, err := VerifyCookie(cookierole, role)
		if err != nil {
			return nil, err
		}
		authDetail := model.Authorize{
			ID:       int(id),
			Username: username,
			Role:     role,
		}

		return &authDetail, nil
	}

	return nil, errors.New("UnAuthorized")
}
