package repository

import (
	"context"

	"github.com/rg-km/final-project-engineering-11/backend/model"
)

type UserInterface interface {
	GetAllUser(ctx context.Context) ([]*model.UserList, error)
	GetByID(ctx context.Context, id int) (*model.UserList, error)
	MentorList(ctx context.Context) ([]*model.MentorList, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.UserRegis) error
	UpdateById(ctx context.Context, user *model.UserUpdate, id int) error
	DeleteUserById(ctx context.Context, id int) error
	FetchUserByRole(ctx context.Context, role string) ([]*model.UserList, error)
	GetIdByUserName(ctx context.Context, username string) (int, error)
}
