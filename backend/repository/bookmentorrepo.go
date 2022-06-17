package repository

import (
	"context"
	"errors"
	"time"

	"github.com/rg-km/final-project-engineering-11/backend/model"
)

func (u *UserRepository) GetAllBookid(ctx context.Context) ([]*model.BookMentor, error) {
	var bookids []*model.BookMentor
	query := "SELECT id,bookid,user_id,mentor_id,status,created_at FROM bookmentor"
	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var bookid model.BookMentor
		err := rows.Scan(&bookid.ID, &bookid.BookID, &bookid.UserID, &bookid.MentorID, &bookid.Status, &bookid.CreatedAt)
		if err != nil {
			return nil, err
		}
		bookids = append(bookids, &bookid)
	}
	return bookids, nil

}

func (u *UserRepository) GetAllBookStatusMemberId(ctx context.Context, userid int) ([]*model.BookListStatus, error) {
	query := "SELECT u.bookid,u.status,m.name FROM bookmentor u INNER JOIN mentor m ON u.mentor_id = m.id WHERE u.member_id = ?"
	var bookids []*model.BookListStatus
	rows, err := u.db.QueryContext(ctx, query, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var bookid model.BookListStatus
		err := rows.Scan(&bookid.BookID, &bookid.Status, &bookid.MentorName)
		if err != nil {
			return nil, err
		}
		bookids = append(bookids, &bookid)
	}
	return bookids, nil
}

func (u *UserRepository) CreateRequest(ctx context.Context, userid, mentorid int, bookid string) error {
	query := "INSERT INTO bookmentor(member_id,mentor_id,bookid,status,created_at) VALUES (?,?,?,?,?)"
	tx, _ := u.db.Begin()
	_, err := tx.ExecContext(ctx, query, userid, mentorid, bookid, "Waiting", time.Now())
	if err != nil {
		tx.Rollback()
		return errors.New("Error creating Booking")
	}
	tx.Commit()

	return nil

}

func (a *UserRepository) UpdateStatusBooking(ctx context.Context, bookid string) error {
	query := "UPDATE bookmentor SET status = ? WHERE bookid = ?"
	tx, _ := a.db.Begin()
	_, err := tx.ExecContext(ctx, query, "Accepted", bookid)
	if err != nil {
		tx.Rollback()
		return errors.New("Error updating Booking")
	}
	tx.Commit()
	return nil
}

func (a *UserRepository) CheckAvailableBookId(ctx context.Context, bookid string) bool {

	query := "SELECT * FROM bookmentor WHERE bookid = ?"
	rows, err := a.db.QueryContext(ctx, query, bookid)
	if err != nil {
		return false
	}
	defer rows.Close()
	if rows.Next() {
		return true
	}
	return false
}
func (a *UserRepository) CheckRowsBookId(ctx context.Context) int {
	query := "SELECT * FROM bookmentor"
	rows, err := a.db.QueryContext(ctx, query)
	if err != nil {
		return 0
	}
	defer rows.Close()
	var count int = 0
	for rows.Next() {
		count++
	}
	return count
}
