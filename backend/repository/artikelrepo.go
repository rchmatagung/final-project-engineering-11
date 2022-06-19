package repository

import (
	"context"
	"errors"
	"time"

	"github.com/rg-km/final-project-engineering-11/backend/model"
)

func (r *UserRepository) CreateArtikel(ctx context.Context, artikel *model.Artikel) error {
	query := `INSERT INTO artikel (title, content, created_at) VALUES (?,?,?)`
	_, err := r.db.ExecContext(ctx, query, artikel.Title, artikel.Content, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) CheckArtikel(ctx context.Context, title string) (bool, error) {
	query := `SELECT title FROM artikel WHERE title = ?`
	row := r.db.QueryRowContext(ctx, query, title)
	var title1 string
	err := row.Scan(&title1)
	if err != nil {
		return false, err
	}
	if len(title1) > 0 {
		return true, nil
	}
	return false, nil
}

func (r *UserRepository) GetAllArtikel(ctx context.Context) ([]*model.ArtikelList, error) {

	query := `SELECT id, title, content, created_at FROM artikel`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artikels []*model.ArtikelList
	for rows.Next() {
		var artikel model.ArtikelList
		err := rows.Scan(&artikel.ID, &artikel.Title, &artikel.Content, &artikel.Date)
		if err != nil {
			return nil, err
		}
		artikels = append(artikels, &artikel)
	}
	return artikels, nil
}

func (r *UserRepository) GetArtikelById(ctx context.Context, id int) (*model.ArtikelDetail, error) {
	var artikel model.ArtikelDetail
	query := `SELECT title,created_at, content FROM artikel WHERE id = ?`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&artikel.Title, &artikel.Date, &artikel.Content)
		if err != nil {
			return nil, err
		}
		return &artikel, nil
	}
	return &artikel, errors.New("Artikel not found")
}
