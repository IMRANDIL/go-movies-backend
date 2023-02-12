package dbrepo

import "database/sql"

type DatabaseRepo struct {
	DB *sql.DB
}