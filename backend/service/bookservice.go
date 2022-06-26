package service

import (
	"context"
	"errors"

	"github.com/rg-km/final-project-engineering-11/backend/mail"
	"github.com/rg-km/final-project-engineering-11/backend/model"
)

func (r *AuthServiceimpl) CreateRequest(memberid, mentorid int) {
	ctx := context.Background()
	mentorEmail, _ := r.mentorRepo.GetMentorEmailById(ctx, mentorid)
	mentorname, _ := r.mentorRepo.GetMentorById(ctx, mentorid)
	data := r.bookRepo.CheckRowsBookId(ctx)
	bookid := mail.CreateBookId(data)
	_ = mail.SendMail(mentorEmail, bookid, mentorname.Name)

	_ = r.bookRepo.CreateRequest(ctx, memberid, mentorid, bookid)

}

func (r *AuthServiceimpl) GetAllBookStatusMember(id int) ([]*model.BookListStatus, error) {
	ctx := context.Background()
	data, err := r.bookRepo.GetAllBookStatusMemberId(ctx, id)
	if err != nil {
		return nil, errors.New("Gagal Mendapatkan Data")
	}
	return data, nil
}

func (r *AuthServiceimpl) UpdateStatusBooking(bookid string) error {
	ctx := context.Background()
	err := r.bookRepo.UpdateStatusBooking(ctx, bookid)
	if err != nil {
		return errors.New("Gagal Mengupdate Status")
	}
	return nil
}

func (r *AuthServiceimpl) CheckFound(bookid string) bool {

	ctx := context.Background()
	isfound := r.bookRepo.CheckAvailableBookId(ctx, bookid)
	if isfound {
		return true
	}
	return false
}
