package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"time"
)

type DatabaseRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3


func (m *DatabaseRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()

query := `
SELECT id, title, release_date,
runtime, mpaa_rating, description,
coalesce(image, ''),



`


	var movies []*models.Movie

	return movies, nil
}