package repo

import "database/sql"

type SQLite3Repo struct {
	DB *sql.DB
}
