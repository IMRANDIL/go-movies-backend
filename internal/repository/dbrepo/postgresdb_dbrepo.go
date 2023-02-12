package dbrepo

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo struct {
	DB *sql.DB
}

func (m *DatabaseRepo) AllMovies() ([]*models.Movie, error) {
	var movies []*models.Movie

	return movies, nil
}