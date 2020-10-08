package repo

import (
	"database/sql"
	"fmt"
)

type IRepo interface {
	LoginInfo(user, password string) (err error)
}

func NewSqlite3Repo(sqlite3FileName string) (SQLite3Repo, error) {
	db, err := sql.Open("sqlite3", sqlite3FileName)
	if err != nil {
		return SQLite3Repo{}, fmt.Errorf("error opening file=%v, error=%v", sqlite3FileName, err)
	}
	return SQLite3Repo{DB: db}, nil
}

func (s SQLite3Repo) LoginInfo(user, password string) (err error) {
	query := fmt.Sprintf(`SELECT user, password FROM login WHERE user='%s' AND password='%s'`, user, password)
	rows, err := s.DB.Query(query)
	if err != nil {
		return fmt.Errorf("error querying db. query=%s, error=%v", query, err)
	}

	var dbUser string
	var dbPassword string
	rows.Next()
	err = rows.Scan(&dbUser, &dbPassword)
	if err != nil {
		return fmt.Errorf("error scanning result from db. query=%s, error=%v", query, err)
	}

	return nil
}
