package service

import "github.com/rg-km/final-project-engineering-11/backend/model"

type UserService interface {
	UpdateUserById(data *model.UserUpdate, id int, cookieid int) error
	RegisMentor(data *model.MentorRegis) error
	GetAllMentor() ([]*model.MentorList, error)
	GetMentorBySkill(skill string) ([]*model.MentorSkill, error)
	GetMentorById(id int) (*model.MentorDetail, error)
	GetUserDataById(id int) (*model.UserDetail, error)
	GetIdByUserName(username string) (int, error)
	CreateRequest(memberid, mentorid int) error
	GetAllBookStatusMember(id int) ([]*model.BookListStatus, error)
	GetRoleByUserName(username string) (string, error)
}
