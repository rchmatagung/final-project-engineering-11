package repository

import (
	"context"

	"github.com/rg-km/final-project-engineering-11/backend/model"
)

type UserRepo interface {
	GetByID(ctx context.Context, id int) (*model.UserDetail, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.UserRegis) error
	UpdateById(ctx context.Context, user *model.UserUpdate, id int) error
	DeleteUserById(ctx context.Context, id int) error
	FetchUserByRole(ctx context.Context, role string) ([]*model.UserList, error)
	GetAllUser(ctx context.Context) ([]*model.UserList, error)
	GetIdByUserName(ctx context.Context, username string) (int, error)
}
