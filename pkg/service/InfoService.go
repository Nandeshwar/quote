package service

import (
	"fmt"
	"quote/pkg/constants"
	"quote/pkg/model"
	"strings"
	"time"
)

type IInfo interface {
	ValidateForm(form model.InfoForm) error
	CreateNewInfo(form model.InfoForm) (int64, error)
	GetInfoByTitleOrInfo(searchTxt string) ([]model.Info, error)
}

func (s QuoteService) ValidateForm(form model.InfoForm) error {
	createdAt := strings.TrimSpace(form.CreatedAt)
	if len(createdAt) > 0 {
		if len(createdAt) != 16 || createdAt[4] != '-' || createdAt[7] != '-' || createdAt[13] != ':' {
			return fmt.Errorf("wrong date and time format. given date=%s, please provide date in this format yyyy-mm-dd tt:mm", createdAt)
		}
		_, err := time.Parse(constants.DATE_FORMAT, createdAt)
		if err != nil {
			return err
		}
	}

	link := strings.TrimSpace(form.Link)
	if len(link) > 0 {
		for _, link := range strings.Split(link, ",") {
			link = strings.TrimSpace(link)
			if len(link) < 4 {
				return fmt.Errorf("pipeline(|) seperated links value must start with http or https. link could not be less than 4")
			}
			if link[0:4] != "http" {
				return fmt.Errorf("pipeline(|) seperated links value must start with http or https")
			}

			if link[len(link)-1] == '"' || link[len(link)-1] == '\'' || link[len(link)-1] == '.' {
				return fmt.Errorf("pipeline(|) seperated link's value should not ended with (\", ', .)")
			}
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
	info := model.Info{
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

func (s QuoteService) GetInfoByTitleOrInfo(searchTxt string) ([]model.Info, error) {
	infoList, err := s.InfoRepo.GetInfoByTitleOrInfo(searchTxt)
	if err != nil {
		return nil, err
	}

	infoListSorted := model.SortInfoByID(infoList)

	var distinctInfoList []model.Info
	var links []string
	found := false
	for i := 0; i < len(infoListSorted); i++ {
		if i+1 < len(infoListSorted) && infoListSorted[i].ID == infoListSorted[i+1].ID {
			links = append(links, infoListSorted[i].Link)
			found = true
		} else {
			found = false
		}

		if !found {
			links = append(links, infoListSorted[i].Link)
			infoListSorted[i].Links = links
			links = nil
			distinctInfoList = append(distinctInfoList, infoListSorted[i])
		}
	}
	return distinctInfoList, nil
}
