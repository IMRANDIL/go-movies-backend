package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"time"
)

type PostgresDatabaseRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3


func (m *PostgresDatabaseRepo) Connection() *sql.DB {
	return m.DB
}




func (m *PostgresDatabaseRepo) AllMovies() ([]*models.Movie, error) {
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
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			
				
		)

		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)

	}

	return movies, nil
}