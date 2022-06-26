package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/rg-km/final-project-engineering-11/backend/model"
	"github.com/rg-km/final-project-engineering-11/backend/repository"
	"github.com/rg-km/final-project-engineering-11/backend/secure"
)

type AuthServiceimpl struct {
	userRepo    repository.UserRepo
	bookRepo    repository.BookRepository
	mentorRepo  repository.MentorRepository
	artikelRepo repository.ArtikelInterface
}

func NewAuthService(userRepo repository.UserRepo, bookRepo repository.BookRepository, mentor repository.MentorRepository, artikel repository.ArtikelInterface) (AuthService, UserService, BookService, AdminService) {
	return &AuthServiceimpl{userRepo, bookRepo, mentor, artikel}, &AuthServiceimpl{userRepo, bookRepo, mentor, artikel}, &AuthServiceimpl{userRepo, bookRepo, mentor, artikel}, &AuthServiceimpl{userRepo, bookRepo, mentor, artikel}
}

func (a *AuthServiceimpl) Login(loginReq model.PayloadUser) (string, error) {
	ctx := context.Background()
	findUser, _ := a.userRepo.GetByUsername(ctx, loginReq.Username)

	if findUser.Username != "" {
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
	data, _ := a.userRepo.GetByUsername(context.Background(), register.Username)
	fmt.Println(data)
	if data.Username != "" {
		return errors.New("Username Sudah Terdaftar")
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
