package repo

import (
	"fmt"
)

type IRepo interface {
	LoginInfo(user, password string) (err error)
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