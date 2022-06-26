package service

import "github.com/rg-km/final-project-engineering-11/backend/model"

type AdminService interface {
	DeleteById(id int) error
	CreateArtikel(artikel *model.Artikel) error
}
