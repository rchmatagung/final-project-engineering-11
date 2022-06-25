package repository

import (
	"context"

	"github.com/rg-km/final-project-engineering-11/backend/model"
)

type MentorRepository interface {
	GetMentorEmailById(ctx context.Context, id int) (string, error)
	GetMentorById(ctx context.Context, id int) (*model.MentorDetail, error)
	MentorList(ctx context.Context) ([]*model.MentorDetail, error)
	CheckMentorBySkill(ctx context.Context, skill string) error
	GetMentorByskill(ctx context.Context, skill string) ([]*model.MentorSkill, error)
	RegisMentor(ctx context.Context, user *model.MentorRegis) error
	CheckEmailMentor(ctx context.Context, email string) error
	GetDataMentorByNoBooking(ctx context.Context, noorder string) (*model.MentorKontak, error)
}
