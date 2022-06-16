package service

import (
	"context"
	"errors"

	"github.com/rg-km/final-project-engineering-11/backend/model"
	"github.com/rg-km/final-project-engineering-11/backend/repository"
	"github.com/rg-km/final-project-engineering-11/backend/secure"
)

type AuthServiceimpl struct {
	userRepo repository.UserInterface
}

func NewAuthService(userRepo repository.UserInterface) AuthService {
	return &AuthServiceimpl{userRepo}
}

func (a *AuthServiceimpl) Login(loginReq model.PayloadUser) (string, error) {
	ctx := context.Background()
	findUser, _ := a.userRepo.GetByUsername(ctx, loginReq.Username)

	if findUser != nil {
		hashPwd := findUser.Password
		pwd := loginReq.Password

		errhash := secure.VerifyPassword(hashPwd, pwd)
		if errhash == nil {
			token, err := secure.GenerateToken(findUser.Username, findUser.ID, findUser.Role)

			if err != nil {
				return "", err
			}

			return token, nil
		} else {
			return "", errors.New("Password Salah")
		}
	} else {
		return "", errors.New("USER Tidak Ditemukan")
	}
}

func (a *AuthServiceimpl) Register(register model.UserRegis) error {
	username, _ := a.userRepo.GetByUsername(context.Background(), register.Username)
	if username != nil {
		return errors.New("USERNAME SUDAH TERDAFTAR")
	}

	hash, err := secure.HashPassword(register.Password)
	if err != nil {
		return err
	}

	register.Password = hash

	err = a.userRepo.CreateUser(context.Background(), &register)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthServiceimpl) GetRoleByUserName(username string) (string, error) {
	ctx := context.Background()
	findUser, err := a.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	return findUser.Role, nil
}

func (r *AuthServiceimpl) GetIdByUserName(username string) (int, error) {
	ctx := context.Background()
	data, err := r.userRepo.GetIdByUserName(ctx, username)
	if err != nil {
		return 0, err
	}
	return data, nil
}
