package repo

import (
	"fmt"
	"quote/pkg/constants"
	"quote/pkg/info"
)

type IInfoRepo interface {
	CreateInfo(info info.Info) (int64, error)
	GetInfoByTitleOrInfo(searchTxt string) ([]info.Info, error)
}

func (s SQLite3Repo) CreateInfo(info info.Info) (int64, error) {
	query := `INSERT INTO info (title, info, created_at, updated_at) VALUES (?, ?, ?, ?)`
	tx, _ := s.DB.Begin()

	defer tx.Commit()

	statement, err := tx.Prepare(query)
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

	if len(info.Links) > 0 {
		query := `INSERT INTO info_link (link_id, link, created_at, updated_at) VALUES (?, ?, ?, ?)`
		for _, l := range info.Links {
			statement, err := tx.Prepare(query)
			if err != nil {
				tx.Rollback()
				return 0, fmt.Errorf("error preparing statements. query=%s, error=%v", query, err)
			}
			_, err = statement.Exec(id, l, info.CreationDate.Format(constants.DATE_FORMAT), info.UpdatedDate.Format(constants.DATE_FORMAT))
			if err != nil {
				tx.Rollback()
				return 0, fmt.Errorf("error executing statements. query=%s, error=%v", query, err)
			}
		}
	}
	return id, nil
}

func (s SQLite3Repo) GetInfoByTitleOrInfo(searchTxt string) ([]info.Info, error) {
	var infoList []info.Info
	query := fmt.Sprintf(`SELECT i.id,
										i.title,
										i.info,
										i.created_at,
										i.updated_at,
										l.link
								FROM info i, info_link l
								WHERE i.id = l.link_id AND (i.title like '%s' OR i.info like '%s')`, "%"+searchTxt+"%", "%"+searchTxt+"%")
	fmt.Println("Query: ", query)
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying db. query=%s, error=%v", query, err)
	}

	var infoDB info.Info
	for rows.Next() {
		err = rows.Scan(&infoDB.ID, &infoDB.Title, &infoDB.Info, &infoDB.CreationDate, &infoDB.UpdatedDate, &infoDB.Link)
		if err != nil {
			return nil, fmt.Errorf("error scanning result from db. query=%s, error=%v", query, err)
		}
		infoList = append(infoList, infoDB)
	}

	return infoList, nil
}
