package service

import (
	"fmt"
	"quote/pkg/constants"
	info2 "quote/pkg/info"
	"quote/pkg/model"
	"strings"
	"time"
)

type IInfo interface {
	ValidateForm(form model.InfoForm) error
	CreateNewInfo(form model.InfoForm) (int64, error)
	GetInfoByTitleOrInfo(searchTxt string) ([]info2.Info, error)
}

func (s QuoteService) ValidateForm(form model.InfoForm) error {
	createdAt := strings.TrimSpace(form.CreatedAt)
	if len(createdAt) > 0 {
		if len(createdAt) != 16 || createdAt[4] != '-' || createdAt[7] != '-' || createdAt[13] != ':' {
			return fmt.Errorf("wrong date and time format. please provide date in this format yyyy-mm-dd tt:mm")
		}
		_, err := time.Parse(constants.DATE_FORMAT, createdAt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s QuoteService) CreateNewInfo(form model.InfoForm) (int64, error) {
	var createdAt time.Time
	var err error

	if len(strings.TrimSpace(form.CreatedAt)) > 0 {
		createdAt, err = time.Parse(constants.DATE_FORMAT, form.CreatedAt)
		if err != nil {
			return 0, err
		}
	} else {
		createdAt = time.Now()
	}

	link := strings.TrimSpace(form.Link)
	var links []string
	if len(link) > 0 {
		links = strings.Split(link, ",")
	}
	info := info2.Info{
		Title:        form.Title,
		Info:         form.Info,
		Links:        links,
		CreationDate: createdAt,
		UpdatedDate:  time.Now(),
	}
	id, err := s.InfoRepo.CreateInfo(info)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s QuoteService) GetInfoByTitleOrInfo(searchTxt string) ([]info2.Info, error) {
	infoList, err := s.InfoRepo.GetInfoByTitleOrInfo(searchTxt)
	if err != nil {
		return nil, err
	}

	var distinctInfoList []info2.Info
	var links []string
	found := false
	for i := 0; i < len(infoList); i++ {
		if i+1 < len(infoList) && infoList[i].Title == infoList[i+1].Title && infoList[i].Info == infoList[i+1].Info {
			links = append(links, infoList[i].Link)
			found = true
		} else {
			found = false
		}

		if !found {
			links = append(links, infoList[i].Link)
			infoList[i].Links = links
			links = nil
			distinctInfoList = append(distinctInfoList, infoList[i])
		}
	}
	return distinctInfoList, nil
}
