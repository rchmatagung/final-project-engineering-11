package service

import (
	"context"
	"errors"

	"github.com/rg-km/final-project-engineering-11/backend/model"
)

func (r *AuthServiceimpl) DeleteById(id int) error {
	ctx := context.Background()
	err := r.userRepo.DeleteUserById(ctx, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *AuthServiceimpl) CreateArtikel(artikel *model.Artikel) error {
	ctx := context.Background()
	available, err := r.artikelRepo.CheckArtikel(ctx, artikel.Title)
	if err != nil {
		return err
	}
	if available {
		return errors.New("Artikel Sudah Terdaftar")
	}
	err = r.artikelRepo.CreateArtikel(ctx, artikel)
	if err != nil {
		return err
	}
	return nil
}
