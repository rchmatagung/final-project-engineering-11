package service

import "github.com/rg-km/final-project-engineering-11/backend/model"

type BookService interface {
	CreateRequest(memberid, mentorid int)
	GetAllBookStatusMember(id int) ([]*model.BookListStatus, error)
	UpdateStatusBooking(bookid string) error
	CheckFound(bookid string) bool
}
