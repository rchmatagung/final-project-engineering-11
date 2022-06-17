package service

import (
	"github.com/rg-km/final-project-engineering-11/backend/model"
)

type AuthService interface {
	Login(payload model.PayloadUser) (string, error)
	Register(payload model.UserRegis) error
}
