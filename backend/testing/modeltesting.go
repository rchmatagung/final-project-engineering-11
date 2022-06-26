package testing_test

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-11/backend/model"
)

type UserMentorDetailTesting struct {
	Data   model.MentorDetail
	Status int
}

type UserProfileTesting struct {
	Data   model.UserDetail
	Status int `json:"status"`
}

type MentorListTesting struct {
	Data   []*model.MentorList
	Status int `json:"status"`
}

type MentorSkillTesting struct {
	Data   []*model.MentorSkill
	Status int `json:"status"`
}
type GeneralTesting struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
type GeneralErrorTesting struct {
	Message string `json:"error"`
	Status  int    `json:"status"`
}

type StatusBookTesting struct {
	Data   []*model.BookListStatus
	Status int `json:"status"`
}

type ArtikelTesting struct {
	Data   []*model.ArtikelList
	Status int `json:"status"`
}
type ArtikelDetailTesting struct {
	Data   *model.ArtikelDetail
	Status int `json:"status"`
}

func GetStatusBookId(sql *sql.DB, bookid string) string {
	var status string
	query := "SELECT status FROM bookmentor WHERE bookid = ?"
	err := sql.QueryRow(query, bookid).Scan(&status)
	if err != nil {
		return ""
	}
	return status

}
