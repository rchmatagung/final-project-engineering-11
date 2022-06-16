package secure

import (
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
		id := claims["id"].(float64)
		username := claims["username"].(string)
		role := claims["role"].(string)
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

	return nil, err
}
