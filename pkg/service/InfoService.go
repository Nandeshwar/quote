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
	CreateNewInfo(form model.InfoForm) error
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

func (s QuoteService) CreateNewInfo(form model.InfoForm) error {
	var createdAt time.Time
	var err error

	if len(strings.TrimSpace(form.CreatedAt)) > 0 {
		createdAt, err = time.Parse(constants.DATE_FORMAT, form.CreatedAt)
		if err != nil {
			return err
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
		Link:         links,
		CreationDate: createdAt,
		UpdatedDate:  time.Now(),
	}
	err = s.InfoRepo.CreateInfo(info)
	if err != nil {
		return err
	}
	return nil
}
