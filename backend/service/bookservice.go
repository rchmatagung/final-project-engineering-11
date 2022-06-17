package service

import (
	"context"
	"errors"

	"github.com/rg-km/final-project-engineering-11/backend/mail"
	"github.com/rg-km/final-project-engineering-11/backend/model"
)

func (r *AuthServiceimpl) CreateRequest(memberid, mentorid int) error {
	ctx := context.Background()
	mentorEmail, err1 := r.mentorRepo.GetMentorEmailById(ctx, mentorid)
	mentorname, err2 := r.mentorRepo.GetMentorById(ctx, mentorid)

	if err2 != nil {
		return err2
	}

	if err1 != nil {
		return err1
	}
	data := r.bookRepo.CheckRowsBookId(ctx)
	bookid := mail.CreateBookId(data)
	err3 := mail.SendMail(mentorEmail, bookid, mentorname.Name)
	if err3 != nil {
		return errors.New("Gagal Mengirim Email")
	}

	err := r.bookRepo.CreateRequest(ctx, memberid, mentorid, bookid)
	if err != nil {
		return errors.New("Gagal Membuat Request")
	}

	return nil

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
