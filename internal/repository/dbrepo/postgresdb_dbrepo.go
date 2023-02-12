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
coalesce(image, ''), created_at, updated_at
from 
movies
order by 
title
`

rows, err := m.DB.QueryContext(ctx, query)

if err != nil {
	return nil, err
}

defer rows.Close()
	var movies []*models.Movie

	for rows.Next(){
		var movie models.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.CreatedAt,
			&movie.Description,
			&movie.Image,
			&movie.MPAARating,
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

	}

	return movies, nil
}