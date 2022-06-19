package repository

import (
	"context"

	"github.com/rg-km/final-project-engineering-11/backend/model"
)

type ArtikelInterface interface {
	CreateArtikel(ctx context.Context, artikel *model.Artikel) error
	CheckArtikel(ctx context.Context, title string) (bool, error)
}
