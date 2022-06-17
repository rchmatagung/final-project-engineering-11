package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/rg-km/final-project-engineering-11/backend/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) (UserRepo, BookRepository, MentorRepository) {
	return &UserRepository{db}, &UserRepository{db}, &UserRepository{db}
}

func (u *UserRepository) GetAllUser(ctx context.Context) ([]*model.UserList, error) {
	var users []*model.UserList
	query := "SELECT id,username,name,role,address,phone,email,created_at FROM users"
	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.UserList
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Role, &user.Address, &user.Phone, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *UserRepository) GetByID(ctx context.Context, id int) (*model.UserDetail, error) {
	var user model.UserDetail
	rows, err := u.db.QueryContext(ctx, "SELECT id,username,name, password,role,address,phone,email FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.Role, &user.Address, &user.Phone, &user.Email)
		if err != nil {
			return nil, err
		}
		return &user, nil

	}
	return &user, errors.New("User not found")
}

func (u *UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	rows, err := u.db.QueryContext(ctx, "SELECT id,username,name,password,role,address,phone,email,created_at FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.Role, &user.Address, &user.Phone, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		return &user, nil

	}
	return &user, errors.New("User not found")
}

func (u *UserRepository) CreateUser(ctx context.Context, user *model.UserRegis) error {

	tx, _ := u.db.Begin()

	Role := "member"
	query := "INSERT INTO users (username,name,password,role,address,phone,email,created_at) VALUES (?,?,?,?,?,?,?,?)"

	_, err := tx.ExecContext(ctx, query, user.Username, user.Name, user.Password, Role, user.Address, user.Phone, user.Email, time.Now())
	if err != nil {
		tx.Rollback()
		return errors.New("Error creating user")
	}

	tx.Commit()
	return nil

}

func (u *UserRepository) UpdateById(ctx context.Context, user *model.UserUpdate, id int) error {
	tx, _ := u.db.Begin()
	query := "UPDATE users SET username = ?, name = ?, password = ?, address = ?, phone = ?, email = ? WHERE id = ?"

	res, err := tx.ExecContext(ctx, query, user.Username, user.Name, user.Password, user.Address, user.Phone, user.Email, id)

	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		tx.Rollback()
		return errors.New("User not found")
	}
	tx.Commit()
	return nil
}

func (u *UserRepository) DeleteUserById(ctx context.Context, id int) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	query := "DELETE FROM users WHERE id = ?"
	_, err = tx.ExecContext(ctx, query, id)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

func (u *UserRepository) FetchUserByRole(ctx context.Context, role string) ([]*model.UserList, error) {
	var users []*model.UserList
	rows, err := u.db.QueryContext(ctx, "SELECT id,username,role,created_at FROM users WHERE role = ?", role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.UserList
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Role, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *UserRepository) GetIdByUserName(ctx context.Context, username string) (int, error) {
	query := "SELECT id FROM users WHERE username = ?"
	var id int
	err := u.db.QueryRowContext(ctx, query, username).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
