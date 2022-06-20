package repository

import (
	"context"
	"errors"
	"time"

	"github.com/rg-km/final-project-engineering-11/backend/model"
)

func (r *UserRepository) CreateArtikel(ctx context.Context, artikel *model.Artikel) error {
	query := `INSERT INTO artikel (judul, content, created_at) VALUES (?,?,?)`
	_, err := r.db.ExecContext(ctx, query, artikel.Title, artikel.Content, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) CheckArtikel(ctx context.Context, title string) (bool, error) {
	query := `SELECT judul FROM artikel WHERE judul = ?`
	row := r.db.QueryRowContext(ctx, query, title)
	var title1 string
	err := row.Scan(&title1)
	if err != nil {
		return false, nil
	}
	if len(title1) > 0 {
		return true, nil
	}
	return false, nil
}

func (r *UserRepository) GetAllArtikel(ctx context.Context) ([]*model.ArtikelList, error) {
	query := `SELECT id, judul, created_at,content FROM artikel`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var artikel []*model.ArtikelList
	for rows.Next() {
		var artikel1 model.ArtikelList
		err := rows.Scan(&artikel1.ID, &artikel1.Title, &artikel1.Date, &artikel1.Content)
		if err != nil {
			return nil, err
		}
		artikel = append(artikel, &artikel1)
	}
	return artikel, nil

}

func (r *UserRepository) GetArtikelById(ctx context.Context, id int) (*model.ArtikelDetail, error) {
	query := `SELECT judul, created_at,content FROM artikel WHERE id = ?`
	row, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var artikel model.ArtikelDetail
	if row.Next() {
		err := row.Scan(&artikel.Title, &artikel.Date, &artikel.Content)
		if err != nil {
			return nil, err
		}
		return &artikel, nil
	}
	return nil, errors.New("Artikel Tidak Ditemukan")

}
