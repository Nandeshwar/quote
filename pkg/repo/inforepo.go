package repo

import (
	"fmt"
	"quote/pkg/constants"
	"quote/pkg/info"
)

type IInfoRepo interface {
	CreateInfo(info info.Info) error
}

func (s SQLite3Repo) CreateInfo(info info.Info) error {
	query := `INSERT INTO info (title, info, created_at, updated_at) VALUES (?, ?, ?, ?)`
	tx, _ := s.DB.Begin()

	defer tx.Commit()

	statement, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error preparing statements. query=%s, error=%v", query, err)
	}
	result, err := statement.Exec(info.Title, info.Info, info.CreationDate.Format(constants.DATE_FORMAT), info.UpdatedDate.Format(constants.DATE_FORMAT))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing statements. query=%s, error=%v", query, err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error getting last inserted id. query=%s, error=%v", query, err)
	}

	if len(info.Link) > 0 {
		query := `INSERT INTO info_link (link_id, link, created_at, updated_at) VALUES (?, ?, ?, ?)`
		for _, l := range info.Link {
			statement, err := tx.Prepare(query)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("error preparing statements. query=%s, error=%v", query, err)
			}
			_, err = statement.Exec(id, l, info.CreationDate.Format(constants.DATE_FORMAT), info.UpdatedDate.Format(constants.DATE_FORMAT))
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("error executing statements. query=%s, error=%v", query, err)
			}
		}
	}
	return nil
}
