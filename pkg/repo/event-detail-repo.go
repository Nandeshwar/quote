package repo

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"quote/pkg/constants"
	"quote/pkg/model"
	"regexp"
	"strings"
)

// Regex for 1 and more white space
var space = regexp.MustCompile(`\s+`)

type IEventDetailRepo interface {
	CreateEventDetail(eventDetail model.EventDetail) (int64, error)
	GetEventDetailByTitleOrInfo(searchTxt string) ([]model.EventDetail, error)
	GetEventDetailByMonthDay(month, day int) ([]model.EventDetail, error)
}

func (s SQLite3Repo) CreateEventDetail(eventDetail model.EventDetail) (int64, error) {
	query := `INSERT INTO event_detail 
    				(day, month, year, title, info, type, created_at, updated_at) 
				VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	q := query
	tx, _ := s.DB.Begin()

	defer tx.Commit()

	logrus.WithFields(logrus.Fields{
		"query":      space.ReplaceAllString(query, " "),
		"day":        eventDetail.Day,
		"month":      eventDetail.Month,
		"year":       eventDetail.Year,
		"title":      eventDetail.Title,
		"info":       eventDetail.Info,
		"type":       eventDetail.Type,
		"created_at": eventDetail.CreationDate.Format(constants.DATE_FORMAT),
		"updated_at": eventDetail.UpdatedAt.Format(constants.DATE_FORMAT),
	}).Debugf("inserting record to db")
	var statement, err = tx.Prepare(q)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error preparing statements. query=%s, error=%v", query, err)
	}
	result, err := statement.Exec(
		eventDetail.Day,
		eventDetail.Month,
		eventDetail.Year,
		eventDetail.Title,
		eventDetail.Info,
		eventDetail.Type,
		eventDetail.CreationDate.Format(constants.DATE_FORMAT),
		eventDetail.UpdatedAt.Format(constants.DATE_FORMAT))
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error executing statements. query=%s, error=%v", query, err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error getting last inserted id. query=%s, error=%v", query, err)
	}

	query = `INSERT INTO event_detail_link (link_id, link, created_at, updated_at) VALUES (?, ?, ?, ?)`
	for _, l := range eventDetail.Links {
		logrus.WithFields(logrus.Fields{
			"query":      space.ReplaceAllString(query, " "),
			"link_id":    id,
			"link":       strings.TrimSpace(l),
			"create_at":  eventDetail.CreationDate.Format(constants.DATE_FORMAT),
			"updated_at": eventDetail.UpdatedAt.Format(constants.DATE_FORMAT),
		}).Debugf("inserting record to db")

		statement, err := tx.Prepare(query)
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("error preparing statements. query=%s, error=%v", query, err)
		}
		_, err = statement.Exec(id, strings.TrimSpace(l), eventDetail.CreationDate.Format(constants.DATE_FORMAT), eventDetail.UpdatedAt.Format(constants.DATE_FORMAT))
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("error executing statements. query=%s, error=%v", query, err)
		}
	}
	if len(eventDetail.Links) == 0 {
		statement, err := tx.Prepare(query)
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("error preparing statements. query=%s, error=%v", query, err)
		}
		_, err = statement.Exec(id, "no-link", eventDetail.CreationDate.Format(constants.DATE_FORMAT), eventDetail.UpdatedAt.Format(constants.DATE_FORMAT))
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

func (s SQLite3Repo) GetEventDetailByTitleOrInfo(searchTxt string) ([]model.EventDetail, error) {
	var eventDetailList []model.EventDetail
	query := `SELECT i.id,
					i.day,
					i.month,
					i.year,
					i.title,
					i.info,
					i.type,
					i.created_at,
					i.updated_at,
					l.link
				FROM 
					event_detail i, event_detail_link l
				WHERE 
					i.id = l.link_id AND (i.title like ? OR i.info like ?)`

	logrus.WithFields(logrus.Fields{
		"query": space.ReplaceAllString(query, " "),
		"arg1":  "%" + searchTxt + "%",
		"arg2":  "%" + searchTxt + "%",
	}).Debugf("fetching data from db")

	rows, err := s.DB.Query(query, "%"+searchTxt+"%", "%"+searchTxt+"%")
	if err != nil {
		return nil, fmt.Errorf("error querying db. query=%s, error=%v", query, err)
	}

	for rows.Next() {
		var eventDetailDB model.EventDetail
		err = rows.Scan(
			&eventDetailDB.ID,
			&eventDetailDB.Day,
			&eventDetailDB.Month,
			&eventDetailDB.Year,
			&eventDetailDB.Title,
			&eventDetailDB.Info,
			&eventDetailDB.Type,
			&eventDetailDB.CreationDate,
			&eventDetailDB.UpdatedAt,
			&eventDetailDB.URL)
		if err != nil {
			return nil, fmt.Errorf("error scanning result from db. query=%s, error=%v", query, err)
		}
		eventDetailList = append(eventDetailList, eventDetailDB)
	}
	logrus.Debugf("data fetch from database=%v", eventDetailList)

	return eventDetailList, nil
}

func (s SQLite3Repo) GetEventDetailByMonthDay(month, day int) ([]model.EventDetail, error) {
	var eventDetailList []model.EventDetail
	query := `SELECT e.id,
					e.day,
					e.month,
					e.year,
					e.title,
					e.info,
					e.type,
					e.created_at,
					e.updated_at,
					l.link
				FROM 
					event_detail e, event_detail_link l
				WHERE 
					e.id = l.link_id AND e.month=? AND e.day=?`
	q := query
	logrus.WithFields(logrus.Fields{
		"Query": space.ReplaceAllString(query, " "),
		"month": month,
		"day":   day,
	}).Debugf("fetching data from database")

	rows, err := s.DB.Query(q, month, day)
	if err != nil {
		return nil, fmt.Errorf("error querying db. query=%s, error=%v", query, err)
	}

	for rows.Next() {
		var eventDetailDB model.EventDetail
		err = rows.Scan(
			&eventDetailDB.ID,
			&eventDetailDB.Day,
			&eventDetailDB.Month,
			&eventDetailDB.Year,
			&eventDetailDB.Title,
			&eventDetailDB.Info,
			&eventDetailDB.Type,
			&eventDetailDB.CreationDate,
			&eventDetailDB.UpdatedAt,
			&eventDetailDB.URL)
		if err != nil {
			return nil, fmt.Errorf("error scanning result from db. query=%s, error=%v", query, err)
		}
		logrus.Debugf("data fetched from database=%v", eventDetailList)
		eventDetailList = append(eventDetailList, eventDetailDB)
	}

	return eventDetailList, nil
}
