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
	var movies []*models.Movie

	return movies, nil
}