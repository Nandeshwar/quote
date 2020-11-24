package repo

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/schema"
)

//go:generate mockgen -destination mock/mock_IRepo.go -source loginrepo.go IRepo
type IRepo interface {
	LoginInfo(user, password string) (err error)
}

/*
func (s SQLite3Repo) LoginInfo(user, password string) (err error) {
	query := fmt.Sprintf(`SELECT user, password FROM login WHERE user='%s' AND password='%s'`, user, password)
	logrus.WithFields(logrus.Fields{
		"query": space.ReplaceAllString(query, " "),
	}).Debugf("querying db")
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
*/

func (s SQLite3Repo) LoginInfo(user, password string) (err error) {
	type Login struct {
		User     string
		Password string
	}
	login := Login{}
	s.GORMDB.Config.NamingStrategy = schema.NamingStrategy{SingularTable: true}
	tx := s.GORMDB.WithContext(context.Background()).Where("user = ? AND password = ?", user, password).Debug().First(&login)
	if tx.Error != nil {
		return fmt.Errorf("error querying db, error=%v", tx.Error)
	}
	if tx.RowsAffected < 1 {
		return fmt.Errorf("invalid login credentials")
	}
	logrus.Info(login)
	return nil
}
