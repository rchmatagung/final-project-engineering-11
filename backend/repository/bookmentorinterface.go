package repository

import (
	"context"

	"github.com/rg-km/final-project-engineering-11/backend/model"
)

type BookRepository interface {
	GetAllBookid(ctx context.Context) ([]*model.BookMentor, error)
	GetAllBookStatusMemberId(ctx context.Context, userid int) ([]*model.BookListStatus, error)
	CreateRequest(ctx context.Context, userid int, mentorid int, bookid string) error
	UpdateStatusBooking(ctx context.Context, bookid string) error
	CheckAvailableBookId(ctx context.Context, bookid string) bool
	CheckRowsBookId(ctx context.Context) int
}
