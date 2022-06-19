package repository

import (
	"context"
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
