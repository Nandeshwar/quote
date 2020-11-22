package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"quote/pkg/constants"
	"quote/pkg/model"
)

//go:generate mockgen -destination mock/mock_iinfo_repo.go -source inforepo.go IInfoRepo
type IInfoRepo interface {
	CreateInfo(ctx context.Context, info model.Info) (int64, error)
	GetInfoByTitleOrInfo(searchTxt string) ([]model.Info, error)
	UpdateInfoByID(info model.Info) error
	GetInfoByID(ctx context.Context, ID int64) ([]model.Info, error)
	GetInfoLinkIDs(links []string) ([]int64, error)
}

/*
func (s SQLite3Repo) CreateInfo(ctx context.Context, info model.Info) (int64, error) {
	query := `INSERT INTO info (title, info, created_at, updated_at) VALUES (?, ?, ?, ?)`
	tx, _ := s.DB.Begin()

	qry := query
	defer tx.Commit()

	logrus.WithFields(logrus.Fields{
		"Query":     space.ReplaceAllString(query, " "),
		"title":     info.Title,
		"info":      info.Info,
		"create_at": info.CreatedAt.Format(constants.DATE_FORMAT),
		"update_at": info.UpdatedAt.Format(constants.DATE_FORMAT),
	}).Debugf("inserting data")
	statement, err := tx.Prepare(qry)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error preparing statements. query=%s, error=%v", query, err)
	}
	result, err := statement.ExecContext(ctx, info.Title, info.Info, info.CreatedAt.Format(constants.DATE_FORMAT), info.UpdatedAt.Format(constants.DATE_FORMAT))
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
			"create_at": info.CreatedAt.Format(constants.DATE_FORMAT),
			"update_at": info.UpdatedAt.Format(constants.DATE_FORMAT),
		}).Debugf("inserting data")

		statement, err := tx.PrepareContext(ctx, query)
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("error preparing statements. query=%s, error=%v", query, err)
		}
		_, err = statement.ExecContext(ctx, id, strings.TrimSpace(l), info.CreatedAt.Format(constants.DATE_FORMAT), info.UpdatedAt.Format(constants.DATE_FORMAT))
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("error executing statements. query=%s, error=%v", query, err)
		}
		statement.Close()
	}
	if len(info.Links) == 0 {
		statement, err := tx.PrepareContext(ctx, query)
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("error preparing statements. query=%s, error=%v", query, err)
		}
		_, err = statement.ExecContext(ctx, id, "no-link", info.CreatedAt.Format(constants.DATE_FORMAT), info.UpdatedAt.Format(constants.DATE_FORMAT))
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("error executing statements. query=%s, error=%v", query, err)
		}
	}
	statement.Close()
	logrus.WithFields(logrus.Fields{
		"id": id,
	}).Debug("event detail  record inserted to db successfully")
	return id, nil
}
*/

func (s SQLite3Repo) CreateInfo(ctx context.Context, info model.Info) (int64, error) {

	infoGORM := model.InfoGORM{
		Title: info.Title,
		Info:  info.Info,
	}

	if len(info.Links) == 0 {
		info.Links = []string{"no-link"}
	}

	for _, link := range info.Links {
		infoGORM.InfoLinks = append(infoGORM.InfoLinks, model.InfoLinkGORM{
			Link: strings.TrimSpace(link),
		})
	}

	tx := s.GORMDB.WithContext(ctx).Create(&infoGORM).Debug()

	if tx.Error != nil {
		return 0, tx.Error
	}

	return infoGORM.ID, nil

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
		err = rows.Scan(&infoDB.ID, &infoDB.Title, &infoDB.Info, &infoDB.CreatedAt, &infoDB.UpdatedAt, &infoDB.Link)
		if err != nil {
			return nil, fmt.Errorf("error scanning result from db. query=%s, error=%v", query, err)
		}
		infoList = append(infoList, infoDB)
	}

	logrus.Debugf("data fetch from database=%v", infoList)

	return infoList, nil
}

func (s SQLite3Repo) GetInfoLinkIDs(links []string) ([]int64, error) {
	query := `SELECT link_id
				FROM info_link
				WHERE link in (?)`

	logrus.WithFields(logrus.Fields{
		"query": space.ReplaceAllString(query, " "),
		"arg1":  links,
	}).Debugf("fetching data from db")

	rows, err := s.DB.Query(query, strings.Join(links, ","))
	if err != nil {
		return nil, fmt.Errorf("error querying db. query=%s, error=%v", query, err)
	}

	var linkID int64
	var linkIDs []int64

	for rows.Next() {
		err = rows.Scan(&linkID)
		if err != nil {
			return nil, fmt.Errorf("error scanning result from db. query=%s, error=%v", query, err)
		}
		linkIDs = append(linkIDs, linkID)
	}

	logrus.Debugf("data fetch from database=%v", linkIDs)

	return linkIDs, nil
}

func (s SQLite3Repo) GetInfoByID(ctx context.Context, ID int64) ([]model.Info, error) {
	var infoList []model.Info
	query := `SELECT i.id,
						i.title,
						i.info,
						i.created_at,
						i.updated_at,
						l.link
				FROM info i, info_link l
				WHERE i.id = l.link_id AND i.id = ? ORDER BY i.id, l.link_id`

	logrus.WithFields(logrus.Fields{
		"query": space.ReplaceAllString(query, " "),
		"arg":   ID,
	}).Debugf("fetching data from db")

	rows, err := s.DB.QueryContext(ctx, query, ID)
	if err != nil {
		return nil, fmt.Errorf("error querying db. query=%s, error=%v", query, err)
	}

	var infoDB model.Info
	for rows.Next() {
		err = rows.Scan(&infoDB.ID, &infoDB.Title, &infoDB.Info, &infoDB.CreatedAt, &infoDB.UpdatedAt, &infoDB.Link)
		if err != nil {
			return nil, fmt.Errorf("error scanning result from db. query=%s, error=%v", query, err)
		}
		infoList = append(infoList, infoDB)
	}

	logrus.Debugf("data fetch from database=%v", infoList)

	return infoList, nil
}

func (s SQLite3Repo) findID(ID int64, tx *sql.Tx) error {
	query := `SELECT id
				FROM info
				WHERE id = ?`

	logrus.WithFields(logrus.Fields{
		"query": space.ReplaceAllString(query, " "),
		"arg":   ID,
	}).Debugf("fetching data from db")

	rows, err := tx.Query(query, ID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error querying db. query=%s, error=%v", query, err)
	}

	if !rows.Next() {
		tx.Rollback()
		logrus.WithError(err).WithField("ID", ID).Errorf("id does not exist in database")
		return fmt.Errorf("id=%d does not exist in database", ID)
	}

	return nil
}

func (s SQLite3Repo) updateInfo(info model.Info, tx *sql.Tx) error {
	updateInfoQuery := `UPDATE info 
						SET title = ?, info = ?, updated_at = ?
						WHERE ID = ?`

	logrus.WithFields(logrus.Fields{
		"Query":      space.ReplaceAllString(updateInfoQuery, " "),
		"ID":         info.ID,
		"title":      info.Title,
		"info":       info.Info,
		"updated_at": info.UpdatedAt.Format(constants.DATE_FORMAT),
	}).Debugf("updating data")

	statement, err := tx.Prepare(updateInfoQuery)
	_, err = tx.Stmt(statement).Exec(info.Title, info.Info, info.UpdatedAt.Format(constants.DATE_FORMAT), info.ID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error preparing statements. query=%s, error=%v", updateInfoQuery, err)
	}
	statement.Close()

	return nil
}

func (s SQLite3Repo) insertLinks(infoID int64, links []string, updatedAt time.Time, tx *sql.Tx) error {
	query := `INSERT INTO info_link (link_id, link, updated_at) 
				VALUES(?, ?, ?)`
	qry := query

	stmt, err := tx.Prepare(qry)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error preparing statement. query=%s, error=%v", qry, err)
	}

	for _, link := range links {
		_, err := stmt.Exec(infoID, link, updatedAt)
		if err != nil {
			return fmt.Errorf("error executing statement. query=%s, error=%v", qry, err)
		}
	}
	stmt.Close()
	return nil
}

func (s SQLite3Repo) deleteLinks(infoID int64, tx *sql.Tx) error {
	query := `DELETE FROM info_link 
				WHERE link_id = ?`
	qry := query

	stmt, err := tx.Prepare(qry)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error preparing statement. query=%s, error=%v", qry, err)
	}

	_, err = stmt.Exec(infoID)
	if err != nil {
		return fmt.Errorf("error executing statement. query=%s, error=%v", qry, err)
	}
	stmt.Close()
	return nil
}

func (s SQLite3Repo) UpdateInfoByID(info model.Info) error {
	tx, err := s.DB.Begin()

	if err != nil {
		return fmt.Errorf("error starting transaction for query while updating info=%v", err)
	}
	defer tx.Commit()

	err = s.findID(info.ID, tx)
	if err != nil {
		return err
	}

	err = s.updateInfo(info, tx)
	if err != nil {
		return err
	}

	err = s.deleteLinks(info.ID, tx)
	if err != nil {
		return err
	}

	err = s.insertLinks(info.ID, info.Links, info.UpdatedAt, tx)
	if err != nil {
		return err
	}

	return nil
}
