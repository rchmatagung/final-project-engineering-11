package service

import (
	"context"
	"errors"

	"github.com/rg-km/final-project-engineering-11/backend/model"
	"github.com/rg-km/final-project-engineering-11/backend/secure"
)

func (r *AuthServiceimpl) GetAllMentor() ([]*model.MentorList, error) {
	ctx := context.Background()
	return r.mentorRepo.MentorList(ctx)
}

func (r *AuthServiceimpl) UpdateUserById(data *model.UserUpdate, id int, cookieid int) error {
	ctx := context.Background()
	if id != cookieid {
		return errors.New("Mau Ngapain Hayo")
	}
	data1, _ := r.userRepo.GetByID(ctx, id)
	alluser, _ := r.userRepo.GetAllUser(ctx)
	if data1.Username != data.Username || data1.Email != data.Email || data1.Phone != data.Phone {
		for _, user := range alluser {
			if user.ID != id && user.Username == data.Username {
				return errors.New("Username Sudah Terdaftar")
			} else if user.ID != id && user.Email == data.Email {
				return errors.New("Email Sudah Terdaftar")
			} else if user.ID != id && user.Phone == data.Phone {
				return errors.New("Phone Sudah Terdaftar")
			}
		}
	}

	if data1.Password != data.Password {
		hash, _ := secure.HashPassword(data.Password)
		data.Password = hash
	} else {
		data.Password = data1.Password
	}
	err := r.userRepo.UpdateById(ctx, data, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthServiceimpl) RegisMentor(data *model.MentorRegis) error {
	ctx := context.Background()
	err := r.mentorRepo.CheckEmailMentor(ctx, data.Email)
	if err != nil {
		return err
	}

	err1 := r.mentorRepo.RegisMentor(ctx, data)
	if err != nil {
		return err1
	}
	return nil
}

func (r *AuthServiceimpl) GetMentorBySkill(skill string) ([]*model.MentorSkill, error) {
	ctx := context.Background()
	data, err := r.mentorRepo.GetMentorByskill(ctx, skill)
	valid := r.mentorRepo.CheckMentorBySkill(ctx, skill)
	if valid != nil {
		return nil, errors.New("Category Skill Tidak Ditemukan")
	}

	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("Mentor Belum Tersedia")
	}
	return data, nil

}

func (r *AuthServiceimpl) GetMentorById(id int) (*model.MentorDetail, error) {
	ctx := context.Background()
	data, err := r.mentorRepo.GetMentorById(ctx, id)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *AuthServiceimpl) GetUserDataById(id int) (*model.UserDetail, error) {
	ctx := context.Background()

	data, err := r.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (r *AuthServiceimpl) GetIdByUserName(username string) (int, error) {
	ctx := context.Background()
	data, err := r.userRepo.GetIdByUserName(ctx, username)
	if err != nil {
		return 0, err
	}
	return data, nil
}

func (a *AuthServiceimpl) GetRoleByUserName(username string) (string, error) {
	ctx := context.Background()
	findUser, err := a.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	return findUser.Role, nil
}
