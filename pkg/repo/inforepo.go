package repo

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"quote/pkg/constants"
	"quote/pkg/model"
	"strings"
)

//go:generate mockgen -destination mock/mock_iinfo_repo.go -source inforepo.go IInfoRepo
type IInfoRepo interface {
	CreateInfo(info model.Info) (int64, error)
	GetInfoByTitleOrInfo(searchTxt string) ([]model.Info, error)
}

func (s SQLite3Repo) CreateInfo(info model.Info) (int64, error) {
	query := `INSERT INTO info (title, info, created_at, updated_at) VALUES (?, ?, ?, ?)`
	tx, _ := s.DB.Begin()

	qry := query
	defer tx.Commit()

	logrus.WithFields(logrus.Fields{
		"Query":     space.ReplaceAllString(query, " "),
		"title":     info.Title,
		"info":      info.Info,
		"create_at": info.CreationDate.Format(constants.DATE_FORMAT),
		"update_at": info.UpdatedDate.Format(constants.DATE_FORMAT),
	}).Debugf("inserting data")
	statement, err := tx.Prepare(qry)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error preparing statements. query=%s, error=%v", query, err)
	}
	result, err := statement.Exec(info.Title, info.Info, info.CreationDate.Format(constants.DATE_FORMAT), info.UpdatedDate.Format(constants.DATE_FORMAT))
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error executing statements. query=%s, error=%v", query, err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error getting last inserted id. query=%s, error=%v", query, err)
	}

	query = `INSERT INTO info_link (link_id, link, created_at, updated_at) VALUES (?, ?, ?, ?)`
	for _, l := range info.Links {
		logrus.WithFields(logrus.Fields{
			"Query":     space.ReplaceAllString(query, " "),
			"link_id":   id,
			"link":      strings.TrimSpace(l),
			"create_at": info.CreationDate.Format(constants.DATE_FORMAT),
			"update_at": info.UpdatedDate.Format(constants.DATE_FORMAT),
		}).Debugf("inserting data")

		statement, err := tx.Prepare(query)
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("error preparing statements. query=%s, error=%v", query, err)
		}
		_, err = statement.Exec(id, strings.TrimSpace(l), info.CreationDate.Format(constants.DATE_FORMAT), info.UpdatedDate.Format(constants.DATE_FORMAT))
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("error executing statements. query=%s, error=%v", query, err)
		}
	}
	if len(info.Links) == 0 {
		statement, err := tx.Prepare(query)
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("error preparing statements. query=%s, error=%v", query, err)
		}
		_, err = statement.Exec(id, "no-link", info.CreationDate.Format(constants.DATE_FORMAT), info.UpdatedDate.Format(constants.DATE_FORMAT))
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("error executing statements. query=%s, error=%v", query, err)
		}
	}
	logrus.WithFields(logrus.Fields{
		"id": id,
	}).Debug("event detail  record inserted to db successfully")
	return id, nil
}

func (s SQLite3Repo) GetInfoByTitleOrInfo(searchTxt string) ([]model.Info, error) {
	var infoList []model.Info
	query := `SELECT i.id,
						i.title,
						i.info,
						i.created_at,
						i.updated_at,
						l.link
				FROM info i, info_link l
				WHERE i.id = l.link_id AND (i.title like ? OR i.info like ?)`

	logrus.WithFields(logrus.Fields{
		"query": space.ReplaceAllString(query, " "),
		"arg1":  "%" + searchTxt + "%",
		"arg2":  "%" + searchTxt + "%",
	}).Debugf("fetching data from db")

	rows, err := s.DB.Query(query, "%"+searchTxt+"%", "%"+searchTxt+"%")
	if err != nil {
		return nil, fmt.Errorf("error querying db. query=%s, error=%v", query, err)
	}

	var infoDB model.Info
	for rows.Next() {
		err = rows.Scan(&infoDB.ID, &infoDB.Title, &infoDB.Info, &infoDB.CreationDate, &infoDB.UpdatedDate, &infoDB.Link)
		if err != nil {
			return nil, fmt.Errorf("error scanning result from db. query=%s, error=%v", query, err)
		}
		infoList = append(infoList, infoDB)
	}

	logrus.Debugf("data fetch from database=%v", infoList)

	return infoList, nil
}
